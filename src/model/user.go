package model

import (
	"mlgs/src/msg"
	"time"
)

func CreateUser(accountID int64, recv *msg.C2S_Login) (m *User, err error) {
	nextSeq, err := NextSeq(TblUser)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	m = Get_User()
	m.ID = int64(nextSeq) * 10000 //+ int64(serverID)
	m.AccountID = accountID
	m.NickName = recv.NickName
	m.Sex = recv.Sex
	m.CreateTime = now.Unix()
	m.LastLoginTime = now.Unix()
	m.AvatarURL = recv.AvatarURL
	m.Level = 1
	return
}

//func (m User) ToMsg(nm *msg.User) {
//	nm.ID = m.ID
//	nm.Name = m.Name
//	nm.Sex = int32(m.Sex)
//}
