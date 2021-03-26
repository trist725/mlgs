package room

import (
	"fmt"
	"github.com/trist725/myleaf/log"
	"github.com/trist725/myleaf/module"
	"mlgs/src/sd"
	"sync"
)

var (
	gRoomId uint64
	//房间玩家上限
	gPlayerLimit int64
	//房间旁观者上限
	gBystanderLimit int64
)

//共4*13=52,不包括大小王
const gCardCount uint32 = 4*13 + 0

type Room struct {
	//对局状态
	stat bool

	sync.RWMutex
	id   uint64
	name string

	//当前房间内的玩家-是否准备
	players map[string]bool
	//当前房间内的旁观者
	bystanders []string
}

func init() {
	gPlayerLimit = sd.InitPlayerPerRoom()
	gBystanderLimit = sd.InitBystanderPerRoom()
}

func (r *Room) PlayerJoin(p string) error {
	if p == "" {
		return fmt.Errorf("PlayerJoin room failed, invalid player")
	}
	if !r.HasPlayerSeat() {
		return fmt.Errorf("PlayerJoin room failed, not enough PlayerSeat")
	}
	r.Lock()
	defer r.Unlock()

	r.players[p] = false

	return nil
}

func (r *Room) bystanderJoin(p string) bool {
	if p == "" {
		log.Error("addBystander failed, invalid player")
		return false
	}
	if !r.HasBystanderSeat() {
		return false
	}

	r.Lock()
	defer r.Unlock()

	r.bystanders = append(r.bystanders, p)

	return true
}

func (r *Room) PlayerLeave(p string) bool {
	if p == "" {
		return false
	}
	r.Lock()
	defer r.Unlock()

	if _, ok := r.players[p]; !ok {
		return false
	}

	if r.stat == false {
		delete(r.players, p)
	}
	return true
}

func (r *Room) bystanderLeave(p string) {
	if p == "" {
		log.Error("BystanderLeave failed, invalid player")
		return
	}
	r.Lock()
	defer r.Unlock()

	for i, b := range r.bystanders {
		if b == p {
			r.bystanders = append(r.bystanders[:i], r.bystanders[i+1:]...)
		}
		return
	}
}

func (r *Room) Destroy() {
	Mgr().delRoom(r)
}

//是否有空闲座位
func (r *Room) HasPlayerSeat() bool {
	r.RLock()
	defer r.RUnlock()

	if int64(len(r.players)) < gPlayerLimit {
		return true
	}
	return false
}

//是否有空闲观众席
func (r *Room) HasBystanderSeat() bool {
	r.RLock()
	defer r.RUnlock()

	if int64(len(r.bystanders)) < gBystanderLimit {
		return true
	}
	return false
}

func (r *Room) PlayerEach(f func(player string, ready bool)) {
	for p, b := range r.players {
		f(p, b)
	}
}

func (r *Room) Name() string {
	return r.name
}

func (r *Room) Id() uint64 {
	return r.id
}

func (r *Room) NewGame(args ...interface{}) bool {
	skeleton := args[0].(*module.Skeleton)

	r.ResetPlayers()
	//异步发
	skeleton.ChanRPCServer.Go("NewGame", r)
	return true
}

//重置玩家状态,
func (r *Room) ResetPlayers() {
}

//todo:保存记录
func (r *Room) GameOver() {
	r.ResetPlayers()
	r.KickOffPlayers()
}

func (r *Room) KickOffPlayers() {
}

func (r *Room) GetPlayerCount() int {
	return len(r.players)
}
