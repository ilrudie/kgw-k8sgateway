package proxy_syncer

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/protobuf/runtime/protoiface"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"

	sologatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	glookubev1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/xds"

	// solokubeclient "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	// solokubecrd "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd/solo.io/v1"
	// "github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/gloo/projects/gloo/pkg/snapshot"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kubesecret"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
	networkingclient "istio.io/client-go/pkg/apis/networking/v1"
	"istio.io/istio/pkg/config/schema/gvr"
	istiogvr "istio.io/istio/pkg/config/schema/gvr"
	"istio.io/istio/pkg/kube"
	"istio.io/istio/pkg/kube/kclient"
	"istio.io/istio/pkg/kube/krt"
	"istio.io/istio/pkg/kube/kubetypes"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/solo-io/gloo/pkg/utils/statsutils"
	"github.com/solo-io/gloo/projects/gateway2/extensions"
	"github.com/solo-io/gloo/projects/gateway2/krtcollections"
	"github.com/solo-io/gloo/projects/gateway2/reports"
	gwplugins "github.com/solo-io/gloo/projects/gateway2/translator/plugins"
	"github.com/solo-io/gloo/projects/gateway2/translator/plugins/registry"
	"github.com/solo-io/gloo/projects/gateway2/translator/translatorutils"
	kubeconverters "github.com/solo-io/gloo/projects/gloo/pkg/api/converters/kube"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	kubeupstreams "github.com/solo-io/gloo/projects/gloo/pkg/upstreams/kubernetes"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	// "github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	// "github.com/solo-io/solo-kit/pkg/utils/protoutils"
	"github.com/solo-io/solo-kit/pkg/utils/statusutils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// ProxySyncer is responsible for translating Kubernetes Gateway CRs into Gloo Proxies
// and syncing the proxyClient with the newly translated proxies.
type ProxySyncer struct {
	controllerName string
	writeNamespace string

	inputs          *GatewayInputChannels
	mgr             manager.Manager
	k8sGwExtensions extensions.K8sGatewayExtensions
	newTranslator   func() translator.Translator

	// proxyReconciler wraps the client that writes Proxy resources into an in-memory cache
	// This cache is utilized by RateLimit and the debug.ProxyEndpointServer
	proxyReconciler gloov1.ProxyReconciler

	proxyTranslator ProxyTranslator
	istioClient     kube.Client

	// used for converting from kube type to gloo type
	// TODO: abstract away the need for this by refactoring convert func()
	// legacyClients map[reflect.Type]*solokubeclient.ResourceClient

	// secret client needed to use existing kube secret -> gloo secret converters
	// the only actually use is to do client.NewResource() to get a gloov1.Secret
	// we can/should probably break this dependency entirely relatively easily
	legacySecretClient gloov1.SecretClient

	uniqlyConnectedClient krt.Collection[krtcollections.UniqlyConnectedClient]
	pods                  krt.Collection[krtcollections.LocalityPod]
}

type GatewayInputChannels struct {
	genericEvent AsyncQueue[struct{}]
	secretEvent  AsyncQueue[SecretInputs]
}

func (x *GatewayInputChannels) Kick(ctx context.Context) {
	x.genericEvent.Enqueue(struct{}{})
}

func (x *GatewayInputChannels) UpdateSecretInputs(ctx context.Context, inputs SecretInputs) {
	x.secretEvent.Enqueue(inputs)
}

func NewGatewayInputChannels() *GatewayInputChannels {
	return &GatewayInputChannels{
		genericEvent: NewAsyncQueue[struct{}](),
		secretEvent:  NewAsyncQueue[SecretInputs](),
	}
}

// labels used to uniquely identify Proxies that are managed by the kube gateway controller
var kubeGatewayProxyLabels = map[string]string{
	// the proxy type key/value must stay in sync with the one defined in projects/gateway2/translator/gateway_translator.go
	utils.ProxyTypeKey: utils.GatewayApiProxyValue,
}

