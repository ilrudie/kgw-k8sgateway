// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/stats/stats.proto

package stats

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This plugin provides additional configuration options to expose statistics.
type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Virtual clusters allow exposing additional statistics for traffic served by a Virtual Host.
	VirtualClusters []*VirtualCluster `protobuf:"bytes,10,rep,name=virtual_clusters,json=virtualClusters,proto3" json:"virtual_clusters,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_rawDescGZIP(), []int{0}
}

func (x *Stats) GetVirtualClusters() []*VirtualCluster {
	if x != nil {
		return x.VirtualClusters
	}
	return nil
}

// Virtual clusters allow you to expose statistics for virtual host traffic that matches certain criteria.
// This is useful because what the application considers to be an endpoint does often not map directly to
// the routing configuration, so Envoy does not emit per endpoint statistics. Using virtual clusters you can define
// logical endpoints and have Envoy emit dedicated statistics for any matching request. Virtual cluster statistics
// are emitted on the downstream side and thus include network level failures.
//
// Please note that virtual clusters add overhead to the processing of each requests and should not be overused.
type VirtualCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the virtual cluster. This value will be used together with the virtual host name to
	// compute the name of the statistics emitted by this virtual cluster. Statistics names will be in the form:
	// vhost.<virtual host name>.vcluster.<virtual cluster name>.<stat name>.
	// See [the official Envoy documentation](https://www.envoyproxy.io/docs/envoy/v1.5.0/configuration/http_filters/router_filter#config-http-filters-router-stats)
	// for more information about the statistics emitted when virtual cluster configurations are specified.
	//
	// Note: This string should not contain any dots ("."), as this is a reserved character for Envoy statistics names.
	// Any dot present in the virtual cluster name will be replaced with an underscore ("_") character by Gloo.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The regex pattern used by Envoy to decide whether to expose statistics for a particular request.
	// Please note that **the entire path** of the request must match the regex (e.g. the regex `/rides/d+` matches
	// the path `/rides/0`, but not `/rides/123/456`).
	// The regex grammar used is defined [here](https://en.cppreference.com/w/cpp/regex/ecmascript).
	Pattern string `protobuf:"bytes,2,opt,name=pattern,proto3" json:"pattern,omitempty"`
	// If specified, statistics will be exposed only for requests matching the given HTTP method.
	Method string `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *VirtualCluster) Reset() {
	*x = VirtualCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualCluster) ProtoMessage() {}

func (x *VirtualCluster) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualCluster.ProtoReflect.Descriptor instead.
func (*VirtualCluster) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_rawDescGZIP(), []int{1}
}

func (x *VirtualCluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VirtualCluster) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *VirtualCluster) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x55, 0x0a, 0x10, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x56, 0x0a, 0x0e, 0x56, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x42, 0x4c, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x73, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_goTypes = []interface{}{
	(*Stats)(nil),          // 0: stats.options.gloo.solo.io.Stats
	(*VirtualCluster)(nil), // 1: stats.options.gloo.solo.io.VirtualCluster
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_depIdxs = []int32{
	1, // 0: stats.options.gloo.solo.io.Stats.virtual_clusters:type_name -> stats.options.gloo.solo.io.VirtualCluster
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualCluster); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_stats_stats_proto_depIdxs = nil
}