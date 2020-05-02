// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_gc_common.proto

package state

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

// Team is either blue or yellow
type Team int32

const (
	// team not set
	Team_UNKNOWN Team = 0
	// yellow team
	Team_YELLOW Team = 1
	// blue team
	Team_BLUE Team = 2
)

var Team_name = map[int32]string{
	0: "UNKNOWN",
	1: "YELLOW",
	2: "BLUE",
}

var Team_value = map[string]int32{
	"UNKNOWN": 0,
	"YELLOW":  1,
	"BLUE":    2,
}

func (x Team) Enum() *Team {
	p := new(Team)
	*p = x
	return p
}

func (x Team) String() string {
	return proto.EnumName(Team_name, int32(x))
}

func (x *Team) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Team_value, data, "Team")
	if err != nil {
		return err
	}
	*x = Team(value)
	return nil
}

func (Team) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1df1d8f3f755967a, []int{0}
}

// Division denotes the current division, which influences some rules
type Division int32

const (
	Division_DIV_UNKNOWN Division = 0
	Division_DIV_A       Division = 1
	Division_DIV_B       Division = 2
)

var Division_name = map[int32]string{
	0: "DIV_UNKNOWN",
	1: "DIV_A",
	2: "DIV_B",
}

var Division_value = map[string]int32{
	"DIV_UNKNOWN": 0,
	"DIV_A":       1,
	"DIV_B":       2,
}

func (x Division) Enum() *Division {
	p := new(Division)
	*p = x
	return p
}

func (x Division) String() string {
	return proto.EnumName(Division_name, int32(x))
}

func (x *Division) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Division_value, data, "Division")
	if err != nil {
		return err
	}
	*x = Division(value)
	return nil
}

func (Division) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1df1d8f3f755967a, []int{1}
}

// RobotId is the combination of a team and a robot id
type RobotId struct {
	// the robot number
	Id *uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// the team that the robot belongs to
	Team                 *Team    `protobuf:"varint,2,opt,name=team,enum=Team" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RobotId) Reset()         { *m = RobotId{} }
func (m *RobotId) String() string { return proto.CompactTextString(m) }
func (*RobotId) ProtoMessage()    {}
func (*RobotId) Descriptor() ([]byte, []int) {
	return fileDescriptor_1df1d8f3f755967a, []int{0}
}

func (m *RobotId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RobotId.Unmarshal(m, b)
}
func (m *RobotId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RobotId.Marshal(b, m, deterministic)
}
func (m *RobotId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RobotId.Merge(m, src)
}
func (m *RobotId) XXX_Size() int {
	return xxx_messageInfo_RobotId.Size(m)
}
func (m *RobotId) XXX_DiscardUnknown() {
	xxx_messageInfo_RobotId.DiscardUnknown(m)
}

var xxx_messageInfo_RobotId proto.InternalMessageInfo

func (m *RobotId) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *RobotId) GetTeam() Team {
	if m != nil && m.Team != nil {
		return *m.Team
	}
	return Team_UNKNOWN
}

func init() {
	proto.RegisterEnum("Team", Team_name, Team_value)
	proto.RegisterEnum("Division", Division_name, Division_value)
	proto.RegisterType((*RobotId)(nil), "RobotId")
}

func init() {
	proto.RegisterFile("ssl_gc_common.proto", fileDescriptor_1df1d8f3f755967a)
}

var fileDescriptor_1df1d8f3f755967a = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xce, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0x05, 0x70, 0x13, 0x3a, 0x37, 0xbf, 0xe1, 0x0c, 0xf1, 0x32, 0x6f, 0xc3, 0xd3, 0x1c, 0xb4,
	0x41, 0xf1, 0x2a, 0x62, 0xdd, 0x0e, 0xc3, 0xd2, 0x41, 0xe7, 0x1c, 0x7a, 0x29, 0x59, 0x1b, 0x6a,
	0x20, 0xc9, 0x57, 0x9a, 0xcc, 0xbf, 0x5f, 0x2a, 0xc8, 0x6e, 0xef, 0xfd, 0xe0, 0xc1, 0x83, 0x6b,
	0xef, 0x4d, 0xd9, 0x54, 0x65, 0x85, 0xd6, 0xa2, 0x4b, 0xda, 0x0e, 0x03, 0xde, 0x3e, 0xc2, 0xb0,
	0xc0, 0x03, 0x86, 0x75, 0xcd, 0x27, 0x40, 0x75, 0x3d, 0x25, 0x33, 0x32, 0xbf, 0x2c, 0xa8, 0xae,
	0xf9, 0x0d, 0x44, 0x41, 0x49, 0x3b, 0xa5, 0x33, 0x32, 0x9f, 0x3c, 0x0c, 0x92, 0x77, 0x25, 0x6d,
	0xf1, 0x47, 0x8b, 0x3b, 0x88, 0xfa, 0xc6, 0xc7, 0x30, 0xdc, 0xe5, 0x6f, 0xf9, 0x66, 0x9f, 0xb3,
	0x33, 0x0e, 0x70, 0xfe, 0xb9, 0xca, 0xb2, 0xcd, 0x9e, 0x11, 0x3e, 0x82, 0x28, 0xcd, 0x76, 0x2b,
	0x46, 0x17, 0xf7, 0x30, 0x5a, 0xea, 0x1f, 0xed, 0x35, 0x3a, 0x7e, 0x05, 0xe3, 0xe5, 0xfa, 0xa3,
	0x3c, 0x4d, 0x2e, 0x60, 0xd0, 0xc3, 0x0b, 0x23, 0xff, 0x31, 0x65, 0x34, 0x7d, 0xfe, 0x7a, 0x6a,
	0x74, 0xf8, 0x3e, 0x1e, 0x92, 0x0a, 0xad, 0xe8, 0xef, 0xbd, 0x1e, 0xdb, 0x78, 0xbb, 0xcd, 0x84,
	0xf7, 0x26, 0x6e, 0xa4, 0x55, 0x71, 0x85, 0x2e, 0x74, 0x68, 0x8c, 0xea, 0x84, 0x76, 0x41, 0x75,
	0x4e, 0x1a, 0x21, 0xdb, 0x56, 0xf8, 0x20, 0x83, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x33,
	0x9b, 0x52, 0xea, 0x00, 0x00, 0x00,
}
