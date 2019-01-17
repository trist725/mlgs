// Code generated by protoc-gen-pbex2-go. DO NOT EDIT IT!!!
// source: match.proto

/*
It has these top-level messages:
	Card
	BestCombo
	Player
	Room
	C2S_QuickMatchStart
	S2C_QuickMatchStart
	C2S_PlayerLeaveRoom
	S2C_PlayerLeaveRoom
	S2C_UpdatePlayerJoinRoom
	S2C_UpdatePlayerLeaveRoom
	S2C_GameStart
	S2C_Turn
	C2S_TurnAction
	S2C_TurnAction
	S2C_PublicCard
	C2S_AutoAction
	S2C_GameOver
	Balance
	S2C_Balance
	C2S_RoomChat
	S2C_RoomChat
*/

package msg

import "sync"
import protocol "github.com/trist725/mgsu/network/protocol/protobuf/v2"

var _ *sync.Pool
var _ = protocol.PH

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [S2C_QuickMatchStart_E_Err_QuickMatchStart] begin

var S2C_QuickMatchStart_E_Err_QuickMatchStart_Slice = []int32{
	0,
	1,
	2,
	3,
	4,
	5,
}

func S2C_QuickMatchStart_E_Err_QuickMatchStart_Len() int {
	return len(S2C_QuickMatchStart_E_Err_QuickMatchStart_Slice)
}

