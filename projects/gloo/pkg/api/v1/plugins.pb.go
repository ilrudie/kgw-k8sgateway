// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plugins.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import aws_plugins_gloo_solo_io "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins/aws"
import azure_plugins_gloo_solo_io "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins/azure"
import kubernetes_plugins_gloo_solo_io "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1/plugins/kubernetes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Plugin-specific configuration that lives on listeners
// Each ListenerPlugin object contains configuration for a specific plugin
type ListenerPlugin struct {
	// Note to developers: new Listener Plugins must be added to this oneof field
	// to be usable by Gloo.
	//
	// Types that are valid to be assigned to PluginType:
	//	*ListenerPlugin_Empty
	PluginType isListenerPlugin_PluginType `protobuf_oneof:"plugin_type"`
}

func (m *ListenerPlugin) Reset()                    { *m = ListenerPlugin{} }
func (m *ListenerPlugin) String() string            { return proto.CompactTextString(m) }
func (*ListenerPlugin) ProtoMessage()               {}
func (*ListenerPlugin) Descriptor() ([]byte, []int) { return fileDescriptorPlugins, []int{0} }

type isListenerPlugin_PluginType interface {
	isListenerPlugin_PluginType()
	Equal(interface{}) bool
}

type ListenerPlugin_Empty struct {
	Empty string `protobuf:"bytes,1,opt,name=empty,proto3,oneof"`
}

func (*ListenerPlugin_Empty) isListenerPlugin_PluginType() {}

func (m *ListenerPlugin) GetPluginType() isListenerPlugin_PluginType {
	if m != nil {
		return m.PluginType
	}
	return nil
}

func (m *ListenerPlugin) GetEmpty() string {
	if x, ok := m.GetPluginType().(*ListenerPlugin_Empty); ok {
		return x.Empty
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ListenerPlugin) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ListenerPlugin_OneofMarshaler, _ListenerPlugin_OneofUnmarshaler, _ListenerPlugin_OneofSizer, []interface{}{
		(*ListenerPlugin_Empty)(nil),
	}
}

func _ListenerPlugin_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ListenerPlugin)
	// plugin_type
	switch x := m.PluginType.(type) {
	case *ListenerPlugin_Empty:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Empty)
	case nil:
	default:
		return fmt.Errorf("ListenerPlugin.PluginType has unexpected type %T", x)
	}
	return nil
}

func _ListenerPlugin_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ListenerPlugin)
	switch tag {
	case 1: // plugin_type.empty
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PluginType = &ListenerPlugin_Empty{x}
		return true, err
	default:
		return false, nil
	}
}

func _ListenerPlugin_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ListenerPlugin)
	// plugin_type
	switch x := m.PluginType.(type) {
	case *ListenerPlugin_Empty:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Empty)))
		n += len(x.Empty)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Plugin-specific configuration that lives on virtual hosts
// Each VirtualHostPlugin object contains configuration for a specific plugin
type VirtualHostPlugin struct {
	// Note to developers: new Virtual Host Plugins must be added to this oneof field
	// to be usable by Gloo.
	//
	// Types that are valid to be assigned to PluginType:
	//	*VirtualHostPlugin_Empty
	PluginType isVirtualHostPlugin_PluginType `protobuf_oneof:"plugin_type"`
}

func (m *VirtualHostPlugin) Reset()                    { *m = VirtualHostPlugin{} }
func (m *VirtualHostPlugin) String() string            { return proto.CompactTextString(m) }
func (*VirtualHostPlugin) ProtoMessage()               {}
func (*VirtualHostPlugin) Descriptor() ([]byte, []int) { return fileDescriptorPlugins, []int{1} }

type isVirtualHostPlugin_PluginType interface {
	isVirtualHostPlugin_PluginType()
	Equal(interface{}) bool
}

type VirtualHostPlugin_Empty struct {
	Empty string `protobuf:"bytes,1,opt,name=empty,proto3,oneof"`
}

func (*VirtualHostPlugin_Empty) isVirtualHostPlugin_PluginType() {}

func (m *VirtualHostPlugin) GetPluginType() isVirtualHostPlugin_PluginType {
	if m != nil {
		return m.PluginType
	}
	return nil
}

