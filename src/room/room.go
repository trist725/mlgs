package room

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
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

func (r *Room) PlayerLeave(p *cache.Player) {
	if p == nil {
		log.Error("PlayerLeave faild, invaild player")
		return
	}
	r.Lock()
	defer r.Unlock()
	delete(r.players, p.Pos())
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

func (r *Room) BoardCast(me *cache.Player) {
	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		//todo: 断线session被销毁但等待重连?
		if session == nil {
			log.Error("use nil session id:[%d]", player.SessionId())
			return
		}
		//自己
		//if player.SessionId() == sid{
		//	return
		//}

		//p := msg.Get_Player()
		//p.Chip = player.Chip()
		//p.NickName = session.UserData().NickName
		//p.UserId = session.UserData().ID
		//p.Pos = player.Pos()
		//p.AvatarURL = session.UserData().AvatarURL
		//
		//send.Room.Players = append(send.Room.Players, p)
	})
}
