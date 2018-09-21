package game

import (
	"mlgs/src/game/internal"
	"github.com/trist725/myleaf/gate"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

func GetAgents() map[gate.Agent]struct{} {
	return internal.Agents
}