// NewProxySyncer returns an implementation of the ProxySyncer
// The provided GatewayInputChannels are used to trigger syncs.
// The proxy sync is triggered by the `genericEvent` which is kicked when
// we reconcile gateway in the gateway controller. The `secretEvent` is kicked when a secret is created, updated,
func NewProxySyncer(
	controllerName string,
	writeNamespace string,
	inputs *GatewayInputChannels,
	mgr manager.Manager,
	client kube.Client,
	uniqlyConnectedClient krt.Collection[krtcollections.UniqlyConnectedClient],
	pods krt.Collection[krtcollections.LocalityPod],
	k8sGwExtensions extensions.K8sGatewayExtensions,
	proxyClient gloov1.ProxyClient,
	newTranslator func() translator.Translator,
	xdsCache envoycache.SnapshotCache,
	settings *gloov1.Settings,
	syncerExtensions []syncer.TranslatorSyncerExtension,
	legacySecretClient gloov1.SecretClient,
	glooReporter reporter.StatusReporter,
) *ProxySyncer {
	// legacyClients := map[reflect.Type]*solokubeclient.ResourceClient{}
	// usc := upstreamClient.BaseClient()
	// if kusc, ok := usc.(*solokubeclient.ResourceClient); !ok {
	// 	panic("upstream base client isn't a kube client!!!")
	// } else {
	// 	legacyClients[reflect.TypeOf(gloov1.Upstream{})] = kusc
	// }
	return &ProxySyncer{
		controllerName:  controllerName,
		writeNamespace:  writeNamespace,
		inputs:          inputs,
		k8sGwExtensions: k8sGwExtensions,
		newTranslator:   newTranslator,
		proxyReconciler: gloov1.NewProxyReconciler(proxyClient, statusutils.NewNoOpStatusClient()),
		proxyTranslator: NewProxyTranslator(newTranslator(), xdsCache, settings, syncerExtensions, glooReporter),
		// legacyClients:   legacyClients,
		legacySecretClient:    legacySecretClient,
		istioClient:           client,
		uniqlyConnectedClient: uniqlyConnectedClient,
	}
}

type ProxyTranslator struct {
	translator       translator.Translator
	settings         *gloov1.Settings
	syncerExtensions []syncer.TranslatorSyncerExtension
	xdsCache         envoycache.SnapshotCache
	// used to no-op during extension syncing as we only do it to get reports
	noopSnapSetter syncer.SnapshotSetter
	// we need to report on upstreams/proxies that we are responsible for translating and syncing
	// so we use this repporter to do so; do we also need to report authconfigs and RLCs...?
	// TODO: possibly consolidate this with the status reporter used in the plugins
	// also TODO: copy the leader election stuff (and maybe leaderStartupAction whatever that is)
	glooReporter reporter.StatusReporter
}

func NewProxyTranslator(translator translator.Translator,
	xdsCache envoycache.SnapshotCache,
	settings *gloov1.Settings,
	syncerExtensions []syncer.TranslatorSyncerExtension,
	glooReporter reporter.StatusReporter,
) ProxyTranslator {
	return ProxyTranslator{
		translator:       translator,
		xdsCache:         xdsCache,
		settings:         settings,
		syncerExtensions: syncerExtensions,
		noopSnapSetter:   &syncer.NoOpSnapshotSetter{},
		glooReporter:     glooReporter,
	}
}

type xdsSnapWrapper struct {
	snap            *xds.EnvoySnapshot
	proxyKey        string
	proxyWithReport translatorutils.ProxyWithReports
	pluginRegistry  registry.PluginRegistry
	fullReports     reporter.ResourceReports
}

var _ krt.ResourceNamer = &xdsSnapWrapper{}

func (p *xdsSnapWrapper) Equals(in *xdsSnapWrapper) bool {
	return p.snap.Equal(in.snap)
}
func (p *xdsSnapWrapper) ResourceName() string {
	return p.proxyKey
}

type glooProxy struct {
	proxy *gloov1.Proxy
	// plugins used to generate this proxy
	pluginRegistry registry.PluginRegistry
}

var _ krt.ResourceNamer = &glooProxy{}

func (p *glooProxy) Equals(in *glooProxy) bool {
	return proto.Equal(p.proxy, in.proxy)
}
func (p *glooProxy) ResourceName() string {
	return xds.SnapshotCacheKey(p.proxy)
}

var _ krt.ResourceNamer = &glooEndpoint{}

// stolen from projects/gloo/pkg/upstreams/serviceentry/krtwrappers.go
// TODO: consolidate this stuff
func UnwrapEps(geps []*glooEndpoint) gloov1.EndpointList {
	out := make(gloov1.EndpointList, 0, len(geps))
	for _, ep := range geps {
		out = append(out, ep.Endpoint)
	}
	return out
}

