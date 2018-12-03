package room

import (
	"fmt"
	"github.com/trist725/mgsu/util"
	"github.com/trist725/myleaf/log"
	"github.com/trist725/myleaf/module"
	q "github.com/yireyun/go-queue"
	"mlgs/src/cache"
	"mlgs/src/msg"
	"mlgs/src/sd"
	s "mlgs/src/session"
	"sync"
	"sync/atomic"
	"time"
)

var (
	gRoomId uint64
	//房间玩家上限
	gPlayerLimit int
	//房间旁观者上限
	gBystanderLimit int
	//公共牌总数
	gCommunityCardCount int
)

//共4*13=52,不包括大小王
const gCardCount uint32 = 4*13 + 0

type Room struct {
	sync.RWMutex
	id   uint64
	name string

	//对局阶段, 0-等待/准备阶段,>0对局开始:
	//1-4第1到4阶段,5-结算
	stage uint32
	//初始小盲注
	sb int64
	//初始大盲注
	bb int64
	//小盲注位置
	sbPos uint32
	//大盲注位置
	bbPos uint32
	//庄家位置
	dPos uint32
	//当前轮询位置
	curPos uint32
	//上次加注位
	raisePos uint32
	//该轮最大下注
	maxBet int64

	//局数
	//对局类型,1-普通赛,2-积分赛
	pType uint32
	//游戏类型,1-德州扑克
	gType uint32
	//当前房间内的玩家, 对局中的位置做key
	players map[uint32]*cache.Player
	//当前房间内的旁观者,用户id做key
	bystanders map[int64]*cache.Player

	//底池
	pot int64

	//牌池
	cardPool *q.EsQueue
	//公共牌
	pc []cache.Card
	//todo:对局循环，make chan
	loopOnce sync.Once

	//循环信号
	stopSig chan struct{}
	//更新准备时间信号
	refreshReadyTimeSig chan struct{}
	//玩家行动信号
	actSig chan TurnAction
}

func init() {
	roomSd := sd.RoomMgr.Get(sd.InitQuickMatchRoomId())
	if roomSd == nil {
		panic("策划坑爹了,读room表快速匹配有误")
	}
	gPlayerLimit = roomSd.Chairlimit
	gBystanderLimit = roomSd.Playerlimit - roomSd.Chairlimit
	gCommunityCardCount = roomSd.Totalpublicpokes
}

func (r *Room) Loop(args ...interface{}) {
	r.loopOnce.Do(func() {
		go r.loop(args)
	})
}

