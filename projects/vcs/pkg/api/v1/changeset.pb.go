// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: changeset.proto

package v1 // import "github.com/solo-io/solo-kit/projects/vcs/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
import v1 "github.com/solo-io/solo-kit/projects/gateway/pkg/api/v1"
import v11 "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
import v12 "github.com/solo-io/solo-kit/projects/sqoop/pkg/api/v1"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
// @solo-kit:resource.short_name=chg
// @solo-kit:resource.plural_name=changesets
// @solo-kit:resource.resource_groups=api.vcs.solo.io
//
// The ChangeSet object represents the current status of a Gloo user's working directory. Each element in the "data"
// element represents the complete snapshot of a resource.
//
type ChangeSet struct {
	// Status indicates the validation status of this resource
	Status core.Status `protobuf:"bytes,1,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,2,opt,name=metadata" json:"metadata"`
	// The name of the git branch the changes will be applied to
	Branch types.StringValue `protobuf:"bytes,3,opt,name=branch" json:"branch"`
	// Indicates whether this changeset has been submitted for commit and is waiting to be pushed
	CommitPending types.BoolValue `protobuf:"bytes,4,opt,name=commit_pending,json=commitPending" json:"commit_pending"`
	// Description of the changeset. This will be the git commit message
	Description types.StringValue `protobuf:"bytes,5,opt,name=description" json:"description"`
	// The number of edits that the user applied to the previous commit.
	// A value greater than zero represents a dirty work tree.
	EditCount types.UInt32Value `protobuf:"bytes,6,opt,name=edit_count,json=editCount" json:"edit_count"`
	// The user who owns this changeset
	// TODO use dedicated message? Also, determine how to handle secrets?
	UserId types.StringValue `protobuf:"bytes,7,opt,name=user_id,json=userId" json:"user_id"`
	// The hash of the commit that the changeset represents an increment upon
	RootCommit types.StringValue `protobuf:"bytes,8,opt,name=root_commit,json=rootCommit" json:"root_commit"`
	// The git commit message for the root commit
	RootDescription types.StringValue `protobuf:"bytes,9,opt,name=root_description,json=rootDescription" json:"root_description"`
	// If a git commit attempt fails, this field will be populated with a user-friendly error message
	// No further git commit attempts will be possible until the user clears this field
	ErrorMsg *types.StringValue `protobuf:"bytes,10,opt,name=error_msg,json=errorMsg" json:"error_msg,omitempty"`
	// A collection of Gloo resources
	Data                 Data     `protobuf:"bytes,11,opt,name=data" json:"data"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeSet) Reset()         { *m = ChangeSet{} }
func (m *ChangeSet) String() string { return proto.CompactTextString(m) }
func (*ChangeSet) ProtoMessage()    {}
func (*ChangeSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_changeset_eb2a83324ead99cc, []int{0}
}
func (m *ChangeSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeSet.Unmarshal(m, b)
}
func (m *ChangeSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeSet.Marshal(b, m, deterministic)
}
func (dst *ChangeSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeSet.Merge(dst, src)
}
func (m *ChangeSet) XXX_Size() int {
	return xxx_messageInfo_ChangeSet.Size(m)
}
func (m *ChangeSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeSet.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeSet proto.InternalMessageInfo

func (m *ChangeSet) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *ChangeSet) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *ChangeSet) GetBranch() types.StringValue {
	if m != nil {
		return m.Branch
	}
	return types.StringValue{}
}

func (m *ChangeSet) GetCommitPending() types.BoolValue {
	if m != nil {
		return m.CommitPending
	}
	return types.BoolValue{}
}

func (m *ChangeSet) GetDescription() types.StringValue {
	if m != nil {
		return m.Description
	}
	return types.StringValue{}
}

func (m *ChangeSet) GetEditCount() types.UInt32Value {
	if m != nil {
		return m.EditCount
	}
	return types.UInt32Value{}
}

func (m *ChangeSet) GetUserId() types.StringValue {
	if m != nil {
		return m.UserId
	}
	return types.StringValue{}
}

func (m *ChangeSet) GetRootCommit() types.StringValue {
	if m != nil {
		return m.RootCommit
	}
	return types.StringValue{}
}

func (m *ChangeSet) GetRootDescription() types.StringValue {
	if m != nil {
		return m.RootDescription
	}
	return types.StringValue{}
}

func (m *ChangeSet) GetErrorMsg() *types.StringValue {
	if m != nil {
		return m.ErrorMsg
	}
	return nil
}

func (m *ChangeSet) GetData() Data {
	if m != nil {
		return m.Data
	}
	return Data{}
}

// A user-specific snapshot of all gloo resources at a given commit plus any non-committed changes made by the user
type Data struct {
	Gateways             []*v1.Gateway        `protobuf:"bytes,1,rep,name=gateways" json:"gateways,omitempty"`
	VirtualServices      []*v1.VirtualService `protobuf:"bytes,2,rep,name=virtual_services,json=virtualServices" json:"virtual_services,omitempty"`
	Proxies              []*v11.Proxy         `protobuf:"bytes,3,rep,name=proxies" json:"proxies,omitempty"`
	Settings             []*v11.Settings      `protobuf:"bytes,4,rep,name=settings" json:"settings,omitempty"`
	Upstreams            []*v11.Upstream      `protobuf:"bytes,5,rep,name=upstreams" json:"upstreams,omitempty"`
	ResolverMaps         []*v12.ResolverMap   `protobuf:"bytes,6,rep,name=resolver_maps,json=resolverMaps" json:"resolver_maps,omitempty"`
	Schemas              []*v12.Schema        `protobuf:"bytes,7,rep,name=schemas" json:"schemas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_changeset_eb2a83324ead99cc, []int{1}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (dst *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(dst, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetGateways() []*v1.Gateway {
	if m != nil {
		return m.Gateways
	}
	return nil
}

