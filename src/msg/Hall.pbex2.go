// Code generated by protoc-gen-pbex2-go. DO NOT EDIT IT!!!
// source: Hall.proto

/*
It has these top-level messages:
	C2S_Login
	S2C_Login
	S2C_LoginInfo
	C2S_DaySign
	S2C_DaySign
*/

package msg

import "sync"
import protocol "github.com/trist725/mgsu/network/protocol/protobuf/v2"

var _ *sync.Pool
var _ = protocol.PH

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [C2S_Login_E_LoginType] begin

var C2S_Login_E_LoginType_Slice = []int32{
	0,
	1,
}

func C2S_Login_E_LoginType_Len() int {
	return len(C2S_Login_E_LoginType_Slice)
}

func Check_C2S_Login_E_LoginType_I(value int32) bool {
	if _, ok := C2S_Login_E_LoginType_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_C2S_Login_E_LoginType(value C2S_Login_E_LoginType) bool {
	return Check_C2S_Login_E_LoginType_I(int32(value))
}

func Each_C2S_Login_E_LoginType(f func(C2S_Login_E_LoginType) bool) {
	for _, value := range C2S_Login_E_LoginType_Slice {
		if !f(C2S_Login_E_LoginType(value)) {
			break
		}
	}
}

func Each_C2S_Login_E_LoginType_I(f func(int32) bool) {
	for _, value := range C2S_Login_E_LoginType_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [C2S_Login_E_LoginType] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [S2C_Login_E_ErrReason] begin

var S2C_Login_E_ErrReason_Slice = []int32{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
}

func S2C_Login_E_ErrReason_Len() int {
	return len(S2C_Login_E_ErrReason_Slice)
}

func Check_S2C_Login_E_ErrReason_I(value int32) bool {
	if _, ok := S2C_Login_E_ErrReason_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_S2C_Login_E_ErrReason(value S2C_Login_E_ErrReason) bool {
	return Check_S2C_Login_E_ErrReason_I(int32(value))
}

func Each_S2C_Login_E_ErrReason(f func(S2C_Login_E_ErrReason) bool) {
	for _, value := range S2C_Login_E_ErrReason_Slice {
		if !f(S2C_Login_E_ErrReason(value)) {
			break
		}
	}
}

func Each_S2C_Login_E_ErrReason_I(f func(int32) bool) {
	for _, value := range S2C_Login_E_ErrReason_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [S2C_Login_E_ErrReason] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [S2C_DaySign_E_Err_DaySign] begin

var S2C_DaySign_E_Err_DaySign_Slice = []int32{
	0,
	1,
	2,
	6,
}

func S2C_DaySign_E_Err_DaySign_Len() int {
	return len(S2C_DaySign_E_Err_DaySign_Slice)
}

func Check_S2C_DaySign_E_Err_DaySign_I(value int32) bool {
	if _, ok := S2C_DaySign_E_Err_DaySign_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_S2C_DaySign_E_Err_DaySign(value S2C_DaySign_E_Err_DaySign) bool {
	return Check_S2C_DaySign_E_Err_DaySign_I(int32(value))
}

func Each_S2C_DaySign_E_Err_DaySign(f func(S2C_DaySign_E_Err_DaySign) bool) {
	for _, value := range S2C_DaySign_E_Err_DaySign_Slice {
		if !f(S2C_DaySign_E_Err_DaySign(value)) {
			break
		}
	}
}

func Each_S2C_DaySign_E_Err_DaySign_I(f func(int32) bool) {
	for _, value := range S2C_DaySign_E_Err_DaySign_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [S2C_DaySign_E_Err_DaySign] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_Login] begin
func (m *C2S_Login) ResetEx() {
	m.UID = ""
	m.NickName = ""
	m.AvatarURL = ""
	m.Sex = ""
	m.Password = ""
	m.Location = ""
	m.Logintype = 0

}

func (m C2S_Login) Clone() *C2S_Login {
	n, ok := g_C2S_Login_Pool.Get().(*C2S_Login)
	if !ok || n == nil {
		n = &C2S_Login{}
	}

	n.UID = m.UID
	n.NickName = m.NickName
	n.AvatarURL = m.AvatarURL
	n.Sex = m.Sex
	n.Password = m.Password
	n.Location = m.Location
	n.Logintype = m.Logintype

	return n
}

func Clone_C2S_Login_Slice(dst []*C2S_Login, src []*C2S_Login) []*C2S_Login {
	for _, i := range dst {
		Put_C2S_Login(i)
	}
	dst = []*C2S_Login{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func (C2S_Login) V2() {
}

func (C2S_Login) MessageID() protocol.MessageID {
	return "msg.C2S_Login"
}

func C2S_Login_MessageID() protocol.MessageID {
	return "msg.C2S_Login"
}

func New_C2S_Login() *C2S_Login {
	m := &C2S_Login{}
	return m
}

var g_C2S_Login_Pool = sync.Pool{}

func Get_C2S_Login() *C2S_Login {
	m, ok := g_C2S_Login_Pool.Get().(*C2S_Login)
	if !ok {
		m = New_C2S_Login()
	} else {
		if m == nil {
			m = New_C2S_Login()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_Login(i interface{}) {
	if m, ok := i.(*C2S_Login); ok && m != nil {
		g_C2S_Login_Pool.Put(i)
	}
}

func init() {
	Protocol.Register(
		&C2S_Login{},
		func() protocol.IMessage { return Get_C2S_Login() },
		func(msg protocol.IMessage) { Put_C2S_Login(msg) },
	)
}

// message [C2S_Login] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_Login] begin
func (m *S2C_Login) ResetEx() {
	m.Reason = 0

}

func (m S2C_Login) Clone() *S2C_Login {
	n, ok := g_S2C_Login_Pool.Get().(*S2C_Login)
	if !ok || n == nil {
		n = &S2C_Login{}
	}

	n.Reason = m.Reason

	return n
}

func Clone_S2C_Login_Slice(dst []*S2C_Login, src []*S2C_Login) []*S2C_Login {
	for _, i := range dst {
		Put_S2C_Login(i)
	}
	dst = []*S2C_Login{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func (S2C_Login) V2() {
}

func (S2C_Login) MessageID() protocol.MessageID {
	return "msg.S2C_Login"
}

func S2C_Login_MessageID() protocol.MessageID {
	return "msg.S2C_Login"
}

func New_S2C_Login() *S2C_Login {
	m := &S2C_Login{}
	return m
}

var g_S2C_Login_Pool = sync.Pool{}

func Get_S2C_Login() *S2C_Login {
	m, ok := g_S2C_Login_Pool.Get().(*S2C_Login)
	if !ok {
		m = New_S2C_Login()
	} else {
		if m == nil {
			m = New_S2C_Login()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_Login(i interface{}) {
	if m, ok := i.(*S2C_Login); ok && m != nil {
		g_S2C_Login_Pool.Put(i)
	}
}

func init() {
	Protocol.Register(
		&S2C_Login{},
		func() protocol.IMessage { return Get_S2C_Login() },
		func(msg protocol.IMessage) { Put_S2C_Login(msg) },
	)
}

// message [S2C_Login] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_LoginInfo] begin
func (m *S2C_LoginInfo) ResetEx() {
	m.ID = 0
	m.NickName = ""
	m.AvatarURL = ""

	for _, i := range m.Monies {
		Put_Money(i)
	}
	m.Monies = []*Money{}
	m.DaySigned = false
	m.SignedDays = 0

	for _, i := range m.SignRewards {
		Put_Item(i)
	}
	m.SignRewards = []*Item{}

}

func (m S2C_LoginInfo) Clone() *S2C_LoginInfo {
	n, ok := g_S2C_LoginInfo_Pool.Get().(*S2C_LoginInfo)
	if !ok || n == nil {
		n = &S2C_LoginInfo{}
	}

	n.ID = m.ID
	n.NickName = m.NickName
	n.AvatarURL = m.AvatarURL

	if len(m.Monies) > 0 {
		for _, i := range m.Monies {
			if i != nil {
				n.Monies = append(n.Monies, i.Clone())
			} else {
				n.Monies = append(n.Monies, nil)
			}
		}
	} else {
		n.Monies = []*Money{}
	}

	n.DaySigned = m.DaySigned
	n.SignedDays = m.SignedDays

	if len(m.SignRewards) > 0 {
		for _, i := range m.SignRewards {
			if i != nil {
				n.SignRewards = append(n.SignRewards, i.Clone())
			} else {
				n.SignRewards = append(n.SignRewards, nil)
			}
		}
	} else {
		n.SignRewards = []*Item{}
	}

	return n
}

func Clone_S2C_LoginInfo_Slice(dst []*S2C_LoginInfo, src []*S2C_LoginInfo) []*S2C_LoginInfo {
	for _, i := range dst {
		Put_S2C_LoginInfo(i)
	}
	dst = []*S2C_LoginInfo{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func (S2C_LoginInfo) V2() {
}

func (S2C_LoginInfo) MessageID() protocol.MessageID {
	return "msg.S2C_LoginInfo"
}

func S2C_LoginInfo_MessageID() protocol.MessageID {
	return "msg.S2C_LoginInfo"
}

func New_S2C_LoginInfo() *S2C_LoginInfo {
	m := &S2C_LoginInfo{
		Monies:      []*Money{},
		SignRewards: []*Item{},
	}
	return m
}

var g_S2C_LoginInfo_Pool = sync.Pool{}

func Get_S2C_LoginInfo() *S2C_LoginInfo {
	m, ok := g_S2C_LoginInfo_Pool.Get().(*S2C_LoginInfo)
	if !ok {
		m = New_S2C_LoginInfo()
	} else {
		if m == nil {
			m = New_S2C_LoginInfo()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_LoginInfo(i interface{}) {
	if m, ok := i.(*S2C_LoginInfo); ok && m != nil {
		g_S2C_LoginInfo_Pool.Put(i)
	}
}

func init() {
	Protocol.Register(
		&S2C_LoginInfo{},
		func() protocol.IMessage { return Get_S2C_LoginInfo() },
		func(msg protocol.IMessage) { Put_S2C_LoginInfo(msg) },
	)
}

// message [S2C_LoginInfo] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_DaySign] begin
func (m *C2S_DaySign) ResetEx() {
	m.Day = 0

}

func (m C2S_DaySign) Clone() *C2S_DaySign {
	n, ok := g_C2S_DaySign_Pool.Get().(*C2S_DaySign)
	if !ok || n == nil {
		n = &C2S_DaySign{}
	}

	n.Day = m.Day

	return n
}

func Clone_C2S_DaySign_Slice(dst []*C2S_DaySign, src []*C2S_DaySign) []*C2S_DaySign {
	for _, i := range dst {
		Put_C2S_DaySign(i)
	}
	dst = []*C2S_DaySign{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func (C2S_DaySign) V2() {
}

func (C2S_DaySign) MessageID() protocol.MessageID {
	return "msg.C2S_DaySign"
}

func C2S_DaySign_MessageID() protocol.MessageID {
	return "msg.C2S_DaySign"
}

func New_C2S_DaySign() *C2S_DaySign {
	m := &C2S_DaySign{}
	return m
}

var g_C2S_DaySign_Pool = sync.Pool{}

func Get_C2S_DaySign() *C2S_DaySign {
	m, ok := g_C2S_DaySign_Pool.Get().(*C2S_DaySign)
	if !ok {
		m = New_C2S_DaySign()
	} else {
		if m == nil {
			m = New_C2S_DaySign()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_DaySign(i interface{}) {
	if m, ok := i.(*C2S_DaySign); ok && m != nil {
		g_C2S_DaySign_Pool.Put(i)
	}
}

func init() {
	Protocol.Register(
		&C2S_DaySign{},
		func() protocol.IMessage { return Get_C2S_DaySign() },
		func(msg protocol.IMessage) { Put_C2S_DaySign(msg) },
	)
}

// message [C2S_DaySign] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_DaySign] begin
func (m *S2C_DaySign) ResetEx() {
	m.Err = 0

	for _, i := range m.Monies {
		Put_Money(i)
	}
	m.Monies = []*Money{}

}

func (m S2C_DaySign) Clone() *S2C_DaySign {
	n, ok := g_S2C_DaySign_Pool.Get().(*S2C_DaySign)
	if !ok || n == nil {
		n = &S2C_DaySign{}
	}

	n.Err = m.Err

	if len(m.Monies) > 0 {
		for _, i := range m.Monies {
			if i != nil {
				n.Monies = append(n.Monies, i.Clone())
			} else {
				n.Monies = append(n.Monies, nil)
			}
		}
	} else {
		n.Monies = []*Money{}
	}

	return n
}

func Clone_S2C_DaySign_Slice(dst []*S2C_DaySign, src []*S2C_DaySign) []*S2C_DaySign {
	for _, i := range dst {
		Put_S2C_DaySign(i)
	}
	dst = []*S2C_DaySign{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func (S2C_DaySign) V2() {
}

func (S2C_DaySign) MessageID() protocol.MessageID {
	return "msg.S2C_DaySign"
}

func S2C_DaySign_MessageID() protocol.MessageID {
	return "msg.S2C_DaySign"
}

func New_S2C_DaySign() *S2C_DaySign {
	m := &S2C_DaySign{
		Monies: []*Money{},
	}
	return m
}

var g_S2C_DaySign_Pool = sync.Pool{}

func Get_S2C_DaySign() *S2C_DaySign {
	m, ok := g_S2C_DaySign_Pool.Get().(*S2C_DaySign)
	if !ok {
		m = New_S2C_DaySign()
	} else {
		if m == nil {
			m = New_S2C_DaySign()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_DaySign(i interface{}) {
	if m, ok := i.(*S2C_DaySign); ok && m != nil {
		g_S2C_DaySign_Pool.Put(i)
	}
}

func init() {
	Protocol.Register(
		&S2C_DaySign{},
		func() protocol.IMessage { return Get_S2C_DaySign() },
		func(msg protocol.IMessage) { Put_S2C_DaySign(msg) },
	)
}

// message [S2C_DaySign] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////