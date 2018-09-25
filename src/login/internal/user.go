package internal

import (
	"fmt"
	"github.com/trist725/myleaf/db/mongodb"
	"mlgs/src/model"
	"mlgs/src/msg"
)

func createUser(dbSession *mongodb.Session, accountId int64, recv *msg.C2S_Login) (*model.User, error) {
	//user关联了accountID
	newUser, err := model.CreateUser(accountId, recv)
	if err != nil {
		return nil, fmt.Errorf("create user fail, %s", err)

	}
	if err := newUser.Insert(dbSession); err != nil {
		return nil, fmt.Errorf("insert new user[%#v] fail, %s", newUser, err)
	}
	return newUser, nil
}
