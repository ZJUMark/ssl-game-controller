// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_vision_detection.proto

package vision

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

type SSL_DetectionBall struct {
	Confidence           *float32 `protobuf:"fixed32,1,req,name=confidence" json:"confidence,omitempty"`
	Area                 *uint32  `protobuf:"varint,2,opt,name=area" json:"area,omitempty"`
	X                    *float32 `protobuf:"fixed32,3,req,name=x" json:"x,omitempty"`
	Y                    *float32 `protobuf:"fixed32,4,req,name=y" json:"y,omitempty"`
	Z                    *float32 `protobuf:"fixed32,5,opt,name=z" json:"z,omitempty"`
	PixelX               *float32 `protobuf:"fixed32,6,req,name=pixel_x,json=pixelX" json:"pixel_x,omitempty"`
	PixelY               *float32 `protobuf:"fixed32,7,req,name=pixel_y,json=pixelY" json:"pixel_y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSL_DetectionBall) Reset()         { *m = SSL_DetectionBall{} }
func (m *SSL_DetectionBall) String() string { return proto.CompactTextString(m) }
func (*SSL_DetectionBall) ProtoMessage()    {}
func (*SSL_DetectionBall) Descriptor() ([]byte, []int) {
	return fileDescriptor_89b1c3228d8f9a76, []int{0}
}

func (m *SSL_DetectionBall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSL_DetectionBall.Unmarshal(m, b)
}
func (m *SSL_DetectionBall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSL_DetectionBall.Marshal(b, m, deterministic)
}
func (m *SSL_DetectionBall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSL_DetectionBall.Merge(m, src)
}
func (m *SSL_DetectionBall) XXX_Size() int {
	return xxx_messageInfo_SSL_DetectionBall.Size(m)
}
func (m *SSL_DetectionBall) XXX_DiscardUnknown() {
	xxx_messageInfo_SSL_DetectionBall.DiscardUnknown(m)
}

var xxx_messageInfo_SSL_DetectionBall proto.InternalMessageInfo

func (m *SSL_DetectionBall) GetConfidence() float32 {
	if m != nil && m.Confidence != nil {
		return *m.Confidence
	}
	return 0
}

func (m *SSL_DetectionBall) GetArea() uint32 {
	if m != nil && m.Area != nil {
		return *m.Area
	}
	return 0
}

func (m *SSL_DetectionBall) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *SSL_DetectionBall) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *SSL_DetectionBall) GetZ() float32 {
	if m != nil && m.Z != nil {
		return *m.Z
	}
	return 0
}

func (m *SSL_DetectionBall) GetPixelX() float32 {
	if m != nil && m.PixelX != nil {
		return *m.PixelX
	}
	return 0
}

func (m *SSL_DetectionBall) GetPixelY() float32 {
	if m != nil && m.PixelY != nil {
		return *m.PixelY
	}
	return 0
}

type SSL_DetectionRobot struct {
	Confidence           *float32 `protobuf:"fixed32,1,req,name=confidence" json:"confidence,omitempty"`
	RobotId              *uint32  `protobuf:"varint,2,opt,name=robot_id,json=robotId" json:"robot_id,omitempty"`
	X                    *float32 `protobuf:"fixed32,3,req,name=x" json:"x,omitempty"`
	Y                    *float32 `protobuf:"fixed32,4,req,name=y" json:"y,omitempty"`
	Orientation          *float32 `protobuf:"fixed32,5,opt,name=orientation" json:"orientation,omitempty"`
	PixelX               *float32 `protobuf:"fixed32,6,req,name=pixel_x,json=pixelX" json:"pixel_x,omitempty"`
	PixelY               *float32 `protobuf:"fixed32,7,req,name=pixel_y,json=pixelY" json:"pixel_y,omitempty"`
	Height               *float32 `protobuf:"fixed32,8,opt,name=height" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSL_DetectionRobot) Reset()         { *m = SSL_DetectionRobot{} }
func (m *SSL_DetectionRobot) String() string { return proto.CompactTextString(m) }
func (*SSL_DetectionRobot) ProtoMessage()    {}
func (*SSL_DetectionRobot) Descriptor() ([]byte, []int) {
	return fileDescriptor_89b1c3228d8f9a76, []int{1}
}

func (m *SSL_DetectionRobot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSL_DetectionRobot.Unmarshal(m, b)
}
func (m *SSL_DetectionRobot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSL_DetectionRobot.Marshal(b, m, deterministic)
}
func (m *SSL_DetectionRobot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSL_DetectionRobot.Merge(m, src)
}
func (m *SSL_DetectionRobot) XXX_Size() int {
	return xxx_messageInfo_SSL_DetectionRobot.Size(m)
}
func (m *SSL_DetectionRobot) XXX_DiscardUnknown() {
	xxx_messageInfo_SSL_DetectionRobot.DiscardUnknown(m)
}

var xxx_messageInfo_SSL_DetectionRobot proto.InternalMessageInfo

func (m *SSL_DetectionRobot) GetConfidence() float32 {
	if m != nil && m.Confidence != nil {
		return *m.Confidence
	}
	return 0
}

