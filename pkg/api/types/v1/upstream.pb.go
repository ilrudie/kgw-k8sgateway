// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: upstream.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import _ "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

type Upstream struct {
	Name              string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type              string                  `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ConnectionTimeout *time.Duration          `protobuf:"bytes,3,opt,name=connection_timeout,json=connectionTimeout,stdduration" json:"connection_timeout,omitempty"`
	Spec              *google_protobuf.Struct `protobuf:"bytes,4,opt,name=spec" json:"spec,omitempty"`
	Functions         []*Function             `protobuf:"bytes,5,rep,name=functions" json:"functions,omitempty"`
	// read only
	Status    *Status `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	*Metadata `protobuf:"bytes,7,opt,name=metadata,embedded=metadata" json:"metadata,omitempty"`
}

func (m *Upstream) Reset()                    { *m = Upstream{} }
func (m *Upstream) String() string            { return proto.CompactTextString(m) }
func (*Upstream) ProtoMessage()               {}
func (*Upstream) Descriptor() ([]byte, []int) { return fileDescriptorUpstream, []int{0} }

func (m *Upstream) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Upstream) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Upstream) GetConnectionTimeout() *time.Duration {
	if m != nil {
		return m.ConnectionTimeout
	}
	return nil
}

func (m *Upstream) GetSpec() *google_protobuf.Struct {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Upstream) GetFunctions() []*Function {
	if m != nil {
		return m.Functions
	}
	return nil
}

func (m *Upstream) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type Function struct {
	Name string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Spec *google_protobuf.Struct `protobuf:"bytes,4,opt,name=spec" json:"spec,omitempty"`
}

func (m *Function) Reset()                    { *m = Function{} }
func (m *Function) String() string            { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()               {}
func (*Function) Descriptor() ([]byte, []int) { return fileDescriptorUpstream, []int{1} }

func (m *Function) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Function) GetSpec() *google_protobuf.Struct {
	if m != nil {
		return m.Spec
	}
	return nil
}

func init() {
	proto.RegisterType((*Upstream)(nil), "v1.Upstream")
	proto.RegisterType((*Function)(nil), "v1.Function")
}
func (this *Upstream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream)
	if !ok {
		that2, ok := that.(Upstream)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.ConnectionTimeout != nil && that1.ConnectionTimeout != nil {
		if *this.ConnectionTimeout != *that1.ConnectionTimeout {
			return false
		}
	} else if this.ConnectionTimeout != nil {
		return false
	} else if that1.ConnectionTimeout != nil {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if len(this.Functions) != len(that1.Functions) {
		return false
	}
	for i := range this.Functions {
		if !this.Functions[i].Equal(that1.Functions[i]) {
			return false
		}
	}
	if !this.Status.Equal(that1.Status) {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	return true
}
func (this *Function) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Function)
	if !ok {
		that2, ok := that.(Function)
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
	if this.Name != that1.Name {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("upstream.proto", fileDescriptorUpstream) }

var fileDescriptorUpstream = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0xc1, 0x4a, 0x2b, 0x31,
	0x14, 0x25, 0x6d, 0x5e, 0x5f, 0x9b, 0x96, 0x82, 0x41, 0x30, 0x16, 0x69, 0x4b, 0x57, 0x45, 0x21,
	0xa5, 0xf5, 0x0f, 0x8a, 0xb8, 0x11, 0x5d, 0xa4, 0xba, 0x96, 0x74, 0x9a, 0x0e, 0x05, 0x67, 0x32,
	0x4c, 0x6e, 0x0a, 0xfe, 0x89, 0x9f, 0xe0, 0xce, 0x5f, 0xf1, 0x0b, 0x14, 0xfc, 0x12, 0x99, 0x9b,
	0xd4, 0x82, 0xba, 0x71, 0x77, 0x72, 0xce, 0xb9, 0x27, 0xe7, 0x5e, 0xd6, 0xf5, 0x85, 0x83, 0xd2,
	0xe8, 0x4c, 0x16, 0xa5, 0x05, 0xcb, 0x6b, 0xdb, 0x69, 0xef, 0x24, 0xb5, 0x36, 0x7d, 0x30, 0x13,
	0x64, 0x96, 0x7e, 0x3d, 0x71, 0x50, 0xfa, 0x04, 0x82, 0xa3, 0xd7, 0xff, 0xae, 0xae, 0x7c, 0xa9,
	0x61, 0x63, 0xf3, 0xa8, 0x1f, 0xa6, 0x36, 0xb5, 0x08, 0x27, 0x15, 0x8a, 0x6c, 0xc7, 0x81, 0x06,
	0xef, 0xe2, 0xab, 0x9b, 0x19, 0xd0, 0x2b, 0x0d, 0x3a, 0xbc, 0x47, 0x2f, 0x35, 0xd6, 0xbc, 0x8b,
	0x45, 0x38, 0x67, 0x34, 0xd7, 0x99, 0x11, 0x64, 0x48, 0xc6, 0x2d, 0x85, 0xb8, 0xe2, 0xe0, 0xb1,
	0x30, 0xa2, 0x16, 0xb8, 0x0a, 0xf3, 0x1b, 0xc6, 0x13, 0x9b, 0xe7, 0x26, 0xa9, 0x3e, 0xbf, 0x87,
	0x4d, 0x66, 0xac, 0x07, 0x51, 0x1f, 0x92, 0x71, 0x7b, 0x76, 0x2c, 0x43, 0x4b, 0xb9, 0x6b, 0x29,
	0x2f, 0x62, 0xcb, 0x39, 0x7d, 0x7a, 0x1f, 0x10, 0x75, 0xb0, 0x1f, 0xbd, 0x0d, 0x93, 0xfc, 0x8c,
	0x51, 0x57, 0x98, 0x44, 0x50, 0x4c, 0x38, 0xfa, 0x91, 0xb0, 0xc0, 0x2b, 0x28, 0x34, 0xf1, 0x53,
	0xd6, 0x5a, 0xfb, 0x1c, 0xe7, 0x9d, 0xf8, 0x37, 0xac, 0x8f, 0xdb, 0xb3, 0x8e, 0xdc, 0x4e, 0xe5,
	0x65, 0x24, 0xd5, 0x5e, 0xe6, 0x23, 0xd6, 0x08, 0xdb, 0x8b, 0x06, 0x46, 0xb3, 0xca, 0xb8, 0x40,
	0x46, 0x45, 0x85, 0x4b, 0xd6, 0xdc, 0xdd, 0x44, 0xfc, 0x47, 0x17, 0xc6, 0x5d, 0x47, 0x6e, 0x4e,
	0x5f, 0xdf, 0x06, 0x44, 0x7d, 0x79, 0x46, 0x57, 0xac, 0xb9, 0xfb, 0xea, 0xd7, 0x83, 0xfd, 0x65,
	0x99, 0x39, 0x7d, 0xfe, 0xe8, 0x93, 0x65, 0x03, 0xc5, 0xf3, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x3e, 0x5d, 0x40, 0x67, 0x13, 0x02, 0x00, 0x00,
}
