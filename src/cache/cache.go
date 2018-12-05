package cache

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/sd"
	"reflect"
	"sort"
	"sync/atomic"
)

var (
	gHandCardCount int
	gPlayerLimit   uint32
)

func init() {
	roomSd := sd.RoomMgr.Get(1)
	if roomSd == nil {
		log.Fatal("策划坑爹?!...read room sd error")
		return
	}
	gHandCardCount = roomSd.Handcard
	gPlayerLimit = uint32(roomSd.Chairlimit)
}

type Op struct {
	//1-让牌,2-弃牌,3-跟注,4-加注,5-Allin,6-大小盲第一轮默认操作
	Op int32
	//操作的筹码数
	Bet int64
	//轮次
	Stage uint32
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
	//对局状态,0-非对局中,1-对局中,2-弃牌,3-无筹码
	stat uint32
	//对局中的位置
	pos uint32
	//所在房间id
	rid uint64
	// 筹码
	chip int64
	//本局已押注
	totalBet int64
	/// 角色, 0-普通玩家,1-庄家,2-小盲,3-大盲,4-占位观战
	role uint32
	//手牌
	cards []Card
	//最大牌组
	nuts CardSlice
	//最大牌组类型
	///10-皇家同花顺,9-同花顺,8-四条(金刚),7-葫芦,6-通话
	///5-顺子,4-三条,3-两队,2-对子,1-高牌
	nutsLevel int32
	//当前勾选的自动操作
	//0-无勾选,1-让牌,2-弃牌,3-跟注num,4-跟任何注
	autoAct int32
	//结算用,赢的筹码,负数为输
	gain int64
	//结算用,可被赢走的筹码
	refundBet int64
	//机器人
	robot bool

	//操作集
	ops []Op
}

func (p *Player) AddOp(op Op) {
	p.ops = append(p.ops, op)
}

func (p *Player) RefundBet() int64 {
	return atomic.LoadInt64(&p.refundBet)
}

func (p *Player) SetRefundBet(b int64) {
	atomic.StoreInt64(&p.refundBet, b)
}

func (p *Player) Robot() bool {
	return p.robot
}

func (p *Player) SetRobot(r bool) {
	p.robot = r
}

func (p *Player) Gain() int64 {
	return atomic.LoadInt64(&p.gain)
}

func (p *Player) SetGain(g int64) {
	atomic.StoreInt64(&p.gain, g)
}

func (p *Player) AddGain(g int64) {
	atomic.StoreInt64(&p.gain, p.gain+g)
}

func (p *Player) ClearOps() {
	p.ops = nil
}

func (p *Player) ClearCards() {
	p.cards = nil
}

func (p *Player) ClearNuts() {
	p.nuts = nil
}

func (p *Player) NutsLevel() int32 {
	return atomic.LoadInt32(&p.nutsLevel)
}

func (p *Player) SetNutsLevel(l int32) {
	atomic.StoreInt32(&p.nutsLevel, l)
}

//todo: 记录每次下注操作
func (p *Player) Bet(b int64) {
	if p.chip < b {
		log.Error("not enough chip:[%d] to bet:[%d]", p.chip, b)
		return
	}
	atomic.StoreInt64(&p.chip, p.chip-b)
	atomic.StoreInt64(&p.totalBet, p.totalBet+b)
	if p.chip == 0 {
		p.stat = 3
	}
}

func (p *Player) AutoAct() int32 {
	return atomic.LoadInt32(&p.autoAct)
}

