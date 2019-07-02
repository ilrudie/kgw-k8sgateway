package virtualservicesvc_test

import (
	"context"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	. "github.com/solo-io/go-utils/testutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	v1 "github.com/solo-io/solo-projects/projects/grpcserver/api/v1"
	mock_settings "github.com/solo-io/solo-projects/projects/grpcserver/server/internal/settings/mocks"
	. "github.com/solo-io/solo-projects/projects/grpcserver/server/service/internal/testutils"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/virtualservicesvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/virtualservicesvc/mocks"
	mock_mutation "github.com/solo-io/solo-projects/projects/grpcserver/server/service/virtualservicesvc/mutation/mocks"
	"google.golang.org/grpc"
)

var (
	grpcServer            *grpc.Server
	conn                  *grpc.ClientConn
	apiserver             v1.VirtualServiceApiServer
	client                v1.VirtualServiceApiClient
	mockCtrl              *gomock.Controller
	virtualServiceClient  *mocks.MockVirtualServiceClient
	mutator               *mock_mutation.MockMutator
	mutationFactory       *mock_mutation.MockMutationFactory
	settingsValues        *mock_settings.MockValuesClient
	testErr               = errors.Errorf("test-err")
	uint32Zero, uint32One = uint32(0), uint32(1)
	metadata              = core.Metadata{
		Namespace: "ns",
		Name:      "name",
	}
	ref = metadata.Ref()
)

