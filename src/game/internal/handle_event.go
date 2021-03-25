package internal

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
	//sender := args[0].(gate.Agent)
	//user := args[1].(*model.User)

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
