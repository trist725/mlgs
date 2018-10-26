package internal

import (
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
	"mlgs/src/msg"
	"mlgs/src/room"
	"mlgs/src/sd"
	s "mlgs/src/session"
	"reflect"
)

func init() {
	regiserMsgHandle(&msg.C2S_DaySign{}, handleDaySign)
	regiserMsgHandle(&msg.C2S_QuickMatchStart{}, handleQuickMatchStart)
}

func regiserMsgHandle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleQuickMatchStart(args []interface{}) {
	// 收到的消息
	//recv := args[0].(*msg.C2S_QuickMatchStart)
	// 消息的发送者
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	send := msg.Get_S2C_QuickMatchStart()
	defer sender.WriteMsg(send)

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleQuickMatchStart return for nil session")
		return
	}

	player := cache.NewPlayer(session.ID(), sd.InitQuickMatchRoomId())
	session.SetPlayer(player)

	success := room.Mgr().PlayerJoin(player)
	//无空房,新建
	if !success {
		nr := room.Mgr().NewRoom(1, 1, sd.InitQuickMatchRoomId())
		success = nr.PlayerJoin(player)
	}
	//新建房间加入还是有可能失败
	if !success {
		send.Err = msg.S2C_QuickMatchStart_E_Err_Room
		return
	}
	log.Error("[%s] player pos:[%d], room id:[%d]", session.Sign(), player.Pos(), player.RoomId())

	r := room.Mgr().GetRoom(player.RoomId())
	send.Room = msg.Get_Room()
	send.Room.Id = r.Id()
	send.Room.Name = r.Name()
	//给自己发所有玩家信息
	//todo:给自己发旁观者信息
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

		p := msg.Get_Player()
		p.Chip = player.Chip()
		p.NickName = session.UserData().NickName
		p.UserId = session.UserData().ID
		p.Pos = player.Pos()
		p.AvatarURL = session.UserData().AvatarURL

		send.Room.Players = append(send.Room.Players, p)
	})
	//todo: 给其它所有玩家发自己信息

	send.Err = msg.S2C_QuickMatchStart_E_Err_Success
}

func handleDaySign(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.C2S_DaySign)
	// 消息的发送者
	sender := args[1].(gate.Agent)
	send := msg.Get_S2C_DaySign()

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleDaySign return for nil session")
		return
	}

	defer sender.WriteMsg(send)
	user := session.UserData()

	//今日已签到
	if user.DaySigned {
		send.Err = msg.S2C_DaySign_E_Err_AlreadySign
		log.Debug("today already signed")
		return
	}

	//签到天数不对
	if recv.Day != user.SignedDays+1 ||
		recv.Day > signCountPerRound ||
		recv.Day < 0 {
		send.Err = msg.S2C_DaySign_E_Err_Unknown
		log.Debug("sign day invaild")
		return
	}

	send.Err = msg.S2C_DaySign_E_Err_Success
	//给签到奖励
	//todo:目前奖励暂写死
	for _, m := range user.Monies {
		if m.Type == 1 {
			m.Num += int64(recv.Day) * 1000
		}
		nm := msg.Get_Money()
		send.Monies = append(send.Monies, m.ToMsg(nm))
	}

	user.DaySigned = true
	user.SignedDays++
	if user.SignedDays == signCountPerRound+1 {
		user.SignedDays = 0
	}
}
