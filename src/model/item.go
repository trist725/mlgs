package model

import (
	"mlgs/src/msg"
	"time"
)

func (m Item) ToMsg(nm *msg.Item) *msg.Item {
	nm.TID = m.TID
	nm.Num = m.Num
	nm.UID = m.UID

	return nm
}

func CreateItem(tid int64) *Item {
	i := Get_Item()
	i.UID = NewObjectId()
	i.TID = tid
	i.CreateTime = time.Now().Unix()
	return i
}
