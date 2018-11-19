package cache

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/sd"
	"sync/atomic"
)

type Card struct {
	//花色,1-黑桃(Spade),2-红桃(Heart),3-方块(Diamond),4-梅花(Club)
	Color uint8
	//牌值,2-14
	Num uint8
}

type Op struct {
	//操作
	Op int32
	//操作的筹码数
	bet int64
}

//todo:保存对局数据,断线重连
type Player struct {
	//session id
	//sid为0表示掉线
	sid uint64
	// user id
	//todo:掉线后根据uid操作
	//todo:重连后判断有无快照数据,是否在对局中
	uid int64
	//对局状态,0-非对局中,1-对局中
	stat uint32
	//对局中的位置
	pos uint32
	//所在房间id
	rid uint64
	/// 筹码
	chip int64
	/// 角色, 0-普通玩家,1-庄家,2-小盲,3-大盲,4-占位观战
	role uint32
	//手牌
	cards []Card
	//当前勾选的自动操作
	//0-无勾选,1-让牌,2-弃牌,3-跟注,4-跟任何注
	autoAct int32
	//是否已自动操作,操作过后的自动操作每轮重置为0
	autoActCount int32

	//操作集
	Ops []Op
}

func (c *Card) Equal(card Card) bool {
	if c.Num == card.Num && c.Color == card.Color {
		return true
	}
	return false
}

func (p *Player) AutoAct() int32 {
	return atomic.LoadInt32(&p.autoAct)
}

func (p *Player) SetAutoAct(act int32) {
	atomic.StoreInt32(&p.autoAct, act)
}

func (p *Player) AutoActCount() int32 {
	return atomic.LoadInt32(&p.autoActCount)
}

func (p *Player) SetAutoActCount(c int32) {
	atomic.StoreInt32(&p.autoActCount, c)
}

func (p *Player) UserId() int64 {
	return atomic.LoadInt64(&p.uid)
}

func (p *Player) SetUserId(uid int64) {
	atomic.StoreInt64(&p.uid, uid)
}

func (p *Player) GetCard(card Card) {
	p.cards = append(p.cards, card)
}

func (p *Player) Cards() []Card {
	return p.cards
}

func NewPlayer(sid uint64, uid int64, t int64) *Player {
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
		uid:  uid,
		chip: rommSd.Chip,
	}
	//todo:扣款

	return p
}

func (p *Player) SetRole(r uint32) {
	if r < 0 || r > 3 {
		log.Error("set role failed, invalid role: [%d]", r)
		return
	}
	p.role = r
}

func (p *Player) Role() uint32 {
	return atomic.LoadUint32(&p.role)
}

func (p *Player) Chip() int64 {
	return atomic.LoadInt64(&p.chip)
}

func (p *Player) SetChip(c int64) {
	atomic.StoreInt64(&p.chip, c)
}

func (p *Player) Pos() uint32 {
	return atomic.LoadUint32(&p.pos)
}

func (p *Player) SetPos(index uint32) {
	atomic.StoreUint32(&p.pos, index)
}

func (p *Player) SetStat(s uint32) {
	atomic.StoreUint32(&p.stat, s)
}

func (p *Player) InRoom() bool {
	if atomic.LoadUint64(&p.rid) == 0 {
		return false
	}
	return true
}

func (p *Player) InTheGame() bool {
	if atomic.LoadUint32(&p.stat) == 0 {
		return false
	}
	return true
}

func (p *Player) SetRoomId(rid uint64) {
	atomic.StoreUint64(&p.rid, rid)
}

func (p *Player) RoomId() uint64 {
	return atomic.LoadUint64(&p.rid)
}

func (p *Player) SetSessionId(sid uint64) {
	atomic.StoreUint64(&p.sid, sid)
}

func (p *Player) SessionId() uint64 {
	return atomic.LoadUint64(&p.sid)
}
