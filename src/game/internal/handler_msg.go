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
	regiserMsgHandle(&msg.C2S_PlayerLeaveRoom{}, handlePlayerLeaveRoom)
	regiserMsgHandle(&msg.C2S_TurnAction{}, handleTurnAction)
	regiserMsgHandle(&msg.C2S_AutoAction{}, handleAutoAction)
	regiserMsgHandle(&msg.C2S_Ping{}, handlePong)
	regiserMsgHandle(&msg.C2S_RoomChat{}, handleRoomChat)
	regiserMsgHandle(&msg.C2S_UpdateUserData{}, handleUpdateUserData)
}

func regiserMsgHandle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlePlayerLeaveRoom(args []interface{}) {
	// 消息的发送者
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	send := msg.Get_S2C_PlayerLeaveRoom()
	defer sender.WriteMsg(send)

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handlePlayerLeaveRoom return for nil session")
		return
	}

	//对局中不能离开
	player := session.Player()
	if player == nil {
		log.Error("session[%d] without player on handlePlayerLeaveRoom", session.ID())
		return
	}
	if player.InTheGame() {
		log.Debug("player:[%d] in the game, can't leave room", player.UserId())
		send.Err = msg.S2C_PlayerLeaveRoom_E_Err_Playing
		return
	}

	r := room.Mgr().GetRoom(session.Player().RoomId())
	if err := r.PlayerLeave(session.Player()); err != nil {
		send.Err = msg.S2C_PlayerLeaveRoom_E_Err_UnKnown
		return
	}

	ChanRPC.Go("PlayerLeaveRoom", session.UserData().ID, r)

	send.Err = msg.S2C_PlayerLeaveRoom_E_Err_Success

	session.Update()
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
	//defer sender.WriteMsg(send)

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleQuickMatchStart return for nil session")
		return
	}

	player := session.Player()
	if player != nil {
		if player.InRoom() {
			log.Debug("player:[%d] already in room", player.UserId())
			return
		}
	}

	//创建游戏内数据
	player = cache.NewPlayer(session.ID(), session.UserData().ID, sd.InitQuickMatchRoomId(), session.UserData())
	session.SetPlayer(player)

	success := room.Mgr().PlayerJoin(player)
	//无空房,新建
	if !success {
		nr := room.Mgr().NewRoom(1, 1, sd.InitQuickMatchRoomId())
		if err := nr.PlayerJoin(player); err == nil {
			success = true
		}
	}
	//新建房间加入还是有可能失败
	if !success {
		send.Err = msg.S2C_QuickMatchStart_E_Err_Room
		return
	}
	log.Debug("[%s] pos:[%d], room id:[%d]", session.Sign(), player.Pos(), player.RoomId())

	r := room.Mgr().GetRoom(player.RoomId())
	send.Room = msg.Get_Room()
	send.Room.Id = r.Id()
	send.Room.Name = r.Name()

	ChanRPC.Go("PlayerJoinRoom", sender, r, send, player)

	session.Update()
}

func handleDaySign(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.C2S_DaySign)
	// 消息的发送者
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}

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
		log.Debug("sign day invalid")
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
		user.SignedDays = 1
	}

	session.Update()
}

func handleTurnAction(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.C2S_TurnAction)
	// 消息的发送者
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleTurnAction return for nil session")
		return
	}
	player := session.Player()
	if player == nil {
		log.Debug("handleTurnAction return for nil player")
		return
	}

	r := room.Mgr().GetRoom(player.RoomId())
	if r == nil {
		log.Debug("player not in room")
		return
	}
	r.SendPlayerActionSig(recv, player)

	session.Update()
}

func handleAutoAction(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.C2S_AutoAction)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleAutoAction return for nil session")
		return
	}
	player := session.Player()
	if player == nil {
		log.Debug("handleAutoAction return for nil player")
		return
	}

	if recv.Act < 0 || recv.Act > 4 {
		log.Debug("[%s] invalid auto action", session.Sign())
	}
	player.SetAutoAct(recv.Act)

	session.Update()
}

func handlePong(args []interface{}) {
	//recv := args[0].(*msg.C2S_Ping)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handlePong return for nil session")
		return
	}
	session.Update()
}

func handleRoomChat(args []interface{}) {
	recv := args[0].(*msg.C2S_RoomChat)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleRoomChat return for nil session")
		return
	}
	player := session.Player()
	if player == nil {
		log.Debug("handleRoomChat return for nil player")
		return
	}
	r := room.Mgr().GetRoom(player.RoomId())
	if r == nil {
		log.Debug("player not in room")
		return
	}

	ChanRPC.Go("RoomChat", r, recv)
	session.Update()
}

func handleUpdateUserData(args []interface{}) {
	//recv := args[0].(*msg.C2S_UpdateUserData)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleUpdateUserData return for nil session")
		return
	}
	send := msg.Get_S2C_UpdateUserData()
	defer sender.WriteMsg(send)

	send.Data = session.UserData().ToMsg(msg.Get_User())
	send.Err = msg.S2C_UpdateUserData_OK
}