func (m *VirtualHostPlugin) GetEmpty() string {
	if x, ok := m.GetPluginType().(*VirtualHostPlugin_Empty); ok {
		return x.Empty
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*VirtualHostPlugin) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _VirtualHostPlugin_OneofMarshaler, _VirtualHostPlugin_OneofUnmarshaler, _VirtualHostPlugin_OneofSizer, []interface{}{
		(*VirtualHostPlugin_Empty)(nil),
	}
}

func _VirtualHostPlugin_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*VirtualHostPlugin)
	// plugin_type
	switch x := m.PluginType.(type) {
	case *VirtualHostPlugin_Empty:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Empty)
	case nil:
	default:
		return fmt.Errorf("VirtualHostPlugin.PluginType has unexpected type %T", x)
	}
	return nil
}

func _VirtualHostPlugin_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*VirtualHostPlugin)
	switch tag {
	case 1: // plugin_type.empty
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PluginType = &VirtualHostPlugin_Empty{x}
		return true, err
	default:
		return false, nil
	}
}

func _VirtualHostPlugin_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*VirtualHostPlugin)
	// plugin_type
	switch x := m.PluginType.(type) {
	case *VirtualHostPlugin_Empty:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Empty)))
		n += len(x.Empty)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Plugin-specific configuration that lives on routes
// Each RoutePlugin object contains configuration for a specific plugin
type RoutePlugin struct {
	// Note to developers: new Route Plugins must be added to this oneof field
	// to be usable by Gloo.
	//
	// Types that are valid to be assigned to PluginType:
	//	*RoutePlugin_Empty
	PluginType isRoutePlugin_PluginType `protobuf_oneof:"plugin_type"`
}

func (m *RoutePlugin) Reset()                    { *m = RoutePlugin{} }
func (m *RoutePlugin) String() string            { return proto.CompactTextString(m) }
func (*RoutePlugin) ProtoMessage()               {}
func (*RoutePlugin) Descriptor() ([]byte, []int) { return fileDescriptorPlugins, []int{2} }

type isRoutePlugin_PluginType interface {
	isRoutePlugin_PluginType()
	Equal(interface{}) bool
}

type RoutePlugin_Empty struct {
	Empty string `protobuf:"bytes,1,opt,name=empty,proto3,oneof"`
}

func (*RoutePlugin_Empty) isRoutePlugin_PluginType() {}

func (m *RoutePlugin) GetPluginType() isRoutePlugin_PluginType {
	if m != nil {
		return m.PluginType
	}
	return nil
}

func (m *RoutePlugin) GetEmpty() string {
	if x, ok := m.GetPluginType().(*RoutePlugin_Empty); ok {
		return x.Empty
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RoutePlugin) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RoutePlugin_OneofMarshaler, _RoutePlugin_OneofUnmarshaler, _RoutePlugin_OneofSizer, []interface{}{
		(*RoutePlugin_Empty)(nil),
	}
}

func _RoutePlugin_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RoutePlugin)
	// plugin_type
	switch x := m.PluginType.(type) {
	case *RoutePlugin_Empty:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Empty)
	case nil:
	default:
		return fmt.Errorf("RoutePlugin.PluginType has unexpected type %T", x)
	}
	return nil
}

func _RoutePlugin_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RoutePlugin)
	switch tag {
	case 1: // plugin_type.empty
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PluginType = &RoutePlugin_Empty{x}
		return true, err
	default:
		return false, nil
	}
}

func _RoutePlugin_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RoutePlugin)
	// plugin_type
	switch x := m.PluginType.(type) {
	case *RoutePlugin_Empty:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Empty)))
		n += len(x.Empty)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Configuration for Destinations that are tied to the UpstreamSpec or ServiceSpec on that destination
type DestinationSpec struct {
	// Note to developers: new DestinationSpecs must be added to this oneof field
	// to be usable by Gloo.
	//
	// Types that are valid to be assigned to DestinationType:
	//	*DestinationSpec_Aws
	//	*DestinationSpec_Azure
	DestinationType isDestinationSpec_DestinationType `protobuf_oneof:"destination_type"`
}

