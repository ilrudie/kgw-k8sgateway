// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gateway2/api/v1alpha1/kube/container.proto

package kube

import (
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Image_PullPolicy int32

const (
	// The image pull policy will be defaulted based on the image tag or digest.
	// See
	// https://kubernetes.io/docs/concepts/containers/images/#imagepullpolicy-defaulting
	// for details.
	Image_Unspecified Image_PullPolicy = 0
	// The image is pulled only if it is not already present locally.
	// See
	// https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy
	// for details.
	Image_IfNotPresent Image_PullPolicy = 1
	// Every time the kubelet launches a container, the kubelet queries the
	// container image registry to resolve the name to an image digest. See
	// https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy
	// for details.
	Image_Always Image_PullPolicy = 2
	// The kubelet does not try fetching the image. If the image is somehow
	// already present locally, the kubelet attempts to start the container;
	// otherwise, startup fails. See
	// https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy
	// for details.
	Image_Never Image_PullPolicy = 3
)

// Enum value maps for Image_PullPolicy.
var (
	Image_PullPolicy_name = map[int32]string{
		0: "Unspecified",
		1: "IfNotPresent",
		2: "Always",
		3: "Never",
	}
	Image_PullPolicy_value = map[string]int32{
		"Unspecified":  0,
		"IfNotPresent": 1,
		"Always":       2,
		"Never":        3,
	}
)

func (x Image_PullPolicy) Enum() *Image_PullPolicy {
	p := new(Image_PullPolicy)
	*p = x
	return p
}

func (x Image_PullPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Image_PullPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_enumTypes[0].Descriptor()
}

func (Image_PullPolicy) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_enumTypes[0]
}

func (x Image_PullPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Image_PullPolicy.Descriptor instead.
func (Image_PullPolicy) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDescGZIP(), []int{0, 0}
}

// A container image. See https://kubernetes.io/docs/concepts/containers/images
// for details.
type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The image registry.
	Registry string `protobuf:"bytes,1,opt,name=registry,proto3" json:"registry,omitempty"`
	// The image repository (name).
	Repository string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	// The image tag.
	Tag string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	// The hash digest of the image, e.g. `sha256:12345...`
	Digest string `protobuf:"bytes,4,opt,name=digest,proto3" json:"digest,omitempty"`
	// The image pull policy for the container. See
	// https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy
	// for details.
	PullPolicy Image_PullPolicy `protobuf:"varint,5,opt,name=pull_policy,json=pullPolicy,proto3,enum=kube.gateway.gloo.solo.io.Image_PullPolicy" json:"pull_policy,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

func (x *Image) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *Image) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Image) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Image) GetPullPolicy() Image_PullPolicy {
	if x != nil {
		return x.PullPolicy
	}
	return Image_Unspecified
}

// Compute resources required by this container. See
// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
// for details.
type ResourceRequirements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum amount of compute resources allowed.
	Limits map[string]string `protobuf:"bytes,1,rep,name=limits,proto3" json:"limits,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The minimum amount of compute resources required.
	Requests map[string]string `protobuf:"bytes,2,rep,name=requests,proto3" json:"requests,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResourceRequirements) Reset() {
	*x = ResourceRequirements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceRequirements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceRequirements) ProtoMessage() {}

func (x *ResourceRequirements) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceRequirements.ProtoReflect.Descriptor instead.
func (*ResourceRequirements) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceRequirements) GetLimits() map[string]string {
	if x != nil {
		return x.Limits
	}
	return nil
}

func (x *ResourceRequirements) GetRequests() map[string]string {
	if x != nil {
		return x.Requests
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6b,
	0x75, 0x62, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x02, 0x0a,
	0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x74, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0b,
	0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0a,
	0x70, 0x75, 0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x46, 0x0a, 0x0a, 0x50, 0x75,
	0x6c, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x66, 0x4e,
	0x6f, 0x74, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41,
	0x6c, 0x77, 0x61, 0x79, 0x73, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x65, 0x76, 0x65, 0x72,
	0x10, 0x03, 0x22, 0xbe, 0x02, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x53, 0x0a, 0x06, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x12, 0x59, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x5e, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0x5a, 0x54, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6b,
	0x75, 0x62, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDescData = file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_goTypes = []interface{}{
	(Image_PullPolicy)(0),        // 0: kube.gateway.gloo.solo.io.Image.PullPolicy
	(*Image)(nil),                // 1: kube.gateway.gloo.solo.io.Image
	(*ResourceRequirements)(nil), // 2: kube.gateway.gloo.solo.io.ResourceRequirements
	nil,                          // 3: kube.gateway.gloo.solo.io.ResourceRequirements.LimitsEntry
	nil,                          // 4: kube.gateway.gloo.solo.io.ResourceRequirements.RequestsEntry
}
var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_depIdxs = []int32{
	0, // 0: kube.gateway.gloo.solo.io.Image.pull_policy:type_name -> kube.gateway.gloo.solo.io.Image.PullPolicy
	3, // 1: kube.gateway.gloo.solo.io.ResourceRequirements.limits:type_name -> kube.gateway.gloo.solo.io.ResourceRequirements.LimitsEntry
	4, // 2: kube.gateway.gloo.solo.io.ResourceRequirements.requests:type_name -> kube.gateway.gloo.solo.io.ResourceRequirements.RequestsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_init() }
func file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_init() {
	if File_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceRequirements); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto = out.File
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_container_proto_depIdxs = nil
}
