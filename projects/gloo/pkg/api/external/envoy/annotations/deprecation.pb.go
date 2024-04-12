// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/annotations/deprecation.proto

package annotations

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         246172783,
		Name:          "solo.io.envoy.annotations.disallowed_by_default",
		Tag:           "varint,246172783,opt,name=disallowed_by_default",
		Filename:      "github.com/solo-io/gloo/projects/gloo/api/external/envoy/annotations/deprecation.proto",
	},
	{
		ExtendedType:  (*descriptor.EnumValueOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         178329844,
		Name:          "solo.io.envoy.annotations.disallowed_by_default_enum",
		Tag:           "varint,178329844,opt,name=disallowed_by_default_enum",
		Filename:      "github.com/solo-io/gloo/projects/gloo/api/external/envoy/annotations/deprecation.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional bool disallowed_by_default = 246172783;
	E_DisallowedByDefault = &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_extTypes[0]
)

// Extension fields to descriptor.EnumValueOptions.
var (
	// optional bool disallowed_by_default_enum = 178329844;
	E_DisallowedByDefaultEnum = &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_extTypes[1]
)

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_rawDesc = []byte{
	0x0a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x54, 0x0a, 0x15, 0x64, 0x69, 0x73,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xef, 0x98, 0xb1, 0x75, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x64, 0x69, 0x73, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x42, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3a,
	0x61, 0x0a, 0x1a, 0x64, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x21, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xf4, 0xb1, 0x84, 0x55, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x64, 0x69, 0x73, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x42, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x6e,
	0x75, 0x6d, 0x42, 0x56, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01,
	0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_goTypes = []interface{}{
	(*descriptor.FieldOptions)(nil),     // 0: google.protobuf.FieldOptions
	(*descriptor.EnumValueOptions)(nil), // 1: google.protobuf.EnumValueOptions
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_depIdxs = []int32{
	0, // 0: solo.io.envoy.annotations.disallowed_by_default:extendee -> google.protobuf.FieldOptions
	1, // 1: solo.io.envoy.annotations.disallowed_by_default_enum:extendee -> google.protobuf.EnumValueOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_depIdxs,
		ExtensionInfos:    file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_extTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_annotations_deprecation_proto_depIdxs = nil
}
