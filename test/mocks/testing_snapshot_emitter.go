package mocks

import (
	"sync"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
)

var (
	mTestingSnapshotIn  = stats.Int64("testing_snap_emitter/snap_in", "The number of snapshots in", "1")
	mTestingSnapshotOut = stats.Int64("testing_snap_emitter/snap_out", "The number of snapshots out", "1")

	testingsnapshotInView = &view.View{
		Name:        "testing_snap_emitter/snap_in",
		Measure:     mTestingSnapshotIn,
		Description: "The number of snapshots updates coming in",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	testingsnapshotOutView = &view.View{
		Name:        "testing_snap_emitter/snap_out",
		Measure:     mTestingSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(testingsnapshotInView, testingsnapshotOutView)
}

type TestingEmitter interface {
	Register() error
	MockResource() MockResourceClient
	FakeResource() FakeResourceClient
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *TestingSnapshot, <-chan error, error)
}

func NewTestingEmitter(mockResourceClient MockResourceClient, fakeResourceClient FakeResourceClient) TestingEmitter {
	return NewTestingEmitterWithEmit(mockResourceClient, fakeResourceClient, make(chan struct{}))
}

func NewTestingEmitterWithEmit(mockResourceClient MockResourceClient, fakeResourceClient FakeResourceClient, emit <-chan struct{}) TestingEmitter {
	return &testingEmitter{
		mockResource: mockResourceClient,
		fakeResource: fakeResourceClient,
		forceEmit:    emit,
	}
}

type testingEmitter struct {
	forceEmit    <-chan struct{}
	mockResource MockResourceClient
	fakeResource FakeResourceClient
}

func (c *testingEmitter) Register() error {
	if err := c.mockResource.Register(); err != nil {
		return err
	}
	if err := c.fakeResource.Register(); err != nil {
		return err
	}
	return nil
}

func (c *testingEmitter) MockResource() MockResourceClient {
	return c.mockResource
}

func (c *testingEmitter) FakeResource() FakeResourceClient {
	return c.fakeResource
}

func (c *testingEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *TestingSnapshot, <-chan error, error) {
	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for MockResource */
	type mockResourceListWithNamespace struct {
		list      MockResourceList
		namespace string
	}
	mockResourceChan := make(chan mockResourceListWithNamespace)
	/* Create channel for FakeResource */
	type fakeResourceListWithNamespace struct {
		list      FakeResourceList
		namespace string
	}
	fakeResourceChan := make(chan fakeResourceListWithNamespace)

	for _, namespace := range watchNamespaces {
		/* Setup watch for MockResource */
		mockResourceNamespacesChan, mockResourceErrs, err := c.mockResource.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting MockResource watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, mockResourceErrs, namespace+"-mocks")
		}(namespace)
		/* Setup watch for FakeResource */
		fakeResourceNamespacesChan, fakeResourceErrs, err := c.fakeResource.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting FakeResource watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, fakeResourceErrs, namespace+"-fakes")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case mockResourceList := <-mockResourceNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case mockResourceChan <- mockResourceListWithNamespace{list: mockResourceList, namespace: namespace}:
					}
				case fakeResourceList := <-fakeResourceNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case fakeResourceChan <- fakeResourceListWithNamespace{list: fakeResourceList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}

	snapshots := make(chan *TestingSnapshot)
	go func() {
		currentSnapshot := TestingSnapshot{}
		sync := func(newSnapshot TestingSnapshot) {
			if currentSnapshot.Hash() == newSnapshot.Hash() {
				return
			}
			currentSnapshot = newSnapshot
			sentSnapshot := currentSnapshot.Clone()

			stats.Record(ctx, mTestingSnapshotOut.M(1))
			snapshots <- &sentSnapshot
		}
		for {
			select {
			case <-ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case mockResourceNamespacedList := <-mockResourceChan:
				namespace := mockResourceNamespacedList.namespace
				mockResourceList := mockResourceNamespacedList.list

				newSnapshot := currentSnapshot.Clone()
				newSnapshot.Mocks.Clear(namespace)
				newSnapshot.Mocks.Add(mockResourceList...)
				sync(newSnapshot)
			case fakeResourceNamespacedList := <-fakeResourceChan:
				namespace := fakeResourceNamespacedList.namespace
				fakeResourceList := fakeResourceNamespacedList.list

				newSnapshot := currentSnapshot.Clone()
				newSnapshot.Fakes.Clear(namespace)
				newSnapshot.Fakes.Add(fakeResourceList...)
				sync(newSnapshot)
			}

			// if we got here its because a new entry in the channel
			stats.Record(ctx, mTestingSnapshotIn.M(1))
		}
	}()
	return snapshots, errs, nil
}
