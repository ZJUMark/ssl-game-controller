// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: ssl_gc_engine.proto

package engine

import (
	geom "github.com/RoboCup-SSL/ssl-game-controller/internal/app/geom"
	state "github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
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

type GcState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamState       map[string]*GcStateTeam    `protobuf:"bytes,1,rep,name=team_state,json=teamState" json:"team_state,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	AutoRefState    map[string]*GcStateAutoRef `protobuf:"bytes,2,rep,name=auto_ref_state,json=autoRefState" json:"auto_ref_state,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	TrackerState    map[string]*GcStateTracker `protobuf:"bytes,3,rep,name=tracker_state,json=trackerState" json:"tracker_state,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	TrackerStateGc  *GcStateTracker            `protobuf:"bytes,4,opt,name=tracker_state_gc,json=trackerStateGc" json:"tracker_state_gc,omitempty"`
	ReadyToContinue *bool                      `protobuf:"varint,5,opt,name=ready_to_continue,json=readyToContinue" json:"ready_to_continue,omitempty"`
}

func (x *GcState) Reset() {
	*x = GcState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcState) ProtoMessage() {}

func (x *GcState) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcState.ProtoReflect.Descriptor instead.
func (*GcState) Descriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{0}
}

func (x *GcState) GetTeamState() map[string]*GcStateTeam {
	if x != nil {
		return x.TeamState
	}
	return nil
}

func (x *GcState) GetAutoRefState() map[string]*GcStateAutoRef {
	if x != nil {
		return x.AutoRefState
	}
	return nil
}

func (x *GcState) GetTrackerState() map[string]*GcStateTracker {
	if x != nil {
		return x.TrackerState
	}
	return nil
}

func (x *GcState) GetTrackerStateGc() *GcStateTracker {
	if x != nil {
		return x.TrackerStateGc
	}
	return nil
}

func (x *GcState) GetReadyToContinue() bool {
	if x != nil && x.ReadyToContinue != nil {
		return *x.ReadyToContinue
	}
	return false
}

type GcStateTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// true: The team is connected
	Connected *bool `protobuf:"varint,1,opt,name=connected" json:"connected,omitempty"`
	// true: The team connected via TLS with a verified certificate
	ConnectionVerified *bool `protobuf:"varint,2,opt,name=connection_verified,json=connectionVerified" json:"connection_verified,omitempty"`
}

func (x *GcStateTeam) Reset() {
	*x = GcStateTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcStateTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcStateTeam) ProtoMessage() {}

func (x *GcStateTeam) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcStateTeam.ProtoReflect.Descriptor instead.
func (*GcStateTeam) Descriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{1}
}

func (x *GcStateTeam) GetConnected() bool {
	if x != nil && x.Connected != nil {
		return *x.Connected
	}
	return false
}

func (x *GcStateTeam) GetConnectionVerified() bool {
	if x != nil && x.ConnectionVerified != nil {
		return *x.ConnectionVerified
	}
	return false
}

type GcStateAutoRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// true: The autoRef connected via TLS with a verified certificate
	ConnectionVerified *bool `protobuf:"varint,1,opt,name=connection_verified,json=connectionVerified" json:"connection_verified,omitempty"`
}

func (x *GcStateAutoRef) Reset() {
	*x = GcStateAutoRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcStateAutoRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcStateAutoRef) ProtoMessage() {}

func (x *GcStateAutoRef) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcStateAutoRef.ProtoReflect.Descriptor instead.
func (*GcStateAutoRef) Descriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{2}
}

func (x *GcStateAutoRef) GetConnectionVerified() bool {
	if x != nil && x.ConnectionVerified != nil {
		return *x.ConnectionVerified
	}
	return false
}

type GcStateTracker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the source
	SourceName *string `protobuf:"bytes,1,opt,name=source_name,json=sourceName" json:"source_name,omitempty"`
	// Current ball
	Ball *Ball `protobuf:"bytes,2,opt,name=ball" json:"ball,omitempty"`
	// Current robots
	Robots []*Robot `protobuf:"bytes,3,rep,name=robots" json:"robots,omitempty"`
}

func (x *GcStateTracker) Reset() {
	*x = GcStateTracker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_engine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcStateTracker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcStateTracker) ProtoMessage() {}

func (x *GcStateTracker) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_engine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcStateTracker.ProtoReflect.Descriptor instead.
func (*GcStateTracker) Descriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{3}
}

func (x *GcStateTracker) GetSourceName() string {
	if x != nil && x.SourceName != nil {
		return *x.SourceName
	}
	return ""
}

func (x *GcStateTracker) GetBall() *Ball {
	if x != nil {
		return x.Ball
	}
	return nil
}

func (x *GcStateTracker) GetRobots() []*Robot {
	if x != nil {
		return x.Robots
	}
	return nil
}

type Ball struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ball position [m]
	Pos *geom.Vector3 `protobuf:"bytes,1,opt,name=pos" json:"pos,omitempty"`
	// ball velocity [m/s]
	Vel *geom.Vector3 `protobuf:"bytes,2,opt,name=vel" json:"vel,omitempty"`
}

func (x *Ball) Reset() {
	*x = Ball{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_engine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ball) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ball) ProtoMessage() {}

func (x *Ball) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_engine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ball.ProtoReflect.Descriptor instead.
func (*Ball) Descriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{4}
}

func (x *Ball) GetPos() *geom.Vector3 {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *Ball) GetVel() *geom.Vector3 {
	if x != nil {
		return x.Vel
	}
	return nil
}

