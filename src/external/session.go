package external

import (
	"github.com/trist725/mgsu/event"
	"mlgs/src/model"
)

type ISession interface {
	//SaveData()
	ID() uint64
	UserData() *model.User
	AccountData() *model.Account

	RegisterEventHandler(id event.ID, handler event.Handler)
	ProcessEvent(ev event.IEvent) error

	GetLogic(id uint8) ILogic
	//Sign() string
	//SetSign(string)
}