func (r *Room) loop(args []interface{}) {
	defer r.SetStage(0)
	skeleton := args[0].(*module.Skeleton)
	if skeleton == nil {
		log.Error("start room loop failed, skeleton is nil")
		return
	}

	//暂写死1,策划蛋疼的配表
	//游戏准备时间
	timeSd := sd.TimeMgr.Get(1)

	//准备阶段
GAME_READY:
	r.SetStage(0)
	for {
		select {
		//todo: 可优化,手动timer.stop
		case <-time.After(time.Duration(timeSd.Value) * time.Second):
			//是否满足最少开局人数
			if len(r.players) >= sd.InitMinStartGamePlayer() {
				if !r.NewGame(skeleton) {
					continue
				}
				goto GAME_STAGE1
			}
		case <-r.refreshReadyTimeSig:
			//人满开
			if len(r.players) == gPlayerLimit {
				if !r.NewGame(skeleton) {
					continue
				}
				goto GAME_STAGE1
			}
			continue
		case <-r.stopSig:
			return
		}
	}

GAME_STAGE1:
	r.SetStage(1)

	//不要直接用sleep
	//等待客户端发手牌动作
	select {
	case <-time.After(time.Duration(sd.InitDealCardTime()) * time.Second):
		//第一轮大小盲自动下注
		r.FirstStageBlindBet()
		//大盲下一位开始行动
		r.curPos = r.bbPos
		r.Turn(skeleton)
	case <-r.stopSig:
		return
	}
	for {
		curPlayer := r.CurPlayer()
		//没人了去结算
		if curPlayer == nil {
			log.Error("invalid cur player pos:[%d]", r.curPos)
			goto GAME_STAGE5
		}

		if curPlayer.AutoAct() == 0 {
			select {
			case <-time.After(time.Duration(sd.InitActionTime_S1()) * time.Second):
				//超时没动作,弃牌处理
				act := msg.Get_C2S_TurnAction()
				act.Act = 2
				r.DoAct(TurnAction{
					act: act,
					p:   curPlayer,
				})
			case act := <-r.actSig:
				r.DoAct(act)
			case <-r.stopSig:
				return
			}
		} else {
			if !r.DoAutoAct(curPlayer) {
				continue
			}
		}
		if r.GameStat() == 1 {
			//阶段结束
			goto GAME_STAGE2
		} else if r.GameStat() == 2 {
			//不发牌,有唯一胜者
			goto GAME_STAGE5
		} else if r.GameStat() == 4 {
			//发完牌结算
			goto GAME_STAGE5
		}
		//GameStat()返回0和3需轮到下一个
		if !r.Turn(skeleton) {
			goto GAME_STAGE5
		}
	}
GAME_STAGE2:
	r.SetStage(2)
	//发三张公共牌
	r.DealCommunityCard(3)
	r.NewStage(skeleton)
	for {
		curPlayer := r.CurPlayer()
		if curPlayer == nil {
			goto GAME_STAGE5
		}
		if curPlayer.AutoAct() == 0 {
			select {
			case <-time.After(time.Duration(sd.InitActionTime_S2()) * time.Second):
				//超时没动作,弃牌处理
				act := msg.Get_C2S_TurnAction()
				act.Act = 2
				r.DoAct(TurnAction{
					act: act,
					p:   curPlayer,
				})
			case act := <-r.actSig:
				r.DoAct(act)
			case <-r.stopSig:
				return
			}
		} else {
			if !r.DoAutoAct(curPlayer) {
				continue
			}
		}
		if r.GameStat() == 1 {
			//阶段结束
			goto GAME_STAGE3
		} else if r.GameStat() == 2 {
			//不发牌,有唯一胜者
			goto GAME_STAGE5
		} else if r.GameStat() == 4 {
			//发完牌结算
			goto GAME_STAGE5
		}
		//GameStat()返回0和3需轮到下一个
		if !r.Turn(skeleton) {
			goto GAME_STAGE5
		}
	}
GAME_STAGE3:
	r.SetStage(3)
	//发1张公共牌
	r.DealCommunityCard(1)
	r.NewStage(skeleton)
	for {
		curPlayer := r.CurPlayer()
		if curPlayer == nil {
			goto GAME_STAGE5
		}
		if curPlayer.AutoAct() == 0 {
			select {
			case <-time.After(time.Duration(sd.InitActionTime_S3()) * time.Second):
				//超时没动作,弃牌处理
				act := msg.Get_C2S_TurnAction()
				act.Act = 2
				r.DoAct(TurnAction{
					act: act,
					p:   curPlayer,
				})
			case act := <-r.actSig:
				r.DoAct(act)
			case <-r.stopSig:
				return
			}
		} else {
			if !r.DoAutoAct(curPlayer) {
				continue
			}
		}
		if r.GameStat() == 1 {
			//阶段结束
			goto GAME_STAGE4
		} else if r.GameStat() == 2 {
			//不发牌,有唯一胜者
			goto GAME_STAGE5
		} else if r.GameStat() == 4 {
			//发完牌结算
			goto GAME_STAGE5
		}
		//GameStat()返回0和3需轮到下一个
		if !r.Turn(skeleton) {
			goto GAME_STAGE5
		}
	}
GAME_STAGE4:
	r.SetStage(4)
	//发1张公共牌
	r.DealCommunityCard(1)
	r.NewStage(skeleton)
	for {
		curPlayer := r.CurPlayer()
		if curPlayer == nil {
			goto GAME_STAGE5
		}
		if curPlayer.AutoAct() == 0 {
			select {
			case <-time.After(time.Duration(sd.InitActionTime_S4()) * time.Second):
				//超时没动作,弃牌处理
				act := msg.Get_C2S_TurnAction()
				act.Act = 2
				r.DoAct(TurnAction{
					act: act,
					p:   curPlayer,
				})
			case act := <-r.actSig:
				r.DoAct(act)
			case <-r.stopSig:
				return
			}
		} else {
			if !r.DoAutoAct(curPlayer) {
				continue
			}
		}
		if r.GameStat() == 1 {
			//阶段结束
			goto GAME_STAGE5
		} else if r.GameStat() == 2 {
			//不发牌,有唯一胜者
			goto GAME_STAGE5
		} else if r.GameStat() == 4 {
			//发完牌结算
			goto GAME_STAGE5
		}
		//GameStat()返回0和3需轮到下一个
		if !r.Turn(skeleton) {
			goto GAME_STAGE5
		}
	}

GAME_STAGE5:
	r.stage = 5
	//发剩下牌
	if r.GameStat() == 4 {
		if len(r.pc) < gCommunityCardCount {
			r.DealCommunityCard(gCommunityCardCount - len(r.pc))
		}
	}
	//todo:结算
	r.Balance()
	time.Sleep(4 * time.Second)
	r.GameOver()
	goto GAME_READY

}

