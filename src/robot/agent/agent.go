package agent

import (
	"github.com/trist725/mgsu/util"
	"github.com/trist725/myleaf/log"
	"github.com/trist725/myleaf/network"
	"github.com/trist725/myleaf/network/protobuf"
	"mlgs/src/msg"
	"reflect"
	"robot/robot"
	"time"
)

type Agent struct {
	conn      *network.TCPConn
	Processor network.Processor
	userData  interface{}
}

func NewAgent(conn *network.TCPConn) network.Agent {
	a := new(Agent)
	a.conn = conn
	return a
}

func (a *Agent) SendSome() {
	//登陆
	{
		send := msg.Get_C2S_Login()
		send.UID = util.GenRandomString(10)
		send.NickName = util.GenRandomString(5)
		send.Sex = "sex"
		send.Location = "loc"
		send.Password = "pwd"
		send.Logintype = msg.C2S_Login_E_LoginType_WeChat
		a.WriteMsg(send)
	}
	//等待登陆成功
	time.Sleep(2 * time.Second)
	//签到
	{
		send := msg.Get_C2S_DaySign()
		send.Day = 1
		a.WriteMsg(send)
	}
	//快速匹配
	{
		send := msg.Get_C2S_QuickMatchStart()
		a.WriteMsg(send)
	}
	//玩家离开房间请求
	{

	}
}

func (a *Agent) Init() {
	a.Processor = protobuf.NewProcessor()
	a.Processor.(*protobuf.Processor).SetByteOrder(true)

	//这里可能导致id和服务端不匹配
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_Ping{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_Pong{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_Login{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_Login{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_LoginInfo{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_DaySign{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_DaySign{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_QuickMatchStart{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_QuickMatchStart{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_PlayerLeaveRoom{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_PlayerLeaveRoom{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_UpdatePlayerJoinRoom{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_UpdatePlayerLeaveRoom{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_GameStart{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_Turn{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_TurnAction{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_TurnAction{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_DisConn{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_AutoAction{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_PublicCard{})

	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_LoginInfo{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_Login{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_DaySign{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_QuickMatchStart{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_PlayerLeaveRoom{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_UpdatePlayerJoinRoom{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_UpdatePlayerLeaveRoom{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_GameStart{}, robot.ChanRPC)

}

func (a *Agent) Run() {
	a.Init()

	a.SendSome()

	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			log.Debug("read message: %v", err)
			break
		}

		if a.Processor != nil {
			msg, err := a.Processor.Unmarshal(data)
			if err != nil {
				log.Debug("unmarshal message error: %v", err)
				break
			}
			err = a.Processor.Route(msg, a)
			if err != nil {
				log.Debug("route message error: %v", err)
				break
			}
		}
	}
}

func (a *Agent) OnClose() {
}

func (a *Agent) WriteMsg(msg interface{}) {
	if a.Processor != nil {
		data, err := a.Processor.Marshal(msg)
		if err != nil {
			log.Error("marshal message %v error: %v", reflect.TypeOf(msg), err)
			return
		}
		err = a.conn.WriteMsg(data...)
		if err != nil {
			log.Error("write message %v error: %v", reflect.TypeOf(msg), err)
		}
	}
}
