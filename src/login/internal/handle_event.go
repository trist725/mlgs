package internal

import (
	"github.com/trist725/mgsu/event"
	ev "mlgs/src/event"
)

func (logic *Logic) registerAllEventHandler() {
	logic.RegisterEventHandler(ev.OnLoginID, handleOnLogin)
}

func handleOnLogin(iEv event.IEvent, args ...interface{}) {
	ev := iEv.(*ev.OnLogin)
	_ = ev
}
