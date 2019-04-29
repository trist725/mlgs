package protocol_v2

import (
	"fmt"
	"log"

	p "github.com/trist725/mgsu/network/protocol"
)

type MessageID = string

type IMessage interface {
	V2()
	MessageID() MessageID
	Size() int
	Unmarshal([]byte) error
	Marshal() ([]byte, error)
	MarshalTo([]byte) (int, error)
	Reset()
	String() string
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type messageProducer func() IMessage
type messageRecycler func(IMessage)

type messageFactory struct {
	id       MessageID
	producer messageProducer
	recycler messageRecycler
}

func newMessageFactory(id MessageID, producer messageProducer, recycler messageRecycler) *messageFactory {
	return &messageFactory{
		id:       id,
		producer: producer,
		recycler: recycler,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type IMessageFactoryManager interface {
	Register(iMsg IMessage, producer messageProducer, recycler messageRecycler)
	Produce(id MessageID) (IMessage, error)
	Recycle(iMsg IMessage) error
}

type messageFactoryManager struct {
	factories map[MessageID]*messageFactory
}

func newMessageFactoryManager() IMessageFactoryManager {
	m := &messageFactoryManager{
		factories: make(map[MessageID]*messageFactory),
	}
	return m
}

func (m *messageFactoryManager) Register(iMsg IMessage, producer messageProducer, recycler messageRecycler) {
	if producer == nil {
		log.Panicf("register message factory fail, producer == nil, id=[%v]", iMsg.MessageID())
	}

	if recycler == nil {
		log.Panicf("register message factory fail, recycler == nil, id=[%v]", iMsg.MessageID())
	}

	if f, ok := m.factories[iMsg.MessageID()]; ok {
		log.Panicf("duplicate message factory, id=[%v], factory=[%+v]", iMsg.MessageID(), f)
	}

	m.factories[iMsg.MessageID()] = newMessageFactory(iMsg.MessageID(), producer, recycler)
}

func (m *messageFactoryManager) Produce(id MessageID) (IMessage, error) {
	if factory := m.factories[id]; factory != nil {
		return factory.producer(), nil
	}
	return nil, fmt.Errorf("unsupported message, id=[%v]", id)
}

func (m *messageFactoryManager) Recycle(iMsg IMessage) error {
	if factory := m.factories[iMsg.MessageID()]; factory != nil {
		factory.recycler(iMsg)
		iMsg = nil
		return nil
	}
	return fmt.Errorf("unsupported message, id=[%v]", iMsg.MessageID())
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type MessageHandler func(IMessage, ...interface{})

type IMessageHandlerManager interface {
	Register(id MessageID, handler MessageHandler)
	Process(iMsg IMessage, args ...interface{}) error
}

type messageHandlerManager struct {
	handlerMap map[MessageID][]MessageHandler
}

func NewMessageHandlerManager() IMessageHandlerManager {
	return &messageHandlerManager{
		handlerMap: make(map[MessageID][]MessageHandler),
	}
}

func (m *messageHandlerManager) Register(id MessageID, handler MessageHandler) {
	if handlers, ok := m.handlerMap[id]; !ok {
		m.handlerMap[id] = []MessageHandler{handler}
	} else {
		m.handlerMap[id] = append(handlers, handler)
	}
}

////////////////////////////////////////////////////////////////////////////////
func (m *messageHandlerManager) Process(iMsg IMessage, args ...interface{}) error {
	handlers, ok := m.handlerMap[iMsg.MessageID()]
	if !ok {
		return &p.ErrNoMessageHandler{
			MessageID: iMsg.MessageID(),
		}
	}

	for _, handler := range handlers {
		handler(iMsg, args...)
	}

	return nil
}
