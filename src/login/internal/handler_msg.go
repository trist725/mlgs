package internal

import (
	"mlgs/src/sd"
	"reflect"

	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"

	"mlgs/src/msg"
)

func init() {
	regiserMsgHandle(&msg.C2S_Login{}, handleLogin)
}

func regiserMsgHandle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.C2S_Login)
	// 消息的发送者
	sender := args[1].(gate.Agent)
	send := &msg.S2C_Login{}
	//close必须在send之后
	closeFlag := false
	defer func() {
		sender.WriteMsg(send)
		if closeFlag {
			log.Debug("close agent...[%v]", sender.RemoteAddr())
			sender.Close()
		}
	}()

	if recv.VerifyCode != sd.InitClientVerifyCode() {
		send.Err = msg.S2C_Login_Failed
		closeFlag = true
	}

}
