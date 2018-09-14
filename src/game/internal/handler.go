package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"mlgs/src/msg"
	"reflect"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.C2S_Login{}, handleLogin)
	handler(&msg.S2C_LoginFailed{}, handleLoginFaild)
	handler(&msg.S2C_LoginSucced{}, handleLoginSuceed)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLoginFaild(args []interface{}) {
	// 收到的消息
	m := args[0].(*msg.S2C_LoginFailed)
	// 消息的发送者
	send := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("get login: %v", m)

	// 给发送者回应一个 Hello 消息
	send.WriteMsg(&msg.S2C_LoginFailed{
		Reason: *proto.Int32(1),
	})
}

func handleLoginSuceed(args []interface{}) {
	// 收到的消息
	m := args[0].(*msg.S2C_LoginSucced)
	// 消息的发送者
	send := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("get login: %v", m)

	// 给发送者回应一个 Hello 消息
	send.WriteMsg(&msg.S2C_LoginSucced{})
}

func handleLogin(args []interface{}) {
	// 收到的消息
	m := args[0].(*msg.C2S_Login)
	// 消息的发送者
	send := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("get login: %v", m)

	// 给发送者回应一个 Hello 消息
	send.WriteMsg(&msg.C2S_Login{})
}
