package model

import (
	"mlgs/src/msg"
)

func (m Card) ToMsg(nm *msg.Card) *msg.Card {
	nm.Color = m.Color
	nm.Num = m.Num
	return nm
}
