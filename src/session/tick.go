package session

import (
	"mlgs/src/sd"
	"time"

	"github.com/trist725/myleaf/module"
)

func (manager *Manager) CheckTick(skeleton *module.Skeleton) {
	manager.tickOnce.Do(func() {
		go manager.checkTick(skeleton)
	})
}

func (manager *Manager) checkTick(skeleton *module.Skeleton) {
	for {
		for i := 0; i < sessionMapNum; i++ {
			smap := &manager.sessionMaps[i]
			smap.Lock()
			for _, session := range smap.sessions {
				if session.lastActiveTime == 0 {
					session.Update()
					continue
				}
				dur := time.Now().Unix() - session.lastActiveTime
				if dur >= sd.InitKickTimeOutClientTime() {
					skeleton.ChanRPCServer.Go("CloseAgent", session.agent)
				}
			}
			smap.Unlock()
		}
		time.Sleep(time.Duration(sd.InitCheckTickInterval()) * time.Second)
	}
}
