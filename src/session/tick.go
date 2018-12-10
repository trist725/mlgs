package session

import (
	"github.com/trist725/myleaf/module"
	"mlgs/src/msg"
	"mlgs/src/sd"
	"time"
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
				} else if dur >= sd.InitCheckTimeOutClientTime() {
					send := msg.Get_S2C_Pong()
					session.agent.WriteMsg(send)
				}
			}
			smap.Unlock()
		}
		time.Sleep(time.Duration(sd.InitCheckTimeOutClientTime()) * time.Second)
	}
}
