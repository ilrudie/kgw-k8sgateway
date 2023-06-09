// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/xff_offset/solo_xff_offset_filter.proto

package xff_offset

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Envoy filter configuration for the solo_xff_offset filter,  not set by the user.
type SoloXffOffset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The amount of addresses from the left of x-forwarded-for (xff) header
	// which should be skipped to get the client address.
	Offset uint32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *SoloXffOffset) Reset() {
	*x = SoloXffOffset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoloXffOffset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoloXffOffset) ProtoMessage() {}

func (x *SoloXffOffset) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoloXffOffset.ProtoReflect.Descriptor instead.
func (*SoloXffOffset) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_rawDescGZIP(), []int{0}
}

func (x *SoloXffOffset) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_rawDesc = []byte{
	0x0a, 0x6b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x78, 0x66, 0x66, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x5f, 0x78, 0x66, 0x66, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x5f, 0x78, 0x66, 0x66,
	0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x22, 0x27, 0x0a, 0x0d, 0x53, 0x6f,
	0x6c, 0x6f, 0x58, 0x66, 0x66, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x42, 0xa5, 0x01, 0x0a, 0x39, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x5f, 0x78, 0x66, 0x66, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x2e, 0x76,
	0x32, 0x42, 0x12, 0x53, 0x6f, 0x6c, 0x6f, 0x58, 0x66, 0x66, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x78, 0x66, 0x66, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_goTypes = []interface{}{
	(*SoloXffOffset)(nil), // 0: envoy.config.filter.http.solo_xff_offset.v2.SoloXffOffset
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoloXffOffset); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_xff_offset_solo_xff_offset_filter_proto_depIdxs = nil
}