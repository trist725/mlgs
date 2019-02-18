package agent

import (
	"github.com/trist725/mgsu/util"
	"github.com/trist725/myleaf/log"
	"github.com/trist725/myleaf/network"
	"github.com/trist725/myleaf/network/protobuf"
	"mlgs/src/msg"
	"mlgs/src/robot/robot"
	"reflect"
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
		//send := msg.Get_C2S_QuickMatchStart()
		//a.WriteMsg(send)
	}
	//玩家离开房间请求
	{

	}
	//获取邮件列表
	{
		send := msg.Get_C2S_GetMailList()
		a.WriteMsg(send)
	}
	//获取邮件列表
	{
		send := msg.Get_C2S_GetMailList()
		a.WriteMsg(send)
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
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_GameOver{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_Balance{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_RoomChat{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_RoomChat{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_UpdateUserData{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_UpdateUserData{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_UpdateItems{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_GetAllQuests{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_GetAllQuests{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_GetQuestReward{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_GetQuestReward{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_GetCompletedAchievements{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_GetCompletedAchievements{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_GetMailList{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_GetMailList{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_GetMailReward{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_GetMailReward{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_GetAllMailReward{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_GetAllMailReward{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_GetOwnItems{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_GetOwnItems{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_GetOwnDealerSkins{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_GetOwnDealerSkins{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_UsingOwnDealerSkins{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_UsingOwnDealerSkins{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_BuyItem{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_BuyItem{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_SwitchHallRoleSex{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_SwitchHallRoleSex{})
	a.Processor.(*protobuf.Processor).Register(&msg.C2S_GetNotices{})
	a.Processor.(*protobuf.Processor).Register(&msg.S2C_GetNotices{})

	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_LoginInfo{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_Login{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_DaySign{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_QuickMatchStart{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_PlayerLeaveRoom{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_UpdatePlayerJoinRoom{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_UpdatePlayerLeaveRoom{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_GameStart{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_GameOver{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_Balance{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_RoomChat{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_UpdateUserData{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_UpdateItems{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_GetAllQuests{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_GetQuestReward{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_GetCompletedAchievements{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_GetMailList{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_GetMailReward{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_GetAllMailReward{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_GetOwnItems{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_GetOwnDealerSkins{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_UsingOwnDealerSkins{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_BuyItem{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_SwitchHallRoleSex{}, robot.ChanRPC)
	a.Processor.(*protobuf.Processor).SetRouter(&msg.S2C_GetNotices{}, robot.ChanRPC)
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