// glooEndpoint provides a krt keying function for Gloo's `v1.Endpoint`
type glooEndpoint struct {
	*gloov1.Endpoint
}

func (ep *glooEndpoint) Equals(in *glooEndpoint) bool {
	return proto.Equal(ep, in)
}

func (ep *glooEndpoint) ResourceName() string {
	return ep.Metadata.GetName() + "/" + ep.Metadata.GetNamespace()
}

type GlooResource interface {
	proto.Message
	snapshot.Md
}

type ResourceWrapper[T GlooResource] struct {
	Inner T
}

func (us ResourceWrapper[T]) ResourceName() string {
	return krt.Named{
		Name:      us.Inner.GetMetadata().GetName(),
		Namespace: us.Inner.GetMetadata().GetNamespace(),
	}.ResourceName()
}
func (us ResourceWrapper[T]) Equals(in UpstreamWrapper) bool {
	return proto.Equal(us.Inner, in.Inner)
}

type KrtWrappedCollection[T GlooResource] struct {
	C    krt.Collection[ResourceWrapper[T]]
	Kctx krt.HandlerContext
}

func (t KrtWrappedCollection[T]) Find(namespace string, name string) (T, error) {
	ret := krt.Fetch(t.Kctx, t.C, krt.FilterObjectName(types.NamespacedName{Name: name, Namespace: namespace}))
	if len(ret) != 1 {
		var zero T
		return zero, fmt.Errorf("list did not find %T %v.%v", zero, namespace, name)
	}
	return ret[0].Inner, nil
}

func (t KrtWrappedCollection[T]) List() []T {
	panic("TODO: this should never be called")
	out := krt.Fetch(t.Kctx, t.C)
	out2 := make([]T, len(out))
	for i, x := range out {
		out2[i] = x.Inner
	}
	return out2
}

func (t KrtWrappedCollection[X]) IsFindEfficient() bool {
	return true
}

type UpstreamWrapper = ResourceWrapper[*gloov1.Upstream]

var _ krt.ResourceNamer = UpstreamWrapper{}

// upstream provides a keying function for Gloo's `v1.Upstream`

type fromKrtSnap struct {
	rtOpts    []*sologatewayv1.RouteOption
	vhostOpts []*sologatewayv1.VirtualHostOption

	upstreams krt.Collection[UpstreamWrapper]
	artifacts krt.Collection[ResourceWrapper[*gloov1.Artifact]]
	secrets   krt.Collection[ResourceWrapper[*gloov1.Secret]]
}

