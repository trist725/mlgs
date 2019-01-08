package room

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
	"mlgs/src/sd"
)

func (r *Room) UpdateMatchQuests() {
	r.PlayerEach(func(player *cache.Player) {
		ud := player.UserData()
		if ud == nil {
			log.Error("UpdateQuests failed, user data is nil")
			return
		}
		for _, q := range ud.Quests {
			taskSd := sd.TaskMgr.Get(q.Id)
			if taskSd == nil {
				log.Error("[%s-%s] get task.xlsx id:[%d] failed", ud.ID, ud.NickName, q.Id)
				continue
			}
			if taskSd.TypeNeed != 1 || q.Completed {
				continue
			}
			rts := taskSd.Room
			for _, rt := range rts {
				//房间类型符合条件
				if r.pType == uint32(rt) {
					//场数
					if q.Progress < taskSd.Need {
						q.Progress += 1
						//完成
						if q.Progress == taskSd.Need {
							//成就任务
							if taskSd.Type == 3 {
								for _, a := range ud.Achieves {
									if a.TaskId == q.Id {
										a.Completed = true
										q.Completed = true
									}
								}
							} else {
								q.Completed = true
							}
						}
					}
					break
				}
			}

		}
	})
}

func (r *Room) UpdateCoinQuests(gain int64) {
	r.PlayerEach(func(player *cache.Player) {
		ud := player.UserData()
		if ud == nil {
			log.Error("UpdateQuests failed, user data is nil")
			return
		}
		for _, q := range ud.Quests {
			taskSd := sd.TaskMgr.Get(q.Id)
			if taskSd == nil {
				log.Error("[%s-%s] get task.xlsx id:[%d] failed", ud.ID, ud.NickName, q.Id)
				continue
			}
			if taskSd.TypeNeed != 2 || q.Completed {
				continue
			}
			rts := taskSd.Room
			for _, rt := range rts {
				//房间类型符合条件
				if r.pType == uint32(rt) {
					if q.Progress < taskSd.Need {
						q.Progress += gain
						//完成
						if q.Progress >= taskSd.Need {
							q.Progress = taskSd.Need
							//成就任务
							if taskSd.Type == 3 {
								for _, a := range ud.Achieves {
									if a.TaskId == q.Id {
										a.Completed = true
										q.Completed = true
									}
								}
							} else {
								q.Completed = true
							}
						}
					}
					break
				}
			}

		}
	})
}
