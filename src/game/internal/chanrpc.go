package internal

import (
	"github.com/trist725/myleaf/gate"
)

var Agents = make(map[gate.Agent]struct{})

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	Agents[a] = struct{}{}
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	delete(Agents, a)
}