func (s *ProxySyncer) Start(ctx context.Context) error {
	ctx = contextutils.WithLogger(ctx, "k8s-gw-syncer")

	// create krt collections needed for building ApiSnapshot
	// RouteOptions := setupCollectionDynamic[sologatewayv1.RouteOption](
	// 	ctx,
	// 	s.istioClient,
	// 	sologatewayv1.SchemeGroupVersion.WithResource("routeoptions"),
	// 	krt.WithName("RouteOptions"),
	// )
	// VirtualHostOptions := setupCollectionDynamic[sologatewayv1.VirtualHostOption](
	// 	ctx,
	// 	s.istioClient,
	// 	sologatewayv1.SchemeGroupVersion.WithResource("virtualhostoptions"),
	// 	krt.WithName("VirtualHostOptions"),
	// )

	// TODO: handle cfgmap noisiness:
	// https://github.com/solo-io/gloo/blob/main/projects/gloo/pkg/api/converters/kube/artifact_converter.go#L31
	configMapClient := kclient.New[*corev1.ConfigMap](s.istioClient)
	ConfigMaps := krt.WrapClient(configMapClient, krt.WithName("ConfigMaps"))
	artifacts := krt.NewCollection(ConfigMaps, func(kctx krt.HandlerContext, cm *corev1.ConfigMap) *ResourceWrapper[*gloov1.Artifact] {
		a := kubeconverters.KubeConfigMapToArtifact(cm)
		if a == nil {
			return nil
		}
		out := ResourceWrapper[*gloov1.Artifact]{a}
		return &out
	})
	secretClient := kclient.New[*corev1.Secret](s.istioClient)
	secrets := krt.WrapClient(secretClient, krt.WithName("Secrets"))
	krtSecrets := krt.NewCollection(secrets, func(kctx krt.HandlerContext, i *corev1.Secret) *ResourceWrapper[*gloov1.Secret] {
		secretResourceClient, ok := s.legacySecretClient.BaseClient().(*kubesecret.ResourceClient)
		if !ok {
			// something is wrong
		}
		secret, err := kubeconverters.GlooSecretConverterChain.FromKubeSecret(ctx, secretResourceClient, i)
		if err != nil {
			// do something
		}
		if secret == nil {
			return nil
		}
		glooSecret, ok := secret.(*gloov1.Secret)
		if !ok {
			// something else is wrong
			return nil
		}

		out := ResourceWrapper[*gloov1.Secret]{glooSecret}
		return &out
	})

	KubeUpstreams := setupCollectionDynamic[glookubev1.Upstream](
		ctx,
		s.istioClient,
		glookubev1.SchemeGroupVersion.WithResource("upstreams"),
		krt.WithName("Upstreams"),
	)
	GlooUpstreams := krt.NewCollection(KubeUpstreams, func(kctx krt.HandlerContext, u *glookubev1.Upstream) *UpstreamWrapper {
		// TODO: not cloning, this is already a copy from the underlying cache, right?!
		glooUs := &u.Spec
		glooUs.Metadata = &core.Metadata{}
		glooUs.Metadata.Name = u.GetName()
		glooUs.Metadata.Namespace = u.GetNamespace()
		us := UpstreamWrapper{glooUs}
		return &us
	}, krt.WithName("InMemoryUpstreams"))

	serviceClient := kclient.New[*corev1.Service](s.istioClient)
	Services := krt.WrapClient(serviceClient, krt.WithName("Services"))
	InMemUpstreams := krt.NewManyCollection(Services, func(kctx krt.HandlerContext, svc *corev1.Service) []UpstreamWrapper {
		uss := []UpstreamWrapper{}
		for _, port := range svc.Spec.Ports {
			us := kubeupstreams.ServiceToUpstream(ctx, svc, port)
			uss = append(uss, UpstreamWrapper{us})
		}
		return uss
	})
	// TODO: get upstream collections from extensions
	FinalUpstreams := krt.JoinCollection([]krt.Collection[UpstreamWrapper]{GlooUpstreams, InMemUpstreams})
	inputs := NewGlooK8sEndpointInputs(s.proxyTranslator.settings, s.istioClient, s.pods, Services, FinalUpstreams)
	GlooEndpoints := NewGlooK8sEndpoints(ctx, inputs)

	kubeGateways := setupCollectionDynamic[gwv1.Gateway](
		ctx,
		s.istioClient,
		istiogvr.KubernetesGateway_v1,
		krt.WithName("KubeGateways"),
	)

	// RouteOptions -> (fetch Gateways from targetRef) -> empty
	// krt.NewCollection(RouteOptions)

	// TODO: figure out the startSynced stuff
	proxyTrigger := krt.NewRecomputeTrigger(true)

	glooProxies := krt.NewCollection(kubeGateways, func(kctx krt.HandlerContext, gw *gwv1.Gateway) **glooProxy {
		proxyTrigger.MarkDependant(kctx)
		proxy := s.buildProxy(ctx, gw)
		return &proxy
	})

	mostXdsSnapshots := krt.NewCollection(glooProxies, func(kctx krt.HandlerContext, proxy *glooProxy) *xdsSnapWrapper {
		krtSnap := &fromKrtSnap{
			artifacts: artifacts,
			// rtOpts:    krt.Fetch(kctx, RouteOptions),
			secrets:   krtSecrets,
			upstreams: FinalUpstreams,
			// vhostOpts: krt.Fetch(kctx, VirtualHostOptions),
		}
		xdsSnap := s.buildXdsSnapshot(ctx, kctx, proxy, krtSnap)
		return xdsSnap
	})

	mostXdsSnapshotsIndex := krt.NewIndex(mostXdsSnapshots, func(snap xdsSnapWrapper) []string {
		// TODO: make sure this matches the gateway name/namespace. or whatever we can correlate to envoy xds node id.
		return []string{snap.proxyKey}
	})
	mostXdsSnapshots.Register(func(e krt.Event[xdsSnapWrapper]) {
		snap := e.Latest()

		err := s.proxyTranslator.syncStatus(ctx, snap.fullReports)
		if err != nil {
			// fixme
		}

		// TODO: handle garbage collection on status plugins
		var proxiesWithReports []translatorutils.ProxyWithReports
		proxiesWithReports = append(proxiesWithReports, snap.proxyWithReport)
		applyStatusPlugins(ctx, proxiesWithReports, snap.pluginRegistry)
	})

	// sync snapshot per client
	destRuleClient := kclient.NewDelayedInformer[*networkingclient.DestinationRule](s.istioClient, gvr.DestinationRule, kubetypes.StandardInformer, kclient.Filter{})
	rawDestrules := krt.WrapClient(destRuleClient, krt.WithName("DestinationRules"))
	destrules := krt.NewCollection(rawDestrules, func(kctx krt.HandlerContext, dr *networkingclient.DestinationRule) *DestinationRuleWrapper {
		return &DestinationRuleWrapper{dr}
	})
	destrulesidx := newDestruleIndex(destrules)
	uccEndpoints := newIndexedEndpoints(s.uniqlyConnectedClient, GlooEndpoints, destrules, destrulesidx)
	perclientSnapCollection := snapshotPerClient(s.uniqlyConnectedClient, mostXdsSnapshots, mostXdsSnapshotsIndex, uccEndpoints)

	perclientSnapCollection.Register(func(e krt.Event[xdsSnapWrapper]) {
		snap := e.Latest()
		s.proxyTranslator.syncXds(ctx, snap.snap, snap.proxyKey, snap.fullReports)
	})

	// wait for caches to sync before accepting events and syncing xds
	if !s.mgr.GetCache().WaitForCacheSync(ctx) {
		return errors.New("kube gateway sync loop waiting for all caches to sync failed")
	}

	for {
		select {
		case <-ctx.Done():
			contextutils.LoggerFrom(ctx).Debug("context done, stopping proxy syncer")
			return nil
		case <-s.inputs.genericEvent.Next():
			proxyTrigger.TriggerRecomputation()
		case <-s.inputs.secretEvent.Next():
			proxyTrigger.TriggerRecomputation()
		}
	}
}

