package room

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
	"mlgs/src/msg"
	s "mlgs/src/session"
)

//广播机器人玩家加入
func (r *Room) BoardCastRPJ(players []*cache.Player) {
	r.PlayerEach(func(player *cache.Player) {
		if player.Robot() {
			return
		}
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Error("use nil session on BoardCastPJ")
			return
		}

		send := msg.Get_S2C_UpdatePlayerJoinRoom()
		for _, p := range players {
			if !p.Robot() {
				continue
			}

			np := msg.Get_Player()
			np.NickName = p.UserData().NickName
			np.Pos = p.Pos()
			np.Chip = p.Chip()
			np.AvatarURL = p.UserData().AvatarURL
			np.UserId = p.UserId()
			np.Sex = p.UserData().Sex

			send.Players = append(send.Players, np)
		}

		session.Agent().WriteMsg(send)
	})

	return
}

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
func (r *Room) BoardCastDisConn(uid int64) {
	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Debug("use nil session on BoardCastDisConn")
			return
		}

		send := msg.Get_S2C_DisConn()
		send.UserId = uid

		session.Agent().WriteMsg(send)
	})

	return
}

//广播玩家离开
func (r *Room) BoardCastPL(id int64, reason msg.S2C_UpdatePlayerLeaveRoom_E_Err) {
	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Debug("use nil session on BoardCastPL")
			return
		}

		send := msg.Get_S2C_UpdatePlayerLeaveRoom()
		send.UserId = id
		send.Reason = reason

		session.Agent().WriteMsg(send)
	})

	return
}

//广播轮到谁
func (r *Room) BoardCastTurn() {
	p, ok := r.players[r.curPos]
	if !ok {
		log.Error("BoardCastTurn: invalid curPos:[%d]", r.curPos)
		return
	}
	send := msg.Get_S2C_Turn()
	send.Pos = r.curPos
	if p.AutoAct() != 0 {
		send.Auto = 1
	}

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
		send.Round = int32(player.Round())
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

//广播房间聊天
func (r *Room) BoardCastRC(m *msg.C2S_RoomChat) {
	send := msg.Get_S2C_RoomChat()
	send.Content = m.Content
	send.DstUserId = m.DstUserId
	send.SrcUserId = m.SrcUserId

	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Debug("use nil session on BoardCastGO")
			return
		}

		session.Agent().WriteMsg(send)
	})

	return
}

//广播结算
func (r *Room) BoardCastBalance() {
	send := msg.Get_S2C_Balance()

	send.Balances = r.MakeBalanceMsg()

	r.PlayerEach(func(player *cache.Player) {
		session := s.Mgr().GetSession(player.SessionId())
		if session == nil {
			log.Debug("use nil session on BoardCastBalnce")
			return
		}

		session.Agent().WriteMsg(send)
	})

	for _, b := range send.Balances {
		log.Debug("=======================: %v", b)
	}
	//for _, p := range r.players {
	//	log.Debug("######################id: %d : %v", p.UserId(), p.Nuts())
	//}
	//return
}

func (r *Room) MakeBalanceMsg() (nb []*msg.Balance) {
	r.PlayerEach(func(player *cache.Player) {
		b := msg.Get_Balance()
		//手牌
		for _, c := range player.Cards() {
			handCard := msg.Get_Card()
			handCard.Num = int32(c.Num)
			handCard.Color = int32(c.Color)
			b.Cards = append(b.Cards, handCard)
		}
		//最大牌
		b.BestCombo = msg.Get_BestCombo()
		for _, n := range player.Nuts() {
			nutCard := msg.Get_Card()
			nutCard.Num = int32(n.Num)
			nutCard.Color = int32(n.Color)
			b.BestCombo.Cards = append(b.BestCombo.Cards, nutCard)
			b.BestCombo.Type = player.NutsLevel()
		}

		b.Gain = player.Gain()
		b.Refund = player.RefundBet()
		b.UserId = player.UserId()
		b.WinRound = int32(player.WinTimes())

		nb = append(nb, b)
	})
	return nb
}
