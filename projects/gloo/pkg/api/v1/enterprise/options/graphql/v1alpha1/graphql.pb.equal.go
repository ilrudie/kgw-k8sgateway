// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/graphql/v1alpha1/graphql.proto

package v1alpha1

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
func (m *RequestTemplate) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RequestTemplate)
	if !ok {
		that2, ok := that.(RequestTemplate)
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

	if len(m.GetHeaders()) != len(target.GetHeaders()) {
		return false
	}
	for k, v := range m.GetHeaders() {

		if strings.Compare(v, target.GetHeaders()[k]) != 0 {
			return false
		}

	}

	if len(m.GetQueryParams()) != len(target.GetQueryParams()) {
		return false
	}
	for k, v := range m.GetQueryParams() {

		if strings.Compare(v, target.GetQueryParams()[k]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetBody()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBody()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBody(), target.GetBody()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ResponseTemplate) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ResponseTemplate)
	if !ok {
		that2, ok := that.(ResponseTemplate)
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

	if strings.Compare(m.GetResultRoot(), target.GetResultRoot()) != 0 {
		return false
	}

	if len(m.GetSetters()) != len(target.GetSetters()) {
		return false
	}
	for k, v := range m.GetSetters() {

		if strings.Compare(v, target.GetSetters()[k]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *RESTResolver) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RESTResolver)
	if !ok {
		that2, ok := that.(RESTResolver)
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

	if h, ok := interface{}(m.GetRequest()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRequest()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRequest(), target.GetRequest()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResponse()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResponse()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResponse(), target.GetResponse()) {
			return false
		}
	}

	if strings.Compare(m.GetSpanName(), target.GetSpanName()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *QueryMatcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*QueryMatcher)
	if !ok {
		that2, ok := that.(QueryMatcher)
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

	switch m.Match.(type) {

	case *QueryMatcher_FieldMatcher_:
		if _, ok := target.Match.(*QueryMatcher_FieldMatcher_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetFieldMatcher()).(equality.Equalizer); ok {
			if !h.Equal(target.GetFieldMatcher()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetFieldMatcher(), target.GetFieldMatcher()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Match != target.Match {
			return false
		}
	}

	return true
}

// Equal function
func (m *Resolution) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Resolution)
	if !ok {
		that2, ok := that.(Resolution)
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

	if h, ok := interface{}(m.GetMatcher()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMatcher()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMatcher(), target.GetMatcher()) {
			return false
		}
	}

	switch m.Resolver.(type) {

	case *Resolution_RestResolver:
		if _, ok := target.Resolver.(*Resolution_RestResolver); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRestResolver()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRestResolver()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRestResolver(), target.GetRestResolver()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Resolver != target.Resolver {
			return false
		}
	}

	return true
}

// Equal function
func (m *GraphQLSchema) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GraphQLSchema)
	if !ok {
		that2, ok := that.(GraphQLSchema)
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

	if h, ok := interface{}(m.GetNamespacedStatuses()).(equality.Equalizer); ok {
		if !h.Equal(target.GetNamespacedStatuses()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetNamespacedStatuses(), target.GetNamespacedStatuses()) {
			return false
		}
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

	if strings.Compare(m.GetSchema(), target.GetSchema()) != 0 {
		return false
	}

	if m.GetEnableIntrospection() != target.GetEnableIntrospection() {
		return false
	}

	if len(m.GetResolutions()) != len(target.GetResolutions()) {
		return false
	}
	for idx, v := range m.GetResolutions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetResolutions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetResolutions()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *QueryMatcher_FieldMatcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*QueryMatcher_FieldMatcher)
	if !ok {
		that2, ok := that.(QueryMatcher_FieldMatcher)
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

	if strings.Compare(m.GetType(), target.GetType()) != 0 {
		return false
	}

	if strings.Compare(m.GetField(), target.GetField()) != 0 {
		return false
	}

	return true
}
