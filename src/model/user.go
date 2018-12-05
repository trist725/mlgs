package model

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/msg"
	"mlgs/src/sd"
	"time"
)

func CreateUser(accountID int64, recv *msg.C2S_Login) (m *User, err error) {
	nextSeq, err := NextSeq(TblUser)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	m = Get_User()
	m.ID = int64(nextSeq) * UserIdTimes //+ int64(serverID)
	m.AccountID = accountID
	m.NickName = recv.NickName
	m.Sex = recv.Sex
	m.CreateTime = now.Unix()
	m.LastLoginTime = now.Unix()
	m.AvatarURL = recv.AvatarURL
	m.Level = 1

	//todo:根据玩家类别初始化
	personSd := sd.PersonMgr.Get(sd.InitUserDataId())
	if personSd == nil {
		log.Fatal("策划坑爹了,读person表有误，id: [%d]", sd.InitUserDataId())
		return
	}

	money := Get_Money()
	money.Type = 1
	money.Num = personSd.Coin
	m.Monies = append(m.Monies, money)
	money = Get_Money()
	money.Type = 2
	money.Num = personSd.Dmd
	m.Monies = append(m.Monies, money)
	money = Get_Money()
	money.Type = 3
	money.Num = personSd.Point
	m.Monies = append(m.Monies, money)

	return
}

//func (m User) ToMsg(nm *msg.User) {
//	nm.ID = m.ID
//	nm.Name = m.Name
//	nm.Sex = int32(m.Sex)
//}
