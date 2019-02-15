package model

import (
	"mlgs/src/msg"
)

func (m Mail) ToMsg(nm *msg.Mail) *msg.Mail {
	nm.Id = m.Id
	nm.RewardType = m.RewardType
	nm.RewardNum = m.RewardNum
	nm.Content = m.Content

	return nm
}
