package msg

import (
	protocol "github.com/trist725/mgsu/network/protocol/protobuf/v2"
	"github.com/trist725/myleaf/network/protobuf"
)

var Protocol = protocol.New(nil, nil, nil)

// 使用 Protobuf 消息处理器
var Processor = protobuf.NewProcessor()

func init() {
	Processor.SetByteOrder(true)

	Processor.Register(&C2S_Ping{})
	Processor.Register(&S2C_Pong{})
	Processor.Register(&C2S_Login{})
	Processor.Register(&S2C_Login{})
	Processor.Register(&S2C_LoginInfo{})
	Processor.Register(&C2S_DaySign{})
	Processor.Register(&S2C_DaySign{})
	Processor.Register(&C2S_QuickMatchStart{})
	Processor.Register(&S2C_QuickMatchStart{})
	Processor.Register(&C2S_PlayerLeaveRoom{})
	Processor.Register(&S2C_PlayerLeaveRoom{})
	Processor.Register(&S2C_UpdatePlayerJoinRoom{})
	Processor.Register(&S2C_UpdatePlayerLeaveRoom{})
	Processor.Register(&S2C_GameStart{})
	Processor.Register(&S2C_Turn{})
	Processor.Register(&C2S_TurnAction{})
	Processor.Register(&S2C_TurnAction{})
	Processor.Register(&S2C_DisConn{})
	Processor.Register(&C2S_AutoAction{})
	Processor.Register(&S2C_PublicCard{})
	Processor.Register(&S2C_GameOver{})
	Processor.Register(&S2C_Balance{})
	Processor.Register(&C2S_RoomChat{})
	Processor.Register(&S2C_RoomChat{})
	Processor.Register(&S2C_UpdateUserData{})
	Processor.Register(&C2S_UpdateUserData{})
	Processor.Register(&S2C_UpdateItems{})
	Processor.Register(&C2S_GetAllQuests{})
	Processor.Register(&S2C_GetAllQuests{})
	Processor.Register(&C2S_GetQuestReward{})
	Processor.Register(&S2C_GetQuestReward{})
	Processor.Register(&C2S_GetCompletedAchievements{})
	Processor.Register(&S2C_GetCompletedAchievements{})
}
