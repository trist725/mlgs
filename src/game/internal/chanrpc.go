package internal

import (
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"mlgs/src/model"
	s "mlgs/src/session"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("LoginAuthPass", rpcHandleLoginAuthPass)

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
		mgr := s.GetSessionMgr()
		if mgr == nil {
			panic("gSessionManager is nil")
		}
		if session := mgr.GetSession(sid); session != nil {
			session.Close()
			log.Debug("session id:[%d] closed", sid)
		}
	}

	//设为0表示conn已断开
	a.SetUserData(uint64(0))
}

func rpcHandleLoginAuthPass(args []interface{}) {
	a := args[0].(gate.Agent)
	if a.UserData() != nil && a.UserData().(uint64) == 0 {
		return
	}
	account := args[1].(*model.Account)
	user := args[2].(*model.User)

	s.NewSession(a, account, user)
	log.Debug("login success, user id:[%d]", user.ID)
}
