// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test.proto

package proto

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FOO int32

const (
	FOO_X FOO = 0
)

var FOO_name = map[int32]string{
	0: "X",
}
var FOO_value = map[string]int32{
	"X": 0,
}

func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}
func (FOO) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_test_c15113eb743169da, []int{0}
}

// message是固定的。UserInfo是类名，可以随意指定，符合规范即可
type UserInfo struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Length               int32    `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	Cnt                  int32    `protobuf:"varint,3,opt,name=cnt,proto3" json:"cnt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_c15113eb743169da, []int{0}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (dst *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(dst, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UserInfo) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *UserInfo) GetCnt() int32 {
	if m != nil {
		return m.Cnt
	}
	return 0
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "proto.UserInfo")
	proto.RegisterEnum("proto.FOO", FOO_name, FOO_value)
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_c15113eb743169da) }

var fileDescriptor_test_c15113eb743169da = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x7e, 0x5c, 0x1c, 0xa1, 0xc5,
	0xa9, 0x45, 0x9e, 0x79, 0x69, 0xf9, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9,
	0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x18, 0x17, 0x5b, 0x4e, 0x6a,
	0x5e, 0x7a, 0x49, 0x86, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x94, 0x27, 0x24, 0xc0, 0xc5,
	0x9c, 0x9c, 0x57, 0x22, 0xc1, 0x0c, 0x16, 0x04, 0x31, 0xb5, 0x78, 0xb8, 0x98, 0xdd, 0xfc, 0xfd,
	0x85, 0x58, 0xb9, 0x18, 0x23, 0x04, 0x18, 0x92, 0xd8, 0xc0, 0x96, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x27, 0x2c, 0x8f, 0x3c, 0x79, 0x00, 0x00, 0x00,
}
