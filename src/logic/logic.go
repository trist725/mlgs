package logic

import (
	"log"
	ext "mlgs/src/external"
)

type ID = uint8

const (
	Login ID = iota + 1
	//Chat
	//Item
)

type Factory struct {
	ID  ID
	New func(s ext.ISession) ext.ILogic
}

var gFactoryMap = map[ID]*Factory{}

func GetFactory(id ID) (*Factory, bool) {
	if factory, ok := gFactoryMap[id]; ok {
		return factory, true
	}
	return nil, false
}

func RegisterFactory(id ID, New func(s ext.ISession) ext.ILogic) {
	if _, ok := GetFactory(id); ok {
		log.Panicf("logic factory already exist, id=[%d]", id)
	}

	factory := &Factory{
		ID:  id,
		New: New,
	}

	gFactoryMap[factory.ID] = factory
}

func GenerateLogicMap(s ext.ISession) map[ID]ext.ILogic {
	logicMap := map[ID]ext.ILogic{}
	for id, factory := range gFactoryMap {
		logicMap[id] = factory.New(s)
	}
	return logicMap
}
