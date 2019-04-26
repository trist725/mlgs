package internal

import (
	"fmt"
	"github.com/trist725/myleaf/db/mongodb"
	"gopkg.in/mgo.v2/bson"
	"mlgs/src/model"
	"mlgs/src/msg"
)

type LoginReq struct {
	MerchantCode string
	Password     string
	PlayerId     string
	ProductCode  string
	Token        string
}

type LoginResp struct {
	Message string `json:"Message, omitempty"`
	UserId  string `json:"UserId, omitempty"`
	Product string `json:"Product, omitempty"`
	Code    int    `json:"Code, omitempty"`
	Token   string
}

func checkAccountExist(dbSession *mongodb.Session, uid string) (*model.Account, error) {
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

func createAccount(dbSession *mongodb.Session, recv *msg.C2S_Login) (*model.Account, error) {
	newAccount, err := model.CreateAccount(recv)
	if err != nil {
		return nil, fmt.Errorf("CreateAccount fail, %s", err)
	}

	if err := newAccount.Insert(dbSession); err != nil {
		return nil, fmt.Errorf("dbsession[%v] insert account[%v] fail, %s", dbSession, newAccount, err)
	}

	return newAccount, nil
}
