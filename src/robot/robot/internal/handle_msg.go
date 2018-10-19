package internal

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/msg"
	"reflect"
)

func init() {
	regiserMsgHandle(&msg.S2C_Login{}, handleLogin)
}

func regiserMsgHandle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.S2C_Login)

	log.Debug("recv [%v]", recv)
}
