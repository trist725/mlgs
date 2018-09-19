package internal

import (
	"fmt"
	"github.com/trist725/myleaf/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"mlgs/src/model"
	"mlgs/src/msg"
)

func checkAccountExist(uid string, dbSession *mongodb.Session) (*model.Account, error) {
	account, err := model.FindOne_Account(dbSession, bson.M{"UID": uid})
	if err == nil {
		return account, fmt.Errorf("account[uid: %s] already exist", uid)
	}
	model.Put_Account(account)
	return nil, nil
}

func checkLoginInfo(account *model.Account, recv *msg.C2S_Login) error {
	if account.UID != recv.Uid {
		return fmt.Errorf("login uid:[%s] not match", recv.Uid)
	}
	if recv.Logintype == msg.C2S_Login_E_LoginType_WeChat {
		if account.Password != recv.Password {
			return fmt.Errorf("login uid:[%s] with password:[%s] not match", recv.Uid, recv.Password)
		}
	}

	return nil
}
