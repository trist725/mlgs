// Code generated by protoc-gen-mgo-go. DO NOT EDIT IT!!!
// source: item.proto

/*
It has these top-level messages:
	Item
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
// collection [Item] begin

/// 用户物品信息
type Item struct {
	/// 唯一id
	UID string `bson:"UID"`
	/// 类型id, item.xlsx中的id字段
	TID int32 `bson:"TID"`
	/// 创建时刻
	CreateTime int64 `bson:"CreateTime"`
	/// 数量
	Num int32 `bson:"Num"`
}

func New_Item() *Item {
	m := &Item{}
	return m
}

func (m Item) String() string {
	ba, _ := json.Marshal(m)
	return fmt.Sprintf("{\"Item\":%s}", string(ba))
}

func (m *Item) Reset() {
	m.UID = ""
	m.TID = 0
	m.CreateTime = 0
	m.Num = 0

}

func (m Item) Clone() *Item {
	n, ok := g_Item_Pool.Get().(*Item)
	if !ok || n == nil {
		n = &Item{}
	}

	n.UID = m.UID
	n.TID = m.TID
	n.CreateTime = m.CreateTime
	n.Num = m.Num

	return n
}

func Clone_Item_Slice(dst []*Item, src []*Item) []*Item {
	for _, i := range dst {
		Put_Item(i)
	}
	dst = []*Item{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

var g_Item_Pool = sync.Pool{}

func Get_Item() *Item {
	m, ok := g_Item_Pool.Get().(*Item)
	if !ok {
		m = New_Item()
	} else {
		if m == nil {
			m = New_Item()
		} else {
			m.Reset()
		}
	}
	return m
}

func Put_Item(i interface{}) {
	if m, ok := i.(*Item); ok && m != nil {
		g_Item_Pool.Put(i)
	}
}

// collection [Item] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
