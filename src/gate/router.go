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
}
