package gate

import (
	"game"
	"login"
	"mlgs/src/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.C2S_Login{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_UpdateUserData{}, game.ChanRPC)
}