func (r *Room) DoAutoAct(player *cache.Player) bool {
	ta := msg.Get_C2S_TurnAction()
	//无自动操作,弃牌
	//if player.AutoAct() == 0 ||
	//	player.AutoActCount() <= 0 {
	//		ta.Act = 2
	//		r.DoAct(TurnAction{
	//		act: ta,
	//		p: player,
	//	})
	//}

	if player.AutoAct() == 4 {
		//跟任何注符合allin条件,让玩家确认
		if player.Chip() <= r.maxBet-player.GetBetByStage(r.stage) {
			player.SetAutoAct(0)
			return false
		}

		ta.Act = 3
	} else {
		ta.Act = player.AutoAct()
	}
	log.Debug("DoAutoAct: Act:[%d]------", ta.Act)
	r.DoAct(TurnAction{
		act: ta,
		p:   player,
	})

	//自动操作只生效一次
	player.SetAutoAct(0)

	return true
}

func (r *Room) DoAct(ta TurnAction) {
	if ta.act == nil ||
		ta.p == nil {
		log.Error("invalid parameter on DoAct")
		return
	}

	//todo:记录每步操作,用于对局回放和断线重连
	//todo: check每部合法性
	var bet int64
	//1-让牌,2-弃牌,3-跟注,4-加注,5-Allin
REACT:
	switch ta.act.Act {
	case 1:
		//本轮有下过注,不能让牌,错误操作当弃牌
		//第一轮特殊
		if r.maxBet != 0 &&
			r.stage != 1 {
			ta.act.Act = 2
			goto REACT
		}
		bet = 0
	case 2:
		ta.p.SetStat(2)
		bet = 0
	case 3:
		bet = r.maxBet - ta.p.GetBetByStage(r.stage)
		//此时可以让或加,优先让牌
		if bet == 0 {
			ta.act.Act = 1
			goto REACT
		}
		if bet < 0 {
			log.Error("invalid max bet")
			return
		}
		//筹码不够,allin
		if ta.p.Chip() < bet {
			ta.act.Act = 5
			ta.act.Bet = ta.p.Chip()
			goto REACT
		}
		r.SetPot(r.Pot() + bet)
		ta.p.Bet(bet)
		//实际跟注值
		ta.act.Bet = bet
	case 4:
		//加注错误(开挂)当弃牌
		if ta.act.Bet <= 0 {
			ta.act.Act = 2
			goto REACT
		}
		bet = r.maxBet - ta.p.GetBetByStage(r.stage) + ta.act.Bet
		//筹码不够,allin
		if ta.p.Chip() < bet {
			ta.act.Act = 5
			ta.act.Bet = ta.p.Chip()
			goto REACT
		}
		r.SetPot(r.pot + bet)
		r.maxBet += ta.act.Bet
		ta.p.Bet(bet)
		//更新上次加注位
		r.raisePos = ta.p.Pos()
		log.Debug("raisePos:[%d]", r.raisePos)
	case 5:
		//钱不够allin,改为实际allin值
		if ta.p.Chip() < ta.act.Bet {
			ta.act.Bet = ta.p.Chip()
		}
		//更新最大下注
		if ta.act.Bet > r.maxBet {
			r.maxBet = ta.act.Bet
		}
		r.SetPot(r.pot + ta.act.Bet)
		ta.p.Bet(ta.act.Bet)
		bet = ta.act.Bet
		ta.p.SetStat(3)
	}

	ta.p.AddOp(cache.Op{
		Op:    ta.act.Act,
		Bet:   bet,
		Stage: r.stage,
	})

	//有人加注或allin后要判断是否大于其他人已下的注,改变其他人的自动操作状态
	if ta.act.Act == 4 || ta.act.Act == 5 {
		r.PlayerEach(func(player *cache.Player) {
			//自动让时候有人加注
			if player.AutoAct() == 1 {
				//本轮下注值小于加注者加注后的值,改为弃
				if player.GetBetByStage(r.stage) < ta.p.GetBetByStage(r.stage) {
					player.SetAutoAct(2)
				} else {
					//否则取消自动操作
					player.SetAutoAct(0)
				}
				//跟num值有变化,取消跟num自动操作
			} else if player.AutoAct() == 3 {
				player.SetAutoAct(0)
			}
		})
	}
	log.Debug("player pos:[%d]-----do act:[%d]------bet:[%d]", ta.p.Pos(), ta.act.Act, bet)
	ta.act.Bet = bet
	r.BoardCastTA(ta)
}

