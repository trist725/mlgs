package internal

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"mlgs/src/model"
	"mlgs/src/msg"
)

func checkAccountExist(uid string) (*model.Account, error) {
	dbSession := model.GetSession()
	defer model.PutSession(dbSession)

	account, err := model.FindOne_Account(dbSession, bson.M{"UID": uid})
	if err == nil {
		return account, fmt.Errorf("account[uid: %s] already exist", uid)
	}
	model.Put_Account(account)
	return nil, nil
}

func checkLoginInfo(account *model.Account, recv *msg.C2S_Login) error {
	if account.UID != recv.UID {
		return fmt.Errorf("login uid:[%s] not match", recv.UID)
	}
	if recv.Logintype == msg.C2S_Login_E_LoginType_WeChat {
		if account.Password != recv.Password {
			return fmt.Errorf("login uid:[%s] with password:[%s] not match", recv.UID, recv.Password)
		}
	}

	return nil
}

func createAccount(recv *msg.C2S_Login) (*model.Account, error) {
	dbSession := model.GetSession()
	defer model.PutSession(dbSession)

	newAccount, err := model.CreateAccount(recv)
	if err != nil {
		return nil, fmt.Errorf("CreateAccount fail, %s", err)
	}

	if err := newAccount.Insert(dbSession); err != nil {
		return nil, fmt.Errorf("dbsession[%v] insert account[%v] fail, %s", dbSession, newAccount, err)
	}

	return newAccount, nil
}
