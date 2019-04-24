// Code generated by protoc-gen-pbex2-go. DO NOT EDIT IT!!!
// source: shop.proto

/*
It has these top-level messages:
	C2S_GetOwnItems
	S2C_GetOwnItems
	C2S_GetOwnDealerSkins
	S2C_GetOwnDealerSkins
	C2S_UsingOwnDealerSkins
	S2C_UsingOwnDealerSkins
	C2S_BuyItem
	S2C_BuyItem
	C2S_Charge
	S2C_Charge
*/

package msg

import "sync"
import protocol "github.com/trist725/mgsu/network/protocol/protobuf/v2"

var _ *sync.Pool
var _ = protocol.PH

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [S2C_GetOwnDealerSkins_E_Err] begin

var S2C_GetOwnDealerSkins_E_Err_Slice = []int32{
	0,
	1,
	2,
}

func S2C_GetOwnDealerSkins_E_Err_Len() int {
	return len(S2C_GetOwnDealerSkins_E_Err_Slice)
}

func Check_S2C_GetOwnDealerSkins_E_Err_I(value int32) bool {
	if _, ok := S2C_GetOwnDealerSkins_E_Err_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_S2C_GetOwnDealerSkins_E_Err(value S2C_GetOwnDealerSkins_E_Err) bool {
	return Check_S2C_GetOwnDealerSkins_E_Err_I(int32(value))
}

func Each_S2C_GetOwnDealerSkins_E_Err(f func(S2C_GetOwnDealerSkins_E_Err) bool) {
	for _, value := range S2C_GetOwnDealerSkins_E_Err_Slice {
		if !f(S2C_GetOwnDealerSkins_E_Err(value)) {
			break
		}
	}
}

func Each_S2C_GetOwnDealerSkins_E_Err_I(f func(int32) bool) {
	for _, value := range S2C_GetOwnDealerSkins_E_Err_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [S2C_GetOwnDealerSkins_E_Err] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [S2C_UsingOwnDealerSkins_E_Err] begin

var S2C_UsingOwnDealerSkins_E_Err_Slice = []int32{
	0,
	1,
	2,
	3,
}

func S2C_UsingOwnDealerSkins_E_Err_Len() int {
	return len(S2C_UsingOwnDealerSkins_E_Err_Slice)
}

func Check_S2C_UsingOwnDealerSkins_E_Err_I(value int32) bool {
	if _, ok := S2C_UsingOwnDealerSkins_E_Err_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_S2C_UsingOwnDealerSkins_E_Err(value S2C_UsingOwnDealerSkins_E_Err) bool {
	return Check_S2C_UsingOwnDealerSkins_E_Err_I(int32(value))
}

func Each_S2C_UsingOwnDealerSkins_E_Err(f func(S2C_UsingOwnDealerSkins_E_Err) bool) {
	for _, value := range S2C_UsingOwnDealerSkins_E_Err_Slice {
		if !f(S2C_UsingOwnDealerSkins_E_Err(value)) {
			break
		}
	}
}

func Each_S2C_UsingOwnDealerSkins_E_Err_I(f func(int32) bool) {
	for _, value := range S2C_UsingOwnDealerSkins_E_Err_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [S2C_UsingOwnDealerSkins_E_Err] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [S2C_BuyItem_E_Err] begin

var S2C_BuyItem_E_Err_Slice = []int32{
	0,
	1,
	2,
	3,
	4,
}

func S2C_BuyItem_E_Err_Len() int {
	return len(S2C_BuyItem_E_Err_Slice)
}

func Check_S2C_BuyItem_E_Err_I(value int32) bool {
	if _, ok := S2C_BuyItem_E_Err_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_S2C_BuyItem_E_Err(value S2C_BuyItem_E_Err) bool {
	return Check_S2C_BuyItem_E_Err_I(int32(value))
}

func Each_S2C_BuyItem_E_Err(f func(S2C_BuyItem_E_Err) bool) {
	for _, value := range S2C_BuyItem_E_Err_Slice {
		if !f(S2C_BuyItem_E_Err(value)) {
			break
		}
	}
}

func Each_S2C_BuyItem_E_Err_I(f func(int32) bool) {
	for _, value := range S2C_BuyItem_E_Err_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [S2C_BuyItem_E_Err] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [S2C_Charge_E_Err] begin

var S2C_Charge_E_Err_Slice = []int32{
	0,
	1,
	3,
	4,
}

func S2C_Charge_E_Err_Len() int {
	return len(S2C_Charge_E_Err_Slice)
}

func Check_S2C_Charge_E_Err_I(value int32) bool {
	if _, ok := S2C_Charge_E_Err_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_S2C_Charge_E_Err(value S2C_Charge_E_Err) bool {
	return Check_S2C_Charge_E_Err_I(int32(value))
}

func Each_S2C_Charge_E_Err(f func(S2C_Charge_E_Err) bool) {
	for _, value := range S2C_Charge_E_Err_Slice {
		if !f(S2C_Charge_E_Err(value)) {
			break
		}
	}
}

func Each_S2C_Charge_E_Err_I(f func(int32) bool) {
	for _, value := range S2C_Charge_E_Err_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [S2C_Charge_E_Err] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_GetOwnItems] begin
func (m *C2S_GetOwnItems) ResetEx() {

}

func (m C2S_GetOwnItems) Clone() *C2S_GetOwnItems {
	n, ok := g_C2S_GetOwnItems_Pool.Get().(*C2S_GetOwnItems)
	if !ok || n == nil {
		n = &C2S_GetOwnItems{}
	}

	return n
}

func Clone_C2S_GetOwnItems_Slice(dst []*C2S_GetOwnItems, src []*C2S_GetOwnItems) []*C2S_GetOwnItems {
	for _, i := range dst {
		Put_C2S_GetOwnItems(i)
	}
	dst = []*C2S_GetOwnItems{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_C2S_GetOwnItems() *C2S_GetOwnItems {
	m := &C2S_GetOwnItems{}
	return m
}

var g_C2S_GetOwnItems_Pool = sync.Pool{}

func Get_C2S_GetOwnItems() *C2S_GetOwnItems {
	m, ok := g_C2S_GetOwnItems_Pool.Get().(*C2S_GetOwnItems)
	if !ok {
		m = New_C2S_GetOwnItems()
	} else {
		if m == nil {
			m = New_C2S_GetOwnItems()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_GetOwnItems(i interface{}) {
	if m, ok := i.(*C2S_GetOwnItems); ok && m != nil {
		g_C2S_GetOwnItems_Pool.Put(i)
	}
}

// message [C2S_GetOwnItems] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_GetOwnItems] begin
func (m *S2C_GetOwnItems) ResetEx() {

	for _, i := range m.Items {
		Put_Item(i)
	}
	m.Items = []*Item{}

}

func (m S2C_GetOwnItems) Clone() *S2C_GetOwnItems {
	n, ok := g_S2C_GetOwnItems_Pool.Get().(*S2C_GetOwnItems)
	if !ok || n == nil {
		n = &S2C_GetOwnItems{}
	}

	if len(m.Items) > 0 {
		for _, i := range m.Items {
			if i != nil {
				n.Items = append(n.Items, i.Clone())
			} else {
				n.Items = append(n.Items, nil)
			}
		}
	} else {
		n.Items = []*Item{}
	}

	return n
}

func Clone_S2C_GetOwnItems_Slice(dst []*S2C_GetOwnItems, src []*S2C_GetOwnItems) []*S2C_GetOwnItems {
	for _, i := range dst {
		Put_S2C_GetOwnItems(i)
	}
	dst = []*S2C_GetOwnItems{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_GetOwnItems() *S2C_GetOwnItems {
	m := &S2C_GetOwnItems{
		Items: []*Item{},
	}
	return m
}

var g_S2C_GetOwnItems_Pool = sync.Pool{}

func Get_S2C_GetOwnItems() *S2C_GetOwnItems {
	m, ok := g_S2C_GetOwnItems_Pool.Get().(*S2C_GetOwnItems)
	if !ok {
		m = New_S2C_GetOwnItems()
	} else {
		if m == nil {
			m = New_S2C_GetOwnItems()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_GetOwnItems(i interface{}) {
	if m, ok := i.(*S2C_GetOwnItems); ok && m != nil {
		g_S2C_GetOwnItems_Pool.Put(i)
	}
}

// message [S2C_GetOwnItems] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_GetOwnDealerSkins] begin
func (m *C2S_GetOwnDealerSkins) ResetEx() {

}

func (m C2S_GetOwnDealerSkins) Clone() *C2S_GetOwnDealerSkins {
	n, ok := g_C2S_GetOwnDealerSkins_Pool.Get().(*C2S_GetOwnDealerSkins)
	if !ok || n == nil {
		n = &C2S_GetOwnDealerSkins{}
	}

	return n
}

func Clone_C2S_GetOwnDealerSkins_Slice(dst []*C2S_GetOwnDealerSkins, src []*C2S_GetOwnDealerSkins) []*C2S_GetOwnDealerSkins {
	for _, i := range dst {
		Put_C2S_GetOwnDealerSkins(i)
	}
	dst = []*C2S_GetOwnDealerSkins{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_C2S_GetOwnDealerSkins() *C2S_GetOwnDealerSkins {
	m := &C2S_GetOwnDealerSkins{}
	return m
}

var g_C2S_GetOwnDealerSkins_Pool = sync.Pool{}

func Get_C2S_GetOwnDealerSkins() *C2S_GetOwnDealerSkins {
	m, ok := g_C2S_GetOwnDealerSkins_Pool.Get().(*C2S_GetOwnDealerSkins)
	if !ok {
		m = New_C2S_GetOwnDealerSkins()
	} else {
		if m == nil {
			m = New_C2S_GetOwnDealerSkins()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_GetOwnDealerSkins(i interface{}) {
	if m, ok := i.(*C2S_GetOwnDealerSkins); ok && m != nil {
		g_C2S_GetOwnDealerSkins_Pool.Put(i)
	}
}

// message [C2S_GetOwnDealerSkins] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_GetOwnDealerSkins] begin
func (m *S2C_GetOwnDealerSkins) ResetEx() {
	m.Ids = []string{}
	m.Id = 0
	m.Err = 0

}

func (m S2C_GetOwnDealerSkins) Clone() *S2C_GetOwnDealerSkins {
	n, ok := g_S2C_GetOwnDealerSkins_Pool.Get().(*S2C_GetOwnDealerSkins)
	if !ok || n == nil {
		n = &S2C_GetOwnDealerSkins{}
	}

	if len(m.Ids) > 0 {
		n.Ids = make([]string, len(m.Ids))
		copy(n.Ids, m.Ids)
	} else {
		n.Ids = []string{}
	}

	n.Id = m.Id
	n.Err = m.Err

	return n
}

func Clone_S2C_GetOwnDealerSkins_Slice(dst []*S2C_GetOwnDealerSkins, src []*S2C_GetOwnDealerSkins) []*S2C_GetOwnDealerSkins {
	for _, i := range dst {
		Put_S2C_GetOwnDealerSkins(i)
	}
	dst = []*S2C_GetOwnDealerSkins{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_GetOwnDealerSkins() *S2C_GetOwnDealerSkins {
	m := &S2C_GetOwnDealerSkins{
		Ids: []string{},
	}
	return m
}

var g_S2C_GetOwnDealerSkins_Pool = sync.Pool{}

func Get_S2C_GetOwnDealerSkins() *S2C_GetOwnDealerSkins {
	m, ok := g_S2C_GetOwnDealerSkins_Pool.Get().(*S2C_GetOwnDealerSkins)
	if !ok {
		m = New_S2C_GetOwnDealerSkins()
	} else {
		if m == nil {
			m = New_S2C_GetOwnDealerSkins()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_GetOwnDealerSkins(i interface{}) {
	if m, ok := i.(*S2C_GetOwnDealerSkins); ok && m != nil {
		g_S2C_GetOwnDealerSkins_Pool.Put(i)
	}
}

// message [S2C_GetOwnDealerSkins] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_UsingOwnDealerSkins] begin
func (m *C2S_UsingOwnDealerSkins) ResetEx() {
	m.Id = 0

}

func (m C2S_UsingOwnDealerSkins) Clone() *C2S_UsingOwnDealerSkins {
	n, ok := g_C2S_UsingOwnDealerSkins_Pool.Get().(*C2S_UsingOwnDealerSkins)
	if !ok || n == nil {
		n = &C2S_UsingOwnDealerSkins{}
	}

	n.Id = m.Id

	return n
}

func Clone_C2S_UsingOwnDealerSkins_Slice(dst []*C2S_UsingOwnDealerSkins, src []*C2S_UsingOwnDealerSkins) []*C2S_UsingOwnDealerSkins {
	for _, i := range dst {
		Put_C2S_UsingOwnDealerSkins(i)
	}
	dst = []*C2S_UsingOwnDealerSkins{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_C2S_UsingOwnDealerSkins() *C2S_UsingOwnDealerSkins {
	m := &C2S_UsingOwnDealerSkins{}
	return m
}

var g_C2S_UsingOwnDealerSkins_Pool = sync.Pool{}

func Get_C2S_UsingOwnDealerSkins() *C2S_UsingOwnDealerSkins {
	m, ok := g_C2S_UsingOwnDealerSkins_Pool.Get().(*C2S_UsingOwnDealerSkins)
	if !ok {
		m = New_C2S_UsingOwnDealerSkins()
	} else {
		if m == nil {
			m = New_C2S_UsingOwnDealerSkins()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_UsingOwnDealerSkins(i interface{}) {
	if m, ok := i.(*C2S_UsingOwnDealerSkins); ok && m != nil {
		g_C2S_UsingOwnDealerSkins_Pool.Put(i)
	}
}

// message [C2S_UsingOwnDealerSkins] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_UsingOwnDealerSkins] begin
func (m *S2C_UsingOwnDealerSkins) ResetEx() {
	m.Id = 0
	m.Err = 0

}

func (m S2C_UsingOwnDealerSkins) Clone() *S2C_UsingOwnDealerSkins {
	n, ok := g_S2C_UsingOwnDealerSkins_Pool.Get().(*S2C_UsingOwnDealerSkins)
	if !ok || n == nil {
		n = &S2C_UsingOwnDealerSkins{}
	}

	n.Id = m.Id
	n.Err = m.Err

	return n
}

func Clone_S2C_UsingOwnDealerSkins_Slice(dst []*S2C_UsingOwnDealerSkins, src []*S2C_UsingOwnDealerSkins) []*S2C_UsingOwnDealerSkins {
	for _, i := range dst {
		Put_S2C_UsingOwnDealerSkins(i)
	}
	dst = []*S2C_UsingOwnDealerSkins{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_UsingOwnDealerSkins() *S2C_UsingOwnDealerSkins {
	m := &S2C_UsingOwnDealerSkins{}
	return m
}

var g_S2C_UsingOwnDealerSkins_Pool = sync.Pool{}

func Get_S2C_UsingOwnDealerSkins() *S2C_UsingOwnDealerSkins {
	m, ok := g_S2C_UsingOwnDealerSkins_Pool.Get().(*S2C_UsingOwnDealerSkins)
	if !ok {
		m = New_S2C_UsingOwnDealerSkins()
	} else {
		if m == nil {
			m = New_S2C_UsingOwnDealerSkins()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_UsingOwnDealerSkins(i interface{}) {
	if m, ok := i.(*S2C_UsingOwnDealerSkins); ok && m != nil {
		g_S2C_UsingOwnDealerSkins_Pool.Put(i)
	}
}

// message [S2C_UsingOwnDealerSkins] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_BuyItem] begin
func (m *C2S_BuyItem) ResetEx() {
	m.Id = 0
	m.Num = 0

}

func (m C2S_BuyItem) Clone() *C2S_BuyItem {
	n, ok := g_C2S_BuyItem_Pool.Get().(*C2S_BuyItem)
	if !ok || n == nil {
		n = &C2S_BuyItem{}
	}

	n.Id = m.Id
	n.Num = m.Num

	return n
}

func Clone_C2S_BuyItem_Slice(dst []*C2S_BuyItem, src []*C2S_BuyItem) []*C2S_BuyItem {
	for _, i := range dst {
		Put_C2S_BuyItem(i)
	}
	dst = []*C2S_BuyItem{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_C2S_BuyItem() *C2S_BuyItem {
	m := &C2S_BuyItem{}
	return m
}

var g_C2S_BuyItem_Pool = sync.Pool{}

func Get_C2S_BuyItem() *C2S_BuyItem {
	m, ok := g_C2S_BuyItem_Pool.Get().(*C2S_BuyItem)
	if !ok {
		m = New_C2S_BuyItem()
	} else {
		if m == nil {
			m = New_C2S_BuyItem()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_BuyItem(i interface{}) {
	if m, ok := i.(*C2S_BuyItem); ok && m != nil {
		g_C2S_BuyItem_Pool.Put(i)
	}
}

// message [C2S_BuyItem] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_BuyItem] begin
func (m *S2C_BuyItem) ResetEx() {
	m.Id = 0
	m.Err = 0

}

func (m S2C_BuyItem) Clone() *S2C_BuyItem {
	n, ok := g_S2C_BuyItem_Pool.Get().(*S2C_BuyItem)
	if !ok || n == nil {
		n = &S2C_BuyItem{}
	}

	n.Id = m.Id
	n.Err = m.Err

	return n
}

func Clone_S2C_BuyItem_Slice(dst []*S2C_BuyItem, src []*S2C_BuyItem) []*S2C_BuyItem {
	for _, i := range dst {
		Put_S2C_BuyItem(i)
	}
	dst = []*S2C_BuyItem{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_BuyItem() *S2C_BuyItem {
	m := &S2C_BuyItem{}
	return m
}

var g_S2C_BuyItem_Pool = sync.Pool{}

func Get_S2C_BuyItem() *S2C_BuyItem {
	m, ok := g_S2C_BuyItem_Pool.Get().(*S2C_BuyItem)
	if !ok {
		m = New_S2C_BuyItem()
	} else {
		if m == nil {
			m = New_S2C_BuyItem()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_BuyItem(i interface{}) {
	if m, ok := i.(*S2C_BuyItem); ok && m != nil {
		g_S2C_BuyItem_Pool.Put(i)
	}
}

// message [S2C_BuyItem] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_Charge] begin
func (m *C2S_Charge) ResetEx() {
	m.Id = 0

}

func (m C2S_Charge) Clone() *C2S_Charge {
	n, ok := g_C2S_Charge_Pool.Get().(*C2S_Charge)
	if !ok || n == nil {
		n = &C2S_Charge{}
	}

	n.Id = m.Id

	return n
}

func Clone_C2S_Charge_Slice(dst []*C2S_Charge, src []*C2S_Charge) []*C2S_Charge {
	for _, i := range dst {
		Put_C2S_Charge(i)
	}
	dst = []*C2S_Charge{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_C2S_Charge() *C2S_Charge {
	m := &C2S_Charge{}
	return m
}

var g_C2S_Charge_Pool = sync.Pool{}

func Get_C2S_Charge() *C2S_Charge {
	m, ok := g_C2S_Charge_Pool.Get().(*C2S_Charge)
	if !ok {
		m = New_C2S_Charge()
	} else {
		if m == nil {
			m = New_C2S_Charge()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_Charge(i interface{}) {
	if m, ok := i.(*C2S_Charge); ok && m != nil {
		g_C2S_Charge_Pool.Put(i)
	}
}

// message [C2S_Charge] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_Charge] begin
func (m *S2C_Charge) ResetEx() {
	m.Err = 0
	m.Diamond = 0

}

func (m S2C_Charge) Clone() *S2C_Charge {
	n, ok := g_S2C_Charge_Pool.Get().(*S2C_Charge)
	if !ok || n == nil {
		n = &S2C_Charge{}
	}

	n.Err = m.Err
	n.Diamond = m.Diamond

	return n
}

func Clone_S2C_Charge_Slice(dst []*S2C_Charge, src []*S2C_Charge) []*S2C_Charge {
	for _, i := range dst {
		Put_S2C_Charge(i)
	}
	dst = []*S2C_Charge{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_Charge() *S2C_Charge {
	m := &S2C_Charge{}
	return m
}

var g_S2C_Charge_Pool = sync.Pool{}

func Get_S2C_Charge() *S2C_Charge {
	m, ok := g_S2C_Charge_Pool.Get().(*S2C_Charge)
	if !ok {
		m = New_S2C_Charge()
	} else {
		if m == nil {
			m = New_S2C_Charge()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_Charge(i interface{}) {
	if m, ok := i.(*S2C_Charge); ok && m != nil {
		g_S2C_Charge_Pool.Put(i)
	}
}

// message [S2C_Charge] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