func (m *Data) GetVirtualServices() []*v1.VirtualService {
	if m != nil {
		return m.VirtualServices
	}
	return nil
}

func (m *Data) GetProxies() []*v11.Proxy {
	if m != nil {
		return m.Proxies
	}
	return nil
}

func (m *Data) GetSettings() []*v11.Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *Data) GetUpstreams() []*v11.Upstream {
	if m != nil {
		return m.Upstreams
	}
	return nil
}

func (m *Data) GetResolverMaps() []*v12.ResolverMap {
	if m != nil {
		return m.ResolverMaps
	}
	return nil
}

func (m *Data) GetSchemas() []*v12.Schema {
	if m != nil {
		return m.Schemas
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeSet)(nil), "vcs.solo.io.ChangeSet")
	proto.RegisterType((*Data)(nil), "vcs.solo.io.Data")
}
func (this *ChangeSet) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ChangeSet)
	if !ok {
		that2, ok := that.(ChangeSet)
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
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.Branch.Equal(&that1.Branch) {
		return false
	}
	if !this.CommitPending.Equal(&that1.CommitPending) {
		return false
	}
	if !this.Description.Equal(&that1.Description) {
		return false
	}
	if !this.EditCount.Equal(&that1.EditCount) {
		return false
	}
	if !this.UserId.Equal(&that1.UserId) {
		return false
	}
	if !this.RootCommit.Equal(&that1.RootCommit) {
		return false
	}
	if !this.RootDescription.Equal(&that1.RootDescription) {
		return false
	}
	if !this.ErrorMsg.Equal(that1.ErrorMsg) {
		return false
	}
	if !this.Data.Equal(&that1.Data) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Data) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Data)
	if !ok {
		that2, ok := that.(Data)
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
	if len(this.Gateways) != len(that1.Gateways) {
		return false
	}
	for i := range this.Gateways {
		if !this.Gateways[i].Equal(that1.Gateways[i]) {
			return false
		}
	}
	if len(this.VirtualServices) != len(that1.VirtualServices) {
		return false
	}
	for i := range this.VirtualServices {
		if !this.VirtualServices[i].Equal(that1.VirtualServices[i]) {
			return false
		}
	}
	if len(this.Proxies) != len(that1.Proxies) {
		return false
	}
	for i := range this.Proxies {
		if !this.Proxies[i].Equal(that1.Proxies[i]) {
			return false
		}
	}
	if len(this.Settings) != len(that1.Settings) {
		return false
	}
	for i := range this.Settings {
		if !this.Settings[i].Equal(that1.Settings[i]) {
			return false
		}
	}
	if len(this.Upstreams) != len(that1.Upstreams) {
		return false
	}
	for i := range this.Upstreams {
		if !this.Upstreams[i].Equal(that1.Upstreams[i]) {
			return false
		}
	}
	if len(this.ResolverMaps) != len(that1.ResolverMaps) {
		return false
	}
	for i := range this.ResolverMaps {
		if !this.ResolverMaps[i].Equal(that1.ResolverMaps[i]) {
			return false
		}
	}
	if len(this.Schemas) != len(that1.Schemas) {
		return false
	}
	for i := range this.Schemas {
		if !this.Schemas[i].Equal(that1.Schemas[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("changeset.proto", fileDescriptor_changeset_eb2a83324ead99cc) }

var fileDescriptor_changeset_eb2a83324ead99cc = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcd, 0x4e, 0x1b, 0x3b,
	0x14, 0xc7, 0x2f, 0x24, 0xe4, 0xc3, 0xe1, 0xd3, 0x97, 0x8b, 0x7c, 0xa3, 0x2b, 0xb8, 0xca, 0x0a,
	0xa9, 0xc2, 0x23, 0x02, 0x8b, 0x42, 0x17, 0x55, 0x13, 0x24, 0x44, 0xa5, 0x48, 0x68, 0x22, 0x58,
	0x74, 0x13, 0x39, 0x13, 0xe3, 0xb8, 0x64, 0xe6, 0x4c, 0x6d, 0x4f, 0x68, 0x9f, 0xa3, 0x2f, 0xd1,
	0x47, 0xe9, 0x0b, 0x74, 0xcb, 0xa2, 0x8f, 0xd0, 0x27, 0xa8, 0xec, 0xf1, 0x84, 0x50, 0xfa, 0x91,
	0x55, 0x72, 0x7c, 0xfe, 0xbf, 0xbf, 0x67, 0xce, 0x9c, 0x73, 0xd0, 0x46, 0x34, 0x66, 0x89, 0xe0,
	0x9a, 0x1b, 0x9a, 0x2a, 0x30, 0x80, 0x1b, 0xd3, 0x48, 0x53, 0x0d, 0x13, 0xa0, 0x12, 0x9a, 0xdb,
	0x02, 0x04, 0xb8, 0xf3, 0xc0, 0xfe, 0xcb, 0x25, 0xcd, 0x5d, 0x01, 0x20, 0x26, 0x3c, 0x70, 0xd1,
	0x30, 0xbb, 0x09, 0xee, 0x14, 0x4b, 0x53, 0xae, 0xb4, 0xcf, 0x1f, 0x0a, 0x69, 0xc6, 0xd9, 0x90,
	0x46, 0x10, 0x07, 0xd6, 0xe9, 0x40, 0x42, 0xfe, 0x7b, 0x2b, 0x4d, 0xc0, 0x52, 0x19, 0x4c, 0x0f,
	0x83, 0x98, 0x1b, 0x36, 0x62, 0x86, 0x79, 0x24, 0x58, 0x00, 0xd1, 0x86, 0x99, 0xac, 0xb8, 0x63,
	0x4d, 0x30, 0xc3, 0xef, 0xd8, 0x07, 0x1f, 0x36, 0x52, 0x05, 0xef, 0x8b, 0x60, 0x5d, 0x73, 0x63,
	0x64, 0x22, 0x0a, 0xed, 0x7a, 0x96, 0x6a, 0xa3, 0x38, 0x8b, 0x7d, 0x8c, 0x15, 0xd7, 0x30, 0x99,
	0x72, 0x35, 0x88, 0x59, 0xea, 0xcf, 0x56, 0x75, 0x34, 0xe6, 0xb1, 0x7f, 0x9c, 0xd6, 0x97, 0x15,
	0x54, 0xef, 0xba, 0xc2, 0xf4, 0xb9, 0xc1, 0xe7, 0xa8, 0x92, 0xdf, 0x4d, 0x96, 0xfe, 0x5f, 0xda,
	0x6f, 0xb4, 0xb7, 0x69, 0x04, 0x8a, 0x17, 0x45, 0xa2, 0x7d, 0x97, 0xeb, 0xfc, 0xfb, 0xf9, 0x7e,
	0xef, 0xaf, 0x6f, 0xf7, 0x7b, 0x5b, 0x86, 0x6b, 0x33, 0x92, 0x37, 0x37, 0xa7, 0x2d, 0x29, 0x12,
	0x50, 0xbc, 0x15, 0x7a, 0x1c, 0x3f, 0x47, 0xb5, 0xe2, 0xbd, 0xc9, 0xb2, 0xb3, 0xda, 0x79, 0x6c,
	0xd5, 0xf3, 0xd9, 0x4e, 0xd9, 0x9a, 0x85, 0x33, 0x35, 0x3e, 0x45, 0x95, 0xa1, 0x62, 0x49, 0x34,
	0x26, 0x25, 0xc7, 0xfd, 0x47, 0xf3, 0x6f, 0x40, 0x8b, 0x6f, 0x40, 0xfb, 0x46, 0xc9, 0x44, 0x5c,
	0xb3, 0x49, 0xc6, 0x3d, 0xed, 0x09, 0x7c, 0x8e, 0xd6, 0x23, 0x88, 0x63, 0x69, 0x06, 0x29, 0x4f,
	0x46, 0x32, 0x11, 0xa4, 0xec, 0x3c, 0x9a, 0x4f, 0x3c, 0x3a, 0x00, 0x93, 0x79, 0x87, 0xb5, 0x9c,
	0xbb, 0xcc, 0x31, 0x7c, 0x86, 0x1a, 0x23, 0xae, 0x23, 0x25, 0x53, 0x23, 0x21, 0x21, 0x2b, 0x0b,
	0x3f, 0xc9, 0x3c, 0x86, 0x5f, 0x21, 0xc4, 0x47, 0xd2, 0x0c, 0x22, 0xc8, 0x12, 0x43, 0x2a, 0xbf,
	0x30, 0xb9, 0xba, 0x48, 0xcc, 0x51, 0x7b, 0xde, 0xa4, 0x6e, 0xa9, 0xae, 0x85, 0xf0, 0x0b, 0x54,
	0xcd, 0x34, 0x57, 0x03, 0x39, 0x22, 0xd5, 0xc5, 0xcb, 0x61, 0x91, 0x8b, 0x11, 0xee, 0xa2, 0x86,
	0x02, 0xb0, 0xf7, 0xdb, 0x77, 0x23, 0xb5, 0x85, 0x0d, 0x90, 0xc5, 0xba, 0x8e, 0xc2, 0x3d, 0xb4,
	0xe9, 0x4c, 0xe6, 0xeb, 0x51, 0x5f, 0xd8, 0x69, 0xc3, 0xb2, 0x67, 0x73, 0x35, 0x39, 0x41, 0x75,
	0xae, 0x14, 0xa8, 0x41, 0xac, 0x05, 0x41, 0x7f, 0xf6, 0x09, 0x6b, 0x4e, 0xde, 0xd3, 0x02, 0x3f,
	0x43, 0x65, 0xd7, 0x4f, 0x0d, 0x47, 0x6d, 0xd1, 0xb9, 0xf1, 0xa5, 0x67, 0x0f, 0xad, 0xe4, 0x44,
	0xad, 0x8f, 0x25, 0x54, 0xb6, 0x87, 0xf8, 0x18, 0xd5, 0xfc, 0x00, 0xd9, 0xa6, 0x2e, 0xed, 0x37,
	0xda, 0x84, 0x16, 0x13, 0x55, 0xd0, 0xe7, 0x79, 0x1c, 0xce, 0x94, 0xf8, 0x35, 0xda, 0x9c, 0x4a,
	0x65, 0x32, 0x36, 0x19, 0x68, 0xae, 0xa6, 0x32, 0xe2, 0x9a, 0x2c, 0x3b, 0x7a, 0xef, 0x09, 0x7d,
	0x9d, 0x0b, 0xfb, 0xb9, 0x2e, 0xdc, 0x98, 0x3e, 0x8a, 0x35, 0x3e, 0x40, 0x55, 0x3b, 0xb3, 0x92,
	0x6b, 0x52, 0x72, 0x16, 0x7f, 0x53, 0x31, 0x01, 0x98, 0xf1, 0x97, 0x76, 0xa0, 0xc3, 0x42, 0x83,
	0xdb, 0xa8, 0x56, 0x4c, 0x35, 0x29, 0x3b, 0xfd, 0xce, 0x63, 0x7d, 0xdf, 0x67, 0xc3, 0x99, 0x0e,
	0x1f, 0xa3, 0x7a, 0x31, 0xf9, 0x9a, 0xac, 0xfc, 0x0c, 0xba, 0xf2, 0xe9, 0xf0, 0x41, 0x88, 0x5f,
	0xa2, 0xb5, 0xf9, 0xfd, 0xa0, 0x49, 0xc5, 0x91, 0x4d, 0xaa, 0xdf, 0x01, 0xa4, 0x33, 0x34, 0xf4,
	0x9a, 0x1e, 0x4b, 0xc3, 0x55, 0xf5, 0x10, 0x68, 0x1c, 0xa0, 0x6a, 0xbe, 0x4c, 0x34, 0xa9, 0x3a,
	0xf4, 0x9f, 0x1f, 0xd0, 0xbe, 0xcb, 0x86, 0x85, 0xaa, 0x73, 0xf2, 0xe9, 0xeb, 0xee, 0xd2, 0x9b,
	0xa3, 0xdf, 0xad, 0xc0, 0x54, 0xc1, 0x5b, 0x1e, 0x19, 0x1d, 0x4c, 0x23, 0x1d, 0xa4, 0xb7, 0xc2,
	0xef, 0xc4, 0x61, 0xc5, 0x75, 0xc7, 0xd1, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x09, 0xff,
	0x57, 0xc7, 0x05, 0x00, 0x00,
}
