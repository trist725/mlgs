package base

import (
	m "mlgs/src/msg"
	"net"
	"reflect"

	"github.com/trist725/myleaf/log"
	"github.com/trist725/myleaf/network"
	"github.com/trist725/myleaf/network/protobuf"
)

type Agent struct {
	conn *network.TCPConn
	//复用gate解包
	processor network.Processor
	userData  interface{}
}

func (a *Agent) Run() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			log.Debug("read message: %v", err)
			break
		}

		if a.processor != nil {
			//截取clientID
			msg, err := a.processor.Unmarshal(data)
			if err != nil {
				log.Debug("unmarshal raw message error: %v", err)
				break
			}
			a.SetUserData(a.processor.(*protobuf.ServerProcessor).ClientID())
			//a.processor.Route(msg, a)
			//实际解包路由
			msg, err = m.Processor.Unmarshal(msg.([]byte))
			if err != nil {
				log.Debug("unmarshal message error: %v", err)
				break
			}
			err = m.Processor.Route(msg, a)
			if err != nil {
				log.Debug("route message error: %v", err)
				break
			}
		}
	}
}

func (a *Agent) WriteMsg(msg interface{}) {
	if m.Processor != nil {
		data, err := m.Processor.Marshal(msg)
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

func (a *Agent) WriteMsgEx(ext interface{}, msg interface{}) {
	if m.Processor != nil {
		data, err := m.Processor.Marshal(msg)
		if err != nil {
			log.Error("marshal message %v error: %v", reflect.TypeOf(msg), err)
			return
		}
		//insert ext
		extByte := ext.([][]byte)
		extByte = append(extByte, data...)
		err = a.conn.WriteMsg(extByte...)
		if err != nil {
			log.Error("write message %v error: %v", reflect.TypeOf(msg), err)
		}
	}
}

func (a *Agent) WriteCmd(cmdType uint16, cmdData interface{}) {
	cmdHead := Int32ToByteArr(0)
	//insert ext
	cmdTypeByte := Uint16ToByteArr(cmdType)
	var cmdDataByte []byte
	switch cmdType {
	case 0:
		cmdDataByte = Int32ToByteArr(cmdData.(int32))
	case 1:
	default:
	}
	err := a.conn.WriteMsg(cmdHead, cmdTypeByte, cmdDataByte)
	if err != nil {
		log.Error("write cmd type:%d error: %v", cmdType, err)
	}
}

func (a *Agent) UserData() interface{} {
	return a.userData
}

func (a *Agent) SetUserData(data interface{}) {
	a.userData = data
}

func (a *Agent) OnClose() {}

func (a *Agent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *Agent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *Agent) Close() {
	a.conn.Close()
}

func (a *Agent) Destroy() {
	a.conn.Destroy()
}

func (a *Agent) SetConn(conn *network.TCPConn) {
	a.conn = conn
}

func (a *Agent) SetProcessor(processor *protobuf.ServerProcessor) {
	a.processor = processor
}

func (a *Agent) Processor() *protobuf.ServerProcessor {
	return a.processor.(*protobuf.ServerProcessor)
}
