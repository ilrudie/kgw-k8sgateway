// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/trace/v3/opencensus.proto

package v3

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *OpenCensusConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.trace.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/trace/v3.OpenCensusConfig")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTraceConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("TraceConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTraceConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("TraceConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetOcagentExporterEnabled())
	if err != nil {
		return 0, err
	}

	for _, v := range m.GetIncomingTraceContext() {

		err = binary.Write(hasher, binary.LittleEndian, v)
		if err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetOutgoingTraceContext() {

		err = binary.Write(hasher, binary.LittleEndian, v)
		if err != nil {
			return 0, err
		}

	}

	switch m.OcagentAddress.(type) {

	case *OpenCensusConfig_HttpAddress:

		if _, err = hasher.Write([]byte(m.GetHttpAddress())); err != nil {
			return 0, err
		}

	case *OpenCensusConfig_GrpcAddress:

		if h, ok := interface{}(m.GetGrpcAddress()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("GrpcAddress")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGrpcAddress(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("GrpcAddress")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *TraceConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.trace.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/trace/v3.TraceConfig")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetMaxNumberOfAttributes())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetMaxNumberOfAnnotations())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetMaxNumberOfMessageEvents())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetMaxNumberOfLinks())
	if err != nil {
		return 0, err
	}

	switch m.Sampler.(type) {

	case *TraceConfig_ProbabilitySampler:

		if h, ok := interface{}(m.GetProbabilitySampler()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ProbabilitySampler")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetProbabilitySampler(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ProbabilitySampler")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *TraceConfig_ConstantSampler:

		if h, ok := interface{}(m.GetConstantSampler()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ConstantSampler")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetConstantSampler(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ConstantSampler")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *TraceConfig_RateLimitingSampler:

		if h, ok := interface{}(m.GetRateLimitingSampler()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RateLimitingSampler")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRateLimitingSampler(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RateLimitingSampler")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *ProbabilitySampler) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.trace.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/trace/v3.ProbabilitySampler")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSamplingProbability())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *ConstantSampler) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.trace.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/trace/v3.ConstantSampler")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetDecision())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *RateLimitingSampler) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.trace.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/trace/v3.RateLimitingSampler")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetQps())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *OpenCensusConfig_OcagentGrpcAddress) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.trace.v3.github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/trace/v3.OpenCensusConfig_OcagentGrpcAddress")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetTargetUri())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetStatPrefix())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
