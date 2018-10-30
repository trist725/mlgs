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

	if len(recv.Room.Players) == 5 {
		for _, p := range recv.Room.Players {
			log.Debug("players :[%v]", p)
		}
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
	//recv := args[0].(*msg.S2C_UpdatePlayerJoinRoom)
	//
	//for _, p := range recv.Players {
	//	log.Debug("players :[%v]", p)
	//}

}

func handleUpdatePlayerLeaveRoom(args []interface{}) {
	// 收到的消息
	//recv := args[0].(*msg.S2C_QuickMatchStart)

	//log.Debug("Err:[%d]", recv.Err)
	//log.Debug("room id:[%d], name:[%s], chip:[%d], maxBet:[%d]",
	//	recv.Room.Id, recv.Room.Name, recv.Room.Chip, recv.Room.MaxBet)
	//
	//for _, p := range recv.Room.Players {
	//	log.Debug("players :[%v]", p)
	//}

}
