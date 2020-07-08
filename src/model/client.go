package model

import (
	"fmt"

	mongodb "github.com/trist725/mgsu/db/mongodb"
)

type IClient interface {
	Init(url string, sessionNum int, dbName string) (err error)
	Release()
	DBName() string
	GetSession() *mongodb.Session
	PutSession(session *mongodb.Session)
	NextSeq(id string) (int, error)
}

var (
	SC *SimpleClient = NewSimpleClient()
)

type SimpleClient struct {
	url        string
	sessionNum int
	dbName     string

	dialContext *mongodb.DialContext
}

func NewSimpleClient() (sc *SimpleClient) {
	sc = &SimpleClient{}
	return
}

func (sc *SimpleClient) Init(url string, sessionNum int, dbName string) (err error) {
	sc.url = url
	sc.sessionNum = sessionNum
	sc.dbName = dbName

	sc.dialContext, err = mongodb.Dial(sc.url, sc.sessionNum)
	if err != nil {
		err = fmt.Errorf("connect to %s fail, %s", sc.url, err)
		return
	}

	for _, seq := range seqs {
		err = sc.dialContext.EnsureCounter(sc.dbName, TblCounters, seq)
		if err != nil {
			err = fmt.Errorf("ensure counters [%s] error, %s", seq, err)
			return
		}
	}

	for tbl, indexes := range uniqueIndexes {
		for _, index := range indexes {
			err = sc.dialContext.EnsureUniqueIndex(sc.dbName, tbl, index)
			if err != nil {
				err = fmt.Errorf("ensure table[%s] unique index[%+v] error, %s", tbl, index, err)
				return
			}
		}
	}

	for tbl, is := range indexes {
		for _, index := range is {
			err = sc.dialContext.EnsureIndex(sc.dbName, tbl, index)
			if err != nil {
				err = fmt.Errorf("ensure table[%s] index[%+v] error, %s", tbl, index, err)
				return
			}
		}
	}

	return
}

func (sc *SimpleClient) Release() {
	if sc.dialContext != nil {
		sc.dialContext.Close()
		sc.dialContext = nil
	}
}

func (sc *SimpleClient) DialContext() *mongodb.DialContext {
	return sc.dialContext
}

func (sc SimpleClient) DBName() string {
	return sc.dbName
}

func (sc *SimpleClient) GetSession() *mongodb.Session {
	return sc.dialContext.Ref()
}

func (sc *SimpleClient) PutSession(session *mongodb.Session) {
	sc.dialContext.UnRef(session)
}

func (sc *SimpleClient) NextSeq(id string) (int, error) {
	return sc.dialContext.NextSeq(sc.dbName, TblCounters, id)
}

type ReplicaClient struct {
	urls       string
	sessionNum int
	dbName     string

	dialContext *mongodb.DialContext
}

func NewReplicaClient() (sc *ReplicaClient) {
	sc = &ReplicaClient{}
	return
}

type ShardClient struct {
	urls       string
	sessionNum int
	dbName     string

	dialContext *mongodb.DialContext
}

func NewShardClient() (sc *ReplicaClient) {
	sc = &ReplicaClient{}
	return
}
