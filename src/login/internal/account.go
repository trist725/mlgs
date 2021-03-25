package internal

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/trist725/mgsu/db/mongodb"

	"mlgs/src/conf"
	"mlgs/src/model"
	"mlgs/src/msg"
)

func checkAccountExist(dbSession *mongodb.Session, uid string) (*model.Account, error) {
	account, err := model.SC.FindOne_Account(dbSession, bson.M{"UID": uid})
	if err == nil {
		return account, fmt.Errorf("account[uid: %s] already exist", uid)
	}
	model.Put_Account(account)
	return nil, nil
}

func createAccount(dbSession *mongodb.Session, recv *msg.C2S_Login) (*model.Account, error) {
	newAccount, err := model.CreateAccount(recv)
	if err != nil {
		return nil, fmt.Errorf("CreateAccount fail, %s", err)
	}

	if err := newAccount.Insert(dbSession, conf.Server.DBName); err != nil {
		return nil, fmt.Errorf("dbsession[%v] insert account[%v] fail, %s", dbSession, newAccount, err)
	}

	return newAccount, nil
}
