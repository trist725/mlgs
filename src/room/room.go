package room

import (
	"fmt"
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
	"mlgs/src/msg"
	s "mlgs/src/session"
	"sd"
	"sync"
)

var (
	gRoomId uint64
	//房间玩家上限
	gPlayerLimit int
	//房间旁观者上限
	gBystanderLimit int
)

type Room struct {
	sync.RWMutex
	id   uint64
	name string

	//初始小盲注
	sb int64
	//初始大盲注
	bb int64
	//局数
	//对局类型,1-普通赛,2-积分赛
	pType uint32
	//游戏类型,1-德州扑克
	gType uint32
	//当前房间内的玩家, 对局中的位置做key
	players map[uint32]*cache.Player
	//当前房间内的旁观者,用户id做key
	bystanders map[int64]*cache.Player

	//todo:对局循环
}

func init() {
	roomSd := sd.RoomMgr.Get(sd.InitQuickMatchRoomId())
	if roomSd == nil {
		panic("策划坑爹了,读room表快速匹配有误")
	}
	gPlayerLimit = roomSd.Chairlimit
	gBystanderLimit = roomSd.Playerlimit - roomSd.Chairlimit
}

func (r *Room) PlayerJoin(p *cache.Player) bool {
	if p == nil {
		log.Error("addPlayer faild, invaild player")
		return false
	}
	if !r.PlayerIdle() {
		log.Debug("room id:[%d] not idle", r.id)
		return false
	}
	//if int(p.Pos()) > gPlayerLimit{
	//	log.Error("faild to join player, invaild pos")
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

func (r *Room) bystanderJoin(p *cache.Player) bool {
	if p == nil {
		log.Error("addBystander faild, invaild player")
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
		return fmt.Errorf("PlayerLeave faild, invaild player")
	}
	r.Lock()
	defer r.Unlock()

	if player, ok := r.players[p.Pos()]; ok && player == p {
		delete(r.players, p.Pos())
		return nil
	}

	return fmt.Errorf("PlayerLeave faild, invalid player or pos:%d", p.Pos())
}

func (r *Room) bystanderLeave(p *cache.Player) {
	if p == nil {
		log.Error("BystanderLeave faild, invaild player")
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

//广播玩家加入
func (r *Room) BoardCastPJ(players []*cache.Player) {
	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Error("use nil session id:[%d]", player.SessionId())
			return
		}

		send := msg.Get_S2C_UpdatePlayerJoinRoom()
		for _, p := range players {
			session := s.Mgr().GetSession(p.SessionId())
			if session == nil {
				log.Error("use nil session id:[%d]", player.SessionId())
				return
			}
			np := msg.Get_Player()
			np.NickName = session.UserData().NickName
			np.Pos = p.Pos()
			np.Chip = p.Chip()
			np.AvatarURL = session.UserData().AvatarURL
			np.UserId = session.UserData().ID

			send.Players = append(send.Players, np)
		}

		session.Agent().WriteMsg(send)
	})

	return
}

//广播玩家离开
func (r *Room) BoardCastPL(ids []int64) {
	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Error("use nil session id:[%d]", player.SessionId())
			return
		}

		send := msg.Get_S2C_UpdatePlayerLeaveRoom()
		for _, id := range ids {
			send.UserIds = append(send.UserIds, id)
		}

		session.Agent().WriteMsg(send)
	})

	return
}
