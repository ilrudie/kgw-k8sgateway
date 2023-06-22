// Code generated by skv2. DO NOT EDIT.

package ratelimit_resource_handler

import (
	"context"
	"sort"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"

	"github.com/solo-io/go-utils/contextutils"
	skv2v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	ratelimit_solo_io_v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"
	rpc_edge_v1 "github.com/solo-io/solo-projects/projects/apiserver/pkg/api/rpc.edge.gloo/v1"
	"github.com/solo-io/solo-projects/projects/apiserver/server/apiserverutils"
	fedv1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.solo.io/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewFedRatelimitResourceHandler(
	instanceClient fedv1.GlooInstanceClient,
	mcRatelimitCRDClientset ratelimit_solo_io_v1alpha1.MulticlusterClientset,

) rpc_edge_v1.RatelimitResourceApiServer {
	return &fedRatelimitResourceHandler{
		instanceClient:          instanceClient,
		mcRatelimitCRDClientset: mcRatelimitCRDClientset,
	}
}

type fedRatelimitResourceHandler struct {
	instanceClient          fedv1.GlooInstanceClient
	mcRatelimitCRDClientset ratelimit_solo_io_v1alpha1.MulticlusterClientset
}

func (k *fedRatelimitResourceHandler) ListRateLimitConfigs(ctx context.Context, request *rpc_edge_v1.ListRateLimitConfigsRequest) (*rpc_edge_v1.ListRateLimitConfigsResponse, error) {
	var rpcRateLimitConfigs []*rpc_edge_v1.RateLimitConfig
	if request.GetGlooInstanceRef() == nil || request.GetGlooInstanceRef().GetName() == "" || request.GetGlooInstanceRef().GetNamespace() == "" {
		// List rateLimitConfigs across all gloo edge instances
		instanceList, err := k.instanceClient.ListGlooInstance(ctx)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list gloo edge instances")
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		for _, instance := range instanceList.Items {
			rpcRateLimitConfigList, err := k.listRateLimitConfigsForGlooInstance(ctx, &instance)
			if err != nil {
				wrapped := eris.Wrapf(err, "Failed to list rateLimitConfigs for gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
				contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
				return nil, wrapped
			}
			rpcRateLimitConfigs = append(rpcRateLimitConfigs, rpcRateLimitConfigList...)
		}
	} else {
		// List rateLimitConfigs for a specific gloo edge instance
		instance, err := k.instanceClient.GetGlooInstance(ctx, types.NamespacedName{
			Name:      request.GetGlooInstanceRef().GetName(),
			Namespace: request.GetGlooInstanceRef().GetNamespace(),
		})
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to get gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		rpcRateLimitConfigs, err = k.listRateLimitConfigsForGlooInstance(ctx, instance)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list rateLimitConfigs for gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
	}

	// Search, Filter
	var filteredRateLimitConfigs []*rpc_edge_v1.RateLimitConfig
	qs := request.GetQueryString()
	sf := request.GetStatusFilter()
	if sf != nil || qs != "" {
		for _, d := range rpcRateLimitConfigs {
			if (sf == nil || sf.State == int32(d.Status.State)) && strings.Contains(d.Metadata.Name, qs) {
				filteredRateLimitConfigs = append(filteredRateLimitConfigs, d)
			}
		}
	} else {
		filteredRateLimitConfigs = rpcRateLimitConfigs
	}
	// Sort
	sortOptions := request.GetSortOptions()
	if sortOptions != nil {
		isDescending := sortOptions.GetDescending()
		sortKey := sortOptions.GetSortKey()
		if isDescending == true {
			switch sortKey {
			case rpc_edge_v1.SortOptions_NAME:
				sort.SliceStable(filteredRateLimitConfigs, func(i, j int) bool {
					a := filteredRateLimitConfigs[i]
					b := filteredRateLimitConfigs[j]
					return a.Metadata.Name+a.Metadata.Namespace+string(a.Status.State) > b.Metadata.Name+b.Metadata.Namespace+string(b.Status.State)
				})
			case rpc_edge_v1.SortOptions_NAMESPACE:
				sort.SliceStable(filteredRateLimitConfigs, func(i, j int) bool {
					a := filteredRateLimitConfigs[i]
					b := filteredRateLimitConfigs[j]
					return a.Metadata.Namespace+a.Metadata.Name+string(a.Status.State) > b.Metadata.Namespace+b.Metadata.Name+string(b.Status.State)
				})
			case rpc_edge_v1.SortOptions_STATUS:
				sort.SliceStable(filteredRateLimitConfigs, func(i, j int) bool {
					a := filteredRateLimitConfigs[i]
					b := filteredRateLimitConfigs[j]
					return string(a.Status.State)+a.Metadata.Name+a.Metadata.Namespace > string(b.Status.State)+b.Metadata.Name+b.Metadata.Namespace
				})
			}
		} else {
			switch sortKey {
			case rpc_edge_v1.SortOptions_NAME:
				sort.SliceStable(filteredRateLimitConfigs, func(i, j int) bool {
					a := filteredRateLimitConfigs[i]
					b := filteredRateLimitConfigs[j]
					return a.Metadata.Name+a.Metadata.Namespace+string(a.Status.State) < b.Metadata.Name+b.Metadata.Namespace+string(b.Status.State)
				})
			case rpc_edge_v1.SortOptions_NAMESPACE:
				sort.SliceStable(filteredRateLimitConfigs, func(i, j int) bool {
					a := filteredRateLimitConfigs[i]
					b := filteredRateLimitConfigs[j]
					return a.Metadata.Namespace+a.Metadata.Name+string(a.Status.State) < b.Metadata.Namespace+b.Metadata.Name+string(b.Status.State)
				})
			case rpc_edge_v1.SortOptions_STATUS:
				sort.SliceStable(filteredRateLimitConfigs, func(i, j int) bool {
					a := filteredRateLimitConfigs[i]
					b := filteredRateLimitConfigs[j]
					return string(a.Status.State)+a.Metadata.Name+a.Metadata.Namespace < string(b.Status.State)+b.Metadata.Name+b.Metadata.Namespace
				})
			}
		}
	}
	// Paginate
	paginatedRateLimitConfigs := filteredRateLimitConfigs
	pagination := request.GetPagination()
	totalCount := int32(len(filteredRateLimitConfigs))
	if pagination.GetLimit() > 0 && pagination.GetOffset() >= 0 {
		start := apiserverutils.Min(pagination.GetOffset(), totalCount)
		end := apiserverutils.Min(pagination.GetOffset()+pagination.GetLimit(), totalCount)
		paginatedRateLimitConfigs = filteredRateLimitConfigs[start:end]
	}
	// Build response
	res := &rpc_edge_v1.ListRateLimitConfigsResponse{
		RateLimitConfigs: paginatedRateLimitConfigs,
		Total:            totalCount,
	}
	return res, nil
}

func (k *fedRatelimitResourceHandler) listRateLimitConfigsForGlooInstance(ctx context.Context, instance *fedv1.GlooInstance) ([]*rpc_edge_v1.RateLimitConfig, error) {

	ratelimitCRDClientset, err := k.mcRatelimitCRDClientset.Cluster(instance.Spec.GetCluster())
	if err != nil {
		return nil, err
	}
	rateLimitConfigClient := ratelimitCRDClientset.RateLimitConfigs()

	var ratelimitRateLimitConfigList []*ratelimit_solo_io_v1alpha1.RateLimitConfig
	watchedNamespaces := instance.Spec.GetControlPlane().GetWatchedNamespaces()
	if len(watchedNamespaces) != 0 {
		for _, ns := range watchedNamespaces {
			list, err := rateLimitConfigClient.ListRateLimitConfig(ctx, client.InNamespace(ns))
			if err != nil {
				return nil, err
			}
			for i, _ := range list.Items {
				ratelimitRateLimitConfigList = append(ratelimitRateLimitConfigList, &list.Items[i])
			}
		}
	} else {
		list, err := rateLimitConfigClient.ListRateLimitConfig(ctx)
		if err != nil {
			return nil, err
		}
		for i, _ := range list.Items {
			ratelimitRateLimitConfigList = append(ratelimitRateLimitConfigList, &list.Items[i])
		}
	}
	sort.Slice(ratelimitRateLimitConfigList, func(i, j int) bool {
		x := ratelimitRateLimitConfigList[i]
		y := ratelimitRateLimitConfigList[j]
		return x.GetNamespace()+x.GetName() < y.GetNamespace()+y.GetName()
	})

	var rpcRateLimitConfigs []*rpc_edge_v1.RateLimitConfig
	for _, rateLimitConfig := range ratelimitRateLimitConfigList {
		rpcRateLimitConfigs = append(rpcRateLimitConfigs, BuildRpcRateLimitConfig(rateLimitConfig, &skv2v1.ObjectRef{
			Name:      instance.GetName(),
			Namespace: instance.GetNamespace(),
		}, instance.Spec.GetCluster()))
	}
	return rpcRateLimitConfigs, nil
}

func BuildRpcRateLimitConfig(rateLimitConfig *ratelimit_solo_io_v1alpha1.RateLimitConfig, glooInstance *skv2v1.ObjectRef, cluster string) *rpc_edge_v1.RateLimitConfig {
	m := &rpc_edge_v1.RateLimitConfig{
		Metadata:     apiserverutils.ToMetadata(rateLimitConfig.ObjectMeta),
		GlooInstance: glooInstance,
		Spec:         &rateLimitConfig.Spec,
		Status:       &rateLimitConfig.Status,
	}
	m.Metadata.ClusterName = cluster
	return m
}

func (k *fedRatelimitResourceHandler) GetRateLimitConfigYaml(ctx context.Context, request *rpc_edge_v1.GetRateLimitConfigYamlRequest) (*rpc_edge_v1.GetRateLimitConfigYamlResponse, error) {
	ratelimitClientSet, err := k.mcRatelimitCRDClientset.Cluster(request.GetRateLimitConfigRef().GetClusterName())
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get ratelimit client set")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	rateLimitConfig, err := ratelimitClientSet.RateLimitConfigs().GetRateLimitConfig(ctx, client.ObjectKey{
		Namespace: request.GetRateLimitConfigRef().GetNamespace(),
		Name:      request.GetRateLimitConfigRef().GetName(),
	})
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get rateLimitConfig")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	content, err := yaml.Marshal(rateLimitConfig)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to marshal kube resource into yaml")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	return &rpc_edge_v1.GetRateLimitConfigYamlResponse{
		YamlData: &rpc_edge_v1.ResourceYaml{
			Yaml: string(content),
		},
	}, nil
}