func (r *Room) Pot() int64 {
	return atomic.LoadInt64(&r.pot)
}

func (r *Room) SetPot(p int64) {
	atomic.StoreInt64(&r.pot, p)
}

func (r *Room) SetStage(s uint32) {
	atomic.StoreUint32(&r.stage, s)
}

func (r *Room) Stage() uint32 {
	return atomic.LoadUint32(&r.stage)
}

func (r *Room) PlayerJoin(p *cache.Player) bool {
	if p == nil {
		log.Error("addPlayer failed, invalid player")
		return false
	}
	if !r.PlayerIdle() {
		log.Debug("room id:[%d] not idle", r.id)
		return false
	}
	//if int(p.Pos()) > gPlayerLimit{
	//	log.Error("failed to join player, invalid pos")
	//	return false
	//}

	r.Lock()
	defer r.Unlock()

	for i := 1; i <= gPlayerLimit; i++ {
		//座位没人,可分配
		if _, ok := r.players[uint32(i)]; !ok {
			r.players[uint32(i)] = p
			p.SetPos(uint32(i))
			p.SetRoomId(r.id)
			return true
		}
	}

	return false
}

func (r *Room) FirstStageBlindBet() bool {
	if _, ok := r.players[r.sbPos]; !ok {
		log.Error("invalid sbPos on FirstStageBlindBet")
		return false
	}
	if _, ok := r.players[r.bbPos]; !ok {
		log.Error("invalid bbPos on FirstStageBlindBet")
		return false
	}

	//ta := msg.Get_C2S_TurnAction()
	//ta.Act = 4
	//ta.Bet = r.sb
	//r.DoAct(TurnAction{
	//	act: ta,
	//	p: sbPlayer,
	//})
	//
	//ta = msg.Get_C2S_TurnAction()
	//ta.Act = 4
	//ta.Bet = r.bb - r.sb
	//r.DoAct(TurnAction{
	//	act: ta,
	//	p: bbPlayer,
	//})

	r.players[r.sbPos].Bet(r.sb)
	r.players[r.sbPos].AddOp(cache.Op{
		Op:    6,
		Bet:   r.sb,
		Stage: 1,
	})

	r.players[r.bbPos].Bet(r.bb)
	r.players[r.bbPos].AddOp(cache.Op{
		Op:    6,
		Bet:   r.bb,
		Stage: 1,
	})

	r.SetPot(r.bb + r.sb)
	r.maxBet = r.bb
	return true
}

