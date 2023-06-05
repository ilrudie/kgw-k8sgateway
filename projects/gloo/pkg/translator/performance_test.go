package translator_test

import (
	"context"
	"fmt"
	"github.com/onsi/gomega/types"
	"github.com/solo-io/gloo/test/gomega/matchers"

	"github.com/golang/mock/gomock"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/validation"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/registry"
	mock_consul "github.com/solo-io/gloo/projects/gloo/pkg/upstreams/consul/mocks"
	glooutils "github.com/solo-io/gloo/projects/gloo/pkg/utils"
	gloohelpers "github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
	. "github.com/solo-io/gloo/projects/gloo/pkg/translator"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	validationutils "github.com/solo-io/gloo/projects/gloo/pkg/utils/validation"
	"github.com/solo-io/gloo/test/ginkgo/labels"
	"time"
)

type benchmarkConfig struct {
	tries             int
	maxDur            time.Duration
	benchmarkMatchers []types.GomegaMatcher
}

var _ = FDescribe("Translation - Benchmarking Tests", Serial, Label(labels.Performance), func() {
	var (
		ctrl       *gomock.Controller
		settings   *v1.Settings
		translator Translator

		registeredPlugins []plugins.Plugin
	)

	BeforeEach(func() {

		ctrl = gomock.NewController(T)

		settings = &v1.Settings{}
		memoryClientFactory := &factory.MemoryResourceClientFactory{
			Cache: memory.NewInMemoryResourceCache(),
		}
		opts := bootstrap.Opts{
			Settings:  settings,
			Secrets:   memoryClientFactory,
			Upstreams: memoryClientFactory,
			Consul: bootstrap.Consul{
				ConsulWatcher: mock_consul.NewMockConsulWatcher(ctrl), // just needed to activate the consul plugin
			},
		}
		registeredPlugins = registry.Plugins(opts)

		pluginRegistry := registry.NewPluginRegistry(registeredPlugins)

		translator = NewTranslatorWithHasher(glooutils.NewSslConfigTranslator(), settings, pluginRegistry, EnvoyCacheResourcesListToFnvHash)
	})

	DescribeTable("Translate",
		func(desc string, apiSnap *v1snap.ApiSnapshot, config benchmarkConfig) {

			params := plugins.Params{
				Ctx:      context.Background(),
				Snapshot: apiSnap,
			}

			var (
				snap   cache.Snapshot
				errs   reporter.ResourceReports
				report *validation.ProxyReport
			)

			experiment := gmeasure.NewExperiment("translate")

			AddReportEntry(experiment.Name, experiment)

			statName := fmt.Sprintf("translating %s", desc)
			experiment.Sample(func(idx int) {

				// Time translation
				experiment.MeasureDuration(statName, func() {
					snap, errs, report = translator.Translate(params, gloohelpers.Proxy())
				})

				// Assert expected results
				Expect(errs.Validate()).NotTo(HaveOccurred())
				Expect(snap).NotTo(BeNil())
				Expect(report).To(Equal(validationutils.MakeReport(gloohelpers.Proxy())))
			}, gmeasure.SamplingConfig{N: config.tries, Duration: config.maxDur})

			durations := experiment.Get(statName).Durations

			Expect(durations).Should(And(config.benchmarkMatchers...))
		},
		Entry("basic", "basic", basicSnap, basicConfig),
		Entry("10 upstreams", "10 upstreams", gloohelpers.ScaledSnapshot(gloohelpers.ScaleConfig{
			Upstreams: 10,
			Endpoints: 1,
		}), basicConfig),
		Entry("100 upstreams", "100 upstreams", gloohelpers.ScaledSnapshot(gloohelpers.ScaleConfig{
			Upstreams: 100,
			Endpoints: 1,
		}), basicConfig),
		Entry("1000 upstreams", "1000 upstreams", gloohelpers.ScaledSnapshot(gloohelpers.ScaleConfig{
			Upstreams: 1000,
			Endpoints: 1,
		}), basicConfig),
		Entry("10 endpoints", "10 endpoints", gloohelpers.ScaledSnapshot(gloohelpers.ScaleConfig{
			Upstreams: 1,
			Endpoints: 10,
		}), basicConfig),
		Entry("100 endpoints", "100 endpoints", gloohelpers.ScaledSnapshot(gloohelpers.ScaleConfig{
			Upstreams: 1,
			Endpoints: 100,
		}), basicConfig),
		Entry("1000 endpoints", "1000 endpoints", gloohelpers.ScaledSnapshot(gloohelpers.ScaleConfig{
			Upstreams: 1,
			Endpoints: 1000,
		}), basicConfig),
		Entry("10 of everything", "10 of everything", gloohelpers.ScaledSnapshot(gloohelpers.ScaleConfig{
			Upstreams: 10,
			Endpoints: 10,
		}), basicConfig),
	)
})

var basicSnap = &v1snap.ApiSnapshot{
	Endpoints: []*v1.Endpoint{gloohelpers.Endpoint},
	Upstreams: []*v1.Upstream{gloohelpers.Upstream},
}

var basicConfig = benchmarkConfig{
	tries:  1000,
	maxDur: time.Second,
	benchmarkMatchers: []types.GomegaMatcher{
		matchers.Median(5 * time.Millisecond),
		matchers.Percentile(90, 10*time.Millisecond),
	},
}
