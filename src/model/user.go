package model

import (
	"time"
)

func CreateUser(accountID int64, serverID int32, name string, sex string) (m *User, err error) {
	nextSeq, err := NextSeq(TblUser)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	m = Get_User()
	m.ID = int64(nextSeq)*10000 + int64(serverID)
	m.AccountID = accountID
	m.NickName = name
	m.Sex = sex
	m.CreateTime = now.Unix()
	m.LastLoginTime = now.Unix()
	return
}

//func (m User) ToMsg(nm *msg.User) {
//	nm.ID = m.ID
//	nm.Name = m.Name
//	nm.Sex = int32(m.Sex)
//}
