// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: ssl_gc_rcon_team.proto

package rcon

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// a registration that must be send by teams and autoRefs to the controller as the very first message
type TeamRegistration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the exact team name as published by the game-controller
	TeamName *string `protobuf:"bytes,1,req,name=team_name,json=teamName" json:"team_name,omitempty"`
	// signature can optionally be specified to enable secure communication
	Signature *Signature `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
}

func (x *TeamRegistration) Reset() {
	*x = TeamRegistration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_team_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamRegistration) ProtoMessage() {}

func (x *TeamRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_team_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamRegistration.ProtoReflect.Descriptor instead.
func (*TeamRegistration) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_team_proto_rawDescGZIP(), []int{0}
}

func (x *TeamRegistration) GetTeamName() string {
	if x != nil && x.TeamName != nil {
		return *x.TeamName
	}
	return ""
}

func (x *TeamRegistration) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// wrapper for all messages from a team's computer to the controller
type TeamToController struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signature can optionally be specified to enable secure communication
	Signature *Signature `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// Types that are assignable to Msg:
	//	*TeamToController_DesiredKeeper
	//	*TeamToController_SubstituteBot
	//	*TeamToController_Ping
	Msg isTeamToController_Msg `protobuf_oneof:"msg"`
}

func (x *TeamToController) Reset() {
	*x = TeamToController{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_team_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamToController) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamToController) ProtoMessage() {}

func (x *TeamToController) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_team_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamToController.ProtoReflect.Descriptor instead.
func (*TeamToController) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_team_proto_rawDescGZIP(), []int{1}
}

func (x *TeamToController) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (m *TeamToController) GetMsg() isTeamToController_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *TeamToController) GetDesiredKeeper() int32 {
	if x, ok := x.GetMsg().(*TeamToController_DesiredKeeper); ok {
		return x.DesiredKeeper
	}
	return 0
}

func (x *TeamToController) GetSubstituteBot() bool {
	if x, ok := x.GetMsg().(*TeamToController_SubstituteBot); ok {
		return x.SubstituteBot
	}
	return false
}

func (x *TeamToController) GetPing() bool {
	if x, ok := x.GetMsg().(*TeamToController_Ping); ok {
		return x.Ping
	}
	return false
}

type isTeamToController_Msg interface {
	isTeamToController_Msg()
}

type TeamToController_DesiredKeeper struct {
	// request a new desired keeper id
	// this is only allowed during STOP and will be rejected otherwise
	DesiredKeeper int32 `protobuf:"varint,2,opt,name=desired_keeper,json=desiredKeeper,oneof"`
}

type TeamToController_SubstituteBot struct {
	// request to substitute a robot at the next possibility
	SubstituteBot bool `protobuf:"varint,4,opt,name=substitute_bot,json=substituteBot,oneof"`
}

type TeamToController_Ping struct {
	// send a ping to the GC to test if the connection is still open.
	// the value is ignored and a reply is sent back
	Ping bool `protobuf:"varint,5,opt,name=ping,oneof"`
}

func (*TeamToController_DesiredKeeper) isTeamToController_Msg() {}

func (*TeamToController_SubstituteBot) isTeamToController_Msg() {}

func (*TeamToController_Ping) isTeamToController_Msg() {}

// wrapper for all messages from controller to a team's computer
type ControllerToTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*ControllerToTeam_ControllerReply
	Msg isControllerToTeam_Msg `protobuf_oneof:"msg"`
}

func (x *ControllerToTeam) Reset() {
	*x = ControllerToTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_team_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerToTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerToTeam) ProtoMessage() {}

func (x *ControllerToTeam) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_team_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerToTeam.ProtoReflect.Descriptor instead.
func (*ControllerToTeam) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_team_proto_rawDescGZIP(), []int{2}
}

