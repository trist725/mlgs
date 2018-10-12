package model

import (
	"mlgs/src/msg"
)

func (m Money) ToMsg(nm *msg.Money) *msg.Money {
	nm.Type = m.Type
	nm.Num = m.Num
	return nm
}
