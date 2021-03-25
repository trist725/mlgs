package session

import (
	"fmt"
	"mlgs/src/conf"
	"mlgs/src/model"
	"sync/atomic"
	"time"

	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"github.com/trist725/myleaf/timer"
)

//非线程安全
type Session struct {
	id uint64
	//定时写库
	timer *timer.Timer
	sign  string // 日志标识

	agent     gate.Agent
	closeFlag int32
	userID    string
	user      *model.User    // 需要保存到数据库的用户数据
	account   *model.Account // 帐号数据
	//cache          *cache.Player  // 玩家游戏过程中的数据
	lastActiveTime int64 //上次活动时间
}

var gSessionId uint64

func New(agent gate.Agent, userID string, account *model.Account, user *model.User) *Session {
	session := &Session{
		agent:   agent,
		account: account,
		user:    user,
		userID:  userID,
		id:      atomic.AddUint64(&gSessionId, 1),
		sign:    fmt.Sprintf("user-%d-%s", user.ID, userID),
	}
	//用于从agent获取到session
	session.agent.SetUserData(session.id)

	if gSessionManager == nil {
		panic("new session failed, because gSessionManager is nil")
	}
	gSessionManager.putSession(session)
	return session
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

func (s *Session) SaveData() {
	if s.user != nil {
		// 保存用户数据
		log.Debug("[%s] save data on [%v]", s.sign, time.Now())
		dbSession := model.SC.GetSession()
		if err := s.user.UpdateByID(dbSession, conf.Server.DBName); err != nil {
			log.Error("[%s], save data error:[%s]", s.sign, err)
		}
		model.SC.PutSession(dbSession)
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
		//if s.cache != nil {
		//	//游戏中,不删session
		//
		//}
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

//func (s *Session) Player() *cache.Player {
//	return s.cache
//}
//
//func (s *Session) SetPlayer(p *cache.Player) {
//	s.cache = p
//}

func (s *Session) Agent() gate.Agent {
	return s.agent
}

func (s *Session) Update() {
	s.lastActiveTime = time.Now().Unix()
}

func (s *Session) SetAgent(a gate.Agent) {
	s.agent = a
}

func (s *Session) SetCloseFlag(f int32) {
	atomic.StoreInt32(&s.closeFlag, f)
}
