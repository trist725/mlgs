package model

import (
	"mlgs/src/msg"
	"strconv"
	"time"
)

func (sc *SimpleClient) CreateUser(accountID int64, serverID int32, name string, sex int32) (m *User, err error) {
	nextSeq, err := sc.NextSeq(TblUser)
	if err != nil {
		return nil, err
	}
	m = Get_User()
	m.ID = int64(nextSeq)*10000 + int64(serverID)
	m.AccountID = accountID
	m.Sex = strconv.Itoa(int(sex))
	m.CreateTime = time.Now().Unix()
	return
}

func (sc *SimpleClient) CreateUserByMsg(accountID int64, recv *msg.C2S_Login) (m *User, err error) {
	nextSeq, err := sc.NextSeq(TblUser)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	m = Get_User()
	m.ID = int64(nextSeq)*UserIdTimes + UserIdOffset //+ int64(serverID)
	m.AccountID = accountID
	m.NickName = recv.NickName
	m.CreateTime = now.Unix()
	m.LastLoginTime = now.Unix()
	m.Level = 1

	return
}
