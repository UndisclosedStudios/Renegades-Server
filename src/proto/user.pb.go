// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/proto.proto

package user_protobuf

import (
	fmt "fmt"
	proto "github.com/golang/proto/proto"
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

type Point struct {
	Username             string   `proto:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_198e8835fe62d806, []int{0}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*Point)(nil), "proto.informations.Point")
}

func init() {
	proto.RegisterFile("proto/proto.proto", fileDescriptor_198e8835fe62d806)
}

var fileDescriptor_198e8835fe62d806 = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x03, 0xf3, 0x84, 0x04, 0xc1, 0xec,
	0xcc, 0xbc, 0xb4, 0xfc, 0xa2, 0xdc, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x25, 0x65, 0x2e, 0xd6,
	0x80, 0xfc, 0xcc, 0xbc, 0x12, 0x21, 0x29, 0x2e, 0x0e, 0x90, 0x6c, 0x5e, 0x62, 0x6e, 0xaa, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0xef, 0xc4, 0x1f, 0xc5, 0x0b, 0x62, 0xc7, 0xc3, 0xcc,
	0x4c, 0x62, 0x03, 0xb3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x52, 0xfe, 0x94, 0x66,
	0x00, 0x00, 0x00,
}