func (p *Player) SetAutoAct(a int32) {
	atomic.StoreInt32(&p.autoAct, a)
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

func NewRobotPlayer(rid int64, t int64) *Player {
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
		uid:  rid,
		chip: rommSd.Chip,
	}
	p.SetRobot(true)
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

func (p *Player) Stat() uint32 {
	return atomic.LoadUint32(&p.stat)
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

func (p *Player) TotalBet() int64 {
	return atomic.LoadInt64(&p.totalBet)
}

func (p *Player) SetTotalBet(tb int64) {
	atomic.StoreInt64(&p.totalBet, tb)
}

func (p *Player) HadAction(stage uint32) bool {
	if stage < 0 || stage > 5 {
		log.Error("HadAction: invalid stage: %d", stage)
		return false
	}
	for _, op := range p.ops {
		if op.Stage == stage {
			//排除大小盲第一轮默认操作
			if op.Op != 6 {
				return true
			}
		}
	}
	return false
}

func (p *Player) GetBetByStage(stage uint32) int64 {
	if stage < 0 || stage > 5 {
		log.Error("GetBetByStage: invalid stage: %d", stage)
		return -1
	}
	var bet int64
	for _, op := range p.ops {
		if op.Stage == stage {
			bet += op.Bet
		}
		if op.Stage > stage {
			break
		}
	}
	return bet
}

func (p *Player) Nuts() CardSlice {
	return p.nuts
}

func (p *Player) UpdateNuts(new CardSlice) {
	p.nuts = nil
	p.nuts = append(p.nuts, new...)
}

func (p *Player) CalNuts(pc CardSlice) {
	switch pc.Len() {
	case 0:
		if len(p.cards) != gHandCardCount {
			log.Error("invalid hand card count")
			return
		}
		p.UpdateNuts(p.cards)
		p.SetNutsLevel(p.nuts.CalLevel())
	case 1:
		log.Error("nani?")
	case 2:
		log.Error("impossible")
	default:
		//>=3
		//7选5
		var calCards CardSlice
		calCards = append(calCards, p.cards...)
		calCards = append(calCards, pc...)
		for i := 0; i < calCards.Len()-2; i++ {
			for j := i + 1; j < calCards.Len()-1; j++ {
				for k := j + 1; k < calCards.Len(); k++ {
					for m := k + 1; m < calCards.Len(); m++ {
						for n := m + 1; n < calCards.Len(); n++ {
							var cards CardSlice
							cards = append(cards, calCards[i])
							cards = append(cards, calCards[j])
							cards = append(cards, calCards[k])
							cards = append(cards, calCards[m])
							cards = append(cards, calCards[n])
							//先排序,降序
							sort.Sort(CardSlice(cards))
							//计算牌型等级
							newLvl := cards.CalLevel()
							curLvl := p.nutsLevel
							if curLvl < newLvl ||
								(p.nuts.Len() == 2 && pc.Len() == 3) {
								p.SetNutsLevel(newLvl)
								p.UpdateNuts(cards)
							} else if curLvl == newLvl {
								//只有同级比较才有意义
								p.CompareCards(cards)
							} //else next
						}
					}
				}
			}
		}
	}
}

func (p *Player) CompareCards(cs2 CardSlice) CardSlice {
	if p.nuts.Len() != cs2.Len() {
		log.Error("diff len CardSlice can't compare")
		return nil
	}

	var bigger CardSlice
	switch p.nutsLevel {
	case 10:
		bigger = p.nuts.RoyalFlushCompare()
	case 9:
		bigger = p.nuts.StraightFlushCompare(cs2)
	case 8:
		bigger = p.nuts.FourOfAKindCompare(cs2)
	case 7:
		bigger = p.nuts.FullHouseCompare(cs2)
	case 6:
		bigger = p.nuts.FlushCompare(cs2)
	case 5:
		bigger = p.nuts.StraightCompare(cs2)
	case 4:
		bigger = p.nuts.TriOfAKindCompare(cs2)
	case 3:
		bigger = p.nuts.TwoPairCompare(cs2)
	case 2:
		bigger = p.nuts.OnePairCompare(cs2)
	case 1:
		bigger = p.nuts.HighCardCompare(cs2)
	}

	if bigger != nil {
		p.UpdateNuts(bigger)
	}
	return bigger
}

type PlayerSlice []*Player

func (ps PlayerSlice) Len() int { // 重写 Len() 方法
	return len(ps)
}
func (ps PlayerSlice) Swap(i, j int) { // 重写 Swap() 方法
	ps[i], ps[j] = ps[j], ps[i]
}

// 重写 Less() 方法，以牌型大小排序
// 从大到小排,i大就是true
func (ps PlayerSlice) Less(i, j int) bool {
	if ps[j].nutsLevel < ps[i].nutsLevel {
		return true
	} else if ps[j].nutsLevel > ps[i].nutsLevel {
		return false
	} else {
		cs := ps[i].CompareCards(ps[j].Nuts())
		//平牌,比位置
		if cs == nil {
			//小盲位先行动
			if ps[i].Role() == 2 {
				return true
			} else if ps[j].Role() == 2 {
				return false
			} else {
				//1,6特例
				if ps[i].Pos() == 1 && ps[j].Pos() == gPlayerLimit {
					return false
				} else if ps[i].Pos() == gPlayerLimit && ps[j].Pos() == 1 {
					return true
				} else { //pos小的先行动,吃亏
					if ps[i].Pos() < ps[j].Pos() {
						return true
					} else {
						return false
					}
				}
			}
		} else if reflect.DeepEqual(cs, ps[i].Nuts()) {
			return true
		} else { //if reflect.DeepEqual(cs, ps[j].Nuts())
			return false
		}
	}
}
