package gate

import (
	"login"
	"mlgs/src/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.C2S_Login{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.C2S_UpdateUserData{}, login.ChanRPC)
}
