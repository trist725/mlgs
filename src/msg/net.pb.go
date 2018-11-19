// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: net.proto

package msg

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// /通知掉线
type S2C_DisConn struct {
	// /玩家id
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (m *S2C_DisConn) Reset()         { *m = S2C_DisConn{} }
func (m *S2C_DisConn) String() string { return proto.CompactTextString(m) }
func (*S2C_DisConn) ProtoMessage()    {}
func (*S2C_DisConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_net_36a1c34e49b39a31, []int{0}
}
func (m *S2C_DisConn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *S2C_DisConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_S2C_DisConn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *S2C_DisConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_DisConn.Merge(dst, src)
}
func (m *S2C_DisConn) XXX_Size() int {
	return m.Size()
}
func (m *S2C_DisConn) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_DisConn.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_DisConn proto.InternalMessageInfo

func (m *S2C_DisConn) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

// / ping心跳
type C2S_Ping struct {
}

func (m *C2S_Ping) Reset()         { *m = C2S_Ping{} }
func (m *C2S_Ping) String() string { return proto.CompactTextString(m) }
func (*C2S_Ping) ProtoMessage()    {}
func (*C2S_Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_net_36a1c34e49b39a31, []int{1}
}
func (m *C2S_Ping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *C2S_Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_C2S_Ping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *C2S_Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S_Ping.Merge(dst, src)
}
func (m *C2S_Ping) XXX_Size() int {
	return m.Size()
}
func (m *C2S_Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_C2S_Ping proto.InternalMessageInfo

// / 回复ping心跳
type S2C_Pong struct {
}

func (m *S2C_Pong) Reset()         { *m = S2C_Pong{} }
func (m *S2C_Pong) String() string { return proto.CompactTextString(m) }
func (*S2C_Pong) ProtoMessage()    {}
func (*S2C_Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_net_36a1c34e49b39a31, []int{2}
}
func (m *S2C_Pong) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *S2C_Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_S2C_Pong.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *S2C_Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_Pong.Merge(dst, src)
}
func (m *S2C_Pong) XXX_Size() int {
	return m.Size()
}
func (m *S2C_Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_Pong proto.InternalMessageInfo

func init() {
	proto.RegisterType((*S2C_DisConn)(nil), "msg.S2C_DisConn")
	proto.RegisterType((*C2S_Ping)(nil), "msg.C2S_Ping")
	proto.RegisterType((*S2C_Pong)(nil), "msg.S2C_Pong")
}
func (m *S2C_DisConn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2C_DisConn) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNet(dAtA, i, uint64(m.UserId))
	}
	return i, nil
}

func (m *C2S_Ping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *C2S_Ping) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *S2C_Pong) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2C_Pong) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintNet(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *S2C_DisConn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserId != 0 {
		n += 1 + sovNet(uint64(m.UserId))
	}
	return n
}

func (m *C2S_Ping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *S2C_Pong) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovNet(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNet(x uint64) (n int) {
	return sovNet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *S2C_DisConn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: S2C_DisConn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: S2C_DisConn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNet
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
func (m *C2S_Ping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: C2S_Ping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: C2S_Ping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNet
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
func (m *S2C_Pong) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: S2C_Pong: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: S2C_Pong: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNet
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
func skipNet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNet
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
					return 0, ErrIntOverflowNet
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNet
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNet
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNet
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNet(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNet = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNet   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("net.proto", fileDescriptor_net_36a1c34e49b39a31) }

var fileDescriptor_net_36a1c34e49b39a31 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x4b, 0x2d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0x52, 0xe5, 0xe2, 0x0e, 0x36,
	0x72, 0x8e, 0x77, 0xc9, 0x2c, 0x76, 0xce, 0xcf, 0xcb, 0x13, 0x12, 0xe3, 0x62, 0x0b, 0x2d, 0x4e,
	0x2d, 0xf2, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0xf2, 0x94, 0xb8, 0xb8, 0x38,
	0x9c, 0x8d, 0x82, 0xe3, 0x03, 0x32, 0xf3, 0xd2, 0x41, 0x6c, 0x90, 0x96, 0x80, 0xfc, 0xbc, 0x74,
	0x27, 0x89, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0x5b, 0x62,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xed, 0xbe, 0x45, 0xdd, 0x71, 0x00, 0x00, 0x00,
}
