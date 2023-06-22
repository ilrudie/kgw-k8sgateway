package sanitize_cluster_header

import (
	"github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/extauth"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
)

var (
	_ plugins.Plugin           = new(plugin)
	_ plugins.HttpFilterPlugin = new(plugin)
)

const (
	ExtensionName      = "sanitize_cluster_header"
	SanitizeFilterName = "io.solo.filters.http.sanitize"
)

var (
	sanitizeFilterStage = plugins.BeforeStage(plugins.AuthNStage)
)

type plugin struct{}

func NewPlugin() *plugin {
	return &plugin{}
}

func (p *plugin) Name() string {
	return ExtensionName
}

func (p *plugin) Init(_ plugins.InitParams) {
}

func (p *plugin) HttpFilters(params plugins.Params, listener *v1.HttpListener) ([]plugins.StagedHttpFilter, error) {
	var filters []plugins.StagedHttpFilter
	// check if we should sanitize for cluster_header
	if sanitizeClusterHeader := listener.GetOptions().GetSanitizeClusterHeader(); sanitizeClusterHeader.GetValue() == false {
		return filters, nil
	}
	headersToRemoveSet := make(map[string]struct{})

	// get all headers used for cluster_header destination
	hosts := listener.GetVirtualHosts()
	for _, host := range hosts {
		routes := host.GetRoutes()
		for _, route := range routes {
			if header := route.GetRouteAction().GetClusterHeader(); header != "" {
				headersToRemoveSet[header] = struct{}{}
			}
		}
	}
	// sanitize those headers from downstreams
	if len(headersToRemoveSet) > 0 {
		var headersToRemove []string
		for header := range headersToRemoveSet {
			headersToRemove = append(headersToRemove, header)
		}
		sanitizeConf := &extauth.Sanitize{HeadersToRemove: headersToRemove}
		stagedFilter, err := plugins.NewStagedFilterWithConfig(SanitizeFilterName, sanitizeConf, sanitizeFilterStage)
		if err != nil {
			return nil, err
		}
		filters = append(filters, stagedFilter)
	}
	return filters, nil
}