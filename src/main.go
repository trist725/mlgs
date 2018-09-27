package main

import (
	"conf"
	"game"
	"gate"
	"github.com/trist725/myleaf"
	lconf "github.com/trist725/myleaf/conf"
	"login"
	"mlgs/src/model"
	"session"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	if err := model.Init(conf.Server.MgoUrl, conf.Server.MgoSessionNum, conf.Server.MgoName); err != nil {
		panic(err)
	}
	defer model.Release()

	initLogic()

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)

	session.GetSessionMgr().Dispose()
}
