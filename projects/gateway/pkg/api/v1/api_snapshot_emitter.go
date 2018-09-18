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
	Gateway() GatewayClient
	VirtualService() VirtualServiceClient
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *ApiSnapshot, <-chan error, error)
}

func NewApiEmitter(gatewayClient GatewayClient, virtualServiceClient VirtualServiceClient) ApiEmitter {
	return NewApiEmitterWithEmit(gatewayClient, virtualServiceClient, make(chan struct{}))
}

func NewApiEmitterWithEmit(gatewayClient GatewayClient, virtualServiceClient VirtualServiceClient, emit <-chan struct{}) ApiEmitter {
	return &apiEmitter{
		gateway:        gatewayClient,
		virtualService: virtualServiceClient,
		forceEmit:      emit,
	}
}

type apiEmitter struct {
	forceEmit      <-chan struct{}
	gateway        GatewayClient
	virtualService VirtualServiceClient
}

func (c *apiEmitter) Register() error {
	if err := c.gateway.Register(); err != nil {
		return err
	}
	if err := c.virtualService.Register(); err != nil {
		return err
	}
	return nil
}

func (c *apiEmitter) Gateway() GatewayClient {
	return c.gateway
}

func (c *apiEmitter) VirtualService() VirtualServiceClient {
	return c.virtualService
}

func (c *apiEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *ApiSnapshot, <-chan error, error) {
	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for Gateway */
	type gatewayListWithNamespace struct {
		list      GatewayList
		namespace string
	}
	gatewayChan := make(chan gatewayListWithNamespace)
	/* Create channel for VirtualService */
	type virtualServiceListWithNamespace struct {
		list      VirtualServiceList
		namespace string
	}
	virtualServiceChan := make(chan virtualServiceListWithNamespace)

	for _, namespace := range watchNamespaces {
		/* Setup watch for Gateway */
		gatewayNamespacesChan, gatewayErrs, err := c.gateway.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Gateway watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, gatewayErrs, namespace+"-gateways")
		}(namespace)
		/* Setup watch for VirtualService */
		virtualServiceNamespacesChan, virtualServiceErrs, err := c.virtualService.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting VirtualService watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, virtualServiceErrs, namespace+"-virtualServices")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case gatewayList := <-gatewayNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case gatewayChan <- gatewayListWithNamespace{list: gatewayList, namespace: namespace}:
					}
				case virtualServiceList := <-virtualServiceNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case virtualServiceChan <- virtualServiceListWithNamespace{list: virtualServiceList, namespace: namespace}:
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
		      gatewayNamespacedList := <- gatewayChan:
		   	namespace := gatewayNamespacedList.namespace
		   	gatewayList := gatewayNamespacedList.list

		   	currentSnapshot.Gateways.Clear(namespace)
		   	currentSnapshot.Gateways.Add(gatewayList...)
		      virtualServiceNamespacedList := <- virtualServiceChan:
		   	namespace := virtualServiceNamespacedList.namespace
		   	virtualServiceList := virtualServiceNamespacedList.list

		   	currentSnapshot.VirtualServices.Clear(namespace)
		   	currentSnapshot.VirtualServices.Add(virtualServiceList...)
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
			case gatewayNamespacedList := <-gatewayChan:
				namespace := gatewayNamespacedList.namespace
				gatewayList := gatewayNamespacedList.list

				currentSnapshot.Gateways.Clear(namespace)
				currentSnapshot.Gateways.Add(gatewayList...)
			case virtualServiceNamespacedList := <-virtualServiceChan:
				namespace := virtualServiceNamespacedList.namespace
				virtualServiceList := virtualServiceNamespacedList.list

				currentSnapshot.VirtualServices.Clear(namespace)
				currentSnapshot.VirtualServices.Add(virtualServiceList...)
			}

			// if we got here its because a new entry in the channel
			stats.Record(ctx, mApiSnapshotIn.M(1))
		}
	}()
	return snapshots, errs, nil
}
