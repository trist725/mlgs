package internal

import (
	"mlgs/src/base"
	"mlgs/src/msg"
	"reflect"
)

func init() {
	regiserMsgHandle(&msg.C2S_Ping{}, handlePong)
}

func regiserMsgHandle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlePong(args []interface{}) {
	//recv := args[0].(*msg.C2S_Ping)
	//test
	send := msg.New_S2C_Pong()
	sender := args[1].(*base.Agent)
	clientID := args[2].(int32)

	ext := [][]byte{base.Int32ToByteArr(clientID)}
	sender.WriteMsgEx(ext, send)
	//sender.WriteCmd(0, clientID)
}
