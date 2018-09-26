package internal

import (
	"github.com/trist725/myleaf/gate"
	//s "mlgs/src/session"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
	//mgr := s.GetSessionMgr()
	//mgr.GetSession(a.UserData().(uint64)).Close()

}
