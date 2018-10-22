package session

import (
	"fmt"
	"github.com/trist725/mgsu/event"
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"github.com/trist725/myleaf/timer"
	ext "mlgs/src/external"
	l "mlgs/src/logic"
	"mlgs/src/model"
	"sync/atomic"
	"time"
)

//todo:心跳处理
//非线程安全
type Session struct {
	id uint64
	//事件管理器
	eventHandlerMgr *event.HandlerManager
	//定时写库
	timer *timer.Timer
	sign  string // 日志标识

	//防循环引用
	logicMap  map[l.ID]ext.ILogic
	agent     gate.Agent
	closeFlag int32
	user      *model.User    // 需要保存到数据库的用户数据
	account   *model.Account // 帐号数据
	//cache       *cache.User    // 不需要保存到数据库的临时数据
}

var gSessionId uint64

func New(agent gate.Agent, account *model.Account, user *model.User) *Session {
	session := &Session{
		agent:           agent,
		account:         account,
		user:            user,
		id:              atomic.AddUint64(&gSessionId, 1),
		eventHandlerMgr: event.NewHandlerManager(),
		sign:            fmt.Sprintf("user-%d-%s", user.ID, user.NickName),
	}
	//用于从agent获取到session
	session.agent.SetUserData(session.id)

	session.logicMap = l.GenerateLogicMap(session)
	if err := session.initLogic(); err != nil {
		log.Fatal("init logicMap failed, %s", err)
		return nil
	}
	session.runLogic()

	if gSessionManager == nil {
		panic("new session failed, because gSessionManager is nil")
	}
	gSessionManager.putSession(session)
	return session
}

func (s *Session) initLogic() error {
	for _, lm := range s.logicMap {
		if err := lm.Init(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Session) runLogic() {
	for _, lm := range s.logicMap {
		lm.Run()
	}
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

func (s *Session) GetLogic(id l.ID) ext.ILogic {
	if lm, ok := s.logicMap[id]; ok {
		return lm
	}
	return nil
}

func (s *Session) SaveData() {
	if s.user != nil {
		// 保存用户数据
		log.Debug("[%s] save data on [%v]", s.sign, time.Now())
		dbSession := model.GetSession()
		s.user.Update(dbSession, s.user.ID, s.user)
		model.PutSession(dbSession)
	}
}

func (session *Session) IsClosed() bool {
	return atomic.LoadInt32(&session.closeFlag) == 1
}

//todo：断线重连,deepcopy保存快照
func (s *Session) Close() error {
	if atomic.CompareAndSwapInt32(&s.closeFlag, 0, 1) {
		if gSessionManager == nil {
			panic("close session failed because gSessionManager is nil")
		}
		s.agent.Close()
		//更新最后登出时间
		if s.user != nil {
			s.user.LastLogoutTime = time.Now().Unix()
		}
		s.SaveData()
		if s.timer != nil {
			s.timer.Stop()
		}
		gSessionManager.delSession(s)
	}
	return nil
}

func (s *Session) Sign() string {
	return s.sign
}

func (s *Session) SetSign(sign string) {
	s.sign = sign
}

func (s *Session) SetTimer(t *timer.Timer) {
	s.timer = t
}

func GetSession(sid uint64) *Session {
	mgr := Mgr()
	if mgr == nil {
		log.Fatal("gSessionManager is nil, session id:[%d]", sid)
		return nil
	}

	session := mgr.GetSession(sid)
	if session == nil {
		log.Debug("session id:[%d] not exist", sid)
		return nil
	}

	return session
}
