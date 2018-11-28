package internal

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/msg"
	"reflect"
)

func init() {
	regiserMsgHandle(&msg.S2C_Login{}, handleLogin)
	regiserMsgHandle(&msg.S2C_QuickMatchStart{}, handleQuickMatchStart)
	regiserMsgHandle(&msg.S2C_PlayerLeaveRoom{}, handlePlayerLeaveRoom)
	regiserMsgHandle(&msg.S2C_UpdatePlayerJoinRoom{}, handleUpdatePlayerJoinRoom)
	regiserMsgHandle(&msg.S2C_UpdatePlayerLeaveRoom{}, handleUpdatePlayerLeaveRoom)
	regiserMsgHandle(&msg.S2C_GameStart{}, handleGameStart)
	regiserMsgHandle(&msg.S2C_GameOver{}, handleGameOver)
	regiserMsgHandle(&msg.S2C_Balance{}, handleBalance)
}

func regiserMsgHandle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.S2C_Login)
	if recv.Reason == msg.S2C_Login_E_Err_LoginSuccess {
		log.Debug("login success")
	}
}

func handleQuickMatchStart(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.S2C_QuickMatchStart)

	//log.Debug("Err:[%d]", recv.Err)
	//log.Debug("room id:[%d], name:[%s], chip:[%d], maxBet:[%d]",
	//	recv.Room.Id, recv.Room.Name, recv.Room.Chip, recv.Room.MaxBet)

	for _, p := range recv.Room.Players {
		log.Debug("players :[%v]", p)
	}
	log.Debug("------------------------------")
}

func handlePlayerLeaveRoom(args []interface{}) {
	// 收到的消息
	//recv := args[0].(*msg.S2C_QuickMatchStart)
	//
	//log.Debug("Err:[%d]", recv.Err)
	//log.Debug("room id:[%d], name:[%s], chip:[%d], maxBet:[%d]",
	//	recv.Room.Id, recv.Room.Name, recv.Room.Chip, recv.Room.MaxBet)
	//
	//for _, p := range recv.Room.Players {
	//	log.Debug("players :[%v]", p)
	//}

}

func handleUpdatePlayerJoinRoom(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.S2C_UpdatePlayerJoinRoom)

	log.Debug("handleUpdatePlayerJoinRoom, %v", recv.Players[0])

}

func handleUpdatePlayerLeaveRoom(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.S2C_UpdatePlayerLeaveRoom)

	log.Debug("handleUpdatePlayerLeaveRoom, %v", recv.UserId)
	//log.Debug("Err:[%d]", recv.Err)
	//log.Debug("room id:[%d], name:[%s], chip:[%d], maxBet:[%d]",
	//	recv.Room.Id, recv.Room.Name, recv.Room.Chip, recv.Room.MaxBet)
	//
	//for _, p := range recv.Room.Players {
	//	log.Debug("players :[%v]", p)
	//}

}

func handleGameStart(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.S2C_GameStart)

	log.Debug("handleGameStart, %v", recv)

}

func handleGameOver(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.S2C_GameStart)

	log.Debug("handleGameOver, %v", recv)

}

func handleBalance(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.S2C_GameStart)

	log.Debug("handleBalance, %v", recv)

}
