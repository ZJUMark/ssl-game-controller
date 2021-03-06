// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_vision_wrapper_tracked.proto

package tracker

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A wrapper packet containing meta data of the source
// Also serves for the possibility to extend the protocol later
type TrackerWrapperPacket struct {
	// A random UUID of the source that is kept constant at the source while running
	// If multiple sources are broadcasting to the same network, this id can be used to identify individual sources
	Uuid *string `protobuf:"bytes,1,req,name=uuid" json:"uuid,omitempty"`
	// The name of the source software that is producing this messages.
	SourceName *string `protobuf:"bytes,2,opt,name=source_name,json=sourceName" json:"source_name,omitempty"`
	// The tracked frame
	TrackedFrame         *TrackedFrame `protobuf:"bytes,3,opt,name=tracked_frame,json=trackedFrame" json:"tracked_frame,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TrackerWrapperPacket) Reset()         { *m = TrackerWrapperPacket{} }
func (m *TrackerWrapperPacket) String() string { return proto.CompactTextString(m) }
func (*TrackerWrapperPacket) ProtoMessage()    {}
func (*TrackerWrapperPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e84f35cc8d505696, []int{0}
}

func (m *TrackerWrapperPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackerWrapperPacket.Unmarshal(m, b)
}
func (m *TrackerWrapperPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackerWrapperPacket.Marshal(b, m, deterministic)
}
func (m *TrackerWrapperPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackerWrapperPacket.Merge(m, src)
}
func (m *TrackerWrapperPacket) XXX_Size() int {
	return xxx_messageInfo_TrackerWrapperPacket.Size(m)
}
func (m *TrackerWrapperPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackerWrapperPacket.DiscardUnknown(m)
}

var xxx_messageInfo_TrackerWrapperPacket proto.InternalMessageInfo

func (m *TrackerWrapperPacket) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *TrackerWrapperPacket) GetSourceName() string {
	if m != nil && m.SourceName != nil {
		return *m.SourceName
	}
	return ""
}

func (m *TrackerWrapperPacket) GetTrackedFrame() *TrackedFrame {
	if m != nil {
		return m.TrackedFrame
	}
	return nil
}

func init() {
	proto.RegisterType((*TrackerWrapperPacket)(nil), "TrackerWrapperPacket")
}

func init() {
	proto.RegisterFile("ssl_vision_wrapper_tracked.proto", fileDescriptor_e84f35cc8d505696)
}

var fileDescriptor_e84f35cc8d505696 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8d, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x46, 0xd9, 0xd3, 0xe6, 0x72, 0x5e, 0xb3, 0x58, 0x1c, 0x36, 0x2e, 0x57, 0x6d, 0xb3, 0x1b,
	0xb8, 0x3f, 0x20, 0x2a, 0x58, 0x89, 0xc8, 0x9e, 0x20, 0xd8, 0x84, 0x5c, 0x76, 0x3c, 0x83, 0xd9,
	0x4c, 0x98, 0x4c, 0xb4, 0xf4, 0xaf, 0x4b, 0x2e, 0x16, 0x6b, 0x37, 0xdf, 0xe3, 0xf1, 0x46, 0x34,
	0x31, 0x3a, 0xf5, 0x65, 0xa3, 0x45, 0xaf, 0xbe, 0x49, 0x87, 0x00, 0xa4, 0x98, 0xb4, 0xf9, 0x84,
	0xb1, 0x0f, 0x84, 0x8c, 0x57, 0xdb, 0x99, 0x31, 0x02, 0x83, 0xe1, 0x7c, 0xfd, 0x73, 0xb6, 0x3f,
	0xe2, 0xf2, 0xe5, 0x04, 0xe8, 0xb5, 0x34, 0x9e, 0xf3, 0xe0, 0xba, 0x16, 0xe7, 0x29, 0xd9, 0x71,
	0x53, 0x35, 0x8b, 0x76, 0x39, 0x9c, 0xee, 0xfa, 0x5a, 0xac, 0x22, 0x26, 0x32, 0xa0, 0xbc, 0x9e,
	0x60, 0xb3, 0x68, 0xaa, 0x76, 0x39, 0x88, 0x82, 0x9e, 0xf4, 0x04, 0xf5, 0x4e, 0xac, 0xff, 0xea,
	0xea, 0x9d, 0xb2, 0x72, 0xd6, 0x54, 0xed, 0x6a, 0xb7, 0xee, 0xcb, 0x8b, 0xf1, 0x21, 0xc3, 0xe1,
	0x82, 0x67, 0xeb, 0xee, 0xf6, 0xed, 0xe6, 0x68, 0xf9, 0x23, 0x1d, 0x7a, 0x83, 0x93, 0x1c, 0xf0,
	0x80, 0xf7, 0x29, 0x74, 0xfb, 0xfd, 0xa3, 0x8c, 0xd1, 0x75, 0x47, 0x3d, 0x41, 0x67, 0xd0, 0x33,
	0xa1, 0x73, 0x40, 0xd2, 0x7a, 0x06, 0xf2, 0xda, 0x49, 0x1d, 0x82, 0x2c, 0x1d, 0xfa, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xed, 0x36, 0x8a, 0x0d, 0x0a, 0x01, 0x00, 0x00,
}
