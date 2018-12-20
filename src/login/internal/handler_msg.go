package internal

import (
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"gopkg.in/mgo.v2/bson"
	"mlgs/src/game"
	"mlgs/src/model"
	"mlgs/src/msg"
	s "mlgs/src/session"
	"reflect"
)

func init() {
	regiserMsgHandle(&msg.C2S_Login{}, handleLoginAuth)
}

func regiserMsgHandle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLoginAuth(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.C2S_Login)
	// 消息的发送者
	sender := args[1].(gate.Agent)
	send := &msg.S2C_Login{}
	//close必须在send之后
	closeFlag := false
	if closeFlag {
		defer sender.Close()
	}
	defer sender.WriteMsg(send)
	dbSession := model.GetSession()
	defer model.PutSession(dbSession)

	if account, err := checkAccountExist(dbSession, recv.UID); err != nil {
		//登陆
		if err = checkLoginInfo(account, recv); err != nil {
			send.Reason = msg.S2C_Login_E_Err_LoginInfoNotMatch
			log.Debug("login err: [%s]", err)
			closeFlag = true
			return
		}

		if recv.Location != account.Location {
			send.Reason = msg.S2C_Login_E_Err_LocationWarn
		}

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
				closeFlag = true
				return
			}
			log.Debug("not found user by accountID=%d", account.ID)
			send.Reason = msg.S2C_Login_E_Err_UserNotExist
			return
		}

		//是否已在线
		session := s.Mgr().GetByUserId(user.ID)
		if session != nil {
			log.Debug("[%d-%s] already online", user.ID, user.NickName)
			send.Reason = msg.S2C_Login_E_Err_AlreadyLogin
			closeFlag = true
			return
		}

		account.Location = recv.Location
		user.Sex = recv.Sex
		user.AvatarURL = recv.AvatarURL
		user.NickName = recv.NickName
		game.ChanRPC.Go("LoginAuthPass", sender, *account, *user)
		send.Reason = msg.S2C_Login_E_Err_LoginSuccess
		return
	}

	//创建新账号
	newAccount, err := createAccount(dbSession, recv)
	if err != nil {
		send.Reason = msg.S2C_Login_E_Err_Unknown
		closeFlag = true
		return
	}
	defer model.Put_Account(newAccount)

	//创建新用户
	newUser, err := createUser(dbSession, newAccount.ID, recv)
	if err != nil {
		send.Reason = msg.S2C_Login_E_Err_Unknown
		closeFlag = true
		return
	}
	defer model.Put_User(newUser)

	game.ChanRPC.Go("LoginAuthPass", sender, *newAccount, *newUser)
	send.Reason = msg.S2C_Login_E_Err_NewAccount
}
