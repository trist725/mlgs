package session

//ref https://github.com/funny/link

import (
	"sync"
)

var gSessionManager = newManager()

func Mgr() *Manager {
	//todo: recover gSessionManager
	if gSessionManager == nil {
		panic("gSessionManager is nil")
	}
	return gSessionManager
}

const sessionMapNum = 32

type Manager struct {
	sessionMaps [sessionMapNum]sessionMap
	disposeOnce sync.Once
	disposeWait sync.WaitGroup
	tickOnce    sync.Once
}

type sessionMap struct {
	sync.RWMutex
	sessions map[uint64]*Session
	disposed bool
}

func newManager() *Manager {
	manager := &Manager{}
	for i := 0; i < len(manager.sessionMaps); i++ {
		manager.sessionMaps[i].sessions = make(map[uint64]*Session)
	}
	return manager
}

func (manager *Manager) Dispose() {
	manager.disposeOnce.Do(func() {
		for i := 0; i < sessionMapNum; i++ {
			smap := &manager.sessionMaps[i]
			smap.Lock()
			smap.disposed = true
			for _, session := range smap.sessions {
				session.Close()
			}
			smap.Unlock()
		}
		manager.disposeWait.Wait()
	})
}

func (manager *Manager) GetSession(sessionID uint64) *Session {
	smap := &manager.sessionMaps[sessionID%sessionMapNum]
	smap.RLock()
	defer smap.RUnlock()

	session, _ := smap.sessions[sessionID]
	return session
}

func (manager *Manager) putSession(session *Session) {
	smap := &manager.sessionMaps[session.id%sessionMapNum]

	smap.Lock()
	defer smap.Unlock()

	if smap.disposed {
		session.Close()
		return
	}

	smap.sessions[session.id] = session
	manager.disposeWait.Add(1)
}

func (manager *Manager) delSession(session *Session) {
	smap := &manager.sessionMaps[session.id%sessionMapNum]

	smap.Lock()
	defer smap.Unlock()

	delete(smap.sessions, session.id)
	manager.disposeWait.Done()
}

func (manager *Manager) DelSessionById(sessionID uint64) {
	session := manager.GetSession(sessionID)
	if session == nil {
		return
	}
	manager.delSession(session)
}

func (manager *Manager) GetByUserId(id int64) *Session {
	for i := 0; i < sessionMapNum; i++ {
		smap := &manager.sessionMaps[i]
		for _, s := range smap.sessions {
			if s.UserData().ID == id {
				return s
			}
		}
	}

	return nil
}

func (manager *Manager) Count() uint64 {
	var count uint64
	for i := 0; i < sessionMapNum; i++ {
		smap := &manager.sessionMaps[i]
		count += uint64(len(smap.sessions))
	}

	return count
}
