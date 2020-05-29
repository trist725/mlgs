package internal

import (
	"mlgs/src/model"
	"mlgs/src/msg"

	"github.com/trist725/myleaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("AfterLoginAuthPass", OnAfterLoginAuthPass)
	skeleton.RegisterChanRPC("NewGame", OnNewGame)
	skeleton.RegisterChanRPC("TurnAction", OnTurnAction)
	skeleton.RegisterChanRPC("DisConn", OnDisConn)
	skeleton.RegisterChanRPC("RoomChat", OnRoomChat)
	skeleton.RegisterChanRPC("UpdateUserData", OnUpdateUserData)
	skeleton.RegisterChanRPC("UpdateItems", OnUpdateItems)
	skeleton.RegisterChanRPC("SyncGameStatus", OnSyncGameStatus)
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

	//for _, m := range user.Monies {
	//	nm := msg.Get_Money()
	//	send.Monies = append(send.Monies, m.ToMsg(nm))
	//}

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
	//user.AllocDayQuests()
	//更新邮件列表
	//user.UpdateMails()
	send.InTheGame = args[2].(bool)
	sender.WriteMsg(send)
	if send.InTheGame {
		ChanRPC.Go("SyncGameStatus", nil, sender)
	}
}

func OnNewGame(args []interface{}) {
}

func OnTurnAction(args []interface{}) {

}

func OnDisConn(args []interface{}) {
	//uid := args[0].(int64)
}

func OnRoomChat(args []interface{}) {
	//recv := args[1].(*msg.C2S_RoomChat)
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

func OnSyncGameStatus(args []interface{}) {

}
