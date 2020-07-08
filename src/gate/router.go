package gate

import (
	"mlgs/src/login"
	"mlgs/src/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.C2S_Login{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_Ping{}, ChanRPC)
}
