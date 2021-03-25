package msg

import (
	//protocol "github.com/trist725/mgsu/network/protocol/protobuf/v2"
	"github.com/trist725/myleaf/network/protobuf"
)

//var Protocol = protocol.New(nil, nil, nil)

// 使用 Protobuf 消息处理器
var (
	PH        = 0
	Processor = protobuf.NewProcessor()
)

func init() {
	Processor.SetByteOrder(true)

	Processor.Register(&C2S_Ping{})
	Processor.Register(&S2C_Pong{})
	Processor.Register(&S2C_DisConn{})

	Processor.Register(&C2S_Login{})
	Processor.Register(&S2C_Login{})
}