type Robot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// robot id and team
	Id *state.RobotId `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// robot position [m]
	Pos *geom.Vector2 `protobuf:"bytes,2,opt,name=pos" json:"pos,omitempty"`
}

func (x *Robot) Reset() {
	*x = Robot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_engine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Robot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Robot) ProtoMessage() {}

func (x *Robot) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_engine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Robot.ProtoReflect.Descriptor instead.
func (*Robot) Descriptor() ([]byte, []int) {
	return file_ssl_gc_engine_proto_rawDescGZIP(), []int{5}
}

func (x *Robot) GetId() *state.RobotId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Robot) GetPos() *geom.Vector2 {
	if x != nil {
		return x.Pos
	}
	return nil
}

var File_ssl_gc_engine_proto protoreflect.FileDescriptor

var file_ssl_gc_engine_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x67, 0x65,
	0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x73,
	0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9b, 0x04, 0x0a, 0x07, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a,
	0x0a, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x74, 0x65, 0x61, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65,
	0x66, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x66, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x52,
	0x65, 0x66, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x67, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x47, 0x63, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x74, 0x6f, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x1a,
	0x4a, 0x0a, 0x0e, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50, 0x0a, 0x11, 0x41,
	0x75, 0x74, 0x6f, 0x52, 0x65, 0x66, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52,
	0x65, 0x66, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50, 0x0a,
	0x11, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x5c, 0x0a, 0x0b, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x13,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x41, 0x0a,
	0x0e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x66, 0x12,
	0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x22, 0x6c, 0x0a, 0x0e, 0x47, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x62, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x42, 0x61, 0x6c, 0x6c, 0x52, 0x04, 0x62, 0x61, 0x6c, 0x6c, 0x12, 0x1e,
	0x0a, 0x06, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x52, 0x06, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x22, 0x3e,
	0x0a, 0x04, 0x42, 0x61, 0x6c, 0x6c, 0x12, 0x1a, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x03, 0x70,
	0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x03, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x03, 0x76, 0x65, 0x6c, 0x22, 0x3d,
	0x0a, 0x05, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x42, 0x40, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x62, 0x6f,
	0x43, 0x75, 0x70, 0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73, 0x6c, 0x2d, 0x67, 0x61, 0x6d, 0x65,
	0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
}

var (
	file_ssl_gc_engine_proto_rawDescOnce sync.Once
	file_ssl_gc_engine_proto_rawDescData = file_ssl_gc_engine_proto_rawDesc
)

func file_ssl_gc_engine_proto_rawDescGZIP() []byte {
	file_ssl_gc_engine_proto_rawDescOnce.Do(func() {
		file_ssl_gc_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_gc_engine_proto_rawDescData)
	})
	return file_ssl_gc_engine_proto_rawDescData
}

var file_ssl_gc_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ssl_gc_engine_proto_goTypes = []interface{}{
	(*GcState)(nil),        // 0: GcState
	(*GcStateTeam)(nil),    // 1: GcStateTeam
	(*GcStateAutoRef)(nil), // 2: GcStateAutoRef
	(*GcStateTracker)(nil), // 3: GcStateTracker
	(*Ball)(nil),           // 4: Ball
	(*Robot)(nil),          // 5: Robot
	nil,                    // 6: GcState.TeamStateEntry
	nil,                    // 7: GcState.AutoRefStateEntry
	nil,                    // 8: GcState.TrackerStateEntry
	(*geom.Vector3)(nil),   // 9: Vector3
	(*state.RobotId)(nil),  // 10: RobotId
	(*geom.Vector2)(nil),   // 11: Vector2
}
var file_ssl_gc_engine_proto_depIdxs = []int32{
	6,  // 0: GcState.team_state:type_name -> GcState.TeamStateEntry
	7,  // 1: GcState.auto_ref_state:type_name -> GcState.AutoRefStateEntry
	8,  // 2: GcState.tracker_state:type_name -> GcState.TrackerStateEntry
	3,  // 3: GcState.tracker_state_gc:type_name -> GcStateTracker
	4,  // 4: GcStateTracker.ball:type_name -> Ball
	5,  // 5: GcStateTracker.robots:type_name -> Robot
	9,  // 6: Ball.pos:type_name -> Vector3
	9,  // 7: Ball.vel:type_name -> Vector3
	10, // 8: Robot.id:type_name -> RobotId
	11, // 9: Robot.pos:type_name -> Vector2
	1,  // 10: GcState.TeamStateEntry.value:type_name -> GcStateTeam
	2,  // 11: GcState.AutoRefStateEntry.value:type_name -> GcStateAutoRef
	3,  // 12: GcState.TrackerStateEntry.value:type_name -> GcStateTracker
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_ssl_gc_engine_proto_init() }
func file_ssl_gc_engine_proto_init() {
	if File_ssl_gc_engine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ssl_gc_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcState); i {
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
		file_ssl_gc_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcStateTeam); i {
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
		file_ssl_gc_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcStateAutoRef); i {
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
		file_ssl_gc_engine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcStateTracker); i {
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
		file_ssl_gc_engine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ball); i {
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
		file_ssl_gc_engine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Robot); i {
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
			RawDescriptor: file_ssl_gc_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_gc_engine_proto_goTypes,
		DependencyIndexes: file_ssl_gc_engine_proto_depIdxs,
		MessageInfos:      file_ssl_gc_engine_proto_msgTypes,
	}.Build()
	File_ssl_gc_engine_proto = out.File
	file_ssl_gc_engine_proto_rawDesc = nil
	file_ssl_gc_engine_proto_goTypes = nil
	file_ssl_gc_engine_proto_depIdxs = nil
}
