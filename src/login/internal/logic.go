package internal

import (
	ext "mlgs/src/external"
	l "mlgs/src/logic"
)

const (
	ID = l.Login
)

type Logic struct {
	ext.ISession

	//cache     *cache.Session
	//reward *ext.IReward
}

func (logic *Logic) Init() error {
	//logic.cache = logic.Cache()
	//logic.reward = logic.GetLogic(l.Reward).(export.IReward)

	logic.registerAllEventHandler()
	return nil
}

func (logic *Logic) Run() {
}

func NewLogic(s ext.ISession) ext.ILogic {
	logic := &Logic{
		ISession: s,
	}
	return logic
}

//注册到gFactoryMap
func Init() {
	l.RegisterFactory(ID, NewLogic)
}
