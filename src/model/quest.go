package model

import (
	"mlgs/src/msg"
)

func (m Quest) ToMsg(nm *msg.Quest) *msg.Quest {
	nm.Received = m.Received
	nm.Progress = m.Progress
	nm.Id = m.Id
	return nm
}
