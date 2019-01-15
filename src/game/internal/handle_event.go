package internal

import (
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
	"mlgs/src/model"
	"mlgs/src/msg"
	r "mlgs/src/room"
	s "mlgs/src/session"
	"sd"
)

func init() {
	skeleton.RegisterChanRPC("AfterLoginAuthPass", OnAfterLoginAuthPass)
	skeleton.RegisterChanRPC("PlayerJoinRoom", OnPlayerJoinRoom)
	skeleton.RegisterChanRPC("PlayerLeaveRoom", OnPlayerLeaveRoom)
	skeleton.RegisterChanRPC("NewGame", OnNewGame)
	skeleton.RegisterChanRPC("Turn", OnTurn)
	skeleton.RegisterChanRPC("TurnAction", OnTurnAction)
	skeleton.RegisterChanRPC("DisConn", OnDisConn)
	skeleton.RegisterChanRPC("RoomChat", OnRoomChat)
	skeleton.RegisterChanRPC("UpdateUserData", OnUpdateUserData)
	skeleton.RegisterChanRPC("UpdateItems", OnUpdateItems)
}

//每轮签到天数
const signCountPerRound = 14

func OnAfterLoginAuthPass(args []interface{}) {
	sender := args[0].(gate.Agent)
	user := args[1].(*model.User)

	send := &msg.S2C_LoginInfo{
		ID:          user.ID,
		NickName:    user.NickName,
		AvatarURL:   user.AvatarURL,
		DaySigned:   user.DaySigned,
		SignedDays:  user.SignedDays,
		UsingDealer: user.UsingDealer,
	}

	for _, m := range user.Monies {
		nm := msg.Get_Money()
		send.Monies = append(send.Monies, m.ToMsg(nm))
	}

	//暂写死,构造14天签到奖励
	//todo：动态获取14天签到奖励
	for i := 1; i <= signCountPerRound; i++ {
		item := &msg.Item{
			TID: 1,
			Num: 1000 * int64(i),
		}
		send.SignRewards = append(send.SignRewards, item)
	}

	//分配每日任务
	user.AllocDayQuests()
	//更新邮件列表
	user.UpdateMails()

	sender.WriteMsg(send)
}

func OnPlayerJoinRoom(args []interface{}) {
	sender := args[0].(gate.Agent)
	room := args[1].(*r.Room)
	send := args[2].(*msg.S2C_QuickMatchStart)
	defer sender.WriteMsg(send)
	player := args[3].(*cache.Player)

	roomSd := sd.RoomMgr.Get(int64(room.GetRoomType()))
	if roomSd == nil {
		log.Error("OnPlayerJoinRoom(): get room sd failed")
		return
	}

	//给自己发所有玩家信息
	//todo:给自己发旁观者信息
	room.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		//todo: 断线session被销毁但等待重连?
		if session == nil && !player.Robot() {
			log.Error("use nil session, on OnPlayerJoinRoom")
			return
		}
		p := msg.Get_Player()
		p.Chip = roomSd.Chip
		p.UserId = player.UserId()
		p.Pos = player.Pos()
		if player.Robot() {
			p.AvatarURL = player.UserData().AvatarURL
			p.NickName = player.UserData().NickName
		} else {
			p.AvatarURL = session.UserData().AvatarURL
			p.NickName = session.UserData().NickName
		}
		send.Room.Players = append(send.Room.Players, p)
	})
	send.Err = msg.S2C_QuickMatchStart_E_Err_Success

	var players []*cache.Player
	players = append(players, player)
	room.BoardCastPJ(players)

	if room.Stage() == 0 {
		room.Loop(skeleton)
		room.SendRefreshReadyTimeSig()
	}
}

func OnPlayerLeaveRoom(args []interface{}) {
	id := args[0].(int64)
	room := args[1].(*r.Room)
	err := args[2].(msg.S2C_UpdatePlayerLeaveRoom_E_Err)
	//var ids []int64
	//ids = append(ids, id)

	room.BoardCastPL(id, err)
	room.SendRefreshReadyTimeSig()
}

func OnNewGame(args []interface{}) {
	room := args[0].(*r.Room)
	room.BoardCastGS()
}

func OnTurn(args []interface{}) {
	room := args[0].(*r.Room)
	room.BoardCastTurn()
}

func OnTurnAction(args []interface{}) {

}

func OnDisConn(args []interface{}) {
	uid := args[0].(int64)
	room := args[1].(*r.Room)

	room.BoardCastDisConn(uid)
}

func OnRoomChat(args []interface{}) {
	recv := args[1].(*msg.C2S_RoomChat)
	room := args[0].(*r.Room)

	room.BoardCastRC(recv)
}

//todo:
func OnUpdateUserData(args []interface{}) {
	//recv := args[1].(*msg.C2S_RoomChat)
	//room := args[0].(*r.Room)
	//
	//room.BoardCastRC(recv)
}

//todo:
func OnUpdateItems(args []interface{}) {
	//recv := args[1].(*msg.C2S_RoomChat)
	//room := args[0].(*r.Room)
	//
	//room.BoardCastRC(recv)
}
