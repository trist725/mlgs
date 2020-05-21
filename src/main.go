package main

import (
	"github.com/trist725/myleaf"
	lconf "github.com/trist725/myleaf/conf"
	"mlgs/src/conf"
	"mlgs/src/game"
	"mlgs/src/gate"
	"mlgs/src/login"
	"mlgs/src/model"
	"mlgs/src/session"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	if err := model.SC.Init(conf.Server.MgoUrl, 1, "db_test"); err != nil {
		panic(err)
	}
	defer model.SC.Release()

	defer session.Mgr().Dispose()

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
