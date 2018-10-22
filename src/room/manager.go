package room

import "sync"

var gRoomManager = newManager()

func Mgr() *Manager {
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

func (manager *Manager) Dispose() {
	manager.disposeOnce.Do(func() {
		for i := 0; i < roomMapNum; i++ {
			smap := &manager.roomMaps[i]
			smap.Lock()
			smap.disposed = true
			for _, room := range smap.rooms {
				room.Destroy()
			}
			smap.Unlock()
		}
		manager.disposeWait.Wait()
	})
}

func (manager *Manager) putRoom(room *Room) {
	smap := &manager.roomMaps[room.id%roomMapNum]

	smap.Lock()
	defer smap.Unlock()

	if smap.disposed {
		room.Destroy()
		return
	}

	smap.rooms[room.id] = room
	manager.disposeWait.Add(1)
}

func (manager *Manager) delRoom(room *Room) {
	smap := &manager.roomMaps[room.id%roomMapNum]

	smap.Lock()
	defer smap.Unlock()

	delete(smap.rooms, room.id)
	manager.disposeWait.Done()
}

func (manager *Manager) GetRoom(roomID uint64) *Room {
	smap := &manager.roomMaps[roomID%roomMapNum]
	smap.RLock()
	defer smap.RUnlock()

	room, _ := smap.rooms[roomID]
	return room
}
