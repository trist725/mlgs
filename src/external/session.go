package external

import (
	"mlgs/src/model"
)

type ISession interface {
	//SaveData()
	ID() uint64
	UserData() *model.User
	AccountData() *model.Account

	//Sign() string
	//SetSign(string)
}
