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
	mApiSnapshotIn  = stats.Int64("api_snap_emitter/snap_in", "The number of snapshots in", "1")
	mApiSnapshotOut = stats.Int64("api_snap_emitter/snap_out", "The number of snapshots out", "1")

	apisnapshotInView = &view.View{
		Name:        "api_snap_emitter/snap_in",
		Measure:     mApiSnapshotIn,
		Description: "The number of snapshots updates coming in",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	apisnapshotOutView = &view.View{
		Name:        "api_snap_emitter/snap_out",
		Measure:     mApiSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(apisnapshotInView, apisnapshotOutView)
}

type ApiEmitter interface {
	Register() error
	ResolverMap() ResolverMapClient
	Schema() SchemaClient
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *ApiSnapshot, <-chan error, error)
}

func NewApiEmitter(resolverMapClient ResolverMapClient, schemaClient SchemaClient) ApiEmitter {
	return NewApiEmitterWithEmit(resolverMapClient, schemaClient, make(chan struct{}))
}

func NewApiEmitterWithEmit(resolverMapClient ResolverMapClient, schemaClient SchemaClient, emit <-chan struct{}) ApiEmitter {
	return &apiEmitter{
		resolverMap: resolverMapClient,
		schema:      schemaClient,
		forceEmit:   emit,
	}
}

type apiEmitter struct {
	forceEmit   <-chan struct{}
	resolverMap ResolverMapClient
	schema      SchemaClient
}

func (c *apiEmitter) Register() error {
	if err := c.resolverMap.Register(); err != nil {
		return err
	}
	if err := c.schema.Register(); err != nil {
		return err
	}
	return nil
}

func (c *apiEmitter) ResolverMap() ResolverMapClient {
	return c.resolverMap
}

func (c *apiEmitter) Schema() SchemaClient {
	return c.schema
}

func (c *apiEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *ApiSnapshot, <-chan error, error) {
	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for ResolverMap */
	type resolverMapListWithNamespace struct {
		list      ResolverMapList
		namespace string
	}
	resolverMapChan := make(chan resolverMapListWithNamespace)
	/* Create channel for Schema */
	type schemaListWithNamespace struct {
		list      SchemaList
		namespace string
	}
	schemaChan := make(chan schemaListWithNamespace)

	for _, namespace := range watchNamespaces {
		/* Setup watch for ResolverMap */
		resolverMapNamespacesChan, resolverMapErrs, err := c.resolverMap.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting ResolverMap watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, resolverMapErrs, namespace+"-resolverMaps")
		}(namespace)
		/* Setup watch for Schema */
		schemaNamespacesChan, schemaErrs, err := c.schema.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Schema watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, schemaErrs, namespace+"-schemas")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case resolverMapList := <-resolverMapNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case resolverMapChan <- resolverMapListWithNamespace{list: resolverMapList, namespace: namespace}:
					}
				case schemaList := <-schemaNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case schemaChan <- schemaListWithNamespace{list: schemaList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}

	snapshots := make(chan *ApiSnapshot)
	go func() {
		originalSnapshot := ApiSnapshot{}
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
		      resolverMapNamespacedList := <- resolverMapChan:
		   	namespace := resolverMapNamespacedList.namespace
		   	resolverMapList := resolverMapNamespacedList.list

		   	currentSnapshot.ResolverMaps.Clear(namespace)
		   	currentSnapshot.ResolverMaps.Add(resolverMapList...)
		      schemaNamespacedList := <- schemaChan:
		   	namespace := schemaNamespacedList.namespace
		   	schemaList := schemaNamespacedList.list

		   	currentSnapshot.Schemas.Clear(namespace)
		   	currentSnapshot.Schemas.Add(schemaList...)
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
			case resolverMapNamespacedList := <-resolverMapChan:
				namespace := resolverMapNamespacedList.namespace
				resolverMapList := resolverMapNamespacedList.list

				currentSnapshot.ResolverMaps.Clear(namespace)
				currentSnapshot.ResolverMaps.Add(resolverMapList...)
			case schemaNamespacedList := <-schemaChan:
				namespace := schemaNamespacedList.namespace
				schemaList := schemaNamespacedList.list

				currentSnapshot.Schemas.Clear(namespace)
				currentSnapshot.Schemas.Add(schemaList...)
			}

			// if we got here its because a new entry in the channel
			stats.Record(ctx, mApiSnapshotIn.M(1))
		}
	}()
	return snapshots, errs, nil
}
