// Code generated by skv2. DO NOT EDIT.

package gloo_resource_handler

import (
	"context"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/stringutils"
	corev1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	types "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	rpc_edge_v1 "github.com/solo-io/solo-projects/projects/apiserver/pkg/api/rpc.edge.gloo/v1"
	"github.com/solo-io/solo-projects/projects/apiserver/server/apiserverutils"
	"go.uber.org/zap"
)

func GetUpstreamSummary(ctx context.Context, upstreamClient types.UpstreamClient, watchedNamespaces []string) *rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary {
	summary := &rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary{}

	upstreamList, err := upstreamClient.ListUpstream(ctx)
	if err != nil {
		contextutils.LoggerFrom(ctx).Warnw("Failed to get Upstream summary", zap.Error(err), zap.Any("watchedNamespaces", watchedNamespaces))
		return summary
	}

	for _, upstream := range upstreamList.Items {
		upstream := upstream

		// If the resource is not in a watched namespace, continue
		if len(watchedNamespaces) > 0 && !stringutils.ContainsString(upstream.Namespace, watchedNamespaces) {
			continue
		}

		summary.Total += 1

		if upstream.Status.GetState() == types.UpstreamStatus_Rejected {
			summary.Errors = append(summary.Errors, &rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      upstream.Name,
					Namespace: upstream.Namespace,
				},
				Message: upstream.Status.Reason,
			})
		}

		if upstream.Status.GetState() == types.UpstreamStatus_Warning {
			summary.Warnings = append(summary.Warnings, &rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      upstream.Name,
					Namespace: upstream.Namespace,
				},
				Message: upstream.Status.Reason,
			})
		}

	}

	apiserverutils.SortCheckSummaryLists(summary)
	return summary
}

func GetUpstreamGroupSummary(ctx context.Context, upstreamGroupClient types.UpstreamGroupClient, watchedNamespaces []string) *rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary {
	summary := &rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary{}

	upstreamGroupList, err := upstreamGroupClient.ListUpstreamGroup(ctx)
	if err != nil {
		contextutils.LoggerFrom(ctx).Warnw("Failed to get UpstreamGroup summary", zap.Error(err), zap.Any("watchedNamespaces", watchedNamespaces))
		return summary
	}

	for _, upstreamGroup := range upstreamGroupList.Items {
		upstreamGroup := upstreamGroup

		// If the resource is not in a watched namespace, continue
		if len(watchedNamespaces) > 0 && !stringutils.ContainsString(upstreamGroup.Namespace, watchedNamespaces) {
			continue
		}

		summary.Total += 1

		if upstreamGroup.Status.GetState() == types.UpstreamGroupStatus_Rejected {
			summary.Errors = append(summary.Errors, &rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      upstreamGroup.Name,
					Namespace: upstreamGroup.Namespace,
				},
				Message: upstreamGroup.Status.Reason,
			})
		}

		if upstreamGroup.Status.GetState() == types.UpstreamGroupStatus_Warning {
			summary.Warnings = append(summary.Warnings, &rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      upstreamGroup.Name,
					Namespace: upstreamGroup.Namespace,
				},
				Message: upstreamGroup.Status.Reason,
			})
		}

	}

	apiserverutils.SortCheckSummaryLists(summary)
	return summary
}

func GetSettingsSummary(ctx context.Context, settingsClient types.SettingsClient, watchedNamespaces []string) *rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary {
	summary := &rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary{}

	settingsList, err := settingsClient.ListSettings(ctx)
	if err != nil {
		contextutils.LoggerFrom(ctx).Warnw("Failed to get Settings summary", zap.Error(err), zap.Any("watchedNamespaces", watchedNamespaces))
		return summary
	}

	for _, settings := range settingsList.Items {
		settings := settings

		// If the resource is not in a watched namespace, continue
		if len(watchedNamespaces) > 0 && !stringutils.ContainsString(settings.Namespace, watchedNamespaces) {
			continue
		}

		summary.Total += 1

		if settings.Status.GetState() == types.SettingsStatus_Rejected {
			summary.Errors = append(summary.Errors, &rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      settings.Name,
					Namespace: settings.Namespace,
				},
				Message: settings.Status.Reason,
			})
		}

		if settings.Status.GetState() == types.SettingsStatus_Warning {
			summary.Warnings = append(summary.Warnings, &rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      settings.Name,
					Namespace: settings.Namespace,
				},
				Message: settings.Status.Reason,
			})
		}

	}

	apiserverutils.SortCheckSummaryLists(summary)
	return summary
}

func GetProxySummary(ctx context.Context, proxyClient types.ProxyClient, watchedNamespaces []string) *rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary {
	summary := &rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary{}

	proxyList, err := proxyClient.ListProxy(ctx)
	if err != nil {
		contextutils.LoggerFrom(ctx).Warnw("Failed to get Proxy summary", zap.Error(err), zap.Any("watchedNamespaces", watchedNamespaces))
		return summary
	}

	for _, proxy := range proxyList.Items {
		proxy := proxy

		// If the resource is not in a watched namespace, continue
		if len(watchedNamespaces) > 0 && !stringutils.ContainsString(proxy.Namespace, watchedNamespaces) {
			continue
		}

		summary.Total += 1

		if proxy.Status.GetState() == types.ProxyStatus_Rejected {
			summary.Errors = append(summary.Errors, &rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      proxy.Name,
					Namespace: proxy.Namespace,
				},
				Message: proxy.Status.Reason,
			})
		}

		if proxy.Status.GetState() == types.ProxyStatus_Warning {
			summary.Warnings = append(summary.Warnings, &rpc_edge_v1.GlooInstance_GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      proxy.Name,
					Namespace: proxy.Namespace,
				},
				Message: proxy.Status.Reason,
			})
		}

	}

	apiserverutils.SortCheckSummaryLists(summary)
	return summary
}