func (r *Room) bystanderJoin(p *cache.Player) bool {
	if p == nil {
		log.Error("addBystander failed, invalid player")
		return false
	}
	if !r.BystanderIdle() {
		return false
	}

	r.Lock()
	defer r.Unlock()

	session := s.Mgr().GetSession(p.SessionId())
	if session == nil {
		log.Error("bystanderJoin，session id:[%d] not exist", p.SessionId())
		return false
	}

	r.bystanders[session.UserData().ID] = p
	return true
}

func (r *Room) PlayerLeave(p *cache.Player) error {
	if p == nil {
		return fmt.Errorf("PlayerLeave failed, invalid player")
	}
	r.Lock()
	defer r.Unlock()

	//已开局,不离开房间
	if r.Stage() != 0 {
		return nil
	}

	if player, ok := r.players[p.Pos()]; ok && player == p {
		delete(r.players, p.Pos())
		player.SetRoomId(0)
		return nil
	}

	return fmt.Errorf("PlayerLeave failed, invalid player or pos:%d", p.Pos())
}

func (r *Room) bystanderLeave(p *cache.Player) {
	if p == nil {
		log.Error("BystanderLeave failed, invalid player")
		return
	}
	r.Lock()
	defer r.Unlock()

	session := s.Mgr().GetSession(p.SessionId())
	if session == nil {
		log.Error("bystanderLeave，session id:[%d] not exist", p.SessionId())
		return
	}

	delete(r.bystanders, session.UserData().ID)
}

func (r *Room) Destroy() {
	r.SendStopLoopSig()
	close(r.refreshReadyTimeSig)
	close(r.stopSig)
	close(r.actSig)

	Mgr().delRoom(r)
}

//是否有空闲座位
func (r *Room) PlayerIdle() bool {
	r.RLock()
	defer r.RUnlock()

	if len(r.players) < gPlayerLimit {
		return true
	}
	return false
}

//是否有空闲观众席
func (r *Room) BystanderIdle() bool {
	r.RLock()
	defer r.RUnlock()

	if len(r.bystanders) < gBystanderLimit {
		return true
	}
	return false
}

func (r *Room) PlayerEach(f func(player *cache.Player)) {
	for _, p := range r.players {
		f(p)
	}
}

func (r *Room) Name() string {
	return r.name
}

func (r *Room) Id() uint64 {
	return r.id
}

func (r *Room) NewGame(args ...interface{}) bool {
	r.SetStage(0)
	r.InitCardPool()
	r.pc = nil

	//todo: 确定玩家是否还有筹码，客户端提示是否兑换，没筹码不兑换踢了或变成游客
	//todo: 如果是大小盲 判断是否筹码大于相应的大小盲注

	skeleton := args[0].(*module.Skeleton)

	r.SetStage(1)
	r.raisePos = 0
	r.ResetPlayers(0)

	//确定庄家
	if !r.AllocRole(1, r.dPos) {

		return false
	}
	//发手牌
	roomSd := sd.RoomMgr.Get(sd.InitQuickMatchRoomId())
	if roomSd == nil {
		log.Error("get room sd failed on NewGame")
		return false
	}
	if !r.DealHandCard(roomSd.Handcard) {
		log.Error("deal hand card failed on NewGame")
		return false
	}

	r.ResetPlayers(1)
	//异步发
	skeleton.ChanRPCServer.Go("NewGame", r)
	return true
}

//重置玩家状态,
func (r *Room) ResetPlayers(stat uint32) {
	r.PlayerEach(func(player *cache.Player) {
		if stat == 1 {
			player.SetStat(stat)
			return
		}
		player.SetStat(0)
		player.ClearOps()
		player.ClearCards()
		player.SetAutoAct(0)
		player.ClearNuts()
		player.SetNutsLevel(0)
		player.SetTotalBet(0)
		player.SetGain(0)
		player.SetRefundBet(0)
	})
}

