package agent

import (
	"mlgs/src/msg"
	"mlgs/src/robot/robot"
	"reflect"

	"github.com/trist725/myleaf/log"
	"github.com/trist725/myleaf/network"
	"github.com/trist725/myleaf/network/protobuf"
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

func (a *Agent) SetUserData(userData interface{}) {
	a.userData = userData
}
func (a *Agent) UserData() interface{} {
	return a.userData
}

func (a *Agent) SendSome() {
	//登陆
	//{
	//	send := msg.Get_C2S_Login()
	//	send.UID = util.GenRandomString(10)
	//	send.NickName = util.GenRandomString(5)
	//	send.Sex = "sex"
	//	send.Location = "loc"
	//	send.Password = "pwd"
	//	send.Logintype = msg.C2S_Login_E_LoginType_WeChat
	//	a.SetUserData(send)
	//	a.WriteMsg(send)
	//}
	{
		send := msg.New_C2S_Ping()
		a.SetUserData(a.conn.LocalAddr().String())
		a.WriteMsg(send)
		//count++
	}
	//等待登陆成功
	//time.Sleep(1 * time.Second)
	//签到
	{
		send := msg.Get_C2S_DaySign()
		send.Day = 1
		//a.WriteMsg(send)
	}
	//快速匹配
	{
		//send := msg.Get_C2S_QuickMatchStart()
		//a.WriteMsg(send)
	}
	//玩家离开房间请求
	{

	}
	//获取邮件列表
	{
		//send := msg.Get_C2S_GetMailList()
		//a.WriteMsg(send)
	}
	{
		//send := msg.Get_S2C_GetAllMailReward()
		//a.WriteMsg(send)
	}
}

func (a *Agent) Init() {
	a.Processor = protobuf.NewProcessor()
	a.Processor.(*protobuf.Processor).SetByteOrder(true)

	a.Processor.(*protobuf.Processor).Register(&msg.C2S_Ping{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_Pong{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_Login{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_Login{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_LoginInfo{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_DaySign{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_DaySign{})

	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_LoginInfo{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_Login{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_DaySign{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_Pong{}, robot.ChanRPC)
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
			err = a.Processor.Route(msg, a.userData)
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
