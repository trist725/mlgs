package gate

import (
	"mlgs/src/game"
	"mlgs/src/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.C2S_Login{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.S2C_LoginFailed{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.S2C_LoginSucced{}, game.ChanRPC)
}