var _ = Describe("ServiceTest", func() {

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		virtualServiceClient = mocks.NewMockVirtualServiceClient(mockCtrl)
		mutator = mock_mutation.NewMockMutator(mockCtrl)
		mutationFactory = mock_mutation.NewMockMutationFactory(mockCtrl)
		settingsValues = mock_settings.NewMockValuesClient(mockCtrl)
		apiserver = virtualservicesvc.NewVirtualServiceGrpcService(context.TODO(), virtualServiceClient, settingsValues, mutator, mutationFactory)

		grpcServer, conn = MustRunGrpcServer(func(s *grpc.Server) { v1.RegisterVirtualServiceApiServer(s, apiserver) })
		client = v1.NewVirtualServiceApiClient(conn)
	})

	AfterEach(func() {
		grpcServer.Stop()
		mockCtrl.Finish()
	})

	Describe("GetVirtualService", func() {
		It("works when the virtual service client works", func() {
			virtualService := gatewayv1.VirtualService{
				Status:   core.Status{State: core.Status_Accepted},
				Metadata: metadata,
			}

			virtualServiceClient.EXPECT().
				Read(metadata.Namespace, metadata.Name, clients.ReadOpts{Ctx: context.TODO()}).
				Return(&virtualService, nil)

			request := &v1.GetVirtualServiceRequest{Ref: &ref}
			actual, err := client.GetVirtualService(context.TODO(), request)
			Expect(err).NotTo(HaveOccurred())
			expected := &v1.GetVirtualServiceResponse{VirtualService: &virtualService}
			ExpectEqualProtoMessages(actual, expected)
		})

		It("errors when the virtual service client errors", func() {
			virtualServiceClient.EXPECT().
				Read(metadata.Namespace, metadata.Name, clients.ReadOpts{Ctx: context.TODO()}).
				Return(nil, testErr)

			request := &v1.GetVirtualServiceRequest{Ref: &ref}
			_, err := client.GetVirtualService(context.TODO(), request)
			Expect(err).To(HaveOccurred())
			expectedErr := virtualservicesvc.FailedToReadVirtualServiceError(testErr, &ref)
			Expect(err.Error()).To(ContainSubstring(expectedErr.Error()))
		})
	})

	Describe("ListVirtualServices", func() {
		It("works when the virtual service client works", func() {
			ns1, ns2 := "one", "two"
			virtualService1 := gatewayv1.VirtualService{
				Status:   core.Status{State: core.Status_Accepted},
				Metadata: core.Metadata{Namespace: ns1},
			}
			virtualService2 := gatewayv1.VirtualService{
				Status:   core.Status{State: core.Status_Pending},
				Metadata: core.Metadata{Namespace: ns2},
			}

			virtualServiceClient.EXPECT().
				List(ns1, clients.ListOpts{Ctx: context.TODO()}).
				Return([]*gatewayv1.VirtualService{&virtualService1}, nil)
			virtualServiceClient.EXPECT().
				List(ns2, clients.ListOpts{Ctx: context.TODO()}).
				Return([]*gatewayv1.VirtualService{&virtualService2}, nil)

			request := &v1.ListVirtualServicesRequest{Namespaces: []string{ns1, ns2}}
			actual, err := client.ListVirtualServices(context.TODO(), request)
			Expect(err).NotTo(HaveOccurred())
			expected := &v1.ListVirtualServicesResponse{VirtualServices: []*gatewayv1.VirtualService{&virtualService1, &virtualService2}}
			ExpectEqualProtoMessages(actual, expected)
		})

		It("errors when the virtual service client errors", func() {
			ns := "ns"

			virtualServiceClient.EXPECT().
				List(ns, clients.ListOpts{Ctx: context.TODO()}).
				Return(nil, testErr)

			request := &v1.ListVirtualServicesRequest{Namespaces: []string{ns}}
			_, err := client.ListVirtualServices(context.TODO(), request)
			Expect(err).To(HaveOccurred())
			expectedErr := virtualservicesvc.FailedToListVirtualServicesError(testErr, ns)
			Expect(err.Error()).To(ContainSubstring(expectedErr.Error()))
		})
	})

	Describe("StreamVirtualServiceList", func() {
		It("works", func() {
			virtualServiceList := []*gatewayv1.VirtualService{
				{
					Metadata: metadata,
				},
				{
					Metadata: metadata,
				},
			}

			refreshRate := time.Minute
			request := v1.StreamVirtualServiceListRequest{
				Namespace: ref.GetNamespace(),
			}
			virtualServiceChan := make(chan gatewayv1.VirtualServiceList, 1)
			virtualServiceChan <- virtualServiceList
			errChan := make(chan error)

			defer func() {
				close(virtualServiceChan)
				close(errChan)
			}()

			settingsValues.EXPECT().GetRefreshRate().Return(refreshRate)
			virtualServiceClient.EXPECT().
				Watch(ref.GetNamespace(), gomock.Any()).
				Return(virtualServiceChan, errChan, nil)

			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()
			streamClient, err := client.StreamVirtualServiceList(ctx, &request)
			Expect(err).NotTo(HaveOccurred())

			wait := make(chan struct{})
			go func() {
				defer GinkgoRecover()
				defer func() {
					close(wait)
				}()

				actual, err := streamClient.Recv()
				Expect(err).NotTo(HaveOccurred())
				expected := &v1.StreamVirtualServiceListResponse{VirtualServices: virtualServiceList}
				ExpectEqualProtoMessages(actual, expected)

				errChan <- testErr
				_, err = streamClient.Recv()
				Expect(err).To(HaveOccurred())
				expectedErr := virtualservicesvc.ErrorWhileWatchingVirtualServices(testErr, metadata.Namespace)
				Expect(err.Error()).To(ContainSubstring(expectedErr.Error()))
			}()

			select {
			case <-wait:
			case <-time.After(time.Second):
				Fail("expected wait to be closed before 1s")
			}
		})
	})

	Describe("CreateVirtualService", func() {
		getInput := func(ref *core.ResourceRef) *v1.VirtualServiceInput {
			return &v1.VirtualServiceInput{
				Ref: ref,
			}
		}

		It("works when the mutator works", func() {
			virtualService := gatewayv1.VirtualService{
				Metadata: metadata,
			}

			mutationFactory.EXPECT().ConfigureVirtualService(getInput(&ref))
			mutator.EXPECT().
				Create(getInput(&ref), gomock.Any()).
				Return(&virtualService, nil)

			actual, err := client.CreateVirtualService(context.TODO(), &v1.CreateVirtualServiceRequest{Input: getInput(&ref)})
			Expect(err).NotTo(HaveOccurred())
			expected := &v1.CreateVirtualServiceResponse{VirtualService: &virtualService}
			ExpectEqualProtoMessages(actual, expected)
		})

		It("errors when the mutator errors", func() {
			mutationFactory.EXPECT().ConfigureVirtualService(getInput(&ref))
			mutator.EXPECT().
				Create(getInput(&ref), gomock.Any()).
				Return(nil, testErr)

			request := &v1.CreateVirtualServiceRequest{
				Input: getInput(&ref),
			}
			_, err := client.CreateVirtualService(context.TODO(), request)
			Expect(err).To(HaveOccurred())
			expectedErr := virtualservicesvc.FailedToCreateVirtualServiceError(testErr, &ref)
			Expect(err.Error()).To(ContainSubstring(expectedErr.Error()))
		})
	})

	Describe("ConfigureVirtualService", func() {
		getInput := func(ref *core.ResourceRef) *v1.VirtualServiceInput {
			return &v1.VirtualServiceInput{
				Ref: ref,
			}
		}

		It("works when the mutator works", func() {
			virtualService := gatewayv1.VirtualService{
				Metadata: metadata,
			}

			mutationFactory.EXPECT().ConfigureVirtualService(getInput(&ref))
			mutator.EXPECT().
				Update(&ref, gomock.Any()).
				Return(&virtualService, nil)

			actual, err := client.UpdateVirtualService(context.TODO(), &v1.UpdateVirtualServiceRequest{Input: getInput(&ref)})
			Expect(err).NotTo(HaveOccurred())
			expected := &v1.UpdateVirtualServiceResponse{VirtualService: &virtualService}
			ExpectEqualProtoMessages(actual, expected)
		})

		It("errors when the mutator errors", func() {
			mutationFactory.EXPECT().ConfigureVirtualService(getInput(&ref))
			mutator.EXPECT().
				Update(&ref, gomock.Any()).
				Return(nil, testErr)

			_, err := client.UpdateVirtualService(context.TODO(), &v1.UpdateVirtualServiceRequest{Input: getInput(&ref)})
			Expect(err).To(HaveOccurred())
			expectedErr := virtualservicesvc.FailedToUpdateVirtualServiceError(testErr, &ref)
			Expect(err.Error()).To(ContainSubstring(expectedErr.Error()))
		})
	})

	Describe("DeleteVirtualService", func() {
		It("works when the virtual service client works", func() {
			virtualServiceClient.EXPECT().
				Delete(ref.Namespace, ref.Name, clients.DeleteOpts{Ctx: context.TODO()}).
				Return(nil)

			request := &v1.DeleteVirtualServiceRequest{Ref: &ref}
			actual, err := client.DeleteVirtualService(context.TODO(), request)
			Expect(err).NotTo(HaveOccurred())
			expected := &v1.DeleteVirtualServiceResponse{}
			ExpectEqualProtoMessages(actual, expected)
		})

		It("errors when the virtual service client errors", func() {
			virtualServiceClient.EXPECT().
				Delete(ref.Namespace, ref.Name, clients.DeleteOpts{Ctx: context.TODO()}).
				Return(testErr)

			request := &v1.DeleteVirtualServiceRequest{Ref: &ref}
			_, err := client.DeleteVirtualService(context.TODO(), request)
			Expect(err).To(HaveOccurred())
			expectedErr := virtualservicesvc.FailedToDeleteVirtualServiceError(testErr, &ref)
			Expect(err.Error()).To(ContainSubstring(expectedErr.Error()))
		})
	})

	Describe("CreateRoute", func() {
		getInput := func(ref *core.ResourceRef) *v1.RouteInput {
			return &v1.RouteInput{
				VirtualServiceRef: ref,
			}
		}

		It("works when the mutator works", func() {
			virtualService := gatewayv1.VirtualService{
				Metadata: metadata,
			}

			mutationFactory.EXPECT().CreateRoute(getInput(&ref))
			mutator.EXPECT().
				Update(&ref, gomock.Any()).
				Return(&virtualService, nil)

			actual, err := client.CreateRoute(context.TODO(), &v1.CreateRouteRequest{Input: getInput(&ref)})
			Expect(err).NotTo(HaveOccurred())
			expected := &v1.CreateRouteResponse{VirtualService: &virtualService}
			ExpectEqualProtoMessages(actual, expected)
		})

		It("errors when the mutator errors", func() {
			mutationFactory.EXPECT().CreateRoute(getInput(&ref))
			mutator.EXPECT().
				Update(&ref, gomock.Any()).
				Return(nil, testErr)

			_, err := client.CreateRoute(context.TODO(), &v1.CreateRouteRequest{Input: getInput(&ref)})
			Expect(err).To(HaveOccurred())
			expectedErr := virtualservicesvc.FailedToCreateRouteError(testErr, &ref)
			Expect(err.Error()).To(ContainSubstring(expectedErr.Error()))
		})
	})

	Describe("UpdateRoute", func() {
		getInput := func(ref *core.ResourceRef) *v1.RouteInput {
			return &v1.RouteInput{
				VirtualServiceRef: ref,
			}
		}

		It("works when the mutator works", func() {
			virtualService := gatewayv1.VirtualService{
				Metadata: metadata,
			}

			mutationFactory.EXPECT().UpdateRoute(getInput(&ref))
			mutator.EXPECT().
				Update(&ref, gomock.Any()).
				Return(&virtualService, nil)

			actual, err := client.UpdateRoute(context.TODO(), &v1.UpdateRouteRequest{Input: getInput(&ref)})
			Expect(err).NotTo(HaveOccurred())
			expected := &v1.UpdateRouteResponse{VirtualService: &virtualService}
			ExpectEqualProtoMessages(actual, expected)
		})

		It("errors when the mutator errors", func() {
			mutationFactory.EXPECT().UpdateRoute(getInput(&ref))
			mutator.EXPECT().
				Update(&ref, gomock.Any()).
				Return(nil, testErr)

			_, err := client.UpdateRoute(context.TODO(), &v1.UpdateRouteRequest{Input: getInput(&ref)})
			Expect(err).To(HaveOccurred())
			expectedErr := virtualservicesvc.FailedToUpdateRouteError(testErr, &ref)
			Expect(err.Error()).To(ContainSubstring(expectedErr.Error()))
		})
	})

	Describe("DeleteRoute", func() {
		getRequest := func(ref *core.ResourceRef) *v1.DeleteRouteRequest {
			return &v1.DeleteRouteRequest{
				VirtualServiceRef: ref,
				Index:             uint32Zero,
			}
		}

		It("works when the mutator works", func() {
			virtualService := gatewayv1.VirtualService{
				Metadata: metadata,
			}

			mutationFactory.EXPECT().DeleteRoute(uint32Zero)
			mutator.EXPECT().
				Update(&ref, gomock.Any()).
				Return(&virtualService, nil)

			actual, err := client.DeleteRoute(context.TODO(), getRequest(&ref))
			Expect(err).NotTo(HaveOccurred())
			expected := &v1.DeleteRouteResponse{VirtualService: &virtualService}
			ExpectEqualProtoMessages(actual, expected)
		})

		It("errors when the mutator errors", func() {
			mutationFactory.EXPECT().DeleteRoute(uint32Zero)
			mutator.EXPECT().
				Update(&ref, gomock.Any()).
				Return(nil, testErr)

			_, err := client.DeleteRoute(context.TODO(), getRequest(&ref))
			Expect(err).To(HaveOccurred())
			expectedErr := virtualservicesvc.FailedToDeleteRouteError(testErr, &ref)
			Expect(err.Error()).To(ContainSubstring(expectedErr.Error()))
		})
	})

	Describe("SwapRoutes", func() {
		getRequest := func(ref *core.ResourceRef) *v1.SwapRoutesRequest {
			return &v1.SwapRoutesRequest{
				VirtualServiceRef: ref,
				Index1:            uint32Zero,
				Index2:            uint32One,
			}
		}

		It("works when the mutator works", func() {
			virtualService := gatewayv1.VirtualService{
				Metadata: metadata,
			}

			mutationFactory.EXPECT().SwapRoutes(uint32Zero, uint32One)
			mutator.EXPECT().
				Update(&ref, gomock.Any()).
				Return(&virtualService, nil)

			actual, err := client.SwapRoutes(context.TODO(), getRequest(&ref))
			Expect(err).NotTo(HaveOccurred())
			expected := &v1.SwapRoutesResponse{VirtualService: &virtualService}
			ExpectEqualProtoMessages(actual, expected)
		})

		It("errors when the mutator errors", func() {
			mutationFactory.EXPECT().SwapRoutes(uint32Zero, uint32One)
			mutator.EXPECT().
				Update(&ref, gomock.Any()).
				Return(nil, testErr)

			_, err := client.SwapRoutes(context.TODO(), getRequest(&ref))
			Expect(err).To(HaveOccurred())
			expectedErr := virtualservicesvc.FailedToSwapRoutesError(testErr, &ref)
			Expect(err.Error()).To(ContainSubstring(expectedErr.Error()))
		})
	})

	Describe("ShiftRoutes", func() {
		getRequest := func(ref *core.ResourceRef) *v1.ShiftRoutesRequest {
			return &v1.ShiftRoutesRequest{
				VirtualServiceRef: ref,
				FromIndex:         uint32Zero,
				ToIndex:           uint32One,
			}
		}

		It("works when the mutator works", func() {
			virtualService := gatewayv1.VirtualService{
				Metadata: metadata,
			}

			mutationFactory.EXPECT().ShiftRoutes(uint32Zero, uint32One)
			mutator.EXPECT().
				Update(&ref, gomock.Any()).
				Return(&virtualService, nil)

			actual, err := client.ShiftRoutes(context.TODO(), getRequest(&ref))
			Expect(err).NotTo(HaveOccurred())
			expected := &v1.ShiftRoutesResponse{VirtualService: &virtualService}
			ExpectEqualProtoMessages(actual, expected)
		})

		It("errors when the mutator errors", func() {
			mutationFactory.EXPECT().ShiftRoutes(uint32Zero, uint32One)
			mutator.EXPECT().
				Update(&ref, gomock.Any()).
				Return(nil, testErr)

			_, err := client.ShiftRoutes(context.TODO(), getRequest(&ref))
			Expect(err).To(HaveOccurred())
			expectedErr := virtualservicesvc.FailedToShiftRoutesError(testErr, &ref)
			Expect(err.Error()).To(ContainSubstring(expectedErr.Error()))
		})
	})
})
