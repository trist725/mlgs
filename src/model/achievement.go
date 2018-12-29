package model

import (
	"mlgs/src/msg"
)

func (m Achievement) ToMsg(nm *msg.Achievement) *msg.Achievement {
	nm.Id = m.Id
	nm.Type = m.Type
	return nm
}
