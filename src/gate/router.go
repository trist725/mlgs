package gate

import (
	"mlgs/src/game"
	"mlgs/src/login"
	"mlgs/src/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.C2S_Login{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_Ping{}, game.ChanRPC)

}
