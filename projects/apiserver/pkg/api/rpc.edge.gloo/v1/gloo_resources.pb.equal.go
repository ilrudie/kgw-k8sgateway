// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/apiserver/api/rpc.edge.gloo/v1/gloo_resources.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *Upstream) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Upstream)
	if !ok {
		that2, ok := that.(Upstream)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSpec(), target.GetSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatus(), target.GetStatus()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGlooInstance()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstance()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstance(), target.GetGlooInstance()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *UpstreamGroup) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamGroup)
	if !ok {
		that2, ok := that.(UpstreamGroup)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSpec(), target.GetSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatus(), target.GetStatus()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGlooInstance()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstance()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstance(), target.GetGlooInstance()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Settings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSpec(), target.GetSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatus(), target.GetStatus()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGlooInstance()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstance()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstance(), target.GetGlooInstance()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Proxy) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Proxy)
	if !ok {
		that2, ok := that.(Proxy)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSpec(), target.GetSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatus()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatus()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatus(), target.GetStatus()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGlooInstance()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstance()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstance(), target.GetGlooInstance()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListUpstreamsRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListUpstreamsRequest)
	if !ok {
		that2, ok := that.(ListUpstreamsRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGlooInstanceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstanceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstanceRef(), target.GetGlooInstanceRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPagination()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPagination()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPagination(), target.GetPagination()) {
			return false
		}
	}

	if strings.Compare(m.GetQueryString(), target.GetQueryString()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetStatusFilter()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatusFilter()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatusFilter(), target.GetStatusFilter()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSortOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSortOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSortOptions(), target.GetSortOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListUpstreamsResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListUpstreamsResponse)
	if !ok {
		that2, ok := that.(ListUpstreamsResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetUpstreams()) != len(target.GetUpstreams()) {
		return false
	}
	for idx, v := range m.GetUpstreams() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetUpstreams()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetUpstreams()[idx]) {
				return false
			}
		}

	}

	if m.GetTotal() != target.GetTotal() {
		return false
	}

	return true
}

