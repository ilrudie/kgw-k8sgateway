// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/tracing/tracing.proto

package tracing

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
func (m *ListenerTracingSettings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListenerTracingSettings)
	if !ok {
		that2, ok := that.(ListenerTracingSettings)
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

	if len(m.GetRequestHeadersForTags()) != len(target.GetRequestHeadersForTags()) {
		return false
	}
	for idx, v := range m.GetRequestHeadersForTags() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequestHeadersForTags()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRequestHeadersForTags()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetVerbose()).(equality.Equalizer); ok {
		if !h.Equal(target.GetVerbose()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetVerbose(), target.GetVerbose()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTracePercentages()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTracePercentages()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTracePercentages(), target.GetTracePercentages()) {
			return false
		}
	}

	if len(m.GetEnvironmentVariablesForTags()) != len(target.GetEnvironmentVariablesForTags()) {
		return false
	}
	for idx, v := range m.GetEnvironmentVariablesForTags() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetEnvironmentVariablesForTags()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetEnvironmentVariablesForTags()[idx]) {
				return false
			}
		}

	}

	if len(m.GetLiteralsForTags()) != len(target.GetLiteralsForTags()) {
		return false
	}
	for idx, v := range m.GetLiteralsForTags() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetLiteralsForTags()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetLiteralsForTags()[idx]) {
				return false
			}
		}

	}

	switch m.ProviderConfig.(type) {

	case *ListenerTracingSettings_ZipkinConfig:
		if _, ok := target.ProviderConfig.(*ListenerTracingSettings_ZipkinConfig); !ok {
			return false
		}

		if h, ok := interface{}(m.GetZipkinConfig()).(equality.Equalizer); ok {
			if !h.Equal(target.GetZipkinConfig()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetZipkinConfig(), target.GetZipkinConfig()) {
				return false
			}
		}

	case *ListenerTracingSettings_DatadogConfig:
		if _, ok := target.ProviderConfig.(*ListenerTracingSettings_DatadogConfig); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDatadogConfig()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDatadogConfig()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDatadogConfig(), target.GetDatadogConfig()) {
				return false
			}
		}

	case *ListenerTracingSettings_OpenTelemetryConfig:
		if _, ok := target.ProviderConfig.(*ListenerTracingSettings_OpenTelemetryConfig); !ok {
			return false
		}

		if h, ok := interface{}(m.GetOpenTelemetryConfig()).(equality.Equalizer); ok {
			if !h.Equal(target.GetOpenTelemetryConfig()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetOpenTelemetryConfig(), target.GetOpenTelemetryConfig()) {
				return false
			}
		}

	case *ListenerTracingSettings_OpenCensusConfig:
		if _, ok := target.ProviderConfig.(*ListenerTracingSettings_OpenCensusConfig); !ok {
			return false
		}

		if h, ok := interface{}(m.GetOpenCensusConfig()).(equality.Equalizer); ok {
			if !h.Equal(target.GetOpenCensusConfig()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetOpenCensusConfig(), target.GetOpenCensusConfig()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ProviderConfig != target.ProviderConfig {
			return false
		}
	}

	return true
}

// Equal function
func (m *RouteTracingSettings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteTracingSettings)
	if !ok {
		that2, ok := that.(RouteTracingSettings)
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

	if strings.Compare(m.GetRouteDescriptor(), target.GetRouteDescriptor()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetTracePercentages()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTracePercentages()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTracePercentages(), target.GetTracePercentages()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPropagate()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPropagate()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPropagate(), target.GetPropagate()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *TracePercentages) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TracePercentages)
	if !ok {
		that2, ok := that.(TracePercentages)
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

	if h, ok := interface{}(m.GetClientSamplePercentage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetClientSamplePercentage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetClientSamplePercentage(), target.GetClientSamplePercentage()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRandomSamplePercentage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRandomSamplePercentage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRandomSamplePercentage(), target.GetRandomSamplePercentage()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetOverallSamplePercentage()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOverallSamplePercentage()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOverallSamplePercentage(), target.GetOverallSamplePercentage()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *TracingTagEnvironmentVariable) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TracingTagEnvironmentVariable)
	if !ok {
		that2, ok := that.(TracingTagEnvironmentVariable)
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

	if h, ok := interface{}(m.GetTag()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTag()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTag(), target.GetTag()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetName()).(equality.Equalizer); ok {
		if !h.Equal(target.GetName()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetName(), target.GetName()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDefaultValue()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDefaultValue()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDefaultValue(), target.GetDefaultValue()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *TracingTagLiteral) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TracingTagLiteral)
	if !ok {
		that2, ok := that.(TracingTagLiteral)
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

	if h, ok := interface{}(m.GetTag()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTag()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTag(), target.GetTag()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetValue()).(equality.Equalizer); ok {
		if !h.Equal(target.GetValue()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetValue(), target.GetValue()) {
			return false
		}
	}

	return true
}