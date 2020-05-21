package model

import (
	"testing"

	"github.com/globalsign/mgo/bson"
)

const (
	url = "mongodb://127.0.0.1:27017/admin"
)

func TestUser(t *testing.T) {
	if err := SC.Init(url, 1, "db_test"); err != nil {
		t.Error(err)
		return
	}
	defer SC.Release()

	session := SC.GetSession()
	defer SC.PutSession(session)

	newUser, err := SC.CreateUser(1, 1, "test", 1)
	if err != nil {
		t.Error(err)
		return
	}

	if err := newUser.Insert(session, SC.DBName()); err != nil {
		t.Error(err)
		return
	}

	u, err := SC.FindOne_User(session, bson.M{"_id": newUser.ID})
	if err != nil {
		t.Error(err)
	}

	t.Logf("%v\n", u)

	_, err = SC.FindOne_User(session, bson.M{"Name": "test"})
	if err != nil {
		t.Error(err)
	}

	some, err := SC.FindSome_User(session, bson.M{"Name": "test"})
	if err != nil || len(some) != 1 {
		t.Error(err)
	}

	err = newUser.RemoveByID(session, SC.DBName())
	if err != nil {
		t.Error(err)
	}
}

func TestClone_User_Slice(t *testing.T) {
	var dst = []*User{
		{ID: 1},
		{ID: 2},
		{ID: 3},
		{ID: 4},
	}

	t.Logf("before, dst=\n")
	for _, i := range dst {
		t.Log(i)
	}

	var src = []*User{
		{ID: 2},
		{ID: 3},
		{ID: 4},
		{ID: 5},
		{ID: 6},
	}

	t.Logf("src=\n")
	for _, i := range src {
		t.Log(i)
	}

	t.Logf("dst=%#v\nsrc=%#v\n", dst, src)

	dst = Clone_User_Slice(dst, src)

	t.Logf("after clone, dst=\n")
	for _, i := range dst {
		t.Log(i)
	}

	t.Logf("src=\n")
	for _, i := range src {
		t.Log(i)
	}

	t.Logf("dst=%#v\nsrc=%#v\n", dst, src)
}
