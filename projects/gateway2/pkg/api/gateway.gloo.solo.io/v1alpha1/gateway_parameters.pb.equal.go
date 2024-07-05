// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway2/api/v1alpha1/gateway_parameters.proto

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
func (m *GatewayParametersSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayParametersSpec)
	if !ok {
		that2, ok := that.(GatewayParametersSpec)
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

	switch m.EnvironmentType.(type) {

	case *GatewayParametersSpec_Kube:
		if _, ok := target.EnvironmentType.(*GatewayParametersSpec_Kube); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKube()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKube()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKube(), target.GetKube()) {
				return false
			}
		}

	case *GatewayParametersSpec_SelfManaged:
		if _, ok := target.EnvironmentType.(*GatewayParametersSpec_SelfManaged); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSelfManaged()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSelfManaged()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSelfManaged(), target.GetSelfManaged()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.EnvironmentType != target.EnvironmentType {
			return false
		}
	}

	return true
}

// Equal function
func (m *KubernetesProxyConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*KubernetesProxyConfig)
	if !ok {
		that2, ok := that.(KubernetesProxyConfig)
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

	if h, ok := interface{}(m.GetEnvoyContainer()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEnvoyContainer()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEnvoyContainer(), target.GetEnvoyContainer()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSdsContainer()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSdsContainer()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSdsContainer(), target.GetSdsContainer()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPodTemplate()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPodTemplate()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPodTemplate(), target.GetPodTemplate()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetService()).(equality.Equalizer); ok {
		if !h.Equal(target.GetService()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetService(), target.GetService()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIstio()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIstio()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIstio(), target.GetIstio()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStats()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStats()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStats(), target.GetStats()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAiExtension()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAiExtension()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAiExtension(), target.GetAiExtension()) {
			return false
		}
	}

	switch m.WorkloadType.(type) {

	case *KubernetesProxyConfig_Deployment:
		if _, ok := target.WorkloadType.(*KubernetesProxyConfig_Deployment); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDeployment()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDeployment()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDeployment(), target.GetDeployment()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.WorkloadType != target.WorkloadType {
			return false
		}
	}

	return true
}

// Equal function
func (m *ProxyDeployment) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ProxyDeployment)
	if !ok {
		that2, ok := that.(ProxyDeployment)
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

	if h, ok := interface{}(m.GetReplicas()).(equality.Equalizer); ok {
		if !h.Equal(target.GetReplicas()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetReplicas(), target.GetReplicas()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *EnvoyContainer) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*EnvoyContainer)
	if !ok {
		that2, ok := that.(EnvoyContainer)
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

	if h, ok := interface{}(m.GetBootstrap()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBootstrap()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBootstrap(), target.GetBootstrap()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetImage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetImage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetImage(), target.GetImage()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSecurityContext()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSecurityContext()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSecurityContext(), target.GetSecurityContext()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResources()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResources()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResources(), target.GetResources()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *EnvoyBootstrap) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*EnvoyBootstrap)
	if !ok {
		that2, ok := that.(EnvoyBootstrap)
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

	if h, ok := interface{}(m.GetLogLevel()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLogLevel()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLogLevel(), target.GetLogLevel()) {
			return false
		}
	}

	if len(m.GetComponentLogLevels()) != len(target.GetComponentLogLevels()) {
		return false
	}
	for k, v := range m.GetComponentLogLevels() {

		if strings.Compare(v, target.GetComponentLogLevels()[k]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *IstioIntegration) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IstioIntegration)
	if !ok {
		that2, ok := that.(IstioIntegration)
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

	if h, ok := interface{}(m.GetIstioProxyContainer()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIstioProxyContainer()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIstioProxyContainer(), target.GetIstioProxyContainer()) {
			return false
		}
	}

	if len(m.GetCustomSidecars()) != len(target.GetCustomSidecars()) {
		return false
	}
	for idx, v := range m.GetCustomSidecars() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetCustomSidecars()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetCustomSidecars()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *SdsContainer) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SdsContainer)
	if !ok {
		that2, ok := that.(SdsContainer)
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

	if h, ok := interface{}(m.GetImage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetImage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetImage(), target.GetImage()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSecurityContext()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSecurityContext()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSecurityContext(), target.GetSecurityContext()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResources()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResources()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResources(), target.GetResources()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetBootstrap()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBootstrap()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBootstrap(), target.GetBootstrap()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *SdsBootstrap) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*SdsBootstrap)
	if !ok {
		that2, ok := that.(SdsBootstrap)
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

	if h, ok := interface{}(m.GetLogLevel()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLogLevel()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLogLevel(), target.GetLogLevel()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *IstioContainer) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*IstioContainer)
	if !ok {
		that2, ok := that.(IstioContainer)
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

	if h, ok := interface{}(m.GetImage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetImage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetImage(), target.GetImage()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSecurityContext()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSecurityContext()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSecurityContext(), target.GetSecurityContext()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResources()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResources()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResources(), target.GetResources()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetLogLevel()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLogLevel()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLogLevel(), target.GetLogLevel()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIstioDiscoveryAddress()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIstioDiscoveryAddress()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIstioDiscoveryAddress(), target.GetIstioDiscoveryAddress()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIstioMetaMeshId()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIstioMetaMeshId()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIstioMetaMeshId(), target.GetIstioMetaMeshId()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIstioMetaClusterId()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIstioMetaClusterId()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIstioMetaClusterId(), target.GetIstioMetaClusterId()) {
			return false
		}
	}

	if len(m.GetComponentLogLevels()) != len(target.GetComponentLogLevels()) {
		return false
	}
	for k, v := range m.GetComponentLogLevels() {

		if strings.Compare(v, target.GetComponentLogLevels()[k]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *GatewayParametersStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GatewayParametersStatus)
	if !ok {
		that2, ok := that.(GatewayParametersStatus)
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

	return true
}

// Equal function
func (m *StatsConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*StatsConfig)
	if !ok {
		that2, ok := that.(StatsConfig)
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

	if h, ok := interface{}(m.GetEnabled()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEnabled()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEnabled(), target.GetEnabled()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRoutePrefixRewrite()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRoutePrefixRewrite()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRoutePrefixRewrite(), target.GetRoutePrefixRewrite()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetEnableStatsRoute()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEnableStatsRoute()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEnableStatsRoute(), target.GetEnableStatsRoute()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatsRoutePrefixRewrite()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatsRoutePrefixRewrite()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatsRoutePrefixRewrite(), target.GetStatsRoutePrefixRewrite()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *AiExtension) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AiExtension)
	if !ok {
		that2, ok := that.(AiExtension)
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

	if h, ok := interface{}(m.GetEnabled()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEnabled()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEnabled(), target.GetEnabled()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetImage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetImage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetImage(), target.GetImage()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSecurityContext()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSecurityContext()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSecurityContext(), target.GetSecurityContext()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResources()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResources()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResources(), target.GetResources()) {
			return false
		}
	}

	if len(m.GetEnv()) != len(target.GetEnv()) {
		return false
	}
	for idx, v := range m.GetEnv() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetEnv()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetEnv()[idx]) {
				return false
			}
		}

	}

	if len(m.GetPorts()) != len(target.GetPorts()) {
		return false
	}
	for idx, v := range m.GetPorts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPorts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPorts()[idx]) {
				return false
			}
		}

	}

	return true
}
