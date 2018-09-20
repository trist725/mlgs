// Code generated by protoc-gen-go. DO NOT EDIT.
// source: money.proto

package msg

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

// / 用户货币信息
type Money struct {
	// / 类型, 1=金币,2=钻石,3=积分,
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	// / 数量
	Num                  int64    `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Money) Reset()         { *m = Money{} }
func (m *Money) String() string { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()    {}
func (*Money) Descriptor() ([]byte, []int) {
	return fileDescriptor_money_658bd7040e078c2a, []int{0}
}
func (m *Money) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Money.Unmarshal(m, b)
}
func (m *Money) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Money.Marshal(b, m, deterministic)
}
func (dst *Money) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Money.Merge(dst, src)
}
func (m *Money) XXX_Size() int {
	return xxx_messageInfo_Money.Size(m)
}
func (m *Money) XXX_DiscardUnknown() {
	xxx_messageInfo_Money.DiscardUnknown(m)
}

var xxx_messageInfo_Money proto.InternalMessageInfo

func (m *Money) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Money) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

// / 更新货币
// @msg
type S2C_UpdateMoney struct {
	// / 发生变更的货币列表
	Monies               []*Money `protobuf:"bytes,1,rep,name=Monies,proto3" json:"Monies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2C_UpdateMoney) Reset()         { *m = S2C_UpdateMoney{} }
func (m *S2C_UpdateMoney) String() string { return proto.CompactTextString(m) }
func (*S2C_UpdateMoney) ProtoMessage()    {}
func (*S2C_UpdateMoney) Descriptor() ([]byte, []int) {
	return fileDescriptor_money_658bd7040e078c2a, []int{1}
}
func (m *S2C_UpdateMoney) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C_UpdateMoney.Unmarshal(m, b)
}
func (m *S2C_UpdateMoney) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C_UpdateMoney.Marshal(b, m, deterministic)
}
func (dst *S2C_UpdateMoney) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_UpdateMoney.Merge(dst, src)
}
func (m *S2C_UpdateMoney) XXX_Size() int {
	return xxx_messageInfo_S2C_UpdateMoney.Size(m)
}
func (m *S2C_UpdateMoney) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_UpdateMoney.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_UpdateMoney proto.InternalMessageInfo

func (m *S2C_UpdateMoney) GetMonies() []*Money {
	if m != nil {
		return m.Monies
	}
	return nil
}

func init() {
	proto.RegisterType((*Money)(nil), "msg.Money")
	proto.RegisterType((*S2C_UpdateMoney)(nil), "msg.S2C_UpdateMoney")
}

func init() { proto.RegisterFile("money.proto", fileDescriptor_money_658bd7040e078c2a) }

var fileDescriptor_money_658bd7040e078c2a = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcd, 0xcf, 0x4b,
	0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0xd2, 0xe5, 0x62,
	0xf5, 0x05, 0x89, 0x09, 0x09, 0x71, 0xb1, 0x84, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0xb0, 0x06, 0x81, 0xd9, 0x42, 0x02, 0x5c, 0xcc, 0x7e, 0xa5, 0xb9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a,
	0xcc, 0x41, 0x20, 0xa6, 0x92, 0x29, 0x17, 0x7f, 0xb0, 0x91, 0x73, 0x7c, 0x68, 0x41, 0x4a, 0x62,
	0x49, 0x2a, 0x44, 0xa3, 0x12, 0x17, 0x9b, 0x6f, 0x7e, 0x5e, 0x66, 0x6a, 0xb1, 0x04, 0xa3, 0x02,
	0xb3, 0x06, 0xb7, 0x11, 0x97, 0x5e, 0x6e, 0x71, 0xba, 0x1e, 0x58, 0x2e, 0x08, 0x2a, 0x93, 0xc4,
	0x06, 0xb6, 0xd1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x3b, 0x43, 0xb8, 0x80, 0x00, 0x00,
	0x00,
}