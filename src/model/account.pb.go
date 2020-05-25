// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: account.proto

package model

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

/// 帐号表
//@collection
type Account struct {
	/// mongodb默认主键_id做账号id
	//@bson=_id
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	/// 客户端传来的唯一id,如微信unionID
	UID string `protobuf:"bytes,2,opt,name=UID,proto3" json:"UID,omitempty"`
	/// 密码
	//string Password = 4;
	/// 注册时间
	RegisterTime int64 `protobuf:"varint,5,opt,name=RegisterTime,proto3" json:"RegisterTime,omitempty"`
	///token校验
	//    string accessToken = 6;
	//    string refreshToken = 7;
	///登陆地理位置
	Location string `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	///密码
	Password string `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
	/// 帐号状态 1=游客, 2=注册, 3=绑定
	//E_AccountState State = 6;
	/// 密钥
	//string Token = 7;
	/// 上次服务器ID
	//int32 LastLoginServerID = 8;
	/// 登录过服务器ID列表
	//repeated int32 LoginList = 9;
	/// 渠道名
	//string ChannelName = 10;
	/// 渠道帐号
	//string ChannelAccount = 11;
	/// 封号标记,1为被封
	Ban int32 `protobuf:"varint,12,opt,name=Ban,proto3" json:"Ban,omitempty"`
	///游客/微信/万博
	Type int32 `protobuf:"varint,13,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Account) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *Account) GetRegisterTime() int64 {
	if m != nil {
		return m.RegisterTime
	}
	return 0
}

func (m *Account) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Account) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Account) GetBan() int32 {
	if m != nil {
		return m.Ban
	}
	return 0
}

func (m *Account) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*Account)(nil), "model.Account")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd,
	0x51, 0x5a, 0xcb, 0xc8, 0xc5, 0xee, 0x08, 0x91, 0x10, 0xe2, 0xe3, 0x62, 0xf2, 0x74, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0xf2, 0x74, 0x11, 0x12, 0xe0, 0x62, 0x0e, 0xf5, 0x74, 0x91,
	0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x94, 0xb8, 0x78, 0x82, 0x52, 0xd3, 0x33,
	0x8b, 0x4b, 0x52, 0x8b, 0x42, 0x32, 0x73, 0x53, 0x25, 0x58, 0xc1, 0x6a, 0x51, 0xc4, 0x84, 0xa4,
	0xb8, 0x38, 0x72, 0xf2, 0x93, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0x24, 0x38, 0xc0, 0x5a, 0xe1, 0x7c,
	0x90, 0x5c, 0x41, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51, 0x8a, 0x04, 0x27, 0x44, 0x0e, 0xc6, 0x07,
	0xd9, 0xe6, 0x94, 0x98, 0x27, 0xc1, 0xa3, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0x62, 0x0a, 0x09, 0x71,
	0xb1, 0x84, 0x54, 0x16, 0xa4, 0x4a, 0xf0, 0x82, 0x85, 0xc0, 0x6c, 0x27, 0x89, 0x13, 0x8f, 0xe4,
	0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f,
	0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0xfb, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x14, 0x10, 0xa1, 0xab, 0xe8, 0x00, 0x00, 0x00,
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x68
	}
	if m.Ban != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.Ban))
		i--
		dAtA[i] = 0x60
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Location) > 0 {
		i -= len(m.Location)
		copy(dAtA[i:], m.Location)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Location)))
		i--
		dAtA[i] = 0x42
	}
	if m.RegisterTime != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.RegisterTime))
		i--
		dAtA[i] = 0x28
	}
	if len(m.UID) > 0 {
		i -= len(m.UID)
		copy(dAtA[i:], m.UID)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.UID)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovAccount(uint64(m.ID))
	}
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.RegisterTime != 0 {
		n += 1 + sovAccount(uint64(m.RegisterTime))
	}
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.Ban != 0 {
		n += 1 + sovAccount(uint64(m.Ban))
	}
	if m.Type != 0 {
		n += 1 + sovAccount(uint64(m.Type))
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisterTime", wireType)
			}
			m.RegisterTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegisterTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ban", wireType)
			}
			m.Ban = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ban |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccount = fmt.Errorf("proto: unexpected end of group")
)