package internal

import (
	"mlgs/src/base"
	"mlgs/src/msg"
	s "mlgs/src/session"
	"reflect"

	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
)

func init() {
	regiserMsgHandle(&msg.C2S_Ping{}, handlePong)

	regiserMsgHandle(&msg.C2S_UpdateUserData{}, handleUpdateUserData)

	regiserMsgHandle(&msg.C2S_GetNotices{}, handleGetNotices)
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
	sender.WriteCmd(0, clientID)
}

func handleUpdateUserData(args []interface{}) {
	//recv := args[0].(*msg.C2S_UpdateUserData)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleUpdateUserData return for nil session")
		return
	}
	send := msg.Get_S2C_UpdateUserData()
	defer sender.WriteMsg(send)

	send.Data = session.UserData().ToMsg(msg.Get_User())
	send.Err = msg.S2C_UpdateUserData_OK
	session.Update()
}

func handleGetNotices(args []interface{}) {
	//recv := args[0].(*msg.C2S_GetNotices)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleGetNotices return for nil session")
		return
	}

	send := msg.Get_S2C_GetNotices()
	defer sender.WriteMsg(send)
	defer session.Update()

	//GetNotices()
	//
	//send.Notices = append(send.Notices, ConvertNotices()...)

}