//洗牌
func (r *Room) InitCardPool() {
	r.cardPool = q.NewQueue(gCardCount)

	arr := make([]int, gCardCount)
	for i := 0; i < int(gCardCount); i++ {
		arr[i] = i
	}
	arr = util.KnuthDurstenfeldShuffle(arr)

	for _, v := range arr {
		var c cache.Card
		num := v + 1
		if uint8(num%4) == 0 {
			c.Color = 4
			c.Num = uint8(num / 4)
		} else {
			c.Color = uint8(num % 4)
			c.Num = uint8(num/4) + 1
		}
		//A当作14
		if c.Num == 1 {
			c.Num = 14
		}
		r.cardPool.Put(c)
	}
}

//role: 0-普通玩家,1-庄家,2-小盲,3-大盲
//p: 原角色位置
func (r *Room) AllocRole(role uint32, p uint32) bool {
	if len(r.players) < 3 {
		log.Error("not enough player to alloc role")
		return false
	}

	//第一局随机取到的第一个玩家为庄家
	if r.dPos == 0 && role == 1 {
		for k, p := range r.players {
			r.dPos = k
			p.SetRole(1)
			r.AllocRole(2, r.dPos)
			return true
		}
	}

	//不是第一局
	for offset := 1; offset < gPlayerLimit; offset++ {
		pos := uint32(offset) + p
		if pos > uint32(gPlayerLimit) {
			pos = pos - uint32(gPlayerLimit)
		}

		player, ok := r.players[pos]
		//该位置没人,下一个
		if !ok {
			continue
		}

		switch role {
		case 1:
			r.dPos = pos
			r.AllocRole(2, r.dPos)
		case 2:
			r.sbPos = pos
			r.AllocRole(3, r.sbPos)
		case 3:
			r.bbPos = pos
			r.AllocRole(0, r.bbPos)
		case 0:
			//设置普通玩家
			if pos == r.dPos {
				return true
			}
			player.SetRole(role)
			continue
		default:
			log.Error("AllocRole failed")
			return false
		}

		player.SetRole(role)
		return true
	}
	//转一圈都没合适的
	log.Error("alloc role failed")
	return false
}

//发手牌
//count: 张数
func (r *Room) DealHandCard(count int) bool {
	var ret = true
	r.PlayerEach(func(p *cache.Player) {
		for i := 0; i < count; i++ {
			card, ok, _ := r.cardPool.Get()
			if !ok {
				log.Error("get card failed")
				ret = false
				return
			}
			p.GetCard(card.(cache.Card))
		}
		p.CalNuts(r.pc)
	})
	return ret
}

//发公共牌
//count: 张数
func (r *Room) DealCommunityCard(count int) {
	var sendCards []cache.Card
	for i := 0; i < count; i++ {
		card, ok, _ := r.cardPool.Get()
		if !ok {
			log.Error("get card failed")
			return
		}
		r.pc = append(r.pc, card.(cache.Card))
		sendCards = append(sendCards, card.(cache.Card))
	}

	//广播
	r.BoardCastDC(sendCards)

	//等待发公共牌动作
	timeSd := sd.TimeMgr.Get(11)
	if timeSd == nil {
		log.Error("策划坑爹。。。。time.xlsx error")
		return
	}
	log.Debug("send CommunityCard---------[%d]", count)
	time.Sleep(time.Duration(timeSd.Value) * time.Second)
}