func Check_S2C_QuickMatchStart_E_Err_QuickMatchStart_I(value int32) bool {
	if _, ok := S2C_QuickMatchStart_E_Err_QuickMatchStart_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_S2C_QuickMatchStart_E_Err_QuickMatchStart(value S2C_QuickMatchStart_E_Err_QuickMatchStart) bool {
	return Check_S2C_QuickMatchStart_E_Err_QuickMatchStart_I(int32(value))
}

func Each_S2C_QuickMatchStart_E_Err_QuickMatchStart(f func(S2C_QuickMatchStart_E_Err_QuickMatchStart) bool) {
	for _, value := range S2C_QuickMatchStart_E_Err_QuickMatchStart_Slice {
		if !f(S2C_QuickMatchStart_E_Err_QuickMatchStart(value)) {
			break
		}
	}
}

func Each_S2C_QuickMatchStart_E_Err_QuickMatchStart_I(f func(int32) bool) {
	for _, value := range S2C_QuickMatchStart_E_Err_QuickMatchStart_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [S2C_QuickMatchStart_E_Err_QuickMatchStart] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom] begin

var S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom_Slice = []int32{
	0,
	1,
	2,
	3,
}

func S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom_Len() int {
	return len(S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom_Slice)
}

func Check_S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom_I(value int32) bool {
	if _, ok := S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom(value S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom) bool {
	return Check_S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom_I(int32(value))
}

func Each_S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom(f func(S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom) bool) {
	for _, value := range S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom_Slice {
		if !f(S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom(value)) {
			break
		}
	}
}

func Each_S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom_I(f func(int32) bool) {
	for _, value := range S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [S2C_PlayerLeaveRoom_E_Err_PlayerLeaveRoom] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// enum [S2C_UpdatePlayerLeaveRoom_E_Err] begin

var S2C_UpdatePlayerLeaveRoom_E_Err_Slice = []int32{
	0,
	1,
	2,
	3,
	4,
	5,
}

func S2C_UpdatePlayerLeaveRoom_E_Err_Len() int {
	return len(S2C_UpdatePlayerLeaveRoom_E_Err_Slice)
}

func Check_S2C_UpdatePlayerLeaveRoom_E_Err_I(value int32) bool {
	if _, ok := S2C_UpdatePlayerLeaveRoom_E_Err_name[value]; ok && value != 0 {
		return true
	}
	return false
}

func Check_S2C_UpdatePlayerLeaveRoom_E_Err(value S2C_UpdatePlayerLeaveRoom_E_Err) bool {
	return Check_S2C_UpdatePlayerLeaveRoom_E_Err_I(int32(value))
}

func Each_S2C_UpdatePlayerLeaveRoom_E_Err(f func(S2C_UpdatePlayerLeaveRoom_E_Err) bool) {
	for _, value := range S2C_UpdatePlayerLeaveRoom_E_Err_Slice {
		if !f(S2C_UpdatePlayerLeaveRoom_E_Err(value)) {
			break
		}
	}
}

func Each_S2C_UpdatePlayerLeaveRoom_E_Err_I(f func(int32) bool) {
	for _, value := range S2C_UpdatePlayerLeaveRoom_E_Err_Slice {
		if !f(value) {
			break
		}
	}
}

// enum [S2C_UpdatePlayerLeaveRoom_E_Err] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [Card] begin
func (m *Card) ResetEx() {
	m.Color = 0
	m.Num = 0

}

func (m Card) Clone() *Card {
	n, ok := g_Card_Pool.Get().(*Card)
	if !ok || n == nil {
		n = &Card{}
	}

	n.Color = m.Color
	n.Num = m.Num

	return n
}

func Clone_Card_Slice(dst []*Card, src []*Card) []*Card {
	for _, i := range dst {
		Put_Card(i)
	}
	dst = []*Card{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_Card() *Card {
	m := &Card{}
	return m
}

var g_Card_Pool = sync.Pool{}

func Get_Card() *Card {
	m, ok := g_Card_Pool.Get().(*Card)
	if !ok {
		m = New_Card()
	} else {
		if m == nil {
			m = New_Card()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_Card(i interface{}) {
	if m, ok := i.(*Card); ok && m != nil {
		g_Card_Pool.Put(i)
	}
}

// message [Card] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [BestCombo] begin
func (m *BestCombo) ResetEx() {

	for _, i := range m.Cards {
		Put_Card(i)
	}
	m.Cards = []*Card{}
	m.Type = 0

}

func (m BestCombo) Clone() *BestCombo {
	n, ok := g_BestCombo_Pool.Get().(*BestCombo)
	if !ok || n == nil {
		n = &BestCombo{}
	}

	if len(m.Cards) > 0 {
		for _, i := range m.Cards {
			if i != nil {
				n.Cards = append(n.Cards, i.Clone())
			} else {
				n.Cards = append(n.Cards, nil)
			}
		}
	} else {
		n.Cards = []*Card{}
	}

	n.Type = m.Type

	return n
}

func Clone_BestCombo_Slice(dst []*BestCombo, src []*BestCombo) []*BestCombo {
	for _, i := range dst {
		Put_BestCombo(i)
	}
	dst = []*BestCombo{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_BestCombo() *BestCombo {
	m := &BestCombo{
		Cards: []*Card{},
	}
	return m
}

var g_BestCombo_Pool = sync.Pool{}

func Get_BestCombo() *BestCombo {
	m, ok := g_BestCombo_Pool.Get().(*BestCombo)
	if !ok {
		m = New_BestCombo()
	} else {
		if m == nil {
			m = New_BestCombo()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_BestCombo(i interface{}) {
	if m, ok := i.(*BestCombo); ok && m != nil {
		g_BestCombo_Pool.Put(i)
	}
}

// message [BestCombo] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [Player] begin
func (m *Player) ResetEx() {
	m.UserId = 0
	m.NickName = ""
	m.AvatarURL = ""
	m.Pos = 0
	m.Role = 0
	m.Chip = 0
	m.BetChip = 0

	for _, i := range m.Cards {
		Put_Card(i)
	}
	m.Cards = []*Card{}
	m.Sex = ""

}

func (m Player) Clone() *Player {
	n, ok := g_Player_Pool.Get().(*Player)
	if !ok || n == nil {
		n = &Player{}
	}

	n.UserId = m.UserId
	n.NickName = m.NickName
	n.AvatarURL = m.AvatarURL
	n.Pos = m.Pos
	n.Role = m.Role
	n.Chip = m.Chip
	n.BetChip = m.BetChip

	if len(m.Cards) > 0 {
		for _, i := range m.Cards {
			if i != nil {
				n.Cards = append(n.Cards, i.Clone())
			} else {
				n.Cards = append(n.Cards, nil)
			}
		}
	} else {
		n.Cards = []*Card{}
	}

	n.Sex = m.Sex

	return n
}

func Clone_Player_Slice(dst []*Player, src []*Player) []*Player {
	for _, i := range dst {
		Put_Player(i)
	}
	dst = []*Player{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_Player() *Player {
	m := &Player{
		Cards: []*Card{},
	}
	return m
}

var g_Player_Pool = sync.Pool{}

func Get_Player() *Player {
	m, ok := g_Player_Pool.Get().(*Player)
	if !ok {
		m = New_Player()
	} else {
		if m == nil {
			m = New_Player()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_Player(i interface{}) {
	if m, ok := i.(*Player); ok && m != nil {
		g_Player_Pool.Put(i)
	}
}

// message [Player] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [Room] begin
func (m *Room) ResetEx() {
	m.Id = 0
	m.Name = ""

	for _, i := range m.Players {
		Put_Player(i)
	}
	m.Players = []*Player{}
	m.Chip = 0
	m.MaxBet = 0

}

func (m Room) Clone() *Room {
	n, ok := g_Room_Pool.Get().(*Room)
	if !ok || n == nil {
		n = &Room{}
	}

	n.Id = m.Id
	n.Name = m.Name

	if len(m.Players) > 0 {
		for _, i := range m.Players {
			if i != nil {
				n.Players = append(n.Players, i.Clone())
			} else {
				n.Players = append(n.Players, nil)
			}
		}
	} else {
		n.Players = []*Player{}
	}

	n.Chip = m.Chip
	n.MaxBet = m.MaxBet

	return n
}

func Clone_Room_Slice(dst []*Room, src []*Room) []*Room {
	for _, i := range dst {
		Put_Room(i)
	}
	dst = []*Room{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_Room() *Room {
	m := &Room{
		Players: []*Player{},
	}
	return m
}

var g_Room_Pool = sync.Pool{}

func Get_Room() *Room {
	m, ok := g_Room_Pool.Get().(*Room)
	if !ok {
		m = New_Room()
	} else {
		if m == nil {
			m = New_Room()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_Room(i interface{}) {
	if m, ok := i.(*Room); ok && m != nil {
		g_Room_Pool.Put(i)
	}
}

// message [Room] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_QuickMatchStart] begin
func (m *C2S_QuickMatchStart) ResetEx() {
	m.Type = 0

}

func (m C2S_QuickMatchStart) Clone() *C2S_QuickMatchStart {
	n, ok := g_C2S_QuickMatchStart_Pool.Get().(*C2S_QuickMatchStart)
	if !ok || n == nil {
		n = &C2S_QuickMatchStart{}
	}

	n.Type = m.Type

	return n
}

func Clone_C2S_QuickMatchStart_Slice(dst []*C2S_QuickMatchStart, src []*C2S_QuickMatchStart) []*C2S_QuickMatchStart {
	for _, i := range dst {
		Put_C2S_QuickMatchStart(i)
	}
	dst = []*C2S_QuickMatchStart{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_C2S_QuickMatchStart() *C2S_QuickMatchStart {
	m := &C2S_QuickMatchStart{}
	return m
}

var g_C2S_QuickMatchStart_Pool = sync.Pool{}

func Get_C2S_QuickMatchStart() *C2S_QuickMatchStart {
	m, ok := g_C2S_QuickMatchStart_Pool.Get().(*C2S_QuickMatchStart)
	if !ok {
		m = New_C2S_QuickMatchStart()
	} else {
		if m == nil {
			m = New_C2S_QuickMatchStart()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_QuickMatchStart(i interface{}) {
	if m, ok := i.(*C2S_QuickMatchStart); ok && m != nil {
		g_C2S_QuickMatchStart_Pool.Put(i)
	}
}

// message [C2S_QuickMatchStart] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_QuickMatchStart] begin
func (m *S2C_QuickMatchStart) ResetEx() {
	m.Err = 0
	m.Room.ResetEx()

}

func (m S2C_QuickMatchStart) Clone() *S2C_QuickMatchStart {
	n, ok := g_S2C_QuickMatchStart_Pool.Get().(*S2C_QuickMatchStart)
	if !ok || n == nil {
		n = &S2C_QuickMatchStart{}
	}

	n.Err = m.Err
	n.Room = m.Room.Clone()

	return n
}

func Clone_S2C_QuickMatchStart_Slice(dst []*S2C_QuickMatchStart, src []*S2C_QuickMatchStart) []*S2C_QuickMatchStart {
	for _, i := range dst {
		Put_S2C_QuickMatchStart(i)
	}
	dst = []*S2C_QuickMatchStart{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_QuickMatchStart() *S2C_QuickMatchStart {
	m := &S2C_QuickMatchStart{
		Room: New_Room(),
	}
	return m
}

var g_S2C_QuickMatchStart_Pool = sync.Pool{}

func Get_S2C_QuickMatchStart() *S2C_QuickMatchStart {
	m, ok := g_S2C_QuickMatchStart_Pool.Get().(*S2C_QuickMatchStart)
	if !ok {
		m = New_S2C_QuickMatchStart()
	} else {
		if m == nil {
			m = New_S2C_QuickMatchStart()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_QuickMatchStart(i interface{}) {
	if m, ok := i.(*S2C_QuickMatchStart); ok && m != nil {
		g_S2C_QuickMatchStart_Pool.Put(i)
	}
}

// message [S2C_QuickMatchStart] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_PlayerLeaveRoom] begin
func (m *C2S_PlayerLeaveRoom) ResetEx() {

}

func (m C2S_PlayerLeaveRoom) Clone() *C2S_PlayerLeaveRoom {
	n, ok := g_C2S_PlayerLeaveRoom_Pool.Get().(*C2S_PlayerLeaveRoom)
	if !ok || n == nil {
		n = &C2S_PlayerLeaveRoom{}
	}

	return n
}

func Clone_C2S_PlayerLeaveRoom_Slice(dst []*C2S_PlayerLeaveRoom, src []*C2S_PlayerLeaveRoom) []*C2S_PlayerLeaveRoom {
	for _, i := range dst {
		Put_C2S_PlayerLeaveRoom(i)
	}
	dst = []*C2S_PlayerLeaveRoom{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_C2S_PlayerLeaveRoom() *C2S_PlayerLeaveRoom {
	m := &C2S_PlayerLeaveRoom{}
	return m
}

var g_C2S_PlayerLeaveRoom_Pool = sync.Pool{}

func Get_C2S_PlayerLeaveRoom() *C2S_PlayerLeaveRoom {
	m, ok := g_C2S_PlayerLeaveRoom_Pool.Get().(*C2S_PlayerLeaveRoom)
	if !ok {
		m = New_C2S_PlayerLeaveRoom()
	} else {
		if m == nil {
			m = New_C2S_PlayerLeaveRoom()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_PlayerLeaveRoom(i interface{}) {
	if m, ok := i.(*C2S_PlayerLeaveRoom); ok && m != nil {
		g_C2S_PlayerLeaveRoom_Pool.Put(i)
	}
}

// message [C2S_PlayerLeaveRoom] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_PlayerLeaveRoom] begin
func (m *S2C_PlayerLeaveRoom) ResetEx() {
	m.Err = 0

}

func (m S2C_PlayerLeaveRoom) Clone() *S2C_PlayerLeaveRoom {
	n, ok := g_S2C_PlayerLeaveRoom_Pool.Get().(*S2C_PlayerLeaveRoom)
	if !ok || n == nil {
		n = &S2C_PlayerLeaveRoom{}
	}

	n.Err = m.Err

	return n
}

func Clone_S2C_PlayerLeaveRoom_Slice(dst []*S2C_PlayerLeaveRoom, src []*S2C_PlayerLeaveRoom) []*S2C_PlayerLeaveRoom {
	for _, i := range dst {
		Put_S2C_PlayerLeaveRoom(i)
	}
	dst = []*S2C_PlayerLeaveRoom{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_PlayerLeaveRoom() *S2C_PlayerLeaveRoom {
	m := &S2C_PlayerLeaveRoom{}
	return m
}

var g_S2C_PlayerLeaveRoom_Pool = sync.Pool{}

func Get_S2C_PlayerLeaveRoom() *S2C_PlayerLeaveRoom {
	m, ok := g_S2C_PlayerLeaveRoom_Pool.Get().(*S2C_PlayerLeaveRoom)
	if !ok {
		m = New_S2C_PlayerLeaveRoom()
	} else {
		if m == nil {
			m = New_S2C_PlayerLeaveRoom()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_PlayerLeaveRoom(i interface{}) {
	if m, ok := i.(*S2C_PlayerLeaveRoom); ok && m != nil {
		g_S2C_PlayerLeaveRoom_Pool.Put(i)
	}
}

// message [S2C_PlayerLeaveRoom] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_UpdatePlayerJoinRoom] begin
func (m *S2C_UpdatePlayerJoinRoom) ResetEx() {

	for _, i := range m.Players {
		Put_Player(i)
	}
	m.Players = []*Player{}

}

func (m S2C_UpdatePlayerJoinRoom) Clone() *S2C_UpdatePlayerJoinRoom {
	n, ok := g_S2C_UpdatePlayerJoinRoom_Pool.Get().(*S2C_UpdatePlayerJoinRoom)
	if !ok || n == nil {
		n = &S2C_UpdatePlayerJoinRoom{}
	}

	if len(m.Players) > 0 {
		for _, i := range m.Players {
			if i != nil {
				n.Players = append(n.Players, i.Clone())
			} else {
				n.Players = append(n.Players, nil)
			}
		}
	} else {
		n.Players = []*Player{}
	}

	return n
}

func Clone_S2C_UpdatePlayerJoinRoom_Slice(dst []*S2C_UpdatePlayerJoinRoom, src []*S2C_UpdatePlayerJoinRoom) []*S2C_UpdatePlayerJoinRoom {
	for _, i := range dst {
		Put_S2C_UpdatePlayerJoinRoom(i)
	}
	dst = []*S2C_UpdatePlayerJoinRoom{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_UpdatePlayerJoinRoom() *S2C_UpdatePlayerJoinRoom {
	m := &S2C_UpdatePlayerJoinRoom{
		Players: []*Player{},
	}
	return m
}

var g_S2C_UpdatePlayerJoinRoom_Pool = sync.Pool{}

func Get_S2C_UpdatePlayerJoinRoom() *S2C_UpdatePlayerJoinRoom {
	m, ok := g_S2C_UpdatePlayerJoinRoom_Pool.Get().(*S2C_UpdatePlayerJoinRoom)
	if !ok {
		m = New_S2C_UpdatePlayerJoinRoom()
	} else {
		if m == nil {
			m = New_S2C_UpdatePlayerJoinRoom()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_UpdatePlayerJoinRoom(i interface{}) {
	if m, ok := i.(*S2C_UpdatePlayerJoinRoom); ok && m != nil {
		g_S2C_UpdatePlayerJoinRoom_Pool.Put(i)
	}
}

// message [S2C_UpdatePlayerJoinRoom] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_UpdatePlayerLeaveRoom] begin
func (m *S2C_UpdatePlayerLeaveRoom) ResetEx() {
	m.UserId = 0
	m.Reason = 0

}

func (m S2C_UpdatePlayerLeaveRoom) Clone() *S2C_UpdatePlayerLeaveRoom {
	n, ok := g_S2C_UpdatePlayerLeaveRoom_Pool.Get().(*S2C_UpdatePlayerLeaveRoom)
	if !ok || n == nil {
		n = &S2C_UpdatePlayerLeaveRoom{}
	}

	n.UserId = m.UserId
	n.Reason = m.Reason

	return n
}

func Clone_S2C_UpdatePlayerLeaveRoom_Slice(dst []*S2C_UpdatePlayerLeaveRoom, src []*S2C_UpdatePlayerLeaveRoom) []*S2C_UpdatePlayerLeaveRoom {
	for _, i := range dst {
		Put_S2C_UpdatePlayerLeaveRoom(i)
	}
	dst = []*S2C_UpdatePlayerLeaveRoom{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_UpdatePlayerLeaveRoom() *S2C_UpdatePlayerLeaveRoom {
	m := &S2C_UpdatePlayerLeaveRoom{}
	return m
}

var g_S2C_UpdatePlayerLeaveRoom_Pool = sync.Pool{}

func Get_S2C_UpdatePlayerLeaveRoom() *S2C_UpdatePlayerLeaveRoom {
	m, ok := g_S2C_UpdatePlayerLeaveRoom_Pool.Get().(*S2C_UpdatePlayerLeaveRoom)
	if !ok {
		m = New_S2C_UpdatePlayerLeaveRoom()
	} else {
		if m == nil {
			m = New_S2C_UpdatePlayerLeaveRoom()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_UpdatePlayerLeaveRoom(i interface{}) {
	if m, ok := i.(*S2C_UpdatePlayerLeaveRoom); ok && m != nil {
		g_S2C_UpdatePlayerLeaveRoom_Pool.Put(i)
	}
}

// message [S2C_UpdatePlayerLeaveRoom] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_GameStart] begin
func (m *S2C_GameStart) ResetEx() {
	m.Pos = 0

	for _, i := range m.Cards {
		Put_Card(i)
	}
	m.Cards = []*Card{}
	m.SmallBlind = 0
	m.Best.ResetEx()
	m.Round = 0

}

func (m S2C_GameStart) Clone() *S2C_GameStart {
	n, ok := g_S2C_GameStart_Pool.Get().(*S2C_GameStart)
	if !ok || n == nil {
		n = &S2C_GameStart{}
	}

	n.Pos = m.Pos

	if len(m.Cards) > 0 {
		for _, i := range m.Cards {
			if i != nil {
				n.Cards = append(n.Cards, i.Clone())
			} else {
				n.Cards = append(n.Cards, nil)
			}
		}
	} else {
		n.Cards = []*Card{}
	}

	n.SmallBlind = m.SmallBlind
	n.Best = m.Best.Clone()
	n.Round = m.Round

	return n
}

func Clone_S2C_GameStart_Slice(dst []*S2C_GameStart, src []*S2C_GameStart) []*S2C_GameStart {
	for _, i := range dst {
		Put_S2C_GameStart(i)
	}
	dst = []*S2C_GameStart{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_GameStart() *S2C_GameStart {
	m := &S2C_GameStart{
		Cards: []*Card{},
		Best:  New_BestCombo(),
	}
	return m
}

var g_S2C_GameStart_Pool = sync.Pool{}

func Get_S2C_GameStart() *S2C_GameStart {
	m, ok := g_S2C_GameStart_Pool.Get().(*S2C_GameStart)
	if !ok {
		m = New_S2C_GameStart()
	} else {
		if m == nil {
			m = New_S2C_GameStart()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_GameStart(i interface{}) {
	if m, ok := i.(*S2C_GameStart); ok && m != nil {
		g_S2C_GameStart_Pool.Put(i)
	}
}

// message [S2C_GameStart] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_Turn] begin
func (m *S2C_Turn) ResetEx() {
	m.Pos = 0
	m.Auto = 0

}

func (m S2C_Turn) Clone() *S2C_Turn {
	n, ok := g_S2C_Turn_Pool.Get().(*S2C_Turn)
	if !ok || n == nil {
		n = &S2C_Turn{}
	}

	n.Pos = m.Pos
	n.Auto = m.Auto

	return n
}

func Clone_S2C_Turn_Slice(dst []*S2C_Turn, src []*S2C_Turn) []*S2C_Turn {
	for _, i := range dst {
		Put_S2C_Turn(i)
	}
	dst = []*S2C_Turn{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_Turn() *S2C_Turn {
	m := &S2C_Turn{}
	return m
}

var g_S2C_Turn_Pool = sync.Pool{}

func Get_S2C_Turn() *S2C_Turn {
	m, ok := g_S2C_Turn_Pool.Get().(*S2C_Turn)
	if !ok {
		m = New_S2C_Turn()
	} else {
		if m == nil {
			m = New_S2C_Turn()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_Turn(i interface{}) {
	if m, ok := i.(*S2C_Turn); ok && m != nil {
		g_S2C_Turn_Pool.Put(i)
	}
}

// message [S2C_Turn] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_TurnAction] begin
func (m *C2S_TurnAction) ResetEx() {
	m.Act = 0
	m.Bet = 0

}

func (m C2S_TurnAction) Clone() *C2S_TurnAction {
	n, ok := g_C2S_TurnAction_Pool.Get().(*C2S_TurnAction)
	if !ok || n == nil {
		n = &C2S_TurnAction{}
	}

	n.Act = m.Act
	n.Bet = m.Bet

	return n
}

func Clone_C2S_TurnAction_Slice(dst []*C2S_TurnAction, src []*C2S_TurnAction) []*C2S_TurnAction {
	for _, i := range dst {
		Put_C2S_TurnAction(i)
	}
	dst = []*C2S_TurnAction{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_C2S_TurnAction() *C2S_TurnAction {
	m := &C2S_TurnAction{}
	return m
}

var g_C2S_TurnAction_Pool = sync.Pool{}

func Get_C2S_TurnAction() *C2S_TurnAction {
	m, ok := g_C2S_TurnAction_Pool.Get().(*C2S_TurnAction)
	if !ok {
		m = New_C2S_TurnAction()
	} else {
		if m == nil {
			m = New_C2S_TurnAction()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_TurnAction(i interface{}) {
	if m, ok := i.(*C2S_TurnAction); ok && m != nil {
		g_C2S_TurnAction_Pool.Put(i)
	}
}

// message [C2S_TurnAction] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_TurnAction] begin
func (m *S2C_TurnAction) ResetEx() {
	m.Act = 0
	m.Bet = 0
	m.Pos = 0

}

func (m S2C_TurnAction) Clone() *S2C_TurnAction {
	n, ok := g_S2C_TurnAction_Pool.Get().(*S2C_TurnAction)
	if !ok || n == nil {
		n = &S2C_TurnAction{}
	}

	n.Act = m.Act
	n.Bet = m.Bet
	n.Pos = m.Pos

	return n
}

func Clone_S2C_TurnAction_Slice(dst []*S2C_TurnAction, src []*S2C_TurnAction) []*S2C_TurnAction {
	for _, i := range dst {
		Put_S2C_TurnAction(i)
	}
	dst = []*S2C_TurnAction{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_TurnAction() *S2C_TurnAction {
	m := &S2C_TurnAction{}
	return m
}

var g_S2C_TurnAction_Pool = sync.Pool{}

func Get_S2C_TurnAction() *S2C_TurnAction {
	m, ok := g_S2C_TurnAction_Pool.Get().(*S2C_TurnAction)
	if !ok {
		m = New_S2C_TurnAction()
	} else {
		if m == nil {
			m = New_S2C_TurnAction()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_TurnAction(i interface{}) {
	if m, ok := i.(*S2C_TurnAction); ok && m != nil {
		g_S2C_TurnAction_Pool.Put(i)
	}
}

// message [S2C_TurnAction] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_PublicCard] begin
func (m *S2C_PublicCard) ResetEx() {

	for _, i := range m.Cards {
		Put_Card(i)
	}
	m.Cards = []*Card{}
	m.Best.ResetEx()

}

func (m S2C_PublicCard) Clone() *S2C_PublicCard {
	n, ok := g_S2C_PublicCard_Pool.Get().(*S2C_PublicCard)
	if !ok || n == nil {
		n = &S2C_PublicCard{}
	}

	if len(m.Cards) > 0 {
		for _, i := range m.Cards {
			if i != nil {
				n.Cards = append(n.Cards, i.Clone())
			} else {
				n.Cards = append(n.Cards, nil)
			}
		}
	} else {
		n.Cards = []*Card{}
	}

	n.Best = m.Best.Clone()

	return n
}

func Clone_S2C_PublicCard_Slice(dst []*S2C_PublicCard, src []*S2C_PublicCard) []*S2C_PublicCard {
	for _, i := range dst {
		Put_S2C_PublicCard(i)
	}
	dst = []*S2C_PublicCard{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_PublicCard() *S2C_PublicCard {
	m := &S2C_PublicCard{
		Cards: []*Card{},
		Best:  New_BestCombo(),
	}
	return m
}

var g_S2C_PublicCard_Pool = sync.Pool{}

func Get_S2C_PublicCard() *S2C_PublicCard {
	m, ok := g_S2C_PublicCard_Pool.Get().(*S2C_PublicCard)
	if !ok {
		m = New_S2C_PublicCard()
	} else {
		if m == nil {
			m = New_S2C_PublicCard()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_PublicCard(i interface{}) {
	if m, ok := i.(*S2C_PublicCard); ok && m != nil {
		g_S2C_PublicCard_Pool.Put(i)
	}
}

// message [S2C_PublicCard] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_AutoAction] begin
func (m *C2S_AutoAction) ResetEx() {
	m.Act = 0

}

func (m C2S_AutoAction) Clone() *C2S_AutoAction {
	n, ok := g_C2S_AutoAction_Pool.Get().(*C2S_AutoAction)
	if !ok || n == nil {
		n = &C2S_AutoAction{}
	}

	n.Act = m.Act

	return n
}

func Clone_C2S_AutoAction_Slice(dst []*C2S_AutoAction, src []*C2S_AutoAction) []*C2S_AutoAction {
	for _, i := range dst {
		Put_C2S_AutoAction(i)
	}
	dst = []*C2S_AutoAction{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_C2S_AutoAction() *C2S_AutoAction {
	m := &C2S_AutoAction{}
	return m
}

var g_C2S_AutoAction_Pool = sync.Pool{}

func Get_C2S_AutoAction() *C2S_AutoAction {
	m, ok := g_C2S_AutoAction_Pool.Get().(*C2S_AutoAction)
	if !ok {
		m = New_C2S_AutoAction()
	} else {
		if m == nil {
			m = New_C2S_AutoAction()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_AutoAction(i interface{}) {
	if m, ok := i.(*C2S_AutoAction); ok && m != nil {
		g_C2S_AutoAction_Pool.Put(i)
	}
}

// message [C2S_AutoAction] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_GameOver] begin
func (m *S2C_GameOver) ResetEx() {

}

func (m S2C_GameOver) Clone() *S2C_GameOver {
	n, ok := g_S2C_GameOver_Pool.Get().(*S2C_GameOver)
	if !ok || n == nil {
		n = &S2C_GameOver{}
	}

	return n
}

func Clone_S2C_GameOver_Slice(dst []*S2C_GameOver, src []*S2C_GameOver) []*S2C_GameOver {
	for _, i := range dst {
		Put_S2C_GameOver(i)
	}
	dst = []*S2C_GameOver{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_GameOver() *S2C_GameOver {
	m := &S2C_GameOver{}
	return m
}

var g_S2C_GameOver_Pool = sync.Pool{}

func Get_S2C_GameOver() *S2C_GameOver {
	m, ok := g_S2C_GameOver_Pool.Get().(*S2C_GameOver)
	if !ok {
		m = New_S2C_GameOver()
	} else {
		if m == nil {
			m = New_S2C_GameOver()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_GameOver(i interface{}) {
	if m, ok := i.(*S2C_GameOver); ok && m != nil {
		g_S2C_GameOver_Pool.Put(i)
	}
}

// message [S2C_GameOver] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [Balance] begin
func (m *Balance) ResetEx() {
	m.UserId = 0
	m.BestCombo.ResetEx()
	m.Gain = 0
	m.Refund = 0

	for _, i := range m.Cards {
		Put_Card(i)
	}
	m.Cards = []*Card{}
	m.WinRound = 0

}

func (m Balance) Clone() *Balance {
	n, ok := g_Balance_Pool.Get().(*Balance)
	if !ok || n == nil {
		n = &Balance{}
	}

	n.UserId = m.UserId
	n.BestCombo = m.BestCombo.Clone()
	n.Gain = m.Gain
	n.Refund = m.Refund

	if len(m.Cards) > 0 {
		for _, i := range m.Cards {
			if i != nil {
				n.Cards = append(n.Cards, i.Clone())
			} else {
				n.Cards = append(n.Cards, nil)
			}
		}
	} else {
		n.Cards = []*Card{}
	}

	n.WinRound = m.WinRound

	return n
}

func Clone_Balance_Slice(dst []*Balance, src []*Balance) []*Balance {
	for _, i := range dst {
		Put_Balance(i)
	}
	dst = []*Balance{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_Balance() *Balance {
	m := &Balance{
		BestCombo: New_BestCombo(),
		Cards:     []*Card{},
	}
	return m
}

var g_Balance_Pool = sync.Pool{}

func Get_Balance() *Balance {
	m, ok := g_Balance_Pool.Get().(*Balance)
	if !ok {
		m = New_Balance()
	} else {
		if m == nil {
			m = New_Balance()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_Balance(i interface{}) {
	if m, ok := i.(*Balance); ok && m != nil {
		g_Balance_Pool.Put(i)
	}
}

// message [Balance] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_Balance] begin
func (m *S2C_Balance) ResetEx() {

	for _, i := range m.Balances {
		Put_Balance(i)
	}
	m.Balances = []*Balance{}

}

func (m S2C_Balance) Clone() *S2C_Balance {
	n, ok := g_S2C_Balance_Pool.Get().(*S2C_Balance)
	if !ok || n == nil {
		n = &S2C_Balance{}
	}

	if len(m.Balances) > 0 {
		for _, i := range m.Balances {
			if i != nil {
				n.Balances = append(n.Balances, i.Clone())
			} else {
				n.Balances = append(n.Balances, nil)
			}
		}
	} else {
		n.Balances = []*Balance{}
	}

	return n
}

func Clone_S2C_Balance_Slice(dst []*S2C_Balance, src []*S2C_Balance) []*S2C_Balance {
	for _, i := range dst {
		Put_S2C_Balance(i)
	}
	dst = []*S2C_Balance{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_Balance() *S2C_Balance {
	m := &S2C_Balance{
		Balances: []*Balance{},
	}
	return m
}

var g_S2C_Balance_Pool = sync.Pool{}

func Get_S2C_Balance() *S2C_Balance {
	m, ok := g_S2C_Balance_Pool.Get().(*S2C_Balance)
	if !ok {
		m = New_S2C_Balance()
	} else {
		if m == nil {
			m = New_S2C_Balance()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_Balance(i interface{}) {
	if m, ok := i.(*S2C_Balance); ok && m != nil {
		g_S2C_Balance_Pool.Put(i)
	}
}

// message [S2C_Balance] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_RoomChat] begin
func (m *C2S_RoomChat) ResetEx() {
	m.SrcUserId = 0
	m.DstUserId = 0
	m.Content = ""

}

func (m C2S_RoomChat) Clone() *C2S_RoomChat {
	n, ok := g_C2S_RoomChat_Pool.Get().(*C2S_RoomChat)
	if !ok || n == nil {
		n = &C2S_RoomChat{}
	}

	n.SrcUserId = m.SrcUserId
	n.DstUserId = m.DstUserId
	n.Content = m.Content

	return n
}

func Clone_C2S_RoomChat_Slice(dst []*C2S_RoomChat, src []*C2S_RoomChat) []*C2S_RoomChat {
	for _, i := range dst {
		Put_C2S_RoomChat(i)
	}
	dst = []*C2S_RoomChat{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_C2S_RoomChat() *C2S_RoomChat {
	m := &C2S_RoomChat{}
	return m
}

var g_C2S_RoomChat_Pool = sync.Pool{}

func Get_C2S_RoomChat() *C2S_RoomChat {
	m, ok := g_C2S_RoomChat_Pool.Get().(*C2S_RoomChat)
	if !ok {
		m = New_C2S_RoomChat()
	} else {
		if m == nil {
			m = New_C2S_RoomChat()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_RoomChat(i interface{}) {
	if m, ok := i.(*C2S_RoomChat); ok && m != nil {
		g_C2S_RoomChat_Pool.Put(i)
	}
}

// message [C2S_RoomChat] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_RoomChat] begin
func (m *S2C_RoomChat) ResetEx() {
	m.SrcUserId = 0
	m.DstUserId = 0
	m.Content = ""

}

func (m S2C_RoomChat) Clone() *S2C_RoomChat {
	n, ok := g_S2C_RoomChat_Pool.Get().(*S2C_RoomChat)
	if !ok || n == nil {
		n = &S2C_RoomChat{}
	}

	n.SrcUserId = m.SrcUserId
	n.DstUserId = m.DstUserId
	n.Content = m.Content

	return n
}

func Clone_S2C_RoomChat_Slice(dst []*S2C_RoomChat, src []*S2C_RoomChat) []*S2C_RoomChat {
	for _, i := range dst {
		Put_S2C_RoomChat(i)
	}
	dst = []*S2C_RoomChat{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_RoomChat() *S2C_RoomChat {
	m := &S2C_RoomChat{}
	return m
}

var g_S2C_RoomChat_Pool = sync.Pool{}

func Get_S2C_RoomChat() *S2C_RoomChat {
	m, ok := g_S2C_RoomChat_Pool.Get().(*S2C_RoomChat)
	if !ok {
		m = New_S2C_RoomChat()
	} else {
		if m == nil {
			m = New_S2C_RoomChat()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_RoomChat(i interface{}) {
	if m, ok := i.(*S2C_RoomChat); ok && m != nil {
		g_S2C_RoomChat_Pool.Put(i)
	}
}

// message [S2C_RoomChat] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
