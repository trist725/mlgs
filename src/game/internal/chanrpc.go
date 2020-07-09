package internal

import (
	"mlgs/src/model"
	s "mlgs/src/session"
	"time"

	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("LoginAuthPass", rpcHandleLoginAuthPass)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

const saveIntervalMinute = 3 // 保存数据间隔（单位：分钟）

//所有handle消息需在此函数执行完后执行
func rpcHandleLoginAuthPass(args []interface{}) {
	a := args[0].(gate.Agent)
	//conn已断开则不必做下一步动作
	if a.UserData() != nil && a.UserData().(uint64) == 0 {
		return
	}
	account := args[1].(model.Account)
	user := args[2].(model.User)
	if account.ID == 0 || user.ID == 0 {
		log.Error("invalid account or user")
	}
	ns := s.New(a, &account, &user)
	sc := s.Mgr().Count()
	log.Debug("current session count: %d", sc)
	s.Mgr().CheckTick(skeleton)

	//下发用户数据
	ChanRPC.Go("AfterLoginAuthPass", a, &user, false)

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
