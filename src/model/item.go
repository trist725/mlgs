package model

import (
	"mlgs/src/msg"
)

func (m Item) ToMsg(nm *msg.Item) *msg.Item {
	nm.TID = m.TID
	nm.Num = m.Num
	nm.UID = m.UID

	return nm
}
