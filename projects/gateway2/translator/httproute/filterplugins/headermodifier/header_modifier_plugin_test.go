package headermodifier_test

import (
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	routev3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	. "github.com/onsi/ginkgo/v2"
	"github.com/solo-io/gloo/projects/gateway2/translator/httproute/filterplugins/filtertests"
	"github.com/solo-io/gloo/projects/gateway2/translator/httproute/filterplugins/headermodifier"
	"google.golang.org/protobuf/types/known/wrapperspb"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

var _ = Describe("HeaderModifierPlugin", func() {
	It("applies header modifier filter", func() {

		plugin := headermodifier.NewPlugin()
		filter := gwv1.HTTPRouteFilter{
			Type: gwv1.HTTPRouteFilterRequestHeaderModifier,
			RequestHeaderModifier: &gwv1.HTTPHeaderFilter{
				Add: []gwv1.HTTPHeader{
					{
						Name:  "foo",
						Value: "bar",
					},
				},
				Set: []gwv1.HTTPHeader{
					{
						Name:  "foo",
						Value: "bar",
					},
				},
				Remove: []string{"foo"},
			},
		}
		expectedRoute := &routev3.Route{
			RequestHeadersToAdd: []*corev3.HeaderValueOption{
				{Header: &corev3.HeaderValue{Key: "foo", Value: "bar"}, Append: wrapperspb.Bool(true)},
				{Header: &corev3.HeaderValue{Key: "foo", Value: "bar"}, Append: wrapperspb.Bool(false)},
			},
			RequestHeadersToRemove: []string{"foo"},
		}
		filtertests.AssertExpectedRoute(
			plugin,
			filter,
			expectedRoute,
			false,
		)
	})
})