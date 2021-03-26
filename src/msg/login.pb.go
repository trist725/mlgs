// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: login.proto

package msg

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

///登录结果
type S2C_Login_E_Error int32

const (
	S2C_Login_Error_ S2C_Login_E_Error = 0
	/// 成功
	S2C_Login_OK S2C_Login_E_Error = 1
	/// 失败
	S2C_Login_Failed S2C_Login_E_Error = 2
)

var S2C_Login_E_Error_name = map[int32]string{
	0: "Error_",
	1: "OK",
	2: "Failed",
}

var S2C_Login_E_Error_value = map[string]int32{
	"Error_": 0,
	"OK":     1,
	"Failed": 2,
}

func (x S2C_Login_E_Error) String() string {
	return proto.EnumName(S2C_Login_E_Error_name, int32(x))
}

func (S2C_Login_E_Error) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{1, 0}
}

///登录请求
//@msg
type C2S_Login struct {
	///玩家id
	UserId string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	///校验码
	VerifyCode string `protobuf:"bytes,2,opt,name=VerifyCode,proto3" json:"VerifyCode,omitempty"`
}

func (m *C2S_Login) Reset()         { *m = C2S_Login{} }
func (m *C2S_Login) String() string { return proto.CompactTextString(m) }
func (*C2S_Login) ProtoMessage()    {}
func (*C2S_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}
func (m *C2S_Login) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *C2S_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_C2S_Login.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *C2S_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S_Login.Merge(m, src)
}
func (m *C2S_Login) XXX_Size() int {
	return m.Size()
}
func (m *C2S_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S_Login.DiscardUnknown(m)
}

var xxx_messageInfo_C2S_Login proto.InternalMessageInfo

func (m *C2S_Login) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *C2S_Login) GetVerifyCode() string {
	if m != nil {
		return m.VerifyCode
	}
	return ""
}

///登录回包
//@msg
type S2C_Login struct {
	///登录结果
	Err S2C_Login_E_Error `protobuf:"varint,1,opt,name=Err,proto3,enum=msg.S2C_Login_E_Error" json:"Err,omitempty"`
}

func (m *S2C_Login) Reset()         { *m = S2C_Login{} }
func (m *S2C_Login) String() string { return proto.CompactTextString(m) }
func (*S2C_Login) ProtoMessage()    {}
func (*S2C_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{1}
}
func (m *S2C_Login) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *S2C_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_S2C_Login.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *S2C_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_Login.Merge(m, src)
}
func (m *S2C_Login) XXX_Size() int {
	return m.Size()
}
func (m *S2C_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_Login.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_Login proto.InternalMessageInfo

func (m *S2C_Login) GetErr() S2C_Login_E_Error {
	if m != nil {
		return m.Err
	}
	return S2C_Login_Error_
}

func init() {
	proto.RegisterEnum("msg.S2C_Login_E_Error", S2C_Login_E_Error_name, S2C_Login_E_Error_value)
	proto.RegisterType((*C2S_Login)(nil), "msg.C2S_Login")
	proto.RegisterType((*S2C_Login)(nil), "msg.S2C_Login")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f) }

var fileDescriptor_67c21677aa7f4e4f = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc9, 0x4f, 0xcf,
	0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0x72, 0xe6, 0xe2,
	0x74, 0x36, 0x0a, 0x8e, 0xf7, 0x01, 0x89, 0x0b, 0x89, 0x71, 0xb1, 0x85, 0x16, 0xa7, 0x16, 0x79,
	0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x42, 0x72, 0x5c, 0x5c, 0x61, 0xa9,
	0x45, 0x99, 0x69, 0x95, 0xce, 0xf9, 0x29, 0xa9, 0x12, 0x4c, 0x60, 0x39, 0x24, 0x11, 0xa5, 0x04,
	0x2e, 0xce, 0x60, 0x23, 0x67, 0xa8, 0x21, 0x1a, 0x5c, 0xcc, 0xae, 0x45, 0x45, 0x60, 0x13, 0xf8,
	0x8c, 0xc4, 0xf4, 0x72, 0x8b, 0xd3, 0xf5, 0xe0, 0x92, 0x7a, 0xae, 0xf1, 0xae, 0x45, 0x45, 0xf9,
	0x45, 0x41, 0x20, 0x25, 0x4a, 0x9a, 0x5c, 0xec, 0x50, 0xbe, 0x10, 0x17, 0x17, 0x1b, 0x98, 0x11,
	0x2f, 0xc0, 0x20, 0xc4, 0xc6, 0xc5, 0xe4, 0xef, 0x2d, 0xc0, 0x08, 0x12, 0x73, 0x4b, 0xcc, 0xcc,
	0x49, 0x4d, 0x11, 0x60, 0x72, 0x92, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86,
	0x24, 0x36, 0xb0, 0x67, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x8c, 0x46, 0xd6, 0xdb,
	0x00, 0x00, 0x00,
}

func (m *C2S_Login) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *C2S_Login) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *C2S_Login) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VerifyCode) > 0 {
		i -= len(m.VerifyCode)
		copy(dAtA[i:], m.VerifyCode)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.VerifyCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UserId) > 0 {
		i -= len(m.UserId)
		copy(dAtA[i:], m.UserId)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.UserId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *S2C_Login) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2C_Login) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *S2C_Login) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Err != 0 {
		i = encodeVarintLogin(dAtA, i, uint64(m.Err))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLogin(dAtA []byte, offset int, v uint64) int {
	offset -= sovLogin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *C2S_Login) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	l = len(m.VerifyCode)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	return n
}

func (m *S2C_Login) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Err != 0 {
		n += 1 + sovLogin(uint64(m.Err))
	}
	return n
}

func sovLogin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLogin(x uint64) (n int) {
	return sovLogin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *C2S_Login) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
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
			return fmt.Errorf("proto: C2S_Login: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: C2S_Login: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
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
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifyCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
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
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerifyCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLogin
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
func (m *S2C_Login) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
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
			return fmt.Errorf("proto: S2C_Login: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: S2C_Login: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			m.Err = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Err |= S2C_Login_E_Error(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLogin
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
func skipLogin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogin
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
					return 0, ErrIntOverflowLogin
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
					return 0, ErrIntOverflowLogin
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
				return 0, ErrInvalidLengthLogin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLogin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLogin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLogin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLogin = fmt.Errorf("proto: unexpected end of group")
)