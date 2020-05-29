// Code generated by protoc-gen-mgo-go. DO NOT EDIT IT!!!
// source: account.proto

package model

import (
	json "encoding/json"
	fmt "fmt"
	math "math"
	msg "mlgs/src/msg"
	sync "sync"

	mgo "github.com/globalsign/mgo"
	proto "github.com/gogo/protobuf/proto"
	mongodb "github.com/trist725/mgsu/db/mongodb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = msg.PH

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// collection [Account] begin

func New_Account() *Account {
	m := &Account{}
	return m
}

func (m Account) JsonString() string {
	bs, _ := json.Marshal(m)
	return fmt.Sprintf("{\"Account\":%s}", string(bs))
}

func (m *Account) ResetEx() {

	m.ID = 0

	m.UID = ""

	m.RegisterTime = 0

	m.Location = ""

	m.Password = ""

	m.Ban = 0

	m.Type = 0

}

func (m Account) Clone() *Account {
	n, ok := g_Account_Pool.Get().(*Account)
	if !ok || n == nil {
		n = &Account{}
	}

	n.ID = m.ID

	n.UID = m.UID

	n.RegisterTime = m.RegisterTime

	n.Location = m.Location

	n.Password = m.Password

	n.Ban = m.Ban

	n.Type = m.Type

	return n
}

func Clone_Account_Slice(dst []*Account, src []*Account) []*Account {
	for _, i := range dst {
		Put_Account(i)
	}
	if len(src) > 0 {
		dst = make([]*Account, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*Account{}
		dst = nil
	}
	return dst
}

func (sc SimpleClient) FindOne_Account(session *mongodb.Session, query interface{}) (one *Account, err error) {
	one = Get_Account()
	err = session.DB(sc.dbName).C(TblAccount).Find(query).One(one)
	if err != nil {
		Put_Account(one)
		return nil, err
	}
	return
}

func (sc SimpleClient) FindSome_Account(session *mongodb.Session, query interface{}) (some []*Account, err error) {
	some = []*Account{}
	err = session.DB(sc.dbName).C(TblAccount).Find(query).All(&some)
	if err != nil {
		return nil, err
	}
	return
}

func (sc SimpleClient) UpdateSome_Account(session *mongodb.Session, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(sc.dbName).C(TblAccount).UpdateAll(selector, update)
	return
}

func (sc SimpleClient) Upsert_Account(session *mongodb.Session, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(sc.dbName).C(TblAccount).Upsert(selector, update)
	return
}

func (sc SimpleClient) UpsertID_Account(session *mongodb.Session, id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = session.DB(sc.dbName).C(TblAccount).UpsertId(id, update)
	return
}

func (m Account) Insert(session *mongodb.Session, dbName string) error {
	return session.DB(dbName).C(TblAccount).Insert(m)
}

func (m Account) Update(session *mongodb.Session, dbName string, selector interface{}, update interface{}) error {
	return session.DB(dbName).C(TblAccount).Update(selector, update)
}

func (m Account) UpdateByID(session *mongodb.Session, dbName string) error {
	return session.DB(dbName).C(TblAccount).UpdateId(m.ID, m)
}

func (m Account) RemoveByID(session *mongodb.Session, dbName string) error {
	return session.DB(dbName).C(TblAccount).RemoveId(m.ID)
}

var g_Account_Pool = sync.Pool{}

func Get_Account() *Account {
	m, ok := g_Account_Pool.Get().(*Account)
	if !ok {
		m = New_Account()
	} else {
		if m == nil {
			m = New_Account()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_Account(i interface{}) {
	if m, ok := i.(*Account); ok && m != nil {
		g_Account_Pool.Put(i)
	}
}

// collection [Account] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
