package base

import (
	"github.com/trist725/myleaf/chanrpc"
	"github.com/trist725/myleaf/module"
	"mlgs/src/conf"
)

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              conf.GoLen,
		TimerDispatcherLen: conf.TimerDispatcherLen,
		AsynCallLen:        conf.AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(conf.ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}
