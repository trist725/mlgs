package internal

import (
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"mlgs/src/model"
	"mlgs/src/msg"
	"reflect"
)

func init() {
	regiserMsgHandle(&msg.C2S_Login{}, handleLogin)
}

func regiserMsgHandle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.C2S_Login)
	// 消息的发送者
	sender := args[1].(gate.Agent)
	send := &msg.S2C_Login{}
	defer sender.WriteMsg(send)

	dbSession := model.GetSession()
	defer model.PutSession(dbSession)
	if account, err := checkAccountExist(recv.Uid, dbSession); err != nil {
		log.Debug("session[%v] : %s", sender, err)

		if err = checkLoginInfo(account, recv); err != nil {
			send.Reason = msg.S2C_Login_E_Err_LoginInfoNotMatch
			log.Debug("login err: [%s]", err)
			return
		}
		//登陆成功
		if recv.Location != account.Location {
			send.Reason = msg.S2C_Login_E_Err_LocationWarn
		}
		send.Reason = msg.S2C_Login_E_Err_SucceeLogin
		//todo:下发数据
		return
	}

	//创建新账号
	newAccount, err := model.CreateAccount(recv)
	if err != nil {
		log.Error("session[%v] CreateAccount fail, %s", sender, err)
		return
	}
	defer model.Put_Account(newAccount)
	send.Reason = msg.S2C_Login_E_Err_NewAccount

	if err := newAccount.Insert(dbSession); err != nil {
		log.Error("session[%v] insert account[%v] fail, %s", dbSession, newAccount, err)
		return
	}
}
