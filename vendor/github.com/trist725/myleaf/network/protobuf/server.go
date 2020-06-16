package protobuf

import (
	"encoding/binary"
	"errors"

	"github.com/trist725/myleaf/chanrpc"
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
	//暂存客户端ID
	clientID int32
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
	if p.defaultRouter != nil {
		switch p.clientID {
		//当clientID为0时,进一步解析作为指令处理
		case 0:
			msgByte := msg.([]byte)
			var cmdType uint16
			if p.littleEndian {
				cmdType = binary.LittleEndian.Uint16(msgByte[0:2])
			} else {
				cmdType = binary.BigEndian.Uint16(msgByte[0:2])
			}
			p.defaultRouter.Go("ServerCommand", msgByte[2:], userData, cmdType)
		default:
			p.defaultRouter.Go("ServerForward", msg, userData, p.clientID)
		}
	}

	return nil
}

// goroutine safe
func (p *ServerProcessor) Unmarshal(data []byte) (interface{}, error) {
	if len(data) < 2 {
		return nil, errors.New("protobuf data too short")
	}
	if p.littleEndian {
		p.clientID = int32(binary.LittleEndian.Uint32(data[:4]))
	} else {
		p.clientID = int32(binary.BigEndian.Uint32(data[:4]))
	}
	return data[4:], nil
}

// goroutine safe
func (p *ServerProcessor) Marshal(msg interface{}) ([][]byte, error) {
	msgByte := msg.([][]byte)
	if len(msgByte) < 2 {
		return nil, errors.New("protobuf data too short")
	}
	return msgByte, nil
}

func (p *ServerProcessor) ClientID() int32 {
	return p.clientID
}
