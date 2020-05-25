package internal

import (
	"fmt"
	"github.com/trist725/mgsu/db/mongodb"
	"mlgs/src/conf"
	"mlgs/src/model"
	"mlgs/src/msg"
)

func createUser(dbSession *mongodb.Session, accountId int64, recv *msg.C2S_Login) (*model.User, error) {
	//user关联了accountID
	newUser, err := model.SC.CreateUserByMsg(accountId, recv)
	if err != nil {
		return nil, fmt.Errorf("create user fail, %s", err)

	}
	if err := newUser.Insert(dbSession, conf.Server.DBName); err != nil {
		return nil, fmt.Errorf("insert new user[%#v] fail, %s", newUser, err)
	}
	return newUser, nil
}