func (s *ProxySyncer) buildXdsSnapshot(ctx context.Context, kctx krt.HandlerContext, proxy *glooProxy, k *fromKrtSnap) *xdsSnapWrapper {
	latestSnap := snapshot.Snapshot{}
	latestSnap.Proxies = snapshot.SliceCollection[*gloov1.Proxy]([]*gloov1.Proxy{proxy.proxy})

	latestSnap.Artifacts = KrtWrappedCollection[*gloov1.Artifact]{
		C:    k.artifacts,
		Kctx: kctx,
	}
	latestSnap.Secrets = KrtWrappedCollection[*gloov1.Secret]{
		C:    k.secrets,
		Kctx: kctx,
	}

	latestSnap.Upstreams = KrtWrappedCollection[*gloov1.Upstream]{
		C:    k.upstreams,
		Kctx: kctx,
	}

	xdsSnapshot, reports, proxyReport := s.proxyTranslator.buildXdsSnapshot(ctx, proxy.proxy, &latestSnap)
	// TODO(Law): now we not able to merge reports after translation!
	filteredReports := reports.FilterByKind("Proxy")

	// build ResourceReports struct containing only this Proxy
	r := make(reporter.ResourceReports)
	r[proxy.proxy] = filteredReports[proxy.proxy]

	// build object used by status plugins
	proxyWithReport := translatorutils.ProxyWithReports{
		Proxy: proxy.proxy,
		Reports: translatorutils.TranslationReports{
			ProxyReport:     proxyReport,
			ResourceReports: r,
		},
	}
	envoySnap, ok := xdsSnapshot.(*xds.EnvoySnapshot)
	if !ok {
		// fixme
	}
	out := xdsSnapWrapper{
		snap:            envoySnap,
		proxyKey:        proxy.ResourceName(),
		proxyWithReport: proxyWithReport,
		// propagate plugins
		pluginRegistry: proxy.pluginRegistry,
		fullReports:    reports,
	}
	return &out
}

