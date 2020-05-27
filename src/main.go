package main

import (
	"fmt"

	"mlgs/src/conf"
	"mlgs/src/game"
	"mlgs/src/gate"
	"mlgs/src/login"
	"mlgs/src/model"
	"mlgs/src/session"

	"github.com/pkg/profile"
	"github.com/trist725/mgsu/util"
	"github.com/trist725/myleaf"
	lconf "github.com/trist725/myleaf/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	if err := model.SC.Init(conf.Server.MgoUrl, conf.Server.MgoSessionNum, conf.Server.DBName); err != nil {
		panic(err)
	}
	defer model.SC.Release()

	defer session.Mgr().Dispose()

	//todo: Multi-modal profiling
	cpuProf := profile.Start(profile.CPUProfile, profile.ProfilePath(conf.Server.ProfilePath), profile.NoShutdownHook)
	defer cpuProf.Stop()

	go func() {
		util.WaitExitSignal()
		fmt.Println("mlgs receive exit signal")
		//todo: do something
	}()

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
