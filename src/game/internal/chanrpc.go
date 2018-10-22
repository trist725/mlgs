package internal

import (
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"mlgs/src/model"
	s "mlgs/src/session"
	"time"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("LoginAuthPass", rpcHandleLoginAuthPass)
	skeleton.RegisterChanRPC("AfterLoginAuthPass", handleAfterLoginAuthPass)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	//_ = a
	a.Destroy()

	//清session
	if a.UserData() != nil {
		sid := a.UserData().(uint64)
		session := s.GetSession(sid)
		if session != nil {
			session.Close()
			log.Debug("[%s] session id:[%d] closed", session.Sign(), sid)
		}
	}

	//设为0表示conn已断开
	a.SetUserData(uint64(0))
}

const saveIntervalMinute = 3 // 保存数据间隔（单位：分钟）

func rpcHandleLoginAuthPass(args []interface{}) {
	a := args[0].(gate.Agent)
	//conn已断开则不必做下一步动作
	if a.UserData() != nil && a.UserData().(uint64) == 0 {
		return
	}
	account := args[1].(*model.Account)
	user := args[2].(*model.User)

	ns := s.New(a, account, user)

	//下发用户数据
	ChanRPC.Go("AfterLoginAuthPass", a, user)

	//定时写库
	var f func()
	f = func() {
		timer := skeleton.AfterFunc(saveIntervalMinute*time.Minute, func() {
			f()
			//实际要做的事
			ns.SaveData()
		})
		ns.SetTimer(timer)
	}
	f()

	log.Debug("[%s] login success", ns.Sign())
}
