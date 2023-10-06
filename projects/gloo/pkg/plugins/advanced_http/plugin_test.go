package advanced_http_test

import (
	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/golang/protobuf/ptypes"
	wrapperspb "github.com/golang/protobuf/ptypes/wrappers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/pkg/utils/api_conversion"
	core2 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core"
	envoy_advanced_http "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/advanced_http"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	gloo_advanced_http "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/advanced_http"
	v1static "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	core3 "github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2/core"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	. "github.com/solo-io/solo-projects/projects/gloo/pkg/plugins/advanced_http"
)

var _ = Describe("Plugin", func() {

	var (
		p               plugins.UpstreamPlugin
		params          plugins.Params
		upstream        *v1.Upstream
		upstreamSpec    *v1static.UpstreamSpec
		out             *envoy_config_cluster_v3.Cluster
		baseHealthCheck *envoy_config_core_v3.HealthCheck_HttpHealthCheck
	)

	BeforeEach(func() {
		p = NewPlugin()
		out = new(envoy_config_cluster_v3.Cluster)

		check := &core2.HealthCheck{
			HealthChecker: &core2.HealthCheck_HttpHealthCheck_{
				HttpHealthCheck: &core2.HealthCheck_HttpHealthCheck{
					Host: "foo",
					Path: "/health",
					RequestHeadersToAdd: []*core3.HeaderValueOption{
						{
							HeaderOption: &core3.HeaderValueOption_Header{
								Header: &core3.HeaderValue{
									Key:   "key",
									Value: "value",
								},
							},
							Append: &wrapperspb.BoolValue{Value: false},
						},
					},
					RequestHeadersToRemove: []string{"test"},
				},
			},
		}

		healthCheck, _ := api_conversion.ToEnvoyHealthCheck(check, nil, api_conversion.HeaderSecretOptions{EnforceNamespaceMatch: false})
		baseHealthCheck = healthCheck.GetHttpHealthCheck()

		out.HealthChecks = []*envoy_config_core_v3.HealthCheck{
			{
				HealthChecker: &envoy_config_core_v3.HealthCheck_HttpHealthCheck_{
					HttpHealthCheck: baseHealthCheck,
				},
			},
		}

		p.Init(plugins.InitParams{})
		upstreamSpec = &v1static.UpstreamSpec{
			Hosts: []*v1static.Host{{
				Addr: "localhost",
				Port: 1234,
				HealthCheckConfig: &v1static.Host_HealthCheckConfig{
					Path: "/foo",
				},
			}},
		}
		upstream = &v1.Upstream{
			Metadata: &core.Metadata{
				Name:      "extauth-server",
				Namespace: "default",
			},
			UpstreamType: &v1.Upstream_Static{
				Static: upstreamSpec,
			},
			HealthChecks: []*core2.HealthCheck{check},
		}
		params = plugins.Params{Snapshot: &v1snap.ApiSnapshot{}}
	})

	It("should create a custom health check when static upstream has a path", func() {
		err := p.ProcessUpstream(params, upstream, out)
		Expect(err).ToNot(HaveOccurred())
		Expect(out.GetHealthChecks()[0].GetCustomHealthCheck().GetName()).To(Equal(HealthCheckerName))
		typedcfg := out.GetHealthChecks()[0].GetCustomHealthCheck().GetTypedConfig()
		var out envoy_advanced_http.AdvancedHttp
		Expect(ptypes.UnmarshalAny(typedcfg, &out)).NotTo(HaveOccurred())

		Expect(out.HttpHealthCheck.Path).To(Equal(baseHealthCheck.Path))
		Expect(out.HttpHealthCheck.Host).To(Equal(baseHealthCheck.Host))
		Expect(out.HttpHealthCheck.RequestHeadersToRemove).To(Equal(baseHealthCheck.RequestHeadersToRemove))
		Expect(out.HttpHealthCheck.RequestHeadersToAdd[0].Header.Key).To(Equal(baseHealthCheck.RequestHeadersToAdd[0].Header.Key))
		Expect(out.HttpHealthCheck.RequestHeadersToAdd[0].Header.Value).To(Equal(baseHealthCheck.RequestHeadersToAdd[0].Header.Value))

		expectedAppend := true
		// We only handle `append_if_exists_or_add` and `overwrite_if_exists_or_add` during translation of `append` to `append_action`, so we need to check for those.
		if baseHealthCheck.RequestHeadersToAdd[0].AppendAction == envoy_config_core_v3.HeaderValueOption_APPEND_IF_EXISTS_OR_ADD {
			expectedAppend = true
		} else if baseHealthCheck.RequestHeadersToAdd[0].AppendAction == envoy_config_core_v3.HeaderValueOption_OVERWRITE_IF_EXISTS_OR_ADD {
			expectedAppend = false
		} else {
			Fail("unexpected append action translation")
		}
		Expect(out.HttpHealthCheck.RequestHeadersToAdd[0].Append.Value).To(Equal(expectedAppend))

	})

	It("should default no_match and match health check properly", func() {
		upstream.HealthChecks = []*core2.HealthCheck{
			{
				HealthChecker: &core2.HealthCheck_HttpHealthCheck_{
					HttpHealthCheck: &core2.HealthCheck_HttpHealthCheck{
						Path: "/health",
						ResponseAssertions: &gloo_advanced_http.ResponseAssertions{
							//NoMatchHealth: leave unset, as we are testing our enum defaulting
							ResponseMatchers: []*gloo_advanced_http.ResponseMatcher{
								{
									ResponseMatch: &gloo_advanced_http.ResponseMatch{
										IgnoreErrorOnParse: false,
										Source:             &gloo_advanced_http.ResponseMatch_Body{},
										Regex:              ".*",
									},
									//MatchHealth: leave unset, as we are testing our enum defaulting
								},
							},
						},
					},
				},
			},
		}
		err := p.ProcessUpstream(params, upstream, out)
		Expect(err).ToNot(HaveOccurred())
		Expect(out.GetHealthChecks()[0].GetCustomHealthCheck().GetName()).To(Equal(HealthCheckerName))
		typedcfg := out.GetHealthChecks()[0].GetCustomHealthCheck().GetTypedConfig()
		var out envoy_advanced_http.AdvancedHttp
		Expect(ptypes.UnmarshalAny(typedcfg, &out)).NotTo(HaveOccurred())

		Expect(out.ResponseAssertions.NoMatchHealth).To(Equal(envoy_advanced_http.HealthCheckResult_unhealthy))
		Expect(out.ResponseAssertions.ResponseMatchers[0].MatchHealth).To(Equal(envoy_advanced_http.HealthCheckResult_healthy))
	})

	It("should error on missing path in health check", func() {
		upstream.HealthChecks = []*core2.HealthCheck{
			{
				HealthChecker: &core2.HealthCheck_HttpHealthCheck_{
					HttpHealthCheck: &core2.HealthCheck_HttpHealthCheck{
						Path: "", // error to be empty
					},
				},
			},
		}
		err := p.ProcessUpstream(params, upstream, out)
		Expect(err).To(HaveOccurred())
	})

	It("should error on missing path in health check", func() {
		upstream.HealthChecks = []*core2.HealthCheck{
			{
				HealthChecker: &core2.HealthCheck_HttpHealthCheck_{
					HttpHealthCheck: &core2.HealthCheck_HttpHealthCheck{
						// path is intentionally nil
						Host: "notempty",
					},
				},
			},
		}
		err := p.ProcessUpstream(params, upstream, out)
		Expect(err).To(HaveOccurred())
	})

	It("should error on nil in health check", func() {
		upstream.HealthChecks = []*core2.HealthCheck{
			{
				HealthChecker: &core2.HealthCheck_HttpHealthCheck_{
					HttpHealthCheck: nil,
				},
			},
		}
		err := p.ProcessUpstream(params, upstream, out)
		Expect(err).To(HaveOccurred())
	})

})