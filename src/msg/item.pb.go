// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: item.proto

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

// / 用户物品信息
type Item struct {
	// / 唯一id
	UID string `protobuf:"bytes,1,opt,name=UID,proto3" json:"UID,omitempty"`
	// / 类型id, item.xlsx中的id字段
	TID int64 `protobuf:"varint,2,opt,name=TID,proto3" json:"TID,omitempty"`
	// / 数量
	Num int64 `protobuf:"varint,3,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_ac24453cfc8351bb, []int{0}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Item.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(dst, src)
}
func (m *Item) XXX_Size() int {
	return m.Size()
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *Item) GetTID() int64 {
	if m != nil {
		return m.TID
	}
	return 0
}

func (m *Item) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

// / 获得物品
// @msg
type S2C_GainItem struct {
	// / 新物品列表
	Items []*Item `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
}

func (m *S2C_GainItem) Reset()         { *m = S2C_GainItem{} }
func (m *S2C_GainItem) String() string { return proto.CompactTextString(m) }
func (*S2C_GainItem) ProtoMessage()    {}
func (*S2C_GainItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_ac24453cfc8351bb, []int{1}
}
func (m *S2C_GainItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *S2C_GainItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_S2C_GainItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *S2C_GainItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_GainItem.Merge(dst, src)
}
func (m *S2C_GainItem) XXX_Size() int {
	return m.Size()
}
func (m *S2C_GainItem) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_GainItem.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_GainItem proto.InternalMessageInfo

func (m *S2C_GainItem) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

// / 失去物品
// @msg
type S2C_LostItem struct {
	// / 失去的物品id列表
	IDs []string `protobuf:"bytes,1,rep,name=IDs" json:"IDs,omitempty"`
}

func (m *S2C_LostItem) Reset()         { *m = S2C_LostItem{} }
func (m *S2C_LostItem) String() string { return proto.CompactTextString(m) }
func (*S2C_LostItem) ProtoMessage()    {}
func (*S2C_LostItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_ac24453cfc8351bb, []int{2}
}
func (m *S2C_LostItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *S2C_LostItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_S2C_LostItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *S2C_LostItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_LostItem.Merge(dst, src)
}
func (m *S2C_LostItem) XXX_Size() int {
	return m.Size()
}
func (m *S2C_LostItem) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_LostItem.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_LostItem proto.InternalMessageInfo

func (m *S2C_LostItem) GetIDs() []string {
	if m != nil {
		return m.IDs
	}
	return nil
}

// / 更新物品
// @msg
type S2C_UpdateItem struct {
	// / 发生变更的物品列表
	Items []*Item `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
}

func (m *S2C_UpdateItem) Reset()         { *m = S2C_UpdateItem{} }
func (m *S2C_UpdateItem) String() string { return proto.CompactTextString(m) }
func (*S2C_UpdateItem) ProtoMessage()    {}
func (*S2C_UpdateItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_ac24453cfc8351bb, []int{3}
}
func (m *S2C_UpdateItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *S2C_UpdateItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_S2C_UpdateItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *S2C_UpdateItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C_UpdateItem.Merge(dst, src)
}
func (m *S2C_UpdateItem) XXX_Size() int {
	return m.Size()
}
func (m *S2C_UpdateItem) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C_UpdateItem.DiscardUnknown(m)
}

var xxx_messageInfo_S2C_UpdateItem proto.InternalMessageInfo

func (m *S2C_UpdateItem) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

// /使用物品请求
// @msg
type C2S_UseItem struct {
	// /物品类型id
	TID int32 `protobuf:"varint,1,opt,name=TID,proto3" json:"TID,omitempty"`
	// /数量
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *C2S_UseItem) Reset()         { *m = C2S_UseItem{} }
func (m *C2S_UseItem) String() string { return proto.CompactTextString(m) }
func (*C2S_UseItem) ProtoMessage()    {}
func (*C2S_UseItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_item_ac24453cfc8351bb, []int{4}
}
func (m *C2S_UseItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *C2S_UseItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_C2S_UseItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *C2S_UseItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S_UseItem.Merge(dst, src)
}
func (m *C2S_UseItem) XXX_Size() int {
	return m.Size()
}
func (m *C2S_UseItem) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S_UseItem.DiscardUnknown(m)
}

var xxx_messageInfo_C2S_UseItem proto.InternalMessageInfo

func (m *C2S_UseItem) GetTID() int32 {
	if m != nil {
		return m.TID
	}
	return 0
}

func (m *C2S_UseItem) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*Item)(nil), "msg.Item")
	proto.RegisterType((*S2C_GainItem)(nil), "msg.S2C_GainItem")
	proto.RegisterType((*S2C_LostItem)(nil), "msg.S2C_LostItem")
	proto.RegisterType((*S2C_UpdateItem)(nil), "msg.S2C_UpdateItem")
	proto.RegisterType((*C2S_UseItem)(nil), "msg.C2S_UseItem")
}
func (m *Item) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Item) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintItem(dAtA, i, uint64(len(m.UID)))
		i += copy(dAtA[i:], m.UID)
	}
	if m.TID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintItem(dAtA, i, uint64(m.TID))
	}
	if m.Num != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintItem(dAtA, i, uint64(m.Num))
	}
	return i, nil
}

func (m *S2C_GainItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2C_GainItem) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintItem(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *S2C_LostItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2C_LostItem) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.IDs) > 0 {
		for _, s := range m.IDs {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *S2C_UpdateItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2C_UpdateItem) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintItem(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *C2S_UseItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *C2S_UseItem) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintItem(dAtA, i, uint64(m.TID))
	}
	if m.Num != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintItem(dAtA, i, uint64(m.Num))
	}
	return i, nil
}

func encodeVarintItem(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Item) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	if m.TID != 0 {
		n += 1 + sovItem(uint64(m.TID))
	}
	if m.Num != 0 {
		n += 1 + sovItem(uint64(m.Num))
	}
	return n
}

func (m *S2C_GainItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovItem(uint64(l))
		}
	}
	return n
}

func (m *S2C_LostItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IDs) > 0 {
		for _, s := range m.IDs {
			l = len(s)
			n += 1 + l + sovItem(uint64(l))
		}
	}
	return n
}

func (m *S2C_UpdateItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovItem(uint64(l))
		}
	}
	return n
}

func (m *C2S_UseItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TID != 0 {
		n += 1 + sovItem(uint64(m.TID))
	}
	if m.Num != 0 {
		n += 1 + sovItem(uint64(m.Num))
	}
	return n
}

func sovItem(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozItem(x uint64) (n int) {
	return sovItem(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Item) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: Item: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Item: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TID", wireType)
			}
			m.TID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Num", wireType)
			}
			m.Num = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Num |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthItem
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
func (m *S2C_GainItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: S2C_GainItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: S2C_GainItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Item{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthItem
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
func (m *S2C_LostItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: S2C_LostItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: S2C_LostItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IDs = append(m.IDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthItem
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
func (m *S2C_UpdateItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: S2C_UpdateItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: S2C_UpdateItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Item{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthItem
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
func (m *C2S_UseItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: C2S_UseItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: C2S_UseItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TID", wireType)
			}
			m.TID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Num", wireType)
			}
			m.Num = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Num |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthItem
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
func skipItem(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowItem
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
					return 0, ErrIntOverflowItem
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
					return 0, ErrIntOverflowItem
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
				return 0, ErrInvalidLengthItem
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowItem
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
				next, err := skipItem(dAtA[start:])
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
	ErrInvalidLengthItem = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowItem   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("item.proto", fileDescriptor_item_ac24453cfc8351bb) }

var fileDescriptor_item_ac24453cfc8351bb = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x2c, 0x49, 0xcd,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0xb2, 0xe1, 0x62, 0xf1,
	0x2c, 0x49, 0xcd, 0x15, 0x12, 0xe0, 0x62, 0x0e, 0xf5, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0x31, 0x41, 0x22, 0x21, 0x9e, 0x2e, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x20,
	0x26, 0x48, 0xc4, 0xaf, 0x34, 0x57, 0x82, 0x19, 0x22, 0xe2, 0x57, 0x9a, 0xab, 0xa4, 0xcf, 0xc5,
	0x13, 0x6c, 0xe4, 0x1c, 0xef, 0x9e, 0x98, 0x99, 0x07, 0x36, 0x45, 0x9e, 0x8b, 0x15, 0x44, 0x17,
	0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x71, 0xea, 0xe5, 0x16, 0xa7, 0xeb, 0x81, 0x44, 0x82,
	0x20, 0xe2, 0x4a, 0x0a, 0x10, 0x0d, 0x3e, 0xf9, 0xc5, 0x25, 0x30, 0x6b, 0x3d, 0x5d, 0x20, 0xca,
	0x39, 0x83, 0x40, 0x4c, 0x25, 0x43, 0x2e, 0x3e, 0x90, 0x8a, 0xd0, 0x82, 0x94, 0xc4, 0x92, 0x54,
	0xe2, 0x0c, 0x35, 0xe4, 0xe2, 0x76, 0x36, 0x0a, 0x8e, 0x0f, 0x2d, 0x4e, 0x85, 0x99, 0x19, 0x02,
	0xf5, 0x0a, 0x2b, 0x8a, 0xc3, 0x99, 0x20, 0x22, 0x7e, 0xa5, 0xb9, 0x4e, 0x12, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0x0e, 0x1c, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x3e, 0x71, 0x18, 0xf8, 0x2a, 0x01, 0x00, 0x00,
}