//return: 0-游戏未结束,阶段不结束,1-游戏未结束,阶段结束,2-有胜者,
//3-出现比牌,轮下一个,4-出现比牌,不轮下一个
func (r *Room) GameStat() int {
	var pRemain, aRemain int
	var remainPlayer *cache.Player
	r.PlayerEach(func(player *cache.Player) {
		//对局中玩家
		if player.Stat() == 1 {
			pRemain++
			remainPlayer = player
			//无筹码(allin)玩家
		} else if player.Stat() == 3 {
			aRemain++
		}
	})
	if pRemain == 1 {
		if aRemain == 0 {
			return 2
		} else if aRemain > 0 {
			if remainPlayer == nil {
				log.Error("GameStat: invalid remainPlayer")
				return -1
			}
			if remainPlayer.GetBetByStage(r.stage) >= r.maxBet {
				return 4
			} else {
				return 3
			}
		}
	} else if pRemain == 0 {
		return 4
	}

	//游戏不结束,判断阶段是否结束
	if r.StageEnd() {
		return 1
	}
	return 0
}

func (r *Room) Turn(skeleton *module.Skeleton) bool {
	pos := r.NextPos()
	if pos == 0 {
		log.Debug("no next pos on Turn")
		return false
	}

	r.curPos = pos
	log.Debug("turn pos: .......%d", pos)
	skeleton.ChanRPCServer.Go("Turn", r)
	return true
}

func (r *Room) CurPlayer() *cache.Player {
	return r.players[r.curPos]
}

func (r *Room) NextPos() uint32 {
	for offset := 1; offset < gPlayerLimit; offset++ {
		pos := uint32(offset) + r.curPos
		if pos > uint32(gPlayerLimit) {
			pos = pos - uint32(gPlayerLimit)
		}

		p, ok := r.players[pos]
		//该位置没人,下一个
		if !ok {
			continue
		}
		//该玩家已弃牌或无筹码,下一个
		if p.Stat() == 2 ||
			p.Stat() == 3 {
			continue
		}

		return pos
	}
	return 0
}

func (r *Room) RaisePrePos() uint32 {
	if r.raisePos <= 0 {
		return 0
	}
	for offset := 1; offset < gPlayerLimit; offset++ {
		pos := r.raisePos - uint32(offset)
		if pos <= 0 {
			pos = pos + uint32(gPlayerLimit)
		}

		_, ok := r.players[pos]
		//该位置没人,下一个
		if !ok {
			continue
		}

		return pos
	}
	return 0
}

func (r *Room) NewStage(skeleton *module.Skeleton) {
	log.Debug("new stage.....")
	//从小盲开始行动
	p, ok := r.players[r.sbPos]
	//该位置没人
	if !ok {
		log.Error("invalid pos on NewStage")
		return
	}
	r.SetCurPos(r.sbPos)
	r.ResetMaxBet()
	r.ResetRaisePos()
	//小盲不在对局中,下一个
	if p.Stat() != 1 {
		r.Turn(skeleton)
		return
	}

	skeleton.ChanRPCServer.Go("Turn", r)
}

//todo:保存记录
func (r *Room) GameOver() {
	r.BoardCastGO()
	r.SetStage(0)
	r.KickOfflinePlayer()
}

func (r *Room) KickOfflinePlayer() {
	r.PlayerEach(func(player *cache.Player) {
		player.SetStat(0)
		if player.SessionId() == 0 {
			r.PlayerLeave(player)
			r.BoardCastPL(player.UserId())
		}
	})
}

func (r *Room) StageEnd() bool {
	//有人加注，到加注者前一位结束
	if r.raisePos != 0 {
		if r.NextPos() == r.raisePos {
			//log.Debug("StageEnd: nextPos:[%d]------raisePos:[%d]", r.NextPos(), r.raisePos)
			return true
		}
		return false
	}
	//无人加注, 所有人轮完结束
	pos := r.NextPos()
	if pos == 0 {
		//log.Error("StageEnd: invalid pos:[%d]", pos)
		return false
	}
	if p, ok := r.players[pos]; ok {
		if p.HadAction(r.stage) {
			return true
		}
	}
	return false
}

func (r *Room) ResetMaxBet() {
	atomic.StoreInt64(&r.maxBet, 0)
}

func (r *Room) ResetRaisePos() {
	atomic.StoreUint32(&r.raisePos, 0)
}

func (r *Room) SetCurPos(cp uint32) {
	atomic.StoreUint32(&r.curPos, cp)
}
