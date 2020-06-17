package protobuf

import (
	"encoding/binary"
	"errors"

	"github.com/trist725/myleaf/chanrpc"
	"github.com/trist725/myleaf/gate"
)

// -------------------------
// | clientID | id | protobuf message |
// -------------------------
// -------------------------
// | 0 | cmdType | cmdData |
// -------------------------
type ServerProcessor struct {
	littleEndian bool
	//默认做转发的router
	defaultRouter *chanrpc.Server
}

func NewServerProcessor() *ServerProcessor {
	p := new(ServerProcessor)
	p.littleEndian = true
	return p
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *ServerProcessor) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

func (p *ServerProcessor) SetDefaultRouter(msgRouter *chanrpc.Server) {
	p.defaultRouter = msgRouter
}

// goroutine safe
func (p *ServerProcessor) Route(msg interface{}, userData interface{}) error {
	//避免加锁,将Unmarshal的工作放到这里
	clientID, msgByte := p.ParseClientID(msg)
	agent := userData.(gate.Agent)
	agent.SetUserData(clientID)
	if p.defaultRouter != nil {
		switch clientID {
		//当clientID为0时,进一步解析作为指令处理
		case 0:
			var cmdType uint16
			if p.littleEndian {
				cmdType = binary.LittleEndian.Uint16(msgByte[:2])
			} else {
				cmdType = binary.BigEndian.Uint16(msgByte[:2])
			}
			p.defaultRouter.Go("ServerCommand", msgByte[2:], userData, cmdType)
		default:
			p.defaultRouter.Go("ServerForward", msgByte, userData, clientID)
		}
	}

	return nil
}

// goroutine safe
func (p *ServerProcessor) Unmarshal(data []byte) (interface{}, error) {
	if len(data) < 2 {
		return nil, errors.New("recv data too short")
	}
	return data, nil
}

// goroutine safe
func (p *ServerProcessor) Marshal(msg interface{}) ([][]byte, error) {
	msgByte := msg.([][]byte)
	if len(msgByte) < 2 {
		return nil, errors.New("send data too short")
	}
	return msgByte, nil
}

//截取clientID
func (p *ServerProcessor) ParseClientID(msg interface{}) (int32, []byte) {
	msgByte := msg.([]byte)
	var clientID int32
	if p.littleEndian {
		clientID = int32(binary.LittleEndian.Uint32(msgByte[:4]))
	} else {
		clientID = int32(binary.BigEndian.Uint32(msgByte[:4]))
	}
	return clientID, msgByte[4:]
}
