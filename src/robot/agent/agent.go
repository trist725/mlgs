package agent

import (
	"github.com/trist725/myleaf/log"
	"github.com/trist725/myleaf/network"
	"reflect"
)

type Agent struct {
	conn      *network.TCPConn
	Processor network.Processor
	userData  interface{}
}

func NewAgent(conn *network.TCPConn) network.Agent {
	a := new(Agent)
	a.conn = conn
	return a
}

func (a *Agent) Run() {
	a.Init()

	a.SendSome()

	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			log.Debug("read message: %v", err)
			break
		}

		if a.Processor != nil {
			msg, err := a.Processor.Unmarshal(data)
			if err != nil {
				log.Debug("unmarshal message error: %v", err)
				break
			}
			err = a.Processor.Route(msg, a)
			if err != nil {
				log.Debug("route message error: %v", err)
				break
			}
		}
	}
}

func (a *Agent) OnClose() {
}

func (a *Agent) WriteMsg(msg interface{}) {
	if a.Processor != nil {
		data, err := a.Processor.Marshal(msg)
		if err != nil {
			log.Error("marshal message %v error: %v", reflect.TypeOf(msg), err)
			return
		}
		err = a.conn.WriteMsg(data...)
		if err != nil {
			log.Error("write message %v error: %v", reflect.TypeOf(msg), err)
		}
	}
}
