package game

import (
	"mlgs/src/game/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

func Init() {
	internal.Init()
}
