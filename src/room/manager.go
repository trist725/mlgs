package room

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
	"mlgs/src/sd"
	"sync"
	"sync/atomic"
)

var gRoomManager = newManager()

func Mgr() *Manager {
	//todo: recover gRoomManager
	if gRoomManager == nil {
		panic("gRoomManager is nil")
	}
	return gRoomManager
}

const roomMapNum = 4

type Manager struct {
	roomMaps    [roomMapNum]roomMap
	disposeOnce sync.Once
	disposeWait sync.WaitGroup
}

type roomMap struct {
	sync.RWMutex
	rooms    map[uint64]*Room
	disposed bool
}

func newManager() *Manager {
	manager := &Manager{}
	for i := 0; i < len(manager.roomMaps); i++ {
		manager.roomMaps[i].rooms = make(map[uint64]*Room)
	}
	return manager
}

func (manager *Manager) NewRoom(pt uint32, gt uint32, t int64) *Room {
	//todo:根据t创建不同房间类型
	var rommSd *sd.Room
	switch t {
	default:
		rommSd = sd.RoomMgr.Get(t)
		if rommSd == nil {
			log.Fatal("策划坑爹了,读room表有误，id: [%d]", t)
			return nil
		}
	}

	room := &Room{
		id:         atomic.AddUint64(&gRoomId, 1),
		pType:      pt,
		gType:      gt,
		sb:         rommSd.Sb,
		bb:         rommSd.Bb,
		players:    make(map[uint32]*cache.Player),
		bystanders: make(map[int64]*cache.Player),
	}

	room.stopSig = make(chan struct{})

	room.refreshReadyTimeSig = make(chan struct{})
	room.actSig = make(chan TurnAction)

	manager.putRoom(room)
	return room
}

func (manager *Manager) Dispose() {
	manager.disposeOnce.Do(func() {
		for i := 0; i < roomMapNum; i++ {
			rmap := &manager.roomMaps[i]
			rmap.Lock()
			rmap.disposed = true
			for _, room := range rmap.rooms {
				room.Destroy()
			}
			rmap.Unlock()
		}
		manager.disposeWait.Wait()
	})
}

func (manager *Manager) putRoom(room *Room) {
	rmap := &manager.roomMaps[room.id%roomMapNum]

	rmap.Lock()
	defer rmap.Unlock()

	if rmap.disposed {
		room.Destroy()
		return
	}

	rmap.rooms[room.id] = room
	manager.disposeWait.Add(1)
}

func (manager *Manager) delRoom(room *Room) {
	rmap := &manager.roomMaps[room.id%roomMapNum]

	rmap.Lock()
	defer rmap.Unlock()

	delete(rmap.rooms, room.id)
	manager.disposeWait.Done()
}

func (manager *Manager) GetRoom(roomID uint64) *Room {
	rmap := &manager.roomMaps[roomID%roomMapNum]
	rmap.RLock()
	defer rmap.RUnlock()

	room, _ := rmap.rooms[roomID]
	return room
}

func (manager *Manager) PlayerJoin(p *cache.Player, t uint32) bool {
	for i := 0; i < roomMapNum; i++ {
		rmap := &manager.roomMaps[i]
		for _, r := range rmap.rooms {
			//对局中的暂不允许加入
			//房间类型匹配
			if r.Stage() > 0 || r.pType != t {
				continue
			}
			if p.RoomId() != 0 {
				log.Debug("[%s] already in room:[%d]", p.UserId(), p.RoomId())
				return false
			}
			if err := r.PlayerJoin(p); err == nil {
				//r.BoardCast(p)
				return true
			}
		}
	}

	return false
}

func (manager *Manager) BystanderJoin(p *cache.Player) bool {
	for i := 0; i < roomMapNum; i++ {
		rmap := &manager.roomMaps[i]
		for _, r := range rmap.rooms {
			if success := r.bystanderJoin(p); success {
				return true
			}
		}
	}

	return false
}