func (m *SSL_DetectionRobot) GetRobotId() uint32 {
	if m != nil && m.RobotId != nil {
		return *m.RobotId
	}
	return 0
}

func (m *SSL_DetectionRobot) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *SSL_DetectionRobot) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *SSL_DetectionRobot) GetOrientation() float32 {
	if m != nil && m.Orientation != nil {
		return *m.Orientation
	}
	return 0
}

func (m *SSL_DetectionRobot) GetPixelX() float32 {
	if m != nil && m.PixelX != nil {
		return *m.PixelX
	}
	return 0
}

func (m *SSL_DetectionRobot) GetPixelY() float32 {
	if m != nil && m.PixelY != nil {
		return *m.PixelY
	}
	return 0
}

func (m *SSL_DetectionRobot) GetHeight() float32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

type SSL_DetectionFrame struct {
	FrameNumber          *uint32               `protobuf:"varint,1,req,name=frame_number,json=frameNumber" json:"frame_number,omitempty"`
	TCapture             *float64              `protobuf:"fixed64,2,req,name=t_capture,json=tCapture" json:"t_capture,omitempty"`
	TSent                *float64              `protobuf:"fixed64,3,req,name=t_sent,json=tSent" json:"t_sent,omitempty"`
	CameraId             *uint32               `protobuf:"varint,4,req,name=camera_id,json=cameraId" json:"camera_id,omitempty"`
	Balls                []*SSL_DetectionBall  `protobuf:"bytes,5,rep,name=balls" json:"balls,omitempty"`
	RobotsYellow         []*SSL_DetectionRobot `protobuf:"bytes,6,rep,name=robots_yellow,json=robotsYellow" json:"robots_yellow,omitempty"`
	RobotsBlue           []*SSL_DetectionRobot `protobuf:"bytes,7,rep,name=robots_blue,json=robotsBlue" json:"robots_blue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SSL_DetectionFrame) Reset()         { *m = SSL_DetectionFrame{} }
func (m *SSL_DetectionFrame) String() string { return proto.CompactTextString(m) }
func (*SSL_DetectionFrame) ProtoMessage()    {}
func (*SSL_DetectionFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_89b1c3228d8f9a76, []int{2}
}

func (m *SSL_DetectionFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSL_DetectionFrame.Unmarshal(m, b)
}
func (m *SSL_DetectionFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSL_DetectionFrame.Marshal(b, m, deterministic)
}
func (m *SSL_DetectionFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSL_DetectionFrame.Merge(m, src)
}
func (m *SSL_DetectionFrame) XXX_Size() int {
	return xxx_messageInfo_SSL_DetectionFrame.Size(m)
}
func (m *SSL_DetectionFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_SSL_DetectionFrame.DiscardUnknown(m)
}

var xxx_messageInfo_SSL_DetectionFrame proto.InternalMessageInfo

func (m *SSL_DetectionFrame) GetFrameNumber() uint32 {
	if m != nil && m.FrameNumber != nil {
		return *m.FrameNumber
	}
	return 0
}

func (m *SSL_DetectionFrame) GetTCapture() float64 {
	if m != nil && m.TCapture != nil {
		return *m.TCapture
	}
	return 0
}

func (m *SSL_DetectionFrame) GetTSent() float64 {
	if m != nil && m.TSent != nil {
		return *m.TSent
	}
	return 0
}

func (m *SSL_DetectionFrame) GetCameraId() uint32 {
	if m != nil && m.CameraId != nil {
		return *m.CameraId
	}
	return 0
}

func (m *SSL_DetectionFrame) GetBalls() []*SSL_DetectionBall {
	if m != nil {
		return m.Balls
	}
	return nil
}

func (m *SSL_DetectionFrame) GetRobotsYellow() []*SSL_DetectionRobot {
	if m != nil {
		return m.RobotsYellow
	}
	return nil
}

func (m *SSL_DetectionFrame) GetRobotsBlue() []*SSL_DetectionRobot {
	if m != nil {
		return m.RobotsBlue
	}
	return nil
}

func init() {
	proto.RegisterType((*SSL_DetectionBall)(nil), "SSL_DetectionBall")
	proto.RegisterType((*SSL_DetectionRobot)(nil), "SSL_DetectionRobot")
	proto.RegisterType((*SSL_DetectionFrame)(nil), "SSL_DetectionFrame")
}

func init() {
	proto.RegisterFile("ssl_vision_detection.proto", fileDescriptor_89b1c3228d8f9a76)
}

var fileDescriptor_89b1c3228d8f9a76 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0xe5, 0x6c, 0x9b, 0x96, 0x69, 0x7b, 0xc0, 0x08, 0x30, 0x20, 0xa1, 0xd0, 0x53, 0x2e,
	0x6d, 0x24, 0xc4, 0x81, 0x13, 0x42, 0x5d, 0x84, 0xb4, 0xd2, 0x8a, 0x43, 0x72, 0x61, 0xb9, 0x58,
	0x4e, 0x32, 0xdb, 0x5a, 0x72, 0xec, 0xc8, 0x71, 0xa0, 0xdd, 0xc7, 0x81, 0x87, 0xe2, 0x75, 0x90,
	0x9d, 0x22, 0x0a, 0x2b, 0x40, 0x7b, 0xf3, 0xf7, 0xcf, 0x3f, 0xce, 0x3f, 0xce, 0xc0, 0xd3, 0xae,
	0x53, 0xfc, 0xb3, 0xec, 0xa4, 0xd1, 0xbc, 0x46, 0x87, 0x95, 0x93, 0x46, 0xaf, 0x5b, 0x6b, 0x9c,
	0x59, 0x7e, 0x23, 0x70, 0xbf, 0x28, 0x2e, 0xf9, 0xbb, 0x9f, 0xfa, 0x46, 0x28, 0x45, 0x9f, 0x03,
	0x54, 0x46, 0x5f, 0xcb, 0x1a, 0x75, 0x85, 0x8c, 0x24, 0x51, 0x1a, 0xe5, 0x27, 0x0a, 0xa5, 0x30,
	0x12, 0x16, 0x05, 0x8b, 0x12, 0x92, 0x2e, 0xf2, 0x70, 0xa6, 0x73, 0x20, 0x7b, 0x76, 0x16, 0xac,
	0x64, 0xef, 0xe9, 0xc0, 0x46, 0x03, 0x1d, 0x3c, 0xdd, 0xb0, 0x71, 0x42, 0x3c, 0xdd, 0xd0, 0xc7,
	0x30, 0x69, 0xe5, 0x1e, 0x15, 0xdf, 0xb3, 0x38, 0x38, 0xe2, 0x80, 0x1f, 0x7f, 0x15, 0x0e, 0x6c,
	0x72, 0x52, 0xb8, 0x5a, 0x7e, 0x27, 0x40, 0x7f, 0x4b, 0x99, 0x9b, 0xd2, 0xb8, 0xff, 0xc6, 0x7c,
	0x02, 0x53, 0xeb, 0x8d, 0x5c, 0xd6, 0xc7, 0xa8, 0x93, 0xc0, 0x17, 0xf5, 0x3f, 0xd3, 0x26, 0x30,
	0x33, 0x56, 0xa2, 0x76, 0xc2, 0x7f, 0xea, 0x98, 0xfb, 0x54, 0xba, 0xfb, 0x04, 0xf4, 0x11, 0xc4,
	0x3b, 0x94, 0xdb, 0x9d, 0x63, 0xd3, 0x70, 0xdd, 0x91, 0x96, 0x5f, 0xa3, 0x3f, 0x26, 0x7b, 0x6f,
	0x45, 0x83, 0xf4, 0x05, 0xcc, 0xaf, 0xfd, 0x81, 0xeb, 0xbe, 0x29, 0xd1, 0x86, 0xd9, 0x16, 0xf9,
	0x2c, 0x68, 0x1f, 0x82, 0x44, 0x9f, 0xc1, 0x3d, 0xc7, 0x2b, 0xd1, 0xba, 0xde, 0x22, 0x8b, 0x92,
	0x28, 0x25, 0xf9, 0xd4, 0x9d, 0x0f, 0x4c, 0x1f, 0x42, 0xec, 0x78, 0x87, 0xda, 0x85, 0x19, 0x49,
	0x3e, 0x76, 0x05, 0x6a, 0xe7, 0x7b, 0x2a, 0xd1, 0xa0, 0x15, 0xfe, 0x45, 0x46, 0xe1, 0xce, 0xe9,
	0x20, 0x5c, 0xd4, 0x34, 0x85, 0x71, 0x29, 0x94, 0xea, 0xd8, 0x38, 0x39, 0x4b, 0x67, 0x2f, 0xe9,
	0xfa, 0xd6, 0x5e, 0xe4, 0x83, 0x81, 0xbe, 0x86, 0x45, 0x78, 0xc7, 0x8e, 0x1f, 0x50, 0x29, 0xf3,
	0x85, 0xc5, 0xa1, 0xe3, 0xc1, 0xfa, 0xf6, 0x3f, 0xca, 0xe7, 0x83, 0xf3, 0x2a, 0x18, 0xe9, 0x2b,
	0x98, 0x1d, 0x3b, 0x4b, 0xd5, 0x23, 0x9b, 0xfc, 0xbd, 0x0f, 0x06, 0xdf, 0x46, 0xf5, 0xb8, 0x79,
	0xfb, 0xe9, 0xcd, 0x56, 0xba, 0x5d, 0x5f, 0xae, 0x2b, 0xd3, 0x64, 0xbe, 0x7e, 0xde, 0xb7, 0xab,
	0xa2, 0xb8, 0xcc, 0xba, 0x4e, 0xad, 0xb6, 0xa2, 0xc1, 0x55, 0x65, 0xb4, 0xb3, 0x46, 0x29, 0xb4,
	0x99, 0xd4, 0x0e, 0xad, 0x16, 0x2a, 0x13, 0x6d, 0x9b, 0x0d, 0x6b, 0xff, 0x23, 0x00, 0x00, 0xff,
	0xff, 0xb8, 0x16, 0x77, 0xa3, 0x03, 0x03, 0x00, 0x00,
}
