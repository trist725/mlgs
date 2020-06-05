package main

import (
	a "mlgs/src/robot/agent"
	"mlgs/src/robot/conf"
	"mlgs/src/robot/robot"
	"time"

	leaf "github.com/trist725/myleaf"
	"github.com/trist725/myleaf/network"
)

var gTcpClient network.TCPClient

func init() {
	gTcpClient = network.TCPClient{
		Addr:            conf.Client.TCPAddr,
		ConnNum:         10000,
		ConnectInterval: 1 * time.Second,
		PendingWriteNum: 1000,
		NewAgent:        a.NewAgent,
		LenMsgLen:       2,
		MaxMsgLen:       4096,
		LittleEndian:    true,
	}
}

func main() {
	gTcpClient.Start()
	defer gTcpClient.Close()

	leaf.Run(robot.Module)
}
