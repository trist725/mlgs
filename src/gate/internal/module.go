package internal

import (
	"mlgs/src/base"
	"mlgs/src/conf"
	"time"

	"github.com/trist725/myleaf/module"
	"github.com/trist725/myleaf/network"
	"github.com/trist725/myleaf/network/protobuf"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	//todo: 配置中心,gate分配id,这里用map
	Gates []*network.TCPClient
)

type Module struct {
	//*gate.Gate
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	for _, addr := range conf.Server.GateAddrs {
		gate := network.TCPClient{
			Addr:            addr,
			ConnNum:         1,
			ConnectInterval: 3 * time.Second,
			PendingWriteNum: conf.PendingWriteNum,
			NewAgent:        newAgent,
			LenMsgLen:       conf.LenMsgLen,
			MaxMsgLen:       conf.MaxMsgLen,
			LittleEndian:    conf.LittleEndian,
			AutoReconnect:   true,
		}
		gate.Start()
		Gates = append(Gates, &gate)
	}

}

func (m *Module) OnDestroy() {
	for _, gate := range Gates {
		gate.Close()
	}
}

func newAgent(conn *network.TCPConn) network.Agent {
	a := new(base.Agent)
	a.SetConn(conn)
	a.SetProcessor(protobuf.NewServerProcessor())
	a.Processor().SetDefaultRouter(ChanRPC)
	return a
}
