// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/static/static.proto

package static

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options"

	google_golang_org_protobuf_types_known_structpb "google.golang.org/protobuf/types/known/structpb"

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
func (m *UpstreamSpec) Clone() proto.Message {
	var target *UpstreamSpec
	if m == nil {
		return target
	}
	target = &UpstreamSpec{}

	if m.GetHosts() != nil {
		target.Hosts = make([]*Host, len(m.GetHosts()))
		for idx, v := range m.GetHosts() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Hosts[idx] = h.Clone().(*Host)
			} else {
				target.Hosts[idx] = proto.Clone(v).(*Host)
			}

		}
	}

	if h, ok := interface{}(m.GetUseTls()).(clone.Cloner); ok {
		target.UseTls = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.UseTls = proto.Clone(m.GetUseTls()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetServiceSpec()).(clone.Cloner); ok {
		target.ServiceSpec = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options.ServiceSpec)
	} else {
		target.ServiceSpec = proto.Clone(m.GetServiceSpec()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options.ServiceSpec)
	}

	if h, ok := interface{}(m.GetAutoSniRewrite()).(clone.Cloner); ok {
		target.AutoSniRewrite = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.AutoSniRewrite = proto.Clone(m.GetAutoSniRewrite()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	return target
}

// Clone function
func (m *Host) Clone() proto.Message {
	var target *Host
	if m == nil {
		return target
	}
	target = &Host{}

	target.Addr = m.GetAddr()

	target.Port = m.GetPort()

	target.SniAddr = m.GetSniAddr()

	if h, ok := interface{}(m.GetLoadBalancingWeight()).(clone.Cloner); ok {
		target.LoadBalancingWeight = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.LoadBalancingWeight = proto.Clone(m.GetLoadBalancingWeight()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetHealthCheckConfig()).(clone.Cloner); ok {
		target.HealthCheckConfig = h.Clone().(*Host_HealthCheckConfig)
	} else {
		target.HealthCheckConfig = proto.Clone(m.GetHealthCheckConfig()).(*Host_HealthCheckConfig)
	}

	if m.GetMetadata() != nil {
		target.Metadata = make(map[string]*google_golang_org_protobuf_types_known_structpb.Struct, len(m.GetMetadata()))
		for k, v := range m.GetMetadata() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Metadata[k] = h.Clone().(*google_golang_org_protobuf_types_known_structpb.Struct)
			} else {
				target.Metadata[k] = proto.Clone(v).(*google_golang_org_protobuf_types_known_structpb.Struct)
			}

		}
	}

	return target
}

// Clone function
func (m *Host_HealthCheckConfig) Clone() proto.Message {
	var target *Host_HealthCheckConfig
	if m == nil {
		return target
	}
	target = &Host_HealthCheckConfig{}

	target.Path = m.GetPath()

	target.Method = m.GetMethod()

	return target
}
