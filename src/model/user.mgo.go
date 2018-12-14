// Code generated by protoc-gen-mgo-go. DO NOT EDIT IT!!!
// source: user.proto

/*
It has these top-level messages:
	User
*/

package model

import "fmt"
import "encoding/json"
import "sync"
import "github.com/trist725/myleaf/db/mongodb"
import "gopkg.in/mgo.v2"

var _ = fmt.Sprintf
var _ = json.Marshal
var _ *sync.Pool
var _ *mongodb.DialContext
var _ *mgo.DBRef

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// collection [User] begin

/// 用户数据 @collection
type User struct {
	/// mongodb默认主键_id做用户id @bson=_id
	ID int64 `bson:"_id"`
	/// 帐号id
	AccountID int64 `bson:"AccountID"`
	///等级
	Level int32 `bson:"Level"`
	///头像url
	AvatarURL string `bson:"AvatarURL"`
	/// 服务器ID int32 ServerID = 3; / 名字
	NickName string `bson:"NickName"`
	/// 性别
	Sex string `bson:"Sex"`
	/// 创建时刻
	CreateTime int64 `bson:"CreateTime"`
	/// 上次登录时刻
	LastLoginTime int64 `bson:"LastLoginTime"`
	/// 上次登出时刻
	LastLogoutTime int64 `bson:"LastLogoutTime"`
	///货币信息
	Monies []*Money `bson:"Monies"`
	///物品
	Items []*Item `bson:"Items"`
	///经验
	Exp int64 `bson:"Exp"`
	///今日是否已签到
	DaySigned bool `bson:"DaySigned"`
	///已签到天数
	SignedDays int32 `bson:"SignedDays"`
	///每日签到奖励,数组第几个代表第几天 repeated Item SignRewards = 15; /历史最大牌型
	BestCombo *BestCombo `bson:"BestCombo"`
}

func New_User() *User {
	m := &User{
		Monies:    []*Money{},
		Items:     []*Item{},
		BestCombo: New_BestCombo(),
	}
	return m
}

func (m User) String() string {
	ba, _ := json.Marshal(m)
	return fmt.Sprintf("{\"User\":%s}", string(ba))
}

func (m *User) Reset() {
	m.ID = 0
	m.AccountID = 0
	m.Level = 0
	m.AvatarURL = ""
	m.NickName = ""
	m.Sex = ""
	m.CreateTime = 0
	m.LastLoginTime = 0
	m.LastLogoutTime = 0

	for _, i := range m.Monies {
		Put_Money(i)
	}
	m.Monies = []*Money{}

	for _, i := range m.Items {
		Put_Item(i)
	}
	m.Items = []*Item{}
	m.Exp = 0
	m.DaySigned = false
	m.SignedDays = 0
	m.BestCombo.Reset()

}

func (m User) Clone() *User {
	n, ok := g_User_Pool.Get().(*User)
	if !ok || n == nil {
		n = &User{}
	}

	n.ID = m.ID
	n.AccountID = m.AccountID
	n.Level = m.Level
	n.AvatarURL = m.AvatarURL
	n.NickName = m.NickName
	n.Sex = m.Sex
	n.CreateTime = m.CreateTime
	n.LastLoginTime = m.LastLoginTime
	n.LastLogoutTime = m.LastLogoutTime

	if len(m.Monies) > 0 {
		for _, i := range m.Monies {
			if i != nil {
				n.Monies = append(n.Monies, i.Clone())
			} else {
				n.Monies = append(n.Monies, nil)
			}
		}
	} else {
		n.Monies = []*Money{}
	}

	if len(m.Items) > 0 {
		for _, i := range m.Items {
			if i != nil {
				n.Items = append(n.Items, i.Clone())
			} else {
				n.Items = append(n.Items, nil)
			}
		}
	} else {
		n.Items = []*Item{}
	}

	n.Exp = m.Exp
	n.DaySigned = m.DaySigned
	n.SignedDays = m.SignedDays
	n.BestCombo = m.BestCombo.Clone()

	return n
}

func Clone_User_Slice(dst []*User, src []*User) []*User {
	for _, i := range dst {
		Put_User(i)
	}
	dst = []*User{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func FindOne_User(session *mongodb.Session, query interface{}) (one *User, err error) {
	one = Get_User()
	err = session.DB(dbName).C(TblUser).Find(query).One(one)
	if err != nil {
		Put_User(one)
		return nil, err
	}
	return
}

func FindSome_User(session *mongodb.Session, query interface{}) (some []*User, err error) {
	some = []*User{}
	err = session.DB(dbName).C(TblUser).Find(query).All(&some)
	if err != nil {
		return nil, err
	}
	return
}

func UpdateSome_User(session *mongodb.Session, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(dbName).C(TblUser).UpdateAll(selector, update)
	return
}

func Upsert_User(session *mongodb.Session, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(dbName).C(TblUser).Upsert(selector, update)
	return
}

func UpsertID_User(session *mongodb.Session, id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(dbName).C(TblUser).UpsertId(id, update)
	return
}

func (m User) Insert(session *mongodb.Session) error {
	return session.DB(dbName).C(TblUser).Insert(m)
}

func (m User) Update(session *mongodb.Session, selector interface{}, update interface{}) error {
	return session.DB(dbName).C(TblUser).Update(selector, update)
}

func (m User) UpdateByID(session *mongodb.Session) error {
	return session.DB(dbName).C(TblUser).UpdateId(m.ID, m)
}

func (m User) RemoveByID(session *mongodb.Session) error {
	return session.DB(dbName).C(TblUser).RemoveId(m.ID)
}

var g_User_Pool = sync.Pool{}

func Get_User() *User {
	m, ok := g_User_Pool.Get().(*User)
	if !ok {
		m = New_User()
	} else {
		if m == nil {
			m = New_User()
		} else {
			m.Reset()
		}
	}
	return m
}

func Put_User(i interface{}) {
	if m, ok := i.(*User); ok && m != nil {
		g_User_Pool.Put(i)
	}
}

// collection [User] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