func (m *DestinationSpec) Reset()                    { *m = DestinationSpec{} }
func (m *DestinationSpec) String() string            { return proto.CompactTextString(m) }
func (*DestinationSpec) ProtoMessage()               {}
func (*DestinationSpec) Descriptor() ([]byte, []int) { return fileDescriptorPlugins, []int{3} }

type isDestinationSpec_DestinationType interface {
	isDestinationSpec_DestinationType()
	Equal(interface{}) bool
}

type DestinationSpec_Aws struct {
	Aws *aws_plugins_gloo_solo_io.DestinationSpec `protobuf:"bytes,1,opt,name=aws,oneof"`
}
type DestinationSpec_Azure struct {
	Azure *azure_plugins_gloo_solo_io.DestinationSpec `protobuf:"bytes,2,opt,name=azure,oneof"`
}

func (*DestinationSpec_Aws) isDestinationSpec_DestinationType()   {}
func (*DestinationSpec_Azure) isDestinationSpec_DestinationType() {}

func (m *DestinationSpec) GetDestinationType() isDestinationSpec_DestinationType {
	if m != nil {
		return m.DestinationType
	}
	return nil
}

func (m *DestinationSpec) GetAws() *aws_plugins_gloo_solo_io.DestinationSpec {
	if x, ok := m.GetDestinationType().(*DestinationSpec_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *DestinationSpec) GetAzure() *azure_plugins_gloo_solo_io.DestinationSpec {
	if x, ok := m.GetDestinationType().(*DestinationSpec_Azure); ok {
		return x.Azure
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DestinationSpec) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DestinationSpec_OneofMarshaler, _DestinationSpec_OneofUnmarshaler, _DestinationSpec_OneofSizer, []interface{}{
		(*DestinationSpec_Aws)(nil),
		(*DestinationSpec_Azure)(nil),
	}
}

func _DestinationSpec_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DestinationSpec)
	// destination_type
	switch x := m.DestinationType.(type) {
	case *DestinationSpec_Aws:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aws); err != nil {
			return err
		}
	case *DestinationSpec_Azure:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Azure); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DestinationSpec.DestinationType has unexpected type %T", x)
	}
	return nil
}

func _DestinationSpec_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DestinationSpec)
	switch tag {
	case 1: // destination_type.aws
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(aws_plugins_gloo_solo_io.DestinationSpec)
		err := b.DecodeMessage(msg)
		m.DestinationType = &DestinationSpec_Aws{msg}
		return true, err
	case 2: // destination_type.azure
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(azure_plugins_gloo_solo_io.DestinationSpec)
		err := b.DecodeMessage(msg)
		m.DestinationType = &DestinationSpec_Azure{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DestinationSpec_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DestinationSpec)
	// destination_type
	switch x := m.DestinationType.(type) {
	case *DestinationSpec_Aws:
		s := proto.Size(x.Aws)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DestinationSpec_Azure:
		s := proto.Size(x.Azure)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Each upstream in Gloo has a type. Supported types include `static`, `kubernetes`, `aws`, `consul`, and more.
// Each upstream type is handled by a corresponding Gloo plugin.
type UpstreamSpec struct {
	// Note to developers: new Upstream Plugins must be added to this oneof field
	// to be usable by Gloo.
	//
	// Types that are valid to be assigned to UpstreamType:
	//	*UpstreamSpec_Kube
	//	*UpstreamSpec_Aws
	//	*UpstreamSpec_Azure
	UpstreamType isUpstreamSpec_UpstreamType `protobuf_oneof:"upstream_type"`
}

func (m *UpstreamSpec) Reset()                    { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string            { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()               {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) { return fileDescriptorPlugins, []int{4} }

type isUpstreamSpec_UpstreamType interface {
	isUpstreamSpec_UpstreamType()
	Equal(interface{}) bool
}

type UpstreamSpec_Kube struct {
	Kube *kubernetes_plugins_gloo_solo_io.UpstreamSpec `protobuf:"bytes,1,opt,name=kube,oneof"`
}
type UpstreamSpec_Aws struct {
	Aws *aws_plugins_gloo_solo_io.UpstreamSpec `protobuf:"bytes,2,opt,name=aws,oneof"`
}
type UpstreamSpec_Azure struct {
	Azure *azure_plugins_gloo_solo_io.UpstreamSpec `protobuf:"bytes,3,opt,name=azure,oneof"`
}

func (*UpstreamSpec_Kube) isUpstreamSpec_UpstreamType()  {}
func (*UpstreamSpec_Aws) isUpstreamSpec_UpstreamType()   {}
func (*UpstreamSpec_Azure) isUpstreamSpec_UpstreamType() {}

func (m *UpstreamSpec) GetUpstreamType() isUpstreamSpec_UpstreamType {
	if m != nil {
		return m.UpstreamType
	}
	return nil
}

func (m *UpstreamSpec) GetKube() *kubernetes_plugins_gloo_solo_io.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*UpstreamSpec_Kube); ok {
		return x.Kube
	}
	return nil
}

func (m *UpstreamSpec) GetAws() *aws_plugins_gloo_solo_io.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*UpstreamSpec_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *UpstreamSpec) GetAzure() *azure_plugins_gloo_solo_io.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*UpstreamSpec_Azure); ok {
		return x.Azure
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UpstreamSpec) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UpstreamSpec_OneofMarshaler, _UpstreamSpec_OneofUnmarshaler, _UpstreamSpec_OneofSizer, []interface{}{
		(*UpstreamSpec_Kube)(nil),
		(*UpstreamSpec_Aws)(nil),
		(*UpstreamSpec_Azure)(nil),
	}
}

