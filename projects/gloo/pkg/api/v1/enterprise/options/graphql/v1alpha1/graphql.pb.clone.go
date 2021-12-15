// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/graphql/v1alpha1/graphql.proto

package v1alpha1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_struct "github.com/golang/protobuf/ptypes/struct"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *RequestTemplate) Clone() proto.Message {
	var target *RequestTemplate
	if m == nil {
		return target
	}
	target = &RequestTemplate{}

	if m.GetHeaders() != nil {
		target.Headers = make(map[string]string, len(m.GetHeaders()))
		for k, v := range m.GetHeaders() {

			target.Headers[k] = v

		}
	}

	if m.GetQueryParams() != nil {
		target.QueryParams = make(map[string]string, len(m.GetQueryParams()))
		for k, v := range m.GetQueryParams() {

			target.QueryParams[k] = v

		}
	}

	if h, ok := interface{}(m.GetBody()).(clone.Cloner); ok {
		target.Body = h.Clone().(*github_com_golang_protobuf_ptypes_struct.Value)
	} else {
		target.Body = proto.Clone(m.GetBody()).(*github_com_golang_protobuf_ptypes_struct.Value)
	}

	return target
}

// Clone function
func (m *ResponseTemplate) Clone() proto.Message {
	var target *ResponseTemplate
	if m == nil {
		return target
	}
	target = &ResponseTemplate{}

	target.ResultRoot = m.GetResultRoot()

	if m.GetSetters() != nil {
		target.Setters = make(map[string]string, len(m.GetSetters()))
		for k, v := range m.GetSetters() {

			target.Setters[k] = v

		}
	}

	return target
}

// Clone function
func (m *RESTResolver) Clone() proto.Message {
	var target *RESTResolver
	if m == nil {
		return target
	}
	target = &RESTResolver{}

	if h, ok := interface{}(m.GetUpstreamRef()).(clone.Cloner); ok {
		target.UpstreamRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.UpstreamRef = proto.Clone(m.GetUpstreamRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if h, ok := interface{}(m.GetRequest()).(clone.Cloner); ok {
		target.Request = h.Clone().(*RequestTemplate)
	} else {
		target.Request = proto.Clone(m.GetRequest()).(*RequestTemplate)
	}

	if h, ok := interface{}(m.GetResponse()).(clone.Cloner); ok {
		target.Response = h.Clone().(*ResponseTemplate)
	} else {
		target.Response = proto.Clone(m.GetResponse()).(*ResponseTemplate)
	}

	target.SpanName = m.GetSpanName()

	return target
}

// Clone function
func (m *QueryMatcher) Clone() proto.Message {
	var target *QueryMatcher
	if m == nil {
		return target
	}
	target = &QueryMatcher{}

	switch m.Match.(type) {

	case *QueryMatcher_FieldMatcher_:

		if h, ok := interface{}(m.GetFieldMatcher()).(clone.Cloner); ok {
			target.Match = &QueryMatcher_FieldMatcher_{
				FieldMatcher: h.Clone().(*QueryMatcher_FieldMatcher),
			}
		} else {
			target.Match = &QueryMatcher_FieldMatcher_{
				FieldMatcher: proto.Clone(m.GetFieldMatcher()).(*QueryMatcher_FieldMatcher),
			}
		}

	}

	return target
}

// Clone function
func (m *Resolution) Clone() proto.Message {
	var target *Resolution
	if m == nil {
		return target
	}
	target = &Resolution{}

	if h, ok := interface{}(m.GetMatcher()).(clone.Cloner); ok {
		target.Matcher = h.Clone().(*QueryMatcher)
	} else {
		target.Matcher = proto.Clone(m.GetMatcher()).(*QueryMatcher)
	}

	switch m.Resolver.(type) {

	case *Resolution_RestResolver:

		if h, ok := interface{}(m.GetRestResolver()).(clone.Cloner); ok {
			target.Resolver = &Resolution_RestResolver{
				RestResolver: h.Clone().(*RESTResolver),
			}
		} else {
			target.Resolver = &Resolution_RestResolver{
				RestResolver: proto.Clone(m.GetRestResolver()).(*RESTResolver),
			}
		}

	}

	return target
}

// Clone function
func (m *GraphQLSchema) Clone() proto.Message {
	var target *GraphQLSchema
	if m == nil {
		return target
	}
	target = &GraphQLSchema{}

	if h, ok := interface{}(m.GetNamespacedStatuses()).(clone.Cloner); ok {
		target.NamespacedStatuses = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	} else {
		target.NamespacedStatuses = proto.Clone(m.GetNamespacedStatuses()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	target.Schema = m.GetSchema()

	target.EnableIntrospection = m.GetEnableIntrospection()

	if m.GetResolutions() != nil {
		target.Resolutions = make([]*Resolution, len(m.GetResolutions()))
		for idx, v := range m.GetResolutions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Resolutions[idx] = h.Clone().(*Resolution)
			} else {
				target.Resolutions[idx] = proto.Clone(v).(*Resolution)
			}

		}
	}

	return target
}

// Clone function
func (m *QueryMatcher_FieldMatcher) Clone() proto.Message {
	var target *QueryMatcher_FieldMatcher
	if m == nil {
		return target
	}
	target = &QueryMatcher_FieldMatcher{}

	target.Type = m.GetType()

	target.Field = m.GetField()

	return target
}
