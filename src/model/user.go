package model

import (
	"fmt"
	"github.com/trist725/myleaf/log"
	"github.com/trist725/myleaf/module"
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

func (m User) ToMsg(nm *msg.User) *msg.User {
	nm.ID = m.ID
	nm.NickName = m.NickName
	for _, money := range m.Monies {
		nm.Monies = append(nm.Monies, money.ToMsg(msg.Get_Money()))
	}
	nm.Sex = m.Sex
	nm.Exp = m.Exp
	nm.Level = m.Level
	nm.BestCombo.Type = m.BestCombo.Type
	for _, card := range m.BestCombo.Cards {
		nm.BestCombo.Cards = append(nm.BestCombo.Cards, card.ToMsg(msg.Get_Card()))
	}
	for _, item := range m.Items {
		nm.Items = append(nm.Items, item.ToMsg(msg.Get_Item()))
	}

	return nm
}

func (user *User) GetMoney(st int32) *Money {
	if st <= 0 {
		log.Error("invalid money type")
		return nil
	}
	for _, m := range user.Monies {
		if m.Type == st {
			return m
		}
	}
	return nil
}

func (user *User) GetNum(tid int64) (num int64) {
	itemSD := sd.ItemMgr.Get(tid)
	if itemSD == nil {
		log.Error("[%d-%s] get item num fail, item[%d] static data not exist", user.ID, user.NickName, tid)
		return 0
	}

	switch sd.E_Item(itemSD.Type) {
	case sd.E_Item_Money:
		st := sd.E_Money(itemSD.SubType)
		if !sd.Check_E_Money_I(itemSD.SubType) || st == sd.E_Money_ {
			log.Error("[%d-%s] get money num fail, invalid type [%v]", user.ID, user.NickName, tid)
			return
		}
		money := user.GetMoney(itemSD.SubType)
		if money == nil {
			log.Error("get money failed")
			return
		}
		num = money.Num

	default:
		for _, i := range user.Items {
			if i.TID == tid {
				num += i.Num
			}
		}
	}

	return
}

func (user *User) Lost(tid int64, num int64, notify bool, skeleton *module.Skeleton) (lostItems []*Item, updateItems []*Item, err error) {
	itemSD := sd.ItemMgr.Get(tid)
	if itemSD == nil {
		err = fmt.Errorf("item static data [%d] not exist", tid)
		return
	}

	if num <= 0 {
		err = fmt.Errorf("item num <= 0")
		return
	}

	switch sd.E_Item(itemSD.Type) {
	case sd.E_Item_Money:
		money := user.GetMoney(itemSD.SubType)
		if money == nil {
			err = fmt.Errorf("get money failed")
			return
		}
		if money.Num < num {
			log.Error("[%d-%s] lost money, lack money, subtype=[%v], own num=[%d], lost num=[%d]", user.ID, user.NickName, itemSD.SubType, money.Num, num)
			err = fmt.Errorf("lack money, subtype=[%v], own num=[%d], lost num=[%d]", itemSD.SubType, money.Num, num)
			return
		}
		// 扣除货币
		money.Num -= num
		if money.Num < 0 {
			money.Num = 0
		}

		if notify {
			// 通知客户端更新
			//skeleton.ChanRPCServer.Go("UpdateUserData", )
		}

		log.Debug("[%d-%s] lost money, subtype=[%v], num=[%d]", user.ID, user.NickName, itemSD.SubType, num)
		return
	}

	return
}

func (user *User) Gain(tid int64, num int64, notify bool, skeleton *module.Skeleton) (lostItems []*Item, updateItems []*Item, err error) {
	itemSD := sd.ItemMgr.Get(tid)
	if itemSD == nil {
		err = fmt.Errorf("item static data [%d] not exist", tid)
		return
	}

	if num <= 0 {
		log.Error("[%d-%s] gain money invalid num [%v]", user.ID, user.NickName, num)
		err = fmt.Errorf("invalid num [%v]", num)
		return
	}

	switch sd.E_Item(itemSD.Type) {
	case sd.E_Item_Money:
		st := sd.E_Money(itemSD.SubType)
		if !sd.Check_E_Money(st) || st == sd.E_Money_ {
			log.Error("[%d-%s] gain money invalid type [%v]", user.ID, user.NickName, tid)
			err = fmt.Errorf("invalid type [%v]", tid)
			return
		}
		money := user.GetMoney(itemSD.SubType)
		if money == nil {
			err = fmt.Errorf("get money failed")
			return
		}
		money.Num += num

		if notify {
			// 通知客户端更新
			//skeleton.ChanRPCServer.Go("UpdateUserData", )
		}

		log.Debug("[%d-%s] gain money, subtype=[%v], num=[%d]", user.ID, user.NickName, st, num)
		return
	}

	return
}
