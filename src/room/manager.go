package room

import (
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

func (manager *Manager) NewRoom(name string) *Room {
	room := &Room{
		id:      atomic.AddUint64(&gRoomId, 1),
		players: make(map[string]bool, gPlayerLimit),
		name:    name,
	}

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

func (manager *Manager) PlayerJoin(player string) uint64 {
	for i := 0; i < roomMapNum; i++ {
		rmap := &manager.roomMaps[i]
		for _, r := range rmap.rooms {
			if err := r.PlayerJoin(player); err == nil {
				return r.id
			}
		}
	}

	return 0
}

func (manager *Manager) BystanderJoin(p string) bool {
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
