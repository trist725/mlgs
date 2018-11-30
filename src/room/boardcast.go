package room

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
	"mlgs/src/msg"
	s "mlgs/src/session"
)

//广播玩家加入
func (r *Room) BoardCastPJ(players []*cache.Player) {
	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Error("use nil session on BoardCastPJ")
			return
		}

		send := msg.Get_S2C_UpdatePlayerJoinRoom()
		for _, p := range players {
			session := s.Mgr().GetSession(p.SessionId())
			if session == nil {
				log.Error("use nil session on BoardCastPJ..2")
				return
			}
			np := msg.Get_Player()
			np.NickName = session.UserData().NickName
			np.Pos = p.Pos()
			np.Chip = p.Chip()
			np.AvatarURL = session.UserData().AvatarURL
			np.UserId = session.UserData().ID
			np.Sex = session.UserData().Sex

			send.Players = append(send.Players, np)
		}

		session.Agent().WriteMsg(send)
	})

	return
}

//广播玩家离开
//func (r *Room) BoardCastPL(ids []int64) {
//	r.PlayerEach(func(player *cache.Player) {
//		session := s.Mgr().GetSession(player.SessionId())
//		if session == nil {
//			log.Error("use nil session id:[%d]", player.SessionId())
//			return
//		}
//
//		send := msg.Get_S2C_UpdatePlayerLeaveRoom()
//		for _, id := range ids {
//			send.UserIds = append(send.UserIds, id)
//		}
//
//		session.Agent().WriteMsg(send)
//	})
//
//	return
//}

//广播玩家掉线
func (r *Room) BoardCastDisConn() {
	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Debug("use nil session on BoardCastDisConn")
			return
		}

		send := msg.Get_S2C_DisConn()
		send.UserId = session.UserData().ID

		session.Agent().WriteMsg(send)
	})

	return
}

//广播玩家离开
func (r *Room) BoardCastPL(id int64) {
	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Error("use nil session on BoardCastPL")
			return
		}

		send := msg.Get_S2C_UpdatePlayerLeaveRoom()
		send.UserId = id

		session.Agent().WriteMsg(send)
	})

	return
}

//广播轮到谁
func (r *Room) BoardCastTurn() {
	r.PlayerEach(func(player *cache.Player) {
		//已掉线
		if player.SessionId() == 0 {
			return
		}
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Error("use nil session on BoardCastTurn")
			return
		}

		send := msg.Get_S2C_Turn()
		send.Pos = r.curPos
		if player.AutoAct() != 0 {
			send.Auto = 1
		}

		session.Agent().WriteMsg(send)
	})

	return
}

//广播游戏开始
func (r *Room) BoardCastGS() {
	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Error("use nil session on BoardCastGS")
			return
		}

		send := msg.Get_S2C_GameStart()
		send.Pos = r.dPos
		cards := player.Cards()
		for _, c := range cards {
			card := msg.Get_Card()
			card.Color = int32(c.Color)
			card.Num = int32(c.Num)
			send.Cards = append(send.Cards, card)
		}
		send.SmallBlind = r.sb

		//最大牌型
		bc := msg.Get_BestCombo()
		bc.Type = player.NutsLevel()
		for _, c := range player.Nuts() {
			card := msg.Get_Card()
			card.Num = int32(c.Num)
			card.Color = int32(c.Color)
			bc.Cards = append(bc.Cards, card)
		}
		send.Best = bc
		session.Agent().WriteMsg(send)
	})

	return
}

//广播玩家操作
func (r *Room) BoardCastTA(ta TurnAction) {
	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Error("use nil session on BoardCastTA")
			return
		}

		send := msg.Get_S2C_TurnAction()
		send.Pos = int32(ta.p.Pos())
		send.Bet = ta.act.Bet
		send.Act = ta.act.Act

		session.Agent().WriteMsg(send)
	})

	return
}

//广播发公共牌
func (r *Room) BoardCastDC(cards []cache.Card) {
	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Error("use nil session on BoardCastTA")
			return
		}

		send := msg.Get_S2C_PublicCard()
		for _, c := range cards {
			card := msg.Get_Card()
			card.Color = int32(c.Color)
			card.Num = int32(c.Num)
			send.Cards = append(send.Cards, card)
		}

		//最大牌型
		player.CalNuts(r.pc)
		bc := msg.Get_BestCombo()
		bc.Type = player.NutsLevel()
		for _, c := range player.Nuts() {
			card := msg.Get_Card()
			card.Num = int32(c.Num)
			card.Color = int32(c.Color)
			bc.Cards = append(bc.Cards, card)
		}
		send.Best = bc

		session.Agent().WriteMsg(send)
	})

	return
}

//广播游戏结束
func (r *Room) BoardCastGO() {
	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Debug("use nil session on BoardCastGO")
			return
		}

		send := msg.Get_S2C_GameOver()

		session.Agent().WriteMsg(send)
	})

	return
}