func (s *ProxySyncer) buildProxy(ctx context.Context, gw *gwv1.Gateway) *glooProxy {
	stopwatch := statsutils.NewTranslatorStopWatch("ProxySyncer")
	stopwatch.Start()
	var (
		proxies gloov1.ProxyList
	)
	defer func() {
		duration := stopwatch.Stop(ctx)
		contextutils.LoggerFrom(ctx).Debugf("translated and wrote %d proxies in %s", len(proxies), duration.String())
	}()

	pluginRegistry := s.k8sGwExtensions.CreatePluginRegistry(ctx)
	rm := reports.NewReportMap()
	r := reports.NewReporter(&rm)

	var translatedGateways []gwplugins.TranslatedGateway
	gatewayTranslator := s.k8sGwExtensions.GetTranslator(ctx, gw, pluginRegistry)
	if gatewayTranslator == nil {
		contextutils.LoggerFrom(ctx).Errorf("no translator found for Gateway %s (gatewayClass %s)", gw.Name, gw.Spec.GatewayClassName)
		return nil
	}
	proxy := gatewayTranslator.TranslateProxy(ctx, gw, s.writeNamespace, r)
	if proxy != nil {
		proxies = append(proxies, proxy)
		translatedGateways = append(translatedGateways, gwplugins.TranslatedGateway{
			Gateway: *gw,
		})
	}

	applyPostTranslationPlugins(ctx, pluginRegistry, &gwplugins.PostTranslationContext{
		TranslatedGateways: translatedGateways,
	})

	// reconcile proxy for extensions
	s.reconcileProxies(ctx, proxies)

	// sync gateway api resource status
	s.syncStatus(ctx, rm, gw)
	s.syncRouteStatus(ctx, rm)

	return &glooProxy{
		proxy:          proxy,
		pluginRegistry: pluginRegistry,
	}
}

type metaObjWithSpec interface {
	metav1.Object
	GetSpec() protoiface.MessageV1
}

type rtOptWrapper struct {
	*sologatewayv1.RouteOption
}

func wrapRtOpts(in []*sologatewayv1.RouteOption) []*rtOptWrapper {
	out := make([]*rtOptWrapper, 0, len(in))
	for _, i := range in {
		out = append(out, &rtOptWrapper{i})
	}
	return out
}

func (x *rtOptWrapper) GetSpec() protoiface.MessageV1 {
	return &x.Spec
}

type vhostOptWrapper struct {
	*sologatewayv1.VirtualHostOption
}

func wrapVhosts(in []*sologatewayv1.VirtualHostOption) []*vhostOptWrapper {
	out := make([]*vhostOptWrapper, 0, len(in))
	for _, i := range in {
		out = append(out, &vhostOptWrapper{i})
	}
	return out
}

func (x *vhostOptWrapper) GetSpec() protoiface.MessageV1 {
	return &x.Spec
}

func unwrapGlooKubeTypes[I metaObjWithSpec, O resources.Resource](in []I) []O {
	out := make([]O, 0, len(in))
	for _, i := range in {
		o := proto.Clone(i.GetSpec()).(O)
		m := &core.Metadata{
			Name:      i.GetName(),
			Namespace: i.GetNamespace(),
		}
		o.SetMetadata(m)
		out = append(out, o)
	}
	return out
}

// setupCollectionDynamic uses the dynamic client to setup an informer for a resource
// and then uses an intermediate krt collection to type the unstructured resource.
// This is a temporary workaround until we update to the latest istio version and can
// uncomment the code below for registering types.
// HACK: we don't want to use this long term, but it's letting me push forward with deveopment
func setupCollectionDynamic[T any](ctx context.Context, client kube.Client, gvr schema.GroupVersionResource, opts ...krt.CollectionOption) krt.Collection[*T] {
	gatewayClient := kclient.NewDelayedInformer[*unstructured.Unstructured](client, gvr, kubetypes.DynamicInformer, kclient.Filter{})
	GatewayMapper := krt.WrapClient(gatewayClient, opts...)
	return krt.NewCollection(GatewayMapper, func(krtctx krt.HandlerContext, i *unstructured.Unstructured) **T {
		var empty T
		out := &empty
		err := runtime.DefaultUnstructuredConverter.FromUnstructured(i.UnstructuredContent(), out)
		if err != nil {
			contextutils.LoggerFrom(ctx).DPanic("failed converting unstructured into %T: %v", empty, i)
			return nil
		}
		return &out
	})
}

func applyStatusPlugins(
	ctx context.Context,
	proxiesWithReports []translatorutils.ProxyWithReports,
	registry registry.PluginRegistry,
) {
	ctx = contextutils.WithLogger(ctx, "k8sGatewayStatusPlugins")
	logger := contextutils.LoggerFrom(ctx)

	statusCtx := &gwplugins.StatusContext{
		ProxiesWithReports: proxiesWithReports,
	}
	for _, plugin := range registry.GetStatusPlugins() {
		err := plugin.ApplyStatusPlugin(ctx, statusCtx)
		if err != nil {
			logger.Errorf("Error applying status plugin: %v", err)
			continue
		}
	}
}

