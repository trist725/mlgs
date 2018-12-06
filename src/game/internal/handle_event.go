package internal

import (
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
	"mlgs/src/model"
	"mlgs/src/msg"
	r "mlgs/src/room"
	s "mlgs/src/session"
)

func init() {
	skeleton.RegisterChanRPC("AfterLoginAuthPass", OnAfterLoginAuthPass)
	skeleton.RegisterChanRPC("PlayerJoinRoom", OnPlayerJoinRoom)
	skeleton.RegisterChanRPC("PlayerLeaveRoom", OnPlayerLeaveRoom)
	skeleton.RegisterChanRPC("NewGame", OnNewGame)
	skeleton.RegisterChanRPC("Turn", OnTurn)
	skeleton.RegisterChanRPC("TurnAction", OnTurnAction)
}

//每轮签到天数
const signCountPerRound = 14

func OnAfterLoginAuthPass(args []interface{}) {
	sender := args[0].(gate.Agent)
	user := args[1].(*model.User)

	send := &msg.S2C_LoginInfo{
		ID:         user.ID,
		NickName:   user.NickName,
		AvatarURL:  user.AvatarURL,
		DaySigned:  user.DaySigned,
		SignedDays: user.SignedDays,
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
			Num: int32(1000 * i),
		}
		send.SignRewards = append(send.SignRewards, item)
	}

	sender.WriteMsg(send)
}

func OnPlayerJoinRoom(args []interface{}) {
	sender := args[0].(gate.Agent)
	room := args[1].(*r.Room)
	send := args[2].(*msg.S2C_QuickMatchStart)
	defer sender.WriteMsg(send)
	player := args[3].(*cache.Player)

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
		p.Chip = player.Chip()
		p.NickName = session.UserData().NickName
		p.UserId = session.UserData().ID
		p.Pos = player.Pos()
		p.AvatarURL = session.UserData().AvatarURL
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
	//var ids []int64
	//ids = append(ids, id)

	//已开局,不离开房间而广播掉线
	//todo:掉线后执行自动操作
	if room.Stage() != 0 {
		room.BoardCastDisConn()
		return
	}

	room.BoardCastPL(id)
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
