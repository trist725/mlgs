package model

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

const (
	url = "mongodb://admin:123456@192.168.101.100:27017/admin"
)

func TestAccount(t *testing.T) {
	if err := Init(url, 1, "hh-test"); err != nil {
		t.Error(err)
		return
	}
	defer Release()

	session := GetSession()
	defer PutSession(session)

	newAccount, err := CreateAccount("test", "123456")
	if err != nil {
		t.Error(err)
		return
	}

	if err := newAccount.Insert(session); err != nil {
		t.Error(err)
		return
	}

	_, err = FindOne_Account(session, bson.M{"_id": newAccount.ID})
	if err != nil {
		t.Error(err)
	}

	_, err = FindOne_Account(session, bson.M{"Name": "test"})
	if err != nil {
		t.Error(err)
	}

	some, err := FindSome_Account(session, bson.M{"Name": "test"})
	if err != nil || len(some) != 1 {
		t.Error(err)
	}

	err = newAccount.RemoveByID(session)
	if err != nil {
		t.Error(err)
	}
}

func TestUser(t *testing.T) {
	if err := Init(url, 1, "hh-test"); err != nil {
		t.Error(err)
		return
	}
	defer Release()

	session := GetSession()
	defer PutSession(session)

	newUser, err := CreateUser(1, 1, "test", 1)
	if err != nil {
		t.Error(err)
		return
	}

	if err := newUser.Insert(session); err != nil {
		t.Error(err)
		return
	}

	_, err = FindOne_User(session, bson.M{"_id": newUser.ID})
	if err != nil {
		t.Error(err)
	}

	_, err = FindOne_User(session, bson.M{"Name": "test"})
	if err != nil {
		t.Error(err)
	}

	some, err := FindSome_User(session, bson.M{"Name": "test"})
	if err != nil || len(some) != 1 {
		t.Error(err)
	}

	err = newUser.RemoveByID(session)
	if err != nil {
		t.Error(err)
	}
}
