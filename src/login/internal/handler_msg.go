package internal

import (
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"gopkg.in/mgo.v2/bson"
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

	if account, err := checkAccountExist(recv.UID); err != nil {
		log.Debug("login: %s", err)

		if err = checkLoginInfo(account, recv); err != nil {
			send.Reason = msg.S2C_Login_E_Err_LoginInfoNotMatch
			log.Debug("login err: [%s]", err)
			return
		}
		//登陆成功
		if recv.Location != account.Location {
			send.Reason = msg.S2C_Login_E_Err_LocationWarn
		}

		//todo:触发登陆事件,下发数据
		//根据关联的accountID查找user
		user, err := model.FindOne_User(
			dbSession,
			bson.M{
				"AccountID": account.ID,
			},
		)
		if err != nil {
			if err.Error() != "not found" {
				log.Error("find user by accountID=%d fail, err:%#v", account.ID, err)
				//加载用户数据失败, 断开连接
				send.Reason = msg.S2C_Login_E_Err_Unknown
				sender.Close()
				return
			}
			log.Debug("not found user by accountID=%d", account.ID)
			send.Reason = msg.S2C_Login_E_Err_UserNotExist
			return
		}

		//agent和user绑定
		sender.SetUserData(user)
		log.Debug("login success")
		send.Reason = msg.S2C_Login_E_Err_LoginSuccess
		return
	}

	//创建新账号
	newAccount, err := createAccount(recv)
	if err != nil {
		send.Reason = msg.S2C_Login_E_Err_Unknown
		sender.Close()
		return
	}
	defer model.Put_Account(newAccount)

	//创建新用户
	newUser, err := createUser(newAccount.ID, recv)
	if err != nil {
		send.Reason = msg.S2C_Login_E_Err_Unknown
		sender.Close()
		return
	}
	defer model.Put_User(newUser)

	//agent和user绑定
	sender.SetUserData(newUser)

	send.Reason = msg.S2C_Login_E_Err_NewAccount
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//logic.SetData(newUser)
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// 同步触发创角事件
	//logic.ProcessEvent(&game_ev.OnCreate{
	//	AccountID: account.ID,
	//	ServerID:  logic.cache.ServerID,
	//	UserID:    newUser.ID,
	//})
}
