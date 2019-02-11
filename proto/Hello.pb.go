// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Hello.proto

package Hello

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

type GetMyCatMessage struct {
	TargetCat            string   `protobuf:"bytes,1,opt,name=target_cat,json=targetCat,proto3" json:"target_cat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMyCatMessage) Reset()         { *m = GetMyCatMessage{} }
func (m *GetMyCatMessage) String() string { return proto.CompactTextString(m) }
func (*GetMyCatMessage) ProtoMessage()    {}
func (*GetMyCatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_Hello_fc0fcd1d18fc60dd, []int{0}
}
func (m *GetMyCatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMyCatMessage.Unmarshal(m, b)
}
func (m *GetMyCatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMyCatMessage.Marshal(b, m, deterministic)
}
func (dst *GetMyCatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMyCatMessage.Merge(dst, src)
}
func (m *GetMyCatMessage) XXX_Size() int {
	return xxx_messageInfo_GetMyCatMessage.Size(m)
}
func (m *GetMyCatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMyCatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetMyCatMessage proto.InternalMessageInfo

func (m *GetMyCatMessage) GetTargetCat() string {
	if m != nil {
		return m.TargetCat
	}
	return ""
}

type MyCatResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind                 string   `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyCatResponse) Reset()         { *m = MyCatResponse{} }
func (m *MyCatResponse) String() string { return proto.CompactTextString(m) }
func (*MyCatResponse) ProtoMessage()    {}
func (*MyCatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_Hello_fc0fcd1d18fc60dd, []int{1}
}
func (m *MyCatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyCatResponse.Unmarshal(m, b)
}
func (m *MyCatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyCatResponse.Marshal(b, m, deterministic)
}
func (dst *MyCatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyCatResponse.Merge(dst, src)
}
func (m *MyCatResponse) XXX_Size() int {
	return xxx_messageInfo_MyCatResponse.Size(m)
}
func (m *MyCatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MyCatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MyCatResponse proto.InternalMessageInfo

func (m *MyCatResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MyCatResponse) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMyCatMessage)(nil), "GetMyCatMessage")
	proto.RegisterType((*MyCatResponse)(nil), "MyCatResponse")
}

func init() { proto.RegisterFile("Hello.proto", fileDescriptor_Hello_fc0fcd1d18fc60dd) }

var fileDescriptor_Hello_fc0fcd1d18fc60dd = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf6, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe0, 0xe2, 0x77, 0x4f, 0x2d, 0xf1, 0xad,
	0x74, 0x4e, 0x2c, 0xf1, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x92, 0xe5, 0xe2, 0x2a, 0x49,
	0x2c, 0x4a, 0x4f, 0x2d, 0x89, 0x4f, 0x4e, 0x2c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2,
	0x84, 0x88, 0x38, 0x27, 0x96, 0x28, 0x99, 0x73, 0xf1, 0x82, 0x95, 0x07, 0xa5, 0x16, 0x17, 0xe4,
	0xe7, 0x15, 0xa7, 0x0a, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x42, 0x55, 0x82, 0xd9, 0x20,
	0xb1, 0xec, 0xcc, 0xbc, 0x14, 0x09, 0x26, 0x88, 0x18, 0x88, 0x6d, 0x64, 0xca, 0xc5, 0xec, 0x9c,
	0x58, 0x22, 0xa4, 0xc7, 0xc5, 0x01, 0xb3, 0x51, 0x48, 0x40, 0x0f, 0xcd, 0x72, 0x29, 0x3e, 0x3d,
	0x14, 0xc3, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x0e, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6c,
	0x89, 0x2e, 0x6c, 0xb7, 0x00, 0x00, 0x00,
}
