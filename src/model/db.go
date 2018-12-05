package model

import (
	"fmt"

	"github.com/trist725/myleaf/db/mongodb"
	"gopkg.in/mgo.v2/bson"
)

const (
	UserIdTimes   = 10000 //真人用户Id倍数
	UserIdOffset  = 0     //真人用户Id偏移,也可用服务器id
	RobotIdTimes  = 1000  //机器人用户Id倍数
	RobotIdOffset = 1     //机器人用户Id偏移
)

const (
	TblCounters = "counters" // 用来生成递增序列的表
	TblAccount  = "account"  // 帐号表
	TblUser     = "user"     // 角色表

	visitorNameSeq = "visitor" // 用来生成游客名的递增序列
)

var (
	dbName      string
	dialContext *mongodb.DialContext
)

// 定义库表的递增序列
var seqs = []string{
	TblAccount,
	TblUser,
	visitorNameSeq,
}

// 定义库表的唯一索引
var uniqueIndexes = map[string][][]string{
	TblAccount: [][]string{
		[]string{"Name"},
	},
	TblUser: [][]string{
		[]string{"Name"},
		[]string{"AccountID", "ServerID"},
	},
}

// 定义库表的索引
var indexes = map[string][][]string{
	TblAccount: [][]string{
		[]string{"VName"},
		[]string{"Token"},
	},
}

func Init(url string, sessionNum int, name string) (err error) {
	dialContext, err = mongodb.Dial(url, sessionNum)
	if err != nil {
		err = fmt.Errorf("connect to %s fail, %s", url, err)
		return
	}

	dbName = name

	for _, seq := range seqs {
		err = dialContext.EnsureCounter(dbName, TblCounters, seq)
		if err != nil {
			err = fmt.Errorf("ensure counters [%s] error, %s", seq, err)
			return
		}
	}

	for tbl, indexes := range uniqueIndexes {
		for _, index := range indexes {
			err = dialContext.EnsureUniqueIndex(dbName, tbl, index)
			if err != nil {
				err = fmt.Errorf("ensure table[%s] unique index[%+v] error, %s", tbl, index, err)
				return
			}
		}
	}

	for tbl, is := range indexes {
		for _, index := range is {
			err = dialContext.EnsureIndex(dbName, tbl, index)
			if err != nil {
				err = fmt.Errorf("ensure table[%s] index[%+v] error, %s", tbl, index, err)
				return
			}
		}
	}

	return
}

func Release() {
	if dialContext != nil {
		dialContext.Close()
		dialContext = nil
	}
}

func DialContext() *mongodb.DialContext {
	return dialContext
}

func GetSession() *mongodb.Session {
	return dialContext.Ref()
}

func PutSession(session *mongodb.Session) {
	dialContext.UnRef(session)
}

func NextSeq(id string) (int, error) {
	return dialContext.NextSeq(dbName, TblCounters, id)
}

func NewObjectId() string {
	return bson.NewObjectId().Hex()
}
