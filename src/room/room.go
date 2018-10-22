package room

import (
	"github.com/trist725/myleaf/log"
	s "mlgs/src/session"
	"sync"
	"sync/atomic"
)

var gRoomId uint64

//todo:保存对局数据,断线重连
type player struct {
	s *s.Session
	//玩家类型,1-对局者,2-旁观者
	t uint32
	//对局状态,0-非对局中,1-对局中
	stat uint32
	//对局中的位置
	pos uint8
	//房间中的角色,1-对局者,2-旁观者
	role uint32
}

type Room struct {
	sync.RWMutex
	id uint64

	//对局类型,1-普通赛,2-积分赛
	pType uint32
	//游戏类型,1-德州扑克
	gType uint32
	//当前房间内的玩家, 对局中的位置做key
	players map[uint8]*player
	//当前房间内的旁观者,session ID做key
	bystanders map[uint64]*player
}

func (r *Room) PlayerJoin(p *player) {
	if p == nil {
		log.Error("addPlayer faild, invaild player")
		return
	}
	r.Lock()
	defer r.Unlock()
	r.players[p.pos] = p
}

func (r *Room) BystanderJoin(p *player) {
	if p == nil {
		log.Error("addBystander faild, invaild player")
		return
	}
	r.Lock()
	defer r.Unlock()
	r.bystanders[p.s.ID()] = p
}

func (r *Room) PlayerLeave(p *player) {
	if p == nil {
		log.Error("PlayerLeave faild, invaild player")
		return
	}
	r.Lock()
	defer r.Unlock()
	delete(r.players, p.pos)
}

func (r *Room) BystanderLeave(p *player) {
	if p == nil {
		log.Error("BystanderLeave faild, invaild player")
		return
	}
	r.Lock()
	defer r.Unlock()
	delete(r.bystanders, p.s.ID())
}

func (r *Room) New(pt uint32, gt uint32, players []*player) *Room {
	room := &Room{
		id:    atomic.AddUint64(&gRoomId, 1),
		pType: pt,
		gType: gt,
	}
	for _, p := range players {
		room.PlayerJoin(p)
	}
	if gRoomManager == nil {
		panic("new room failed, because gRoomManager is nil")
	}
	gRoomManager.putRoom(room)
	return room
}

func (r *Room) Destroy() {

}
