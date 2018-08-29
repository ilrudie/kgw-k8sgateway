// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: endpoint.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core_solo_io "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//
// @solo-kit:resource
// @solo-kit:resource.short_name=ep
// @solo-kit:resource.plural_name=Endpoints
// @solo-kit:resource.group_name=gloo.solo.io
// @solo-kit:resource.version=v1
//
// Endpoints represent dynamically discovered address/ports where an upstream service is listening
type Endpoint struct {
	// Name of the upstream the endpoint belongs to
	UpstreamName string `protobuf:"bytes,1,opt,name=upstream_name,json=upstreamName,proto3" json:"upstream_name,omitempty"`
	// Address of the endpoint (ip or hostname)
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// listening port for the endpoint
	Port uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// Metadata contains the object metadata for this resource
	Metadata core_solo_io.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptorEndpoint, []int{0} }

func (m *Endpoint) GetUpstreamName() string {
	if m != nil {
		return m.UpstreamName
	}
	return ""
}

func (m *Endpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Endpoint) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Endpoint) GetMetadata() core_solo_io.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core_solo_io.Metadata{}
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "gloo.solo.io.Endpoint")
}
func (this *Endpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Endpoint)
	if !ok {
		that2, ok := that.(Endpoint)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.UpstreamName != that1.UpstreamName {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("endpoint.proto", fileDescriptorEndpoint) }

var fileDescriptorEndpoint = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcd, 0x4b, 0x29,
	0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0xcf, 0xc9, 0xcf,
	0xd7, 0x2b, 0xce, 0xcf, 0xc9, 0xd7, 0xcb, 0xcc, 0x97, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x4b,
	0xe8, 0x83, 0x58, 0x10, 0x35, 0x52, 0x86, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9,
	0xb9, 0xfa, 0x20, 0x95, 0xba, 0x99, 0xf9, 0x10, 0x3a, 0x3b, 0xb3, 0x44, 0x3f, 0xb1, 0x20, 0x53,
	0xbf, 0xcc, 0x50, 0x3f, 0x37, 0xb5, 0x24, 0x31, 0x25, 0xb1, 0x24, 0x11, 0xa2, 0x45, 0x69, 0x3a,
	0x23, 0x17, 0x87, 0x2b, 0xd4, 0x26, 0x21, 0x65, 0x2e, 0xde, 0xd2, 0x82, 0xe2, 0x92, 0xa2, 0xd4,
	0xc4, 0xdc, 0xf8, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x1e, 0x98,
	0xa0, 0x5f, 0x62, 0x6e, 0xaa, 0x90, 0x04, 0x17, 0x7b, 0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1,
	0x04, 0x13, 0x58, 0x1a, 0xc6, 0x15, 0x12, 0xe2, 0x62, 0x29, 0xc8, 0x2f, 0x2a, 0x91, 0x60, 0x56,
	0x60, 0xd4, 0xe0, 0x0d, 0x02, 0xb3, 0x85, 0x2c, 0xb8, 0x38, 0x60, 0x36, 0x4a, 0xb0, 0x2b, 0x30,
	0x6a, 0x70, 0x1b, 0x89, 0xe9, 0x25, 0xe7, 0x17, 0xa5, 0xc2, 0x7c, 0xa2, 0xe7, 0x0b, 0x95, 0x75,
	0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08, 0xae, 0xda, 0xc9, 0x6a, 0xc5, 0x23, 0x39, 0xc6, 0x28,
	0x13, 0x7c, 0x5e, 0x2a, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0x07, 0x85, 0x8d, 0x7e,
	0x41, 0x76, 0x3a, 0xd4, 0x93, 0x49, 0x6c, 0x60, 0xcf, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x49, 0xa0, 0xff, 0x9f, 0x45, 0x01, 0x00, 0x00,
}
