// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RPi3.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RPi3_Model int32

const (
	RPi3_ThreeB     RPi3_Model = 0
	RPi3_ThreeBPlus RPi3_Model = 1
)

var RPi3_Model_name = map[int32]string{
	0: "ThreeB",
	1: "ThreeBPlus",
}
var RPi3_Model_value = map[string]int32{
	"ThreeB":     0,
	"ThreeBPlus": 1,
}

func (x RPi3_Model) String() string {
	return proto.EnumName(RPi3_Model_name, int32(x))
}
func (RPi3_Model) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_RPi3_82f764f0b2ef2d94, []int{0, 0}
}

type RPi3_PXE int32

const (
	RPi3_NONE RPi3_PXE = 0
	RPi3_WAIT RPi3_PXE = 1
	RPi3_INIT RPi3_PXE = 2
)

var RPi3_PXE_name = map[int32]string{
	0: "NONE",
	1: "WAIT",
	2: "INIT",
}
var RPi3_PXE_value = map[string]int32{
	"NONE": 0,
	"WAIT": 1,
	"INIT": 2,
}

func (x RPi3_PXE) String() string {
	return proto.EnumName(RPi3_PXE_name, int32(x))
}
func (RPi3_PXE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_RPi3_82f764f0b2ef2d94, []int{0, 1}
}

type RPi3 struct {
	Chassis              string     `protobuf:"bytes,1,opt,name=chassis,proto3" json:"chassis,omitempty"`
	Rank                 uint32     `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Model                RPi3_Model `protobuf:"varint,3,opt,name=model,proto3,enum=proto.RPi3_Model" json:"model,omitempty"`
	Pxe                  RPi3_PXE   `protobuf:"varint,4,opt,name=pxe,proto3,enum=proto.RPi3_PXE" json:"pxe,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RPi3) Reset()         { *m = RPi3{} }
func (m *RPi3) String() string { return proto.CompactTextString(m) }
func (*RPi3) ProtoMessage()    {}
func (*RPi3) Descriptor() ([]byte, []int) {
	return fileDescriptor_RPi3_82f764f0b2ef2d94, []int{0}
}
func (m *RPi3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPi3.Unmarshal(m, b)
}
func (m *RPi3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPi3.Marshal(b, m, deterministic)
}
func (dst *RPi3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPi3.Merge(dst, src)
}
func (m *RPi3) XXX_Size() int {
	return xxx_messageInfo_RPi3.Size(m)
}
func (m *RPi3) XXX_DiscardUnknown() {
	xxx_messageInfo_RPi3.DiscardUnknown(m)
}

var xxx_messageInfo_RPi3 proto.InternalMessageInfo

func (m *RPi3) GetChassis() string {
	if m != nil {
		return m.Chassis
	}
	return ""
}

func (m *RPi3) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *RPi3) GetModel() RPi3_Model {
	if m != nil {
		return m.Model
	}
	return RPi3_ThreeB
}

func (m *RPi3) GetPxe() RPi3_PXE {
	if m != nil {
		return m.Pxe
	}
	return RPi3_NONE
}

func init() {
	proto.RegisterType((*RPi3)(nil), "proto.RPi3")
	proto.RegisterEnum("proto.RPi3_Model", RPi3_Model_name, RPi3_Model_value)
	proto.RegisterEnum("proto.RPi3_PXE", RPi3_PXE_name, RPi3_PXE_value)
}

func init() { proto.RegisterFile("RPi3.proto", fileDescriptor_RPi3_82f764f0b2ef2d94) }

var fileDescriptor_RPi3_82f764f0b2ef2d94 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0a, 0x0a, 0xc8, 0x34,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xa7, 0x18, 0xb9, 0x58, 0x40,
	0xa2, 0x42, 0x12, 0x5c, 0xec, 0xc9, 0x19, 0x89, 0xc5, 0xc5, 0x99, 0xc5, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x10, 0x17, 0x4b, 0x51, 0x62, 0x5e, 0xb6, 0x04, 0x93, 0x02,
	0xa3, 0x06, 0x6f, 0x10, 0x98, 0x2d, 0xa4, 0xce, 0xc5, 0x9a, 0x9b, 0x9f, 0x92, 0x9a, 0x23, 0xc1,
	0xac, 0xc0, 0xa8, 0xc1, 0x67, 0x24, 0x08, 0x31, 0x54, 0x0f, 0x6c, 0xbe, 0x2f, 0x48, 0x22, 0x08,
	0x22, 0x2f, 0xa4, 0xc8, 0xc5, 0x5c, 0x50, 0x91, 0x2a, 0xc1, 0x02, 0x56, 0xc6, 0x8f, 0xac, 0x2c,
	0x20, 0xc2, 0x35, 0x08, 0x24, 0xa7, 0xa4, 0xcc, 0xc5, 0x0a, 0xd6, 0x22, 0xc4, 0xc5, 0xc5, 0x16,
	0x92, 0x51, 0x94, 0x9a, 0xea, 0x24, 0xc0, 0x20, 0xc4, 0xc7, 0xc5, 0x05, 0x61, 0x07, 0xe4, 0x94,
	0x16, 0x0b, 0x30, 0x2a, 0x29, 0x73, 0x31, 0x07, 0x44, 0xb8, 0x0a, 0x71, 0x70, 0xb1, 0xf8, 0xf9,
	0xfb, 0xb9, 0x0a, 0x30, 0x80, 0x58, 0xe1, 0x8e, 0x9e, 0x21, 0x02, 0x8c, 0x20, 0x96, 0xa7, 0x9f,
	0x67, 0x88, 0x00, 0x53, 0x12, 0x1b, 0xd8, 0x78, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33,
	0x7c, 0xa8, 0xa6, 0xe8, 0x00, 0x00, 0x00,
}
