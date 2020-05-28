// Code generated by protoc-gen-mgo-go. DO NOT EDIT IT!!!
// source: user.proto

package model

import (
	json "encoding/json"
	fmt "fmt"
	mgo "github.com/globalsign/mgo"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	mongodb "github.com/trist725/mgsu/db/mongodb"
	math "math"
	msg "mlgs/src/msg"
	sync "sync"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = msg.PH

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// collection [User] begin

func New_User() *User {
	m := &User{}
	return m
}

func (m User) JsonString() string {
	bs, _ := json.Marshal(m)
	return fmt.Sprintf("{\"User\":%s}", string(bs))
}

func (m *User) ResetEx() {

	m.ID = 0

	m.AccountID = 0

	m.Level = 0

	m.AvatarURL = ""

	m.NickName = ""

	m.Sex = ""

	m.CreateTime = 0

	m.LastLoginTime = 0

	m.LastLogoutTime = 0

	m.Exp = 0

	m.DaySigned = false

	m.SignedDays = 0

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

	n.Exp = m.Exp

	n.DaySigned = m.DaySigned

	n.SignedDays = m.SignedDays

	return n
}

func Clone_User_Slice(dst []*User, src []*User) []*User {
	for _, i := range dst {
		Put_User(i)
	}
	if len(src) > 0 {
		dst = make([]*User, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*User{}
		dst = nil
	}
	return dst
}

func (sc SimpleClient) FindOne_User(session *mongodb.Session, query interface{}) (one *User, err error) {
	one = Get_User()
	err = session.DB(sc.dbName).C(TblUser).Find(query).One(one)
	if err != nil {
		Put_User(one)
		return nil, err
	}
	return
}

func (sc SimpleClient) FindSome_User(session *mongodb.Session, query interface{}) (some []*User, err error) {
	some = []*User{}
	err = session.DB(sc.dbName).C(TblUser).Find(query).All(&some)
	if err != nil {
		return nil, err
	}
	return
}

func (sc SimpleClient) UpdateSome_User(session *mongodb.Session, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(sc.dbName).C(TblUser).UpdateAll(selector, update)
	return
}

func (sc SimpleClient) Upsert_User(session *mongodb.Session, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(sc.dbName).C(TblUser).Upsert(selector, update)
	return
}

func (sc SimpleClient) UpsertID_User(session *mongodb.Session, id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(sc.dbName).C(TblUser).UpsertId(id, update)
	return
}

func (m User) Insert(session *mongodb.Session, dbName string) error {
	return session.DB(dbName).C(TblUser).Insert(m)
}

func (m User) Update(session *mongodb.Session, dbName string, selector interface{}, update interface{}) error {
	return session.DB(dbName).C(TblUser).Update(selector, update)
}

func (m User) UpdateByID(session *mongodb.Session, dbName string) error {
	return session.DB(dbName).C(TblUser).UpdateId(m.ID, m)
}

func (m User) RemoveByID(session *mongodb.Session, dbName string) error {
	return session.DB(dbName).C(TblUser).RemoveId(m.ID)
}

func (m User) ToMsg(n *msg.User) *msg.User {
	if n == nil {
		n = msg.Get_User()
	}

	n.ID = m.ID

	n.Level = m.Level

	n.NickName = m.NickName

	n.Sex = m.Sex

	n.Exp = m.Exp

	return n
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
			m.ResetEx()
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
