// Code generated by protoc-gen-mgo-go. DO NOT EDIT IT!!!
// source: achievement.proto

/*
It has these top-level messages:
	Achievement
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
// collection [Achievement] begin

type Achievement struct {
	///在成就表中的id
	Id int64 `bson:"Id"`
	///对应成就任务在任务表中的id
	TaskId int64 `bson:"TaskId"`
	///是否已完成
	Completed bool `bson:"Completed"`
	///成就表中的类型
	Type int32 `bson:"Type"`
}

func New_Achievement() *Achievement {
	m := &Achievement{}
	return m
}

func (m Achievement) String() string {
	ba, _ := json.Marshal(m)
	return fmt.Sprintf("{\"Achievement\":%s}", string(ba))
}

func (m *Achievement) Reset() {
	m.Id = 0
	m.TaskId = 0
	m.Completed = false
	m.Type = 0

}

func (m Achievement) Clone() *Achievement {
	n, ok := g_Achievement_Pool.Get().(*Achievement)
	if !ok || n == nil {
		n = &Achievement{}
	}

	n.Id = m.Id
	n.TaskId = m.TaskId
	n.Completed = m.Completed
	n.Type = m.Type

	return n
}

func Clone_Achievement_Slice(dst []*Achievement, src []*Achievement) []*Achievement {
	for _, i := range dst {
		Put_Achievement(i)
	}
	dst = []*Achievement{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

var g_Achievement_Pool = sync.Pool{}

func Get_Achievement() *Achievement {
	m, ok := g_Achievement_Pool.Get().(*Achievement)
	if !ok {
		m = New_Achievement()
	} else {
		if m == nil {
			m = New_Achievement()
		} else {
			m.Reset()
		}
	}
	return m
}

func Put_Achievement(i interface{}) {
	if m, ok := i.(*Achievement); ok && m != nil {
		g_Achievement_Pool.Put(i)
	}
}

// collection [Achievement] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
