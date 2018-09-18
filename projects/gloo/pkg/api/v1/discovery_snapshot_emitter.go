package v1

import (
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
)

var (
	mDiscoverySnapshotIn  = stats.Int64("discovery_snap_emitter/snap_in", "The number of snapshots in", "1")
	mDiscoverySnapshotOut = stats.Int64("discovery_snap_emitter/snap_out", "The number of snapshots out", "1")

	discoverysnapshotInView = &view.View{
		Name:        "discovery_snap_emitter/snap_in",
		Measure:     mDiscoverySnapshotIn,
		Description: "The number of snapshots updates coming in",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	discoverysnapshotOutView = &view.View{
		Name:        "discovery_snap_emitter/snap_out",
		Measure:     mDiscoverySnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(discoverysnapshotInView, discoverysnapshotOutView)
}

type DiscoveryEmitter interface {
	Register() error
	Secret() SecretClient
	Upstream() UpstreamClient
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *DiscoverySnapshot, <-chan error, error)
}

func NewDiscoveryEmitter(secretClient SecretClient, upstreamClient UpstreamClient) DiscoveryEmitter {
	return NewDiscoveryEmitterWithEmit(secretClient, upstreamClient, make(chan struct{}))
}

func NewDiscoveryEmitterWithEmit(secretClient SecretClient, upstreamClient UpstreamClient, emit <-chan struct{}) DiscoveryEmitter {
	return &discoveryEmitter{
		secret:    secretClient,
		upstream:  upstreamClient,
		forceEmit: emit,
	}
}

type discoveryEmitter struct {
	forceEmit <-chan struct{}
	secret    SecretClient
	upstream  UpstreamClient
}

func (c *discoveryEmitter) Register() error {
	if err := c.secret.Register(); err != nil {
		return err
	}
	if err := c.upstream.Register(); err != nil {
		return err
	}
	return nil
}

func (c *discoveryEmitter) Secret() SecretClient {
	return c.secret
}

func (c *discoveryEmitter) Upstream() UpstreamClient {
	return c.upstream
}

func (c *discoveryEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *DiscoverySnapshot, <-chan error, error) {
	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for Secret */
	type secretListWithNamespace struct {
		list      SecretList
		namespace string
	}
	secretChan := make(chan secretListWithNamespace)
	/* Create channel for Upstream */
	type upstreamListWithNamespace struct {
		list      UpstreamList
		namespace string
	}
	upstreamChan := make(chan upstreamListWithNamespace)

	for _, namespace := range watchNamespaces {
		/* Setup watch for Secret */
		secretNamespacesChan, secretErrs, err := c.secret.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Secret watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, secretErrs, namespace+"-secrets")
		}(namespace)
		/* Setup watch for Upstream */
		upstreamNamespacesChan, upstreamErrs, err := c.upstream.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Upstream watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, upstreamErrs, namespace+"-upstreams")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case secretList := <-secretNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case secretChan <- secretListWithNamespace{list: secretList, namespace: namespace}:
					}
				case upstreamList := <-upstreamNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case upstreamChan <- upstreamListWithNamespace{list: upstreamList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}

	snapshots := make(chan *DiscoverySnapshot)
	go func() {
		originalSnapshot := DiscoverySnapshot{}
		currentSnapshot := originalSnapshot.Clone()
		timer := time.NewTicker(time.Second * 5)
		sync := func() {
			if originalSnapshot.Hash() == currentSnapshot.Hash() {
				return
			}
			originalSnapshot = currentSnapshot.Clone()
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}

		/* TODO (yuval-k): figure out how to make this work to avoid a stale snapshot.
		   		// construct the first snapshot from all the configs that are currently there
		   		// that guarantees that the first snapshot contains all the data.
		   		for range watchNamespaces {
		      secretNamespacedList := <- secretChan:
		   	namespace := secretNamespacedList.namespace
		   	secretList := secretNamespacedList.list

		   	currentSnapshot.Secrets.Clear(namespace)
		   	currentSnapshot.Secrets.Add(secretList...)
		      upstreamNamespacedList := <- upstreamChan:
		   	namespace := upstreamNamespacedList.namespace
		   	upstreamList := upstreamNamespacedList.list

		   	currentSnapshot.Upstreams.Clear(namespace)
		   	currentSnapshot.Upstreams.Add(upstreamList...)
		   		}
		*/

		for {
			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case secretNamespacedList := <-secretChan:
				namespace := secretNamespacedList.namespace
				secretList := secretNamespacedList.list

				currentSnapshot.Secrets.Clear(namespace)
				currentSnapshot.Secrets.Add(secretList...)
			case upstreamNamespacedList := <-upstreamChan:
				namespace := upstreamNamespacedList.namespace
				upstreamList := upstreamNamespacedList.list

				currentSnapshot.Upstreams.Clear(namespace)
				currentSnapshot.Upstreams.Add(upstreamList...)
			}

			// if we got here its because a new entry in the channel
			stats.Record(ctx, mDiscoverySnapshotIn.M(1))
		}
	}()
	return snapshots, errs, nil
}
