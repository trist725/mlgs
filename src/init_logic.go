package main

import (
	"mlgs/src/login"
)

//所有logic首先全部注册到gFactoryMap
func initLogic() {
	login.Init()
}
