package session

import (
	"fmt"
	"github.com/trist725/mgsu/event"
	"github.com/trist725/myleaf/gate"
	"log"
	"mlgs/src/model"
	"sync/atomic"
)

type Session struct {
	id uint64
	//事件管理器
	eventHandlerMgr *event.HandlerManager
	//定时写库
	//timer
	//sign string // 日志标识
	agent     gate.Agent
	closeFlag int32
	user      *model.User    // 需要保存到数据库的用户数据
	account   *model.Account // 帐号数据
	//cache       *cache.User    // 不需要保存到数据库的临时数据
}

var gSessionId uint64

func NewSession(agent gate.Agent, account *model.Account, user *model.User) *Session {
	session := &Session{
		agent:           agent,
		account:         account,
		user:            user,
		id:              atomic.AddUint64(&gSessionId, 1),
		eventHandlerMgr: event.NewHandlerManager(),
	}
	//用于从agent获取到session
	session.agent.SetUserData(session.id)
	if gSessionManager == nil {
		log.Fatal("gSessionManager is nil")
	}
	gSessionManager.putSession(session)
	return session
}

func (s *Session) RegisterEventHandler(id event.ID, handler event.Handler) {
	s.eventHandlerMgr.Register(id, handler)
}

func (s *Session) ProcessEvent(ev event.IEvent) error {
	return s.eventHandlerMgr.Process(ev)
}

func (s *Session) ID() uint64 {
	return s.id
}

func (s *Session) AccountData() *model.Account {
	return s.account
}

func (s *Session) UserData() *model.User {
	return s.user
}

func (s *Session) SetAccountData(account *model.Account) {
	s.account = account
}

func (s *Session) SetUserData(user *model.User) {
	s.user = user
}

//func (s *Session) SetLeafAgent(a *gate.Agent) {
//	s.agent = a
//}
//
//func (s *Session) LeafAgent() *gate.Agent{
//	return s.agent
//}

func (session *Session) IsClosed() bool {
	return atomic.LoadInt32(&session.closeFlag) == 1
}

func (s *Session) Close() error {
	if atomic.CompareAndSwapInt32(&s.closeFlag, 0, 1) {
		if gSessionManager == nil {
			return fmt.Errorf("close session faild because gSessionManager is nil")
		}
		s.agent.Close()
		gSessionManager.delSession(s)
	}
	return nil
}
