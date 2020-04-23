package agent

import (
	"github.com/trist725/myleaf/network/protobuf"
	"mlgs/src/msg"
	"mlgs/src/robot/robot"
)

func (a *Agent) SendSome() {
	//登陆
	{
		send := msg.Get_C2S_Login()
		send.UID = "testclient"
		send.NickName = "testclient"
		send.Sex = "sex"
		send.Location = "loc"
		send.Password = "pwd"
		send.Logintype = msg.C2S_Login_E_LoginType_WeChat
		a.WriteMsg(send)
	}
	//签到
	{

	}
}

func (a *Agent) Init() {
	a.Processor = protobuf.NewProcessor()
	a.Processor.(*protobuf.Processor).SetByteOrder(true)

	//这里可能导致id和服务端不匹配
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_Login{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_Login{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_LoginInfo{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_DaySign{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_DaySign{})

	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_LoginInfo{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_Login{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_DaySign{}, robot.ChanRPC)
}
