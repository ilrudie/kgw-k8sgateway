// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"context"
	"net"

	"github.com/solo-io/go-utils/envutils"
	"github.com/solo-io/solo-projects/pkg/license"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/helpers/rawgetter"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/helpers/status"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/internal/client"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/internal/kube"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/internal/settings"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/artifactsvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/configsvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/envoysvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/envoysvc/envoydetails"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/gatewaysvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/proxysvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/routetablesvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/secretsvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/secretsvc/scrub"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/upstreamgroupsvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/upstreamsvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/upstreamsvc/mutation"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/upstreamsvc/search"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/virtualservicesvc"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/virtualservicesvc/converter"
	mutation2 "github.com/solo-io/solo-projects/projects/grpcserver/server/service/virtualservicesvc/mutation"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/service/virtualservicesvc/selection"
	"github.com/solo-io/solo-projects/projects/grpcserver/server/setup"
)

// Injectors from wire.go:

func InitializeServer(ctx context.Context, listener net.Listener) (*GlooGrpcService, error) {
	string2 := envutils.MustGetPodNamespace(ctx)
	v1Settings := setup.MustSettings(ctx, string2)
	config, err := setup.NewKubeConfig()
	if err != nil {
		return nil, err
	}
	token, err := setup.GetToken(config)
	if err != nil {
		return nil, err
	}
	clientCache, err := client.NewClientCache(ctx, v1Settings, config, token, string2)
	if err != nil {
		return nil, err
	}
	licenseClient := license.NewClient(ctx)
	valuesClient := settings.NewSettingsValuesClient(ctx, clientCache, string2)
	mutator := mutation.NewMutator(clientCache)
	factory := mutation.NewFactory()
	rawGetter := rawgetter.NewKubeYamlRawGetter()
	upstreamSearcher := search.NewUpstreamSearcher(clientCache, valuesClient)
	upstreamApiServer := upstreamsvc.NewUpstreamGrpcService(ctx, clientCache, licenseClient, valuesClient, mutator, factory, rawGetter, upstreamSearcher)
	upstreamGroupApiServer := upstreamgroupsvc.NewUpstreamGroupGrpcService(ctx, clientCache, licenseClient, valuesClient, rawGetter, upstreamSearcher)
	artifactApiServer := artifactsvc.NewArtifactGrpcService(ctx, clientCache, licenseClient, valuesClient)
	coreV1Interface, err := setup.GetK8sCoreInterface(config)
	if err != nil {
		return nil, err
	}
	namespacesGetter := setup.NewNamespacesGetter(coreV1Interface)
	namespaceClient := kube.NewNamespaceClient(namespacesGetter)
	oAuthEndpoint := setup.NewOAuthEndpoint()
	buildVersion := setup.GetBuildVersion()
	configApiServer, err := configsvc.NewConfigGrpcService(ctx, clientCache, licenseClient, namespaceClient, oAuthEndpoint, buildVersion, string2)
	if err != nil {
		return nil, err
	}
	scrubber := scrub.NewScrubber()
	secretApiServer := secretsvc.NewSecretGrpcService(ctx, clientCache, scrubber, licenseClient, valuesClient)
	mutationMutator := mutation2.NewMutator(ctx, clientCache, licenseClient)
	mutationFactory := mutation2.NewMutationFactory()
	virtualServiceDetailsConverter := converter.NewVirtualServiceDetailsConverter(rawGetter)
	virtualServiceSelector := selection.NewVirtualServiceSelector(clientCache, namespaceClient, string2)
	virtualServiceApiServer := virtualservicesvc.NewVirtualServiceGrpcService(ctx, string2, clientCache, licenseClient, valuesClient, mutationMutator, mutationFactory, virtualServiceDetailsConverter, virtualServiceSelector, rawGetter)
	routeTableApiServer := routetablesvc.NewRouteTableGrpcService(ctx, clientCache, licenseClient, valuesClient, rawGetter)
	inputResourceStatusGetter := status.NewInputResourceStatusGetter()
	gatewayApiServer := gatewaysvc.NewGatewayGrpcService(ctx, clientCache, rawGetter, inputResourceStatusGetter, licenseClient, valuesClient)
	proxyApiServer := proxysvc.NewProxyGrpcService(ctx, clientCache, rawGetter, inputResourceStatusGetter, valuesClient)
	podsGetter := setup.NewPodsGetter(coreV1Interface)
	httpGetter := envoydetails.NewHttpGetter()
	proxyStatusGetter := envoydetails.NewProxyStatusGetter(clientCache)
	envoydetailsClient := envoydetails.NewClient(podsGetter, httpGetter, proxyStatusGetter)
	envoyApiServer := envoysvc.NewEnvoyGrpcService(ctx, envoydetailsClient, string2)
	updater := client.NewClientUpdater(clientCache, config, token, string2)
	glooGrpcService := NewGlooGrpcService(ctx, listener, upstreamApiServer, upstreamGroupApiServer, artifactApiServer, configApiServer, secretApiServer, virtualServiceApiServer, routeTableApiServer, gatewayApiServer, proxyApiServer, envoyApiServer, updater)
	return glooGrpcService, nil
}
