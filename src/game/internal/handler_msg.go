package internal

import (
	"mlgs/src/msg"
	s "mlgs/src/session"
	"reflect"

	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
)

func init() {
	regiserMsgHandle(&msg.C2S_Ping{}, handlePing)

}

func regiserMsgHandle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlePing(args []interface{}) {
	//recv := args[0].(*msg.C2S_Ping)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handlePing return for nil session")
		return
	}
	session.Update()
}