func (s *ProxySyncer) syncRouteStatus(ctx context.Context, rm reports.ReportMap) {
	ctx = contextutils.WithLogger(ctx, "routeStatusSyncer")
	logger := contextutils.LoggerFrom(ctx)
	logger.Debugf("syncing k8s gateway route status")
	stopwatch := statsutils.NewTranslatorStopWatch("HTTPRouteStatusSyncer")
	stopwatch.Start()
	defer stopwatch.Stop(ctx)

	rl := gwv1.HTTPRouteList{}
	err := s.mgr.GetClient().List(ctx, &rl)
	if err != nil {
		logger.Error(err)
		return
	}

	for _, route := range rl.Items {
		route := route // pike
		if status := rm.BuildRouteStatus(ctx, route, s.controllerName); status != nil {
			if !isHTTPRouteStatusEqual(&route.Status, status) {
				route.Status = *status
				if err := s.mgr.GetClient().Status().Update(ctx, &route); err != nil {
					logger.Error(err)
				}
			}
		}
	}
}

// syncStatus updates the status of the Gateway CRs
func (s *ProxySyncer) syncStatus(ctx context.Context, rm reports.ReportMap, gw *gwv1.Gateway) {
	ctx = contextutils.WithLogger(ctx, "statusSyncer")
	logger := contextutils.LoggerFrom(ctx)
	stopwatch := statsutils.NewTranslatorStopWatch("GatewayStatusSyncer")
	stopwatch.Start()
	defer stopwatch.Stop(ctx)

	if status := rm.BuildGWStatus(ctx, *gw); status != nil {
		if !isGatewayStatusEqual(&gw.Status, status) {
			gw.Status = *status
			if err := s.mgr.GetClient().Status().Patch(ctx, gw, client.Merge); err != nil {
				logger.Error(err)
			}
		}
	}
}

// reconcileProxies persists the proxies that were generated during translations and stores them in an in-memory cache
// The Gloo Xds Translator will receive these proxies via List() using a MultiResourceClient; two reasons it is needed there:
// 1. To allow Rate Limit extensions to work, as it only syncs RL configs it finds used on Proxies in the snapshots
// 2. This cache is utilized by the debug.ProxyEndpointServer
func (s *ProxySyncer) reconcileProxies(ctx context.Context, proxyList gloov1.ProxyList) {
	ctx = contextutils.WithLogger(ctx, "proxyCache")
	logger := contextutils.LoggerFrom(ctx)

	// Proxy CR is located in the writeNamespace, which may be different from the originating Gateway CR
	err := s.proxyReconciler.Reconcile(
		s.writeNamespace,
		proxyList,
		func(original, desired *gloov1.Proxy) (bool, error) {
			// only reconcile if proxies are not equal
			// we reconcile so ggv2 proxies can be used in extension syncing and debug snap storage
			return !proto.Equal(original, desired), nil
		},
		clients.ListOpts{
			Ctx:      ctx,
			Selector: kubeGatewayProxyLabels,
		})
	if err != nil {
		// A write error to our cache should not impact translation
		// We will emit a message, and continue
		logger.Error(err)
	}
}

func applyPostTranslationPlugins(ctx context.Context, pluginRegistry registry.PluginRegistry, translationContext *gwplugins.PostTranslationContext) {
	ctx = contextutils.WithLogger(ctx, "postTranslation")
	logger := contextutils.LoggerFrom(ctx)

	for _, postTranslationPlugin := range pluginRegistry.GetPostTranslationPlugins() {
		err := postTranslationPlugin.ApplyPostTranslationPlugin(ctx, translationContext)
		if err != nil {
			logger.Errorf("Error applying post-translation plugin: %v", err)
			continue
		}
	}
}

var opts = cmp.Options{
	cmpopts.IgnoreFields(metav1.Condition{}, "LastTransitionTime"),
	cmpopts.IgnoreMapEntries(func(k string, _ any) bool {
		return k == "lastTransitionTime"
	}),
}

func isGatewayStatusEqual(objA, objB *gwv1.GatewayStatus) bool {
	return cmp.Equal(objA, objB, opts)
}

func isHTTPRouteStatusEqual(objA, objB *gwv1.HTTPRouteStatus) bool {
	return cmp.Equal(objA, objB, opts)
}
