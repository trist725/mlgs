package cache

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/sd"
)

//todo:保存对局数据,断线重连
type Player struct {
	//session id
	sid uint64
	//对局状态,0-非对局中,1-对局中
	stat uint32
	//对局中的位置
	pos uint32
	//所在房间id
	rid uint64
	/// 筹码
	chip int64
}

func NewPlayer(sid uint64, t int64) *Player {
	//todo:根据t进入不同房间类型
	var rommSd *sd.Room
	switch t {
	default:
		rommSd = sd.RoomMgr.Get(t)
		if rommSd == nil {
			log.Fatal("策划坑爹了,读room表有误，id: [%d]", t)
			return nil
		}
	}

	p := &Player{
		sid:  sid,
		chip: rommSd.Chip,
	}

	return p
}

func (p *Player) Chip() int64 {
	return p.chip
}

func (p *Player) Pos() uint32 {
	return p.pos
}

func (p *Player) SetPos(index uint32) {
	p.pos = index
}

func (p *Player) InRoom() bool {
	if p.rid == 0 {
		return false
	}
	return true
}

func (p *Player) SetRoomId(rid uint64) {
	p.rid = rid
}

func (p *Player) RoomId() uint64 {
	return p.rid
}

func (p *Player) SetSessionId(sid uint64) {
	p.sid = sid
}

func (p *Player) SessionId() uint64 {
	return p.sid
}
