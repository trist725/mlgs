package main

import (
	"github.com/trist725/myleaf"
	"github.com/trist725/myleaf/network"
	a "robot/agent"
	"robot/conf"
	"robot/robot"
	"time"
)

var gTcpClient network.TCPClient

func init() {
	gTcpClient = network.TCPClient{
		Addr:            conf.Client.TCPAddr,
		ConnNum:         1,
		ConnectInterval: 3 * time.Second,
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