func (k *fedRatelimitResourceHandler) GetRateLimitConfigDetails(ctx context.Context, request *rpc_edge_v1.GetRateLimitConfigDetailsRequest) (*rpc_edge_v1.GetRateLimitConfigDetailsResponse, error) {
	RateLimitConfigRef := request.GetRateLimitConfigRef()
	if RateLimitConfigRef == nil {
		return nil, eris.Errorf("RateLimitConfig ref missing from request: %v", request)
	}
	clientset, err := k.mcRatelimitCRDClientset.Cluster(RateLimitConfigRef.GetClusterName())
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get RateLimitConfig client set")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	RateLimitConfig, err := clientset.RateLimitConfigs().GetRateLimitConfig(ctx, client.ObjectKey{
		Namespace: RateLimitConfigRef.GetNamespace(),
		Name:      RateLimitConfigRef.GetName(),
	})
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get RateLimitConfig")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	// find which gloo instance this RateLimitConfig belongs to, by finding a gloo instance that is watching
	// the RateLimitConfig's namespace
	instanceList, err := k.instanceClient.ListGlooInstance(ctx)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to list gloo edge instances")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	var glooInstance *fedv1.GlooInstance
	for _, instance := range instanceList.Items {
		watchedNamespaces := instance.Spec.GetControlPlane().GetWatchedNamespaces()
		if len(watchedNamespaces) == 0 {
			glooInstance = &instance
			break
		}
		for _, ns := range watchedNamespaces {
			if ns == RateLimitConfigRef.GetNamespace() {
				glooInstance = &instance
				break
			}
		}
		if glooInstance != nil {
			break
		}
	}
	if glooInstance == nil {
		return nil, eris.Errorf("Failed to find a gloo edge instance for namespace %s", RateLimitConfig.GetNamespace())
	}
	rpcRateLimitConfig := &rpc_edge_v1.RateLimitConfig{
		Metadata: apiserverutils.ToMetadata(RateLimitConfig.ObjectMeta),
		GlooInstance: &skv2v1.ObjectRef{
			Name:      glooInstance.GetName(),
			Namespace: glooInstance.GetNamespace(),
		},
		Spec:   &RateLimitConfig.Spec,
		Status: &RateLimitConfig.Status,
	}
	rpcRateLimitConfig.Metadata.ClusterName = glooInstance.Spec.GetCluster()
	return &rpc_edge_v1.GetRateLimitConfigDetailsResponse{
		RateLimitConfig: rpcRateLimitConfig,
	}, nil
}