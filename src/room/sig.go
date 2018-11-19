package room

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/msg"
)

func (r *Room) SendPlayerActionSig(act *msg.C2S_TurnAction) {
	if r.stage == 0 {
		return
	}

	select {
	case r.actSig <- act:
	default:
		log.Debug("no PlayerActionSig sent")
		return
	}
}

func (r *Room) SendRefreshReadyTimeSig() {
	if r.stage != 0 {
		return
	}
	select {
	case r.refreshReadyTimeSig <- struct{}{}:
	default:
		log.Debug("no RefreshReadyTimeSig sent")
		return
	}

}

func (r *Room) SendStopLoopSig() {
	select {
	case r.stopSig <- struct{}{}:
	default:
		log.Error("stop room:[%d] loop failed", r.id)
	}
}
