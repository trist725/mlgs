package model

import (
	"github.com/globalsign/mgo/bson"
)

const (
	UserIdTimes   = 1000  //真人用户Id倍数
	UserIdOffset  = 0     //真人用户Id偏移,也可用服务器id
	RobotIdTimes  = 10000 //机器人用户Id倍数
	RobotIdOffset = 1     //机器人用户Id偏移
)

const (
	TblCounters = "counters" // 用来生成递增序列的表
	TblAccount  = "account"  // 帐号表
	TblUser     = "user"     // 角色表

	visitorNameSeq = "visitor" // 用来生成游客名的递增序列
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

func NewObjectID() string {
	return bson.NewObjectId().Hex()
}
