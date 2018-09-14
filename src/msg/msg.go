package msg

import (
	"github.com/trist725/myleaf/network/protobuf"
)

// 使用 Protobuf 消息处理器
var Processor = protobuf.NewProcessor()

func init() {
	Processor.SetByteOrder(true)
	Processor.Register(&C2S_Login{})
	Processor.Register(&S2C_LoginFailed{})
	Processor.Register(&S2C_LoginSucced{})
}
