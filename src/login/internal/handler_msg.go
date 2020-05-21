package internal

import (
	"encoding/base64"
	"encoding/json"
	"github.com/trist725/mgsu/util"
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"github.com/globalsign/mgo/bson"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"mlgs/src/conf"
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
	defer func() {
		sender.WriteMsg(send)
		if closeFlag {
			log.Debug("close agent...[%v]", sender.RemoteAddr())
			sender.Close()
		}
	}()
	dbSession := model.SC.GetSession()
	defer model.SC.PutSession(dbSession)

	if !CheckCltVer(&CltVer{
		recv.BigVer,
		recv.SmallVer,
		recv.FixVer,
	}, recv.CltType) {
		log.Debug("client version not match...")
		closeFlag = true
		send.Reason = msg.S2C_Login_E_Err_LowVersion
		return
	}

	if recv.Logintype == msg.C2S_Login_E_LoginType_WanBo {
		pubKey, err := ioutil.ReadFile("publicKey.keystore.txt")
		if err != nil {
			log.Error("read public key file error")
			send.Reason = msg.S2C_Login_E_Err_Unknown
			return
		}

		mer, err := util.RsaPubEncrypt(pubKey, []byte(conf.Server.MerchantCode))
		pwd, err := util.RsaPubEncrypt(pubKey, []byte(recv.Password))
		uid, err := util.RsaPubEncrypt(pubKey, []byte(recv.UID))
		pc, err := util.RsaPubEncrypt(pubKey, []byte(conf.Server.ProductCode))

		resp, err := resty.R().SetFormData(map[string]string{
			"MerchantCode": base64.StdEncoding.EncodeToString(mer),
			"Password":     base64.StdEncoding.EncodeToString(pwd),
			"PlayerId":     base64.StdEncoding.EncodeToString(uid),
			"ProductCode":  base64.StdEncoding.EncodeToString(pc),
			"Token":        recv.Token,
		}).Post(conf.Server.UnionPlatUrl + "Player/Login")
		//SetBody(LoginReq{MerchantCode: conf.Server.MerchantCode, Password: recv.Password, PlayerId: recv.UID, ProductCode: conf.Server.ProductCode}).Post(conf.Server.UnionPlatUrl + "Player/Login")
		if err != nil {
			log.Debug("wanbo login failed: %s", err)
			send.Reason = msg.S2C_Login_E_Err_Unknown
			return
		}
		var respLogin LoginResp
		err = json.Unmarshal(resp.Body(), &respLogin)
		if err != nil {
			log.Error("failed to unmarshal login resp")
			send.Reason = msg.S2C_Login_E_Err_Unknown
			return
		}
		send.WanboRes = int32(respLogin.Code)
		send.Token = respLogin.Token
		if respLogin.Code == 0 {
			send.Reason = msg.S2C_Login_E_Err_LoginSuccess
		} else if respLogin.Code >= 501 && respLogin.Code <= 509 {
			send.Reason = msg.S2C_Login_E_Err_LoginInfoNotMatch
			return
		} else {
			send.Reason = msg.S2C_Login_E_Err_Unknown
			return
		}

	}

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
		user, err := model.SC.FindOne_User(
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
			send.Reason = msg.S2C_Login_E_Err_LoginSuccess
			session.SetAgent(sender)
			session.Agent().SetUserData(session.ID())
			session.SetCloseFlag(0)
			//if p := session.Player(); p != nil {
			//	if p.SessionId() == 0 {
			//		p.SetSessionId(session.ID())
			//		game.ChanRPC.Go("AfterLoginAuthPass", sender, user, p.InRoom())
			//	} else {
			//		log.Debug("[%d-%s] already online", user.ID, user.NickName)
			//		send.Reason = msg.S2C_Login_E_Err_AlreadyLogin
			//		closeFlag = true
			//	}
			//}
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