// Equal function
func (m *GetUpstreamYamlRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetUpstreamYamlRequest)
	if !ok {
		that2, ok := that.(GetUpstreamYamlRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetUpstreamRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstreamRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstreamRef(), target.GetUpstreamRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetUpstreamYamlResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetUpstreamYamlResponse)
	if !ok {
		that2, ok := that.(GetUpstreamYamlResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetYamlData()).(equality.Equalizer); ok {
		if !h.Equal(target.GetYamlData()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetYamlData(), target.GetYamlData()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetUpstreamDetailsRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetUpstreamDetailsRequest)
	if !ok {
		that2, ok := that.(GetUpstreamDetailsRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetUpstreamRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstreamRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstreamRef(), target.GetUpstreamRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetUpstreamDetailsResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetUpstreamDetailsResponse)
	if !ok {
		that2, ok := that.(GetUpstreamDetailsResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetUpstream()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstream()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstream(), target.GetUpstream()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListUpstreamGroupsRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListUpstreamGroupsRequest)
	if !ok {
		that2, ok := that.(ListUpstreamGroupsRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGlooInstanceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstanceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstanceRef(), target.GetGlooInstanceRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPagination()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPagination()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPagination(), target.GetPagination()) {
			return false
		}
	}

	if strings.Compare(m.GetQueryString(), target.GetQueryString()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetStatusFilter()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatusFilter()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatusFilter(), target.GetStatusFilter()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSortOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSortOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSortOptions(), target.GetSortOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListUpstreamGroupsResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListUpstreamGroupsResponse)
	if !ok {
		that2, ok := that.(ListUpstreamGroupsResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetUpstreamGroups()) != len(target.GetUpstreamGroups()) {
		return false
	}
	for idx, v := range m.GetUpstreamGroups() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetUpstreamGroups()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetUpstreamGroups()[idx]) {
				return false
			}
		}

	}

	if m.GetTotal() != target.GetTotal() {
		return false
	}

	return true
}

// Equal function
func (m *GetUpstreamGroupYamlRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetUpstreamGroupYamlRequest)
	if !ok {
		that2, ok := that.(GetUpstreamGroupYamlRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetUpstreamGroupRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstreamGroupRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstreamGroupRef(), target.GetUpstreamGroupRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetUpstreamGroupYamlResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetUpstreamGroupYamlResponse)
	if !ok {
		that2, ok := that.(GetUpstreamGroupYamlResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetYamlData()).(equality.Equalizer); ok {
		if !h.Equal(target.GetYamlData()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetYamlData(), target.GetYamlData()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetUpstreamGroupDetailsRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetUpstreamGroupDetailsRequest)
	if !ok {
		that2, ok := that.(GetUpstreamGroupDetailsRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetUpstreamGroupRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstreamGroupRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstreamGroupRef(), target.GetUpstreamGroupRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetUpstreamGroupDetailsResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetUpstreamGroupDetailsResponse)
	if !ok {
		that2, ok := that.(GetUpstreamGroupDetailsResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetUpstreamGroup()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstreamGroup()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstreamGroup(), target.GetUpstreamGroup()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListSettingsRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListSettingsRequest)
	if !ok {
		that2, ok := that.(ListSettingsRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGlooInstanceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstanceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstanceRef(), target.GetGlooInstanceRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPagination()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPagination()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPagination(), target.GetPagination()) {
			return false
		}
	}

	if strings.Compare(m.GetQueryString(), target.GetQueryString()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetStatusFilter()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatusFilter()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatusFilter(), target.GetStatusFilter()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSortOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSortOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSortOptions(), target.GetSortOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListSettingsResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListSettingsResponse)
	if !ok {
		that2, ok := that.(ListSettingsResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetSettings()) != len(target.GetSettings()) {
		return false
	}
	for idx, v := range m.GetSettings() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSettings()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSettings()[idx]) {
				return false
			}
		}

	}

	if m.GetTotal() != target.GetTotal() {
		return false
	}

	return true
}

// Equal function
func (m *GetSettingsYamlRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetSettingsYamlRequest)
	if !ok {
		that2, ok := that.(GetSettingsYamlRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetSettingsRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSettingsRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSettingsRef(), target.GetSettingsRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetSettingsYamlResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetSettingsYamlResponse)
	if !ok {
		that2, ok := that.(GetSettingsYamlResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetYamlData()).(equality.Equalizer); ok {
		if !h.Equal(target.GetYamlData()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetYamlData(), target.GetYamlData()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetSettingsDetailsRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetSettingsDetailsRequest)
	if !ok {
		that2, ok := that.(GetSettingsDetailsRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetSettingsRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSettingsRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSettingsRef(), target.GetSettingsRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetSettingsDetailsResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetSettingsDetailsResponse)
	if !ok {
		that2, ok := that.(GetSettingsDetailsResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetSettings()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSettings()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSettings(), target.GetSettings()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListProxiesRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListProxiesRequest)
	if !ok {
		that2, ok := that.(ListProxiesRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetGlooInstanceRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGlooInstanceRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGlooInstanceRef(), target.GetGlooInstanceRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPagination()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPagination()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPagination(), target.GetPagination()) {
			return false
		}
	}

	if strings.Compare(m.GetQueryString(), target.GetQueryString()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetStatusFilter()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatusFilter()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatusFilter(), target.GetStatusFilter()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSortOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSortOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSortOptions(), target.GetSortOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ListProxiesResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListProxiesResponse)
	if !ok {
		that2, ok := that.(ListProxiesResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetProxies()) != len(target.GetProxies()) {
		return false
	}
	for idx, v := range m.GetProxies() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetProxies()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetProxies()[idx]) {
				return false
			}
		}

	}

	if m.GetTotal() != target.GetTotal() {
		return false
	}

	return true
}

// Equal function
func (m *GetProxyYamlRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetProxyYamlRequest)
	if !ok {
		that2, ok := that.(GetProxyYamlRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetProxyRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProxyRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProxyRef(), target.GetProxyRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetProxyYamlResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetProxyYamlResponse)
	if !ok {
		that2, ok := that.(GetProxyYamlResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetYamlData()).(equality.Equalizer); ok {
		if !h.Equal(target.GetYamlData()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetYamlData(), target.GetYamlData()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetProxyDetailsRequest) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetProxyDetailsRequest)
	if !ok {
		that2, ok := that.(GetProxyDetailsRequest)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetProxyRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProxyRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProxyRef(), target.GetProxyRef()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *GetProxyDetailsResponse) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GetProxyDetailsResponse)
	if !ok {
		that2, ok := that.(GetProxyDetailsResponse)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetProxy()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProxy()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProxy(), target.GetProxy()) {
			return false
		}
	}

	return true
}