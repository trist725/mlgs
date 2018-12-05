package room

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
	"mlgs/src/model"
	"mlgs/src/sd"
)

func (r *Room) AddTestRobot() {
	//机器人暂不用account,只用user
	ns, err := model.NextSeq(model.TblUser)
	if err != nil {
		log.Error("AddTestRobot: get robot user id failed")
		return
	}
	rid := int64(ns)*model.RobotIdTimes + model.RobotIdOffset
	player := cache.NewRobotPlayer(rid, sd.InitQuickMatchRoomId())

	r.PlayerJoin(player)
	var players []*cache.Player
	players = append(players, player)
	r.BoardCastRPJ(players)
}