func _UpstreamSpec_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UpstreamSpec)
	// upstream_type
	switch x := m.UpstreamType.(type) {
	case *UpstreamSpec_Kube:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Kube); err != nil {
			return err
		}
	case *UpstreamSpec_Aws:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aws); err != nil {
			return err
		}
	case *UpstreamSpec_Azure:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Azure); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UpstreamSpec.UpstreamType has unexpected type %T", x)
	}
	return nil
}

func _UpstreamSpec_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UpstreamSpec)
	switch tag {
	case 1: // upstream_type.kube
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes_plugins_gloo_solo_io.UpstreamSpec)
		err := b.DecodeMessage(msg)
		m.UpstreamType = &UpstreamSpec_Kube{msg}
		return true, err
	case 2: // upstream_type.aws
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(aws_plugins_gloo_solo_io.UpstreamSpec)
		err := b.DecodeMessage(msg)
		m.UpstreamType = &UpstreamSpec_Aws{msg}
		return true, err
	case 3: // upstream_type.azure
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(azure_plugins_gloo_solo_io.UpstreamSpec)
		err := b.DecodeMessage(msg)
		m.UpstreamType = &UpstreamSpec_Azure{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UpstreamSpec_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UpstreamSpec)
	// upstream_type
	switch x := m.UpstreamType.(type) {
	case *UpstreamSpec_Kube:
		s := proto.Size(x.Kube)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UpstreamSpec_Aws:
		s := proto.Size(x.Aws)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UpstreamSpec_Azure:
		s := proto.Size(x.Azure)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ListenerPlugin)(nil), "gloo.solo.io.ListenerPlugin")
	proto.RegisterType((*VirtualHostPlugin)(nil), "gloo.solo.io.VirtualHostPlugin")
	proto.RegisterType((*RoutePlugin)(nil), "gloo.solo.io.RoutePlugin")
	proto.RegisterType((*DestinationSpec)(nil), "gloo.solo.io.DestinationSpec")
	proto.RegisterType((*UpstreamSpec)(nil), "gloo.solo.io.UpstreamSpec")
}
func (this *ListenerPlugin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListenerPlugin)
	if !ok {
		that2, ok := that.(ListenerPlugin)
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
	if that1.PluginType == nil {
		if this.PluginType != nil {
			return false
		}
	} else if this.PluginType == nil {
		return false
	} else if !this.PluginType.Equal(that1.PluginType) {
		return false
	}
	return true
}
func (this *ListenerPlugin_Empty) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListenerPlugin_Empty)
	if !ok {
		that2, ok := that.(ListenerPlugin_Empty)
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
	if this.Empty != that1.Empty {
		return false
	}
	return true
}
func (this *VirtualHostPlugin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualHostPlugin)
	if !ok {
		that2, ok := that.(VirtualHostPlugin)
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
	if that1.PluginType == nil {
		if this.PluginType != nil {
			return false
		}
	} else if this.PluginType == nil {
		return false
	} else if !this.PluginType.Equal(that1.PluginType) {
		return false
	}
	return true
}
func (this *VirtualHostPlugin_Empty) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualHostPlugin_Empty)
	if !ok {
		that2, ok := that.(VirtualHostPlugin_Empty)
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
	if this.Empty != that1.Empty {
		return false
	}
	return true
}
func (this *RoutePlugin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RoutePlugin)
	if !ok {
		that2, ok := that.(RoutePlugin)
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
	if that1.PluginType == nil {
		if this.PluginType != nil {
			return false
		}
	} else if this.PluginType == nil {
		return false
	} else if !this.PluginType.Equal(that1.PluginType) {
		return false
	}
	return true
}
func (this *RoutePlugin_Empty) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RoutePlugin_Empty)
	if !ok {
		that2, ok := that.(RoutePlugin_Empty)
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
	if this.Empty != that1.Empty {
		return false
	}
	return true
}
func (this *DestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec)
	if !ok {
		that2, ok := that.(DestinationSpec)
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
	if that1.DestinationType == nil {
		if this.DestinationType != nil {
			return false
		}
	} else if this.DestinationType == nil {
		return false
	} else if !this.DestinationType.Equal(that1.DestinationType) {
		return false
	}
	return true
}
func (this *DestinationSpec_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec_Aws)
	if !ok {
		that2, ok := that.(DestinationSpec_Aws)
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
	if !this.Aws.Equal(that1.Aws) {
		return false
	}
	return true
}
func (this *DestinationSpec_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DestinationSpec_Azure)
	if !ok {
		that2, ok := that.(DestinationSpec_Azure)
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
	if !this.Azure.Equal(that1.Azure) {
		return false
	}
	return true
}
func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
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
	if that1.UpstreamType == nil {
		if this.UpstreamType != nil {
			return false
		}
	} else if this.UpstreamType == nil {
		return false
	} else if !this.UpstreamType.Equal(that1.UpstreamType) {
		return false
	}
	return true
}
func (this *UpstreamSpec_Kube) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_Kube)
	if !ok {
		that2, ok := that.(UpstreamSpec_Kube)
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
	if !this.Kube.Equal(that1.Kube) {
		return false
	}
	return true
}
func (this *UpstreamSpec_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_Aws)
	if !ok {
		that2, ok := that.(UpstreamSpec_Aws)
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
	if !this.Aws.Equal(that1.Aws) {
		return false
	}
	return true
}
func (this *UpstreamSpec_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec_Azure)
	if !ok {
		that2, ok := that.(UpstreamSpec_Azure)
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
	if !this.Azure.Equal(that1.Azure) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("plugins.proto", fileDescriptorPlugins) }

