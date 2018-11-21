package room

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
	"mlgs/src/msg"
)

type TurnAction struct {
	act *msg.C2S_TurnAction
	p   *cache.Player
}

func (r *Room) SendPlayerActionSig(act *msg.C2S_TurnAction, player *cache.Player) {
	if r.stage == 0 {
		return
	}

	select {
	case r.actSig <- TurnAction{
		act: act,
		p:   player,
	}:
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
