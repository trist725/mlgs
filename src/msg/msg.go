package msg

import (
	//protocol "github.com/trist725/mgsu/network/protocol/protobuf/v2"
	"github.com/trist725/myleaf/network/protobuf"
)

//var Protocol = protocol.New(nil, nil, nil)

// 使用 Protobuf 消息处理器
var (
	PH        = 0
	Processor = protobuf.NewProcessor()
)

func init() {
	Processor.SetByteOrder(true)

	Processor.Register(&C2S_Ping{})
	Processor.Register(&S2C_Pong{})
	Processor.Register(&C2S_Login{})
	Processor.Register(&S2C_Login{})
	Processor.Register(&S2C_LoginInfo{})
	Processor.Register(&C2S_DaySign{})
	Processor.Register(&S2C_DaySign{})
	Processor.Register(&S2C_DisConn{})
	Processor.Register(&S2C_UpdateUserData{})
	Processor.Register(&C2S_UpdateUserData{})
	Processor.Register(&S2C_UpdateItems{})
	Processor.Register(&C2S_GetAllQuests{})
	Processor.Register(&S2C_GetAllQuests{})
	Processor.Register(&C2S_GetQuestReward{})
	Processor.Register(&S2C_GetQuestReward{})
	Processor.Register(&C2S_GetCompletedAchievements{})
	Processor.Register(&S2C_GetCompletedAchievements{})
	Processor.Register(&C2S_GetMailList{})
	Processor.Register(&S2C_GetMailList{})
	Processor.Register(&C2S_GetMailReward{})
	Processor.Register(&S2C_GetMailReward{})
	Processor.Register(&C2S_GetAllMailReward{})
	Processor.Register(&S2C_GetAllMailReward{})
	Processor.Register(&C2S_GetOwnItems{})
	Processor.Register(&S2C_GetOwnItems{})
	Processor.Register(&C2S_BuyItem{})
	Processor.Register(&S2C_BuyItem{})
	Processor.Register(&C2S_SwitchHallRoleSex{})
	Processor.Register(&S2C_SwitchHallRoleSex{})
	Processor.Register(&C2S_GetNotices{})
	Processor.Register(&S2C_GetNotices{})
	Processor.Register(&S2C_UpdateMoney{})
}
