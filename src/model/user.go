package model

import (
	"time"

	"msg"
)

func CreateUser(accountID int64, serverID int32, name string, sex int32) (m *User, err error) {
	nextSeq, err := NextSeq(TblUser)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	m = Get_User()
	m.ID = int64(nextSeq)*10000 + int64(serverID)
	m.AccountID = accountID
	m.ServerID = serverID
	m.Name = name
	m.Sex = E_UserSex(sex)
	m.CreateTime = now.Unix()
	m.LastLoginTime = now.Unix()
	return
}

func (m User) ToMsg(nm *msg.User) {
	nm.ID = m.ID
	nm.Name = m.Name
	nm.Sex = int32(m.Sex)
}
