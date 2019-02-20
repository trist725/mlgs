package gate

import (
	"mlgs/src/game"
	"mlgs/src/login"
	"mlgs/src/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.C2S_Login{}, login.ChanRPC)

	msg.Processor.SetRouter(&msg.C2S_DaySign{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_QuickMatchStart{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_PlayerLeaveRoom{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_AutoAction{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_TurnAction{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_Ping{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_RoomChat{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_UpdateUserData{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_GetCompletedAchievements{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_GetAllQuests{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_GetQuestReward{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_GetMailList{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_GetMailReward{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_GetAllMailReward{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_GetOwnItems{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_GetOwnDealerSkins{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_UsingOwnDealerSkins{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_BuyItem{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_SwitchHallRoleSex{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_GetNotices{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_SyncGameStatus{}, game.ChanRPC)
}
