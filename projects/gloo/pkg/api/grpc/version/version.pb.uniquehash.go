// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/grpc/version/version.proto

package version

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"

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
	_ = strconv.Itoa
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ServerVersion) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/version.ServerVersion")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Type")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetType())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Enterprise")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetEnterprise())
	if err != nil {
		return 0, err
	}

	switch m.VersionType.(type) {

	case *ServerVersion_Kubernetes:

		if h, ok := interface{}(m.GetKubernetes()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Kubernetes")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetKubernetes(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Kubernetes")); err != nil {
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

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Kubernetes) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/version.Kubernetes")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Containers")); err != nil {
		return 0, err
	}
	for i, v := range m.GetContainers() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if _, err = hasher.Write([]byte("Namespace")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetNamespace())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ClientVersion) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/version.ClientVersion")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Version")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetVersion())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *KubernetesClusterVersion) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/version.KubernetesClusterVersion")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Major")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetMajor())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Minor")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetMinor())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("GitVersion")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetGitVersion())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("BuildDate")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetBuildDate())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Platform")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetPlatform())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Version) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/version.Version")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetClient()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Client")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetClient(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Client")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Server")); err != nil {
		return 0, err
	}
	for i, v := range m.GetServer() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetKubernetesCluster()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("KubernetesCluster")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetKubernetesCluster(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("KubernetesCluster")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Kubernetes_Container) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/grpc/version.Kubernetes_Container")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Tag")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetTag())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Name")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Registry")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetRegistry())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("OssTag")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetOssTag())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
