package protobuf

import (
	"errors"

	"github.com/trist725/myleaf/chanrpc"
)

// -------------------------
// | id | protobuf message |
// -------------------------
type ClientProcessor struct {
	littleEndian bool
	msgRouter    *chanrpc.Server
}

func NewClientProcessor() *ClientProcessor {
	p := new(ClientProcessor)
	p.littleEndian = true
	return p
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *ClientProcessor) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *ClientProcessor) SetRouter(msgRouter *chanrpc.Server) {
	p.msgRouter = msgRouter
}

// goroutine safe
func (p *ClientProcessor) Route(msg interface{}, userData interface{}) error {
	if p.msgRouter != nil {
		p.msgRouter.Go("ClientForward", msg, userData)
	}
	return nil
}

// goroutine safe
func (p *ClientProcessor) Unmarshal(data []byte) (interface{}, error) {
	if len(data) < 2 {
		return nil, errors.New("protobuf data too short")
	}

	return data, nil
}

// goroutine safe
func (p *ClientProcessor) Marshal(msg interface{}) ([][]byte, error) {
	msgByte := msg.([]byte)
	if len(msgByte) < 2 {
		return nil, errors.New("protobuf data too short")
	}
	return [][]byte{msgByte[:2], msgByte[2:]}, nil
}