var fileDescriptorPlugins = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xdf, 0x4a, 0xeb, 0x40,
	0x10, 0xc6, 0x9b, 0xd3, 0xd3, 0x03, 0x67, 0xdb, 0x5a, 0x0d, 0x22, 0xa5, 0x17, 0x22, 0xbd, 0x90,
	0x8a, 0x74, 0x17, 0xb5, 0x20, 0x14, 0x04, 0x69, 0x45, 0x82, 0x78, 0x21, 0xf1, 0xcf, 0x85, 0x37,
	0x92, 0xd6, 0x25, 0xae, 0x4d, 0x33, 0x4b, 0x76, 0x62, 0xa9, 0x4f, 0xa4, 0xaf, 0xd5, 0x27, 0x91,
	0xec, 0xae, 0x18, 0x8a, 0x4a, 0xdb, 0x8b, 0x24, 0x1b, 0x66, 0x7e, 0xdf, 0xce, 0x37, 0x33, 0xa4,
	0x2a, 0xa3, 0x34, 0x14, 0xb1, 0xa2, 0x32, 0x01, 0x04, 0xb7, 0x12, 0x46, 0x00, 0x54, 0x41, 0x04,
	0x54, 0x40, 0x63, 0x33, 0x84, 0x10, 0x74, 0x80, 0x65, 0x27, 0x93, 0xd3, 0x38, 0x0f, 0x05, 0x3e,
	0xa5, 0x03, 0x3a, 0x84, 0x31, 0xcb, 0x32, 0xdb, 0x02, 0xcc, 0x77, 0x24, 0x90, 0xc9, 0x04, 0x9e,
	0xf9, 0x10, 0x15, 0xcb, 0x84, 0x58, 0x20, 0x05, 0x7b, 0x39, 0x60, 0xf6, 0x0e, 0x16, 0x4c, 0xf4,
	0x63, 0x75, 0x2e, 0x56, 0xd7, 0x79, 0x4d, 0x13, 0x6e, 0xde, 0x56, 0xeb, 0x66, 0x65, 0xad, 0x51,
	0x3a, 0xe0, 0x49, 0xcc, 0x91, 0xe7, 0x8f, 0x46, 0xb5, 0x79, 0x4c, 0xd6, 0x2e, 0x85, 0x42, 0x1e,
	0xf3, 0xe4, 0x4a, 0xa7, 0xbb, 0x5b, 0xa4, 0xc4, 0xc7, 0x12, 0xa7, 0x75, 0x67, 0xc7, 0x69, 0xfd,
	0xf7, 0x0a, 0xbe, 0xf9, 0xed, 0x55, 0x49, 0xd9, 0x08, 0x3e, 0xe0, 0x54, 0xf2, 0x66, 0x97, 0x6c,
	0xdc, 0x89, 0x04, 0xd3, 0x20, 0xf2, 0x40, 0xe1, 0x72, 0x6c, 0x87, 0x94, 0x7d, 0x48, 0x91, 0x2f,
	0x47, 0xbd, 0x3b, 0xa4, 0x76, 0xc6, 0x15, 0x8a, 0x38, 0x40, 0x01, 0xf1, 0xb5, 0xe4, 0x43, 0xf7,
	0x84, 0x14, 0x83, 0x89, 0xd2, 0x60, 0xf9, 0x70, 0x8f, 0xea, 0xce, 0xdb, 0x69, 0xe7, 0xc7, 0x4c,
	0xe7, 0x38, 0xaf, 0xe0, 0x67, 0x9c, 0xdb, 0x27, 0x25, 0xdd, 0xe2, 0xfa, 0x1f, 0x2d, 0xb0, 0x4f,
	0x6d, 0xc3, 0x17, 0x93, 0x30, 0x6c, 0xcf, 0x25, 0xeb, 0x8f, 0x5f, 0x31, 0x53, 0xeb, 0xcc, 0x21,
	0x95, 0x5b, 0xa9, 0x30, 0xe1, 0xc1, 0x58, 0x17, 0xda, 0x27, 0x7f, 0xb3, 0xde, 0xdb, 0x4a, 0xdb,
	0x34, 0x3f, 0x88, 0xef, 0x6e, 0xcb, 0xc3, 0x5e, 0xc1, 0xd7, 0xb0, 0xdb, 0x35, 0x6e, 0x4d, 0xb1,
	0xbb, 0x3f, 0xbb, 0x9d, 0x83, 0xb5, 0xd5, 0xd3, 0x4f, 0xab, 0x45, 0x4d, 0xb7, 0x7e, 0xb3, 0x3a,
	0xc7, 0x5b, 0x9f, 0x35, 0x52, 0x4d, 0x6d, 0x40, 0x9b, 0xec, 0x75, 0xdf, 0x66, 0xdb, 0xce, 0x7d,
	0x67, 0xf1, 0xbd, 0x94, 0xa3, 0xd0, 0xee, 0xe6, 0xe0, 0x9f, 0x5e, 0xbf, 0xa3, 0x8f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x8c, 0x04, 0xf1, 0xbc, 0x9d, 0x03, 0x00, 0x00,
}