func (m *ControllerToTeam) GetMsg() isControllerToTeam_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *ControllerToTeam) GetControllerReply() *ControllerReply {
	if x, ok := x.GetMsg().(*ControllerToTeam_ControllerReply); ok {
		return x.ControllerReply
	}
	return nil
}

type isControllerToTeam_Msg interface {
	isControllerToTeam_Msg()
}

type ControllerToTeam_ControllerReply struct {
	// a reply from the controller
	ControllerReply *ControllerReply `protobuf:"bytes,1,opt,name=controller_reply,json=controllerReply,oneof"`
}

func (*ControllerToTeam_ControllerReply) isControllerToTeam_Msg() {}

var File_ssl_gc_rcon_team_proto protoreflect.FileDescriptor

var file_ssl_gc_rcon_team_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x72, 0x63, 0x6f, 0x6e, 0x5f, 0x74, 0x65,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63,
	0x5f, 0x72, 0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x10, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x10, 0x54, 0x65, 0x61, 0x6d, 0x54,
	0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x27, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64,
	0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x0d, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x12, 0x27,
	0x0a, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x5f, 0x62, 0x6f, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69,
	0x74, 0x75, 0x74, 0x65, 0x42, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x05, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x5e, 0x0a, 0x10, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x54, 0x6f, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x3d,
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x05, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x62, 0x6f, 0x43, 0x75, 0x70,
	0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73, 0x6c, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x63, 0x6f, 0x6e,
}

var (
	file_ssl_gc_rcon_team_proto_rawDescOnce sync.Once
	file_ssl_gc_rcon_team_proto_rawDescData = file_ssl_gc_rcon_team_proto_rawDesc
)

func file_ssl_gc_rcon_team_proto_rawDescGZIP() []byte {
	file_ssl_gc_rcon_team_proto_rawDescOnce.Do(func() {
		file_ssl_gc_rcon_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_gc_rcon_team_proto_rawDescData)
	})
	return file_ssl_gc_rcon_team_proto_rawDescData
}

var file_ssl_gc_rcon_team_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ssl_gc_rcon_team_proto_goTypes = []interface{}{
	(*TeamRegistration)(nil), // 0: TeamRegistration
	(*TeamToController)(nil), // 1: TeamToController
	(*ControllerToTeam)(nil), // 2: ControllerToTeam
	(*Signature)(nil),        // 3: Signature
	(*ControllerReply)(nil),  // 4: ControllerReply
}
var file_ssl_gc_rcon_team_proto_depIdxs = []int32{
	3, // 0: TeamRegistration.signature:type_name -> Signature
	3, // 1: TeamToController.signature:type_name -> Signature
	4, // 2: ControllerToTeam.controller_reply:type_name -> ControllerReply
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ssl_gc_rcon_team_proto_init() }
func file_ssl_gc_rcon_team_proto_init() {
	if File_ssl_gc_rcon_team_proto != nil {
		return
	}
	file_ssl_gc_rcon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ssl_gc_rcon_team_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamRegistration); i {
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
		file_ssl_gc_rcon_team_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamToController); i {
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
		file_ssl_gc_rcon_team_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerToTeam); i {
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
	file_ssl_gc_rcon_team_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TeamToController_DesiredKeeper)(nil),
		(*TeamToController_SubstituteBot)(nil),
		(*TeamToController_Ping)(nil),
	}
	file_ssl_gc_rcon_team_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ControllerToTeam_ControllerReply)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ssl_gc_rcon_team_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_gc_rcon_team_proto_goTypes,
		DependencyIndexes: file_ssl_gc_rcon_team_proto_depIdxs,
		MessageInfos:      file_ssl_gc_rcon_team_proto_msgTypes,
	}.Build()
	File_ssl_gc_rcon_team_proto = out.File
	file_ssl_gc_rcon_team_proto_rawDesc = nil
	file_ssl_gc_rcon_team_proto_goTypes = nil
	file_ssl_gc_rcon_team_proto_depIdxs = nil
}
