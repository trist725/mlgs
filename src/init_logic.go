package main

import (
	"mlgs/src/game"
)

//所有logic首先全部注册到gFactoryMap
func initLogic() {
	game.Init()
}
