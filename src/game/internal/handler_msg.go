package internal

import (
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
	"mlgs/src/cost"
	"mlgs/src/msg"
	"mlgs/src/room"
	"mlgs/src/sd"
	s "mlgs/src/session"
	"reflect"
	"sort"
	"strconv"
)

func init() {
	regiserMsgHandle(&msg.C2S_DaySign{}, handleDaySign)
	regiserMsgHandle(&msg.C2S_QuickMatchStart{}, handleQuickMatchStart)
	regiserMsgHandle(&msg.C2S_PlayerLeaveRoom{}, handlePlayerLeaveRoom)
	regiserMsgHandle(&msg.C2S_TurnAction{}, handleTurnAction)
	regiserMsgHandle(&msg.C2S_AutoAction{}, handleAutoAction)
	regiserMsgHandle(&msg.C2S_Ping{}, handlePong)
	regiserMsgHandle(&msg.C2S_RoomChat{}, handleRoomChat)
	regiserMsgHandle(&msg.C2S_UpdateUserData{}, handleUpdateUserData)
	regiserMsgHandle(&msg.C2S_GetAllQuests{}, handleGetAllQuests)
	regiserMsgHandle(&msg.C2S_GetQuestReward{}, handleGetQuestReward)
	regiserMsgHandle(&msg.C2S_GetCompletedAchievements{}, handleGetCompletedAchievements)
	regiserMsgHandle(&msg.C2S_GetMailList{}, handleGetMailList)
	regiserMsgHandle(&msg.C2S_GetMailReward{}, handleGetMailReward)
	regiserMsgHandle(&msg.C2S_GetAllMailReward{}, handleGetAllMailReward)

	regiserMsgHandle(&msg.C2S_GetOwnDealerSkins{}, handleGetOwnDealerSkins)
	regiserMsgHandle(&msg.C2S_UsingOwnDealerSkins{}, handleUsingOwnDealerSkins)
	regiserMsgHandle(&msg.C2S_BuyItem{}, handleBuyItem)
}

func regiserMsgHandle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlePlayerLeaveRoom(args []interface{}) {
	// 消息的发送者
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	send := msg.Get_S2C_PlayerLeaveRoom()
	defer sender.WriteMsg(send)

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handlePlayerLeaveRoom return for nil session")
		send.Err = msg.S2C_PlayerLeaveRoom_E_Err_UnKnown
		return
	}

	//对局中不能离开
	player := session.Player()
	if player == nil {
		log.Error("session[%d] without player on handlePlayerLeaveRoom", session.ID())
		send.Err = msg.S2C_PlayerLeaveRoom_E_Err_UnKnown
		return
	}
	if player.InTheGame() {
		log.Debug("player:[%d] in the game, can't leave room", player.UserId())
		send.Err = msg.S2C_PlayerLeaveRoom_E_Err_Playing
		return
	}

	r := room.Mgr().GetRoom(session.Player().RoomId())
	if err := r.PlayerLeave(session.Player(), msg.S2C_UpdatePlayerLeaveRoom_E_Err_Normal); err != nil {
		send.Err = msg.S2C_PlayerLeaveRoom_E_Err_UnKnown
		return
	}

	send.Err = msg.S2C_PlayerLeaveRoom_E_Err_Success
	//ChanRPC.Go("PlayerLeaveRoom", session.UserData().ID, r, msg.S2C_UpdatePlayerLeaveRoom_E_Err_Normal)

	session.Update()
}

func handleQuickMatchStart(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.C2S_QuickMatchStart)
	// 消息的发送者
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	send := msg.Get_S2C_QuickMatchStart()
	//defer sender.WriteMsg(send)

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleQuickMatchStart return for nil session")
		return
	}

	player := session.Player()
	if player != nil {
		if player.InRoom() {
			log.Debug("player:[%d] already in room:[%d]", player.UserId(), player.RoomId())
			return
		}
	}

	//创建游戏内数据
	player = cache.NewPlayer(session.ID(), session.UserData().ID, sd.InitQuickMatchRoomId(), session.UserData())
	session.SetPlayer(player)

	success := room.Mgr().PlayerJoin(player, uint32(recv.Type))
	//无空房,新建
	if !success {
		nr := room.Mgr().NewRoom(uint32(recv.Type), 1, sd.InitQuickMatchRoomId())
		if err := nr.PlayerJoin(player); err == nil {
			success = true
		}
	}
	//新建房间加入还是有可能失败
	if !success {
		send.Err = msg.S2C_QuickMatchStart_E_Err_Room
		log.Error("new room failed")
		return
	}
	log.Debug("[%s] pos:[%d], room id:[%d]", session.Sign(), player.Pos(), player.RoomId())

	r := room.Mgr().GetRoom(player.RoomId())
	send.Room = msg.Get_Room()
	send.Room.Id = r.Id()
	send.Room.Name = r.Name()

	ChanRPC.Go("PlayerJoinRoom", sender, r, send, player)

	session.Update()
}

func handleDaySign(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.C2S_DaySign)
	// 消息的发送者
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}

	send := msg.Get_S2C_DaySign()
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleDaySign return for nil session")
		return
	}

	defer sender.WriteMsg(send)
	user := session.UserData()

	//今日已签到
	if user.DaySigned {
		send.Err = msg.S2C_DaySign_E_Err_AlreadySign
		log.Debug("today already signed")
		return
	}

	//签到天数不对
	if recv.Day != user.SignedDays+1 ||
		recv.Day > signCountPerRound ||
		recv.Day < 0 {
		send.Err = msg.S2C_DaySign_E_Err_Unknown
		log.Debug("sign day invalid")
		return
	}

	send.Err = msg.S2C_DaySign_E_Err_Success
	//给签到奖励
	//todo:目前奖励暂写死
	for _, m := range user.Monies {
		if m.Type == 1 {
			m.Num += int64(recv.Day) * 1000
		}
		nm := msg.Get_Money()
		send.Monies = append(send.Monies, m.ToMsg(nm))
	}

	user.DaySigned = true
	user.SignedDays++
	if user.SignedDays == signCountPerRound+1 {
		user.SignedDays = 1
	}

	session.Update()
}

func handleTurnAction(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.C2S_TurnAction)
	// 消息的发送者
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleTurnAction return for nil session")
		return
	}
	player := session.Player()
	if player == nil {
		log.Debug("handleTurnAction return for nil player")
		return
	}

	r := room.Mgr().GetRoom(player.RoomId())
	if r == nil {
		log.Debug("player not in room")
		return
	}
	r.SendPlayerActionSig(recv, player)

	session.Update()
}

func handleAutoAction(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.C2S_AutoAction)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleAutoAction return for nil session")
		return
	}
	player := session.Player()
	if player == nil {
		log.Debug("handleAutoAction return for nil player")
		return
	}

	if recv.Act < 0 || recv.Act > 4 {
		log.Debug("[%s] invalid auto action", session.Sign())
	}
	player.SetAutoAct(recv.Act)

	session.Update()
}

func handlePong(args []interface{}) {
	//recv := args[0].(*msg.C2S_Ping)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handlePong return for nil session")
		return
	}
	session.Update()
}

func handleRoomChat(args []interface{}) {
	recv := args[0].(*msg.C2S_RoomChat)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}

	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleRoomChat return for nil session")
		return
	}
	player := session.Player()
	if player == nil {
		log.Debug("handleRoomChat return for nil player")
		return
	}
	r := room.Mgr().GetRoom(player.RoomId())
	if r == nil {
		log.Debug("player not in room")
		return
	}

	ChanRPC.Go("RoomChat", r, recv)
	session.Update()
}

func handleUpdateUserData(args []interface{}) {
	//recv := args[0].(*msg.C2S_UpdateUserData)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleUpdateUserData return for nil session")
		return
	}
	send := msg.Get_S2C_UpdateUserData()
	defer sender.WriteMsg(send)

	send.Data = session.UserData().ToMsg(msg.Get_User())
	send.Err = msg.S2C_UpdateUserData_OK
	session.Update()
}

func handleGetAllQuests(args []interface{}) {
	// recv := args[0].(*msg.C2S_GetAllQuests)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleGetAllQuests return for nil session")
		return
	}

	send := msg.Get_S2C_GetAllQuests()
	defer sender.WriteMsg(send)

	ud := session.UserData()
	if ud == nil {
		log.Error("[%s] userData in session:[%d] is nil", session.Sign(), session.ID())
		return
	}

	for _, q := range ud.Quests {
		send.Quests = append(send.Quests, q.ToMsg(msg.Get_Quest()))
	}
	sort.Sort(msg.QuestSlice(send.Quests))

	session.Update()
}

func handleGetQuestReward(args []interface{}) {
	recv := args[0].(*msg.C2S_GetQuestReward)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleGetAllQuests return for nil session")
		return
	}

	send := msg.Get_S2C_GetQuestReward()
	send.CltPath = recv.CltPath
	send.Id = recv.Id
	defer sender.WriteMsg(send)
	defer session.Update()

	ud := session.UserData()
	if ud == nil {
		log.Error("[%s] userData in session:[%d] is nil", session.Sign(), session.ID())
		return
	}

	for _, q := range ud.Quests {
		if recv.Id == q.Id {
			if q.Completed {
				if q.Received {
					send.Err = msg.S2C_GetQuestReward_E_Err_Received
					return
				} else {
					taskSd := sd.TaskMgr.Get(q.Id)
					if taskSd == nil {
						send.Err = msg.S2C_GetQuestReward_E_Err_UnKnown
						return
					}
					//获得奖励
					if _, _, err := ud.Gain(taskSd.Reward, taskSd.Rewardnum, false, nil); err != nil {
						send.Err = msg.S2C_GetQuestReward_E_Err_UnKnown
						return
					}
					if sd.E_Money(taskSd.Reward) == sd.E_Money_Gold {
						ud.GainCoin += taskSd.Rewardnum
					}
					send.Err = msg.S2C_GetQuestReward_E_Err_Success
					q.Received = true
					return
				}
			} else {
				send.Err = msg.S2C_GetQuestReward_E_Err_Not_Completed
				return
			}
		}
	}
}

func handleGetCompletedAchievements(args []interface{}) {
	//recv := args[0].(*msg.C2S_GetCompletedAchievements)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleGetCompletedAchievements return for nil session")
		return
	}

	send := msg.Get_S2C_GetCompletedAchievements()
	defer sender.WriteMsg(send)

	ud := session.UserData()
	if ud == nil {
		log.Error("[%s] userData in session:[%d] is nil", session.Sign(), session.ID())
		return
	}

	for _, a := range ud.Achieves {
		if a.Completed {
			send.Achievements = append(send.Achievements, a.ToMsg(msg.Get_Achievement()))
		}
	}

	session.Update()
}

func handleGetMailList(args []interface{}) {
	//recv := args[0].(*msg.C2S_Get)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleGetMailList return for nil session")
		return
	}

	send := msg.Get_S2C_GetMailList()
	defer sender.WriteMsg(send)

	ud := session.UserData()
	if ud == nil {
		log.Error("[%s] userData in session:[%d] is nil", session.Sign(), session.ID())
		return
	}

	for _, mail := range ud.Mails {
		if mail.Received == false {
			send.Ids = append(send.Ids, strconv.FormatInt(mail.Id, 10))
		}
	}

	session.Update()
}

func handleGetMailReward(args []interface{}) {
	recv := args[0].(*msg.C2S_GetMailReward)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleGetMailList return for nil session")
		return
	}

	send := msg.Get_S2C_GetMailReward()
	defer sender.WriteMsg(send)

	ud := session.UserData()
	if ud == nil {
		log.Error("[%s] userData in session:[%d] is nil", session.Sign(), session.ID())
		return
	}

	for _, mail := range ud.Mails {
		if recv.Id == mail.Id {
			if mail.Received == false {
				//获得奖励
				mailSd := sd.EmailMgr.Get(mail.Id)
				if mailSd == nil {
					log.Error("get mail sd failed on handleGetMailReward")
					send.Err = msg.S2C_GetMailReward_E_Err_UnKnown
					break
				}
				if _, _, err := ud.Gain(mailSd.Reward, mailSd.RewardNumber, false, nil); err != nil {
					send.Err = msg.S2C_GetMailReward_E_Err_UnKnown
					break
				}
				if sd.E_Money(mailSd.Reward) == sd.E_Money_Gold {
					ud.GainCoin += mailSd.RewardNumber
				}
				send.Err = msg.S2C_GetMailReward_E_Err_Succeed
				mail.Received = true
				break
			} else {
				send.Err = msg.S2C_GetMailReward_E_Err_Already_Receive
				break
			}
		}
	}
	send.Id = recv.Id

	session.Update()
}

func handleGetAllMailReward(args []interface{}) {
	//recv := args[0].(*msg.C2S_GetAllMailReward)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleGetAllMailList return for nil session")
		return
	}

	send := msg.Get_S2C_GetAllMailReward()
	defer sender.WriteMsg(send)

	ud := session.UserData()
	if ud == nil {
		log.Error("[%s] userData in session:[%d] is nil", session.Sign(), session.ID())
		return
	}

	for _, mail := range ud.Mails {
		if mail.Received == false {
			//获得奖励
			mailSd := sd.EmailMgr.Get(mail.Id)
			if mailSd == nil {
				log.Error("get mail sd failed on handleGetAllMailReward")
				continue
			}
			if _, _, err := ud.Gain(mailSd.Reward, mailSd.RewardNumber, false, nil); err != nil {
				continue
			}
			if sd.E_Money(mailSd.Reward) == sd.E_Money_Gold {
				ud.GainCoin += mailSd.RewardNumber
			}
			send.Ids = append(send.Ids, strconv.FormatInt(mail.Id, 10))
			mail.Received = true
		}
	}

	session.Update()
}

func handleGetOwnDealerSkins(args []interface{}) {
	//recv := args[0].(*msg.C2S_Get)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleGetOwnDealerSkins return for nil session")
		return
	}

	send := msg.Get_S2C_GetOwnDealerSkins()
	defer sender.WriteMsg(send)
	defer session.Update()

	ud := session.UserData()
	if ud == nil {
		log.Error("[%s] userData in session:[%d] is nil", session.Sign(), session.ID())
		return
	}
	send.Id = ud.UsingDealer
	for _, i := range ud.Items {
		itemSd := sd.ItemMgr.Get(i.TID)
		if itemSd == nil {
			log.Error("get item sd failed on handleGetOwnDealerSkins")
			send.Err = msg.S2C_GetOwnDealerSkins_E_Err_UnKnown
			return
		}
		if sd.E_Item(itemSd.Type) == sd.E_Item_DealerSkin {
			send.Ids = append(send.Ids, strconv.FormatInt(i.TID, 10))
		}
	}
	send.Err = msg.S2C_GetOwnDealerSkins_E_Err_Succeed
}

func handleUsingOwnDealerSkins(args []interface{}) {
	recv := args[0].(*msg.C2S_UsingOwnDealerSkins)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleUsingOwnDealerSkins return for nil session")
		return
	}

	send := msg.Get_S2C_UsingOwnDealerSkins()
	defer sender.WriteMsg(send)
	defer session.Update()

	ud := session.UserData()
	if ud == nil {
		log.Error("[%s] userData in session:[%d] is nil", session.Sign(), session.ID())
		return
	}
	if recv.Id < 0 {
		send.Err = msg.S2C_UsingOwnDealerSkins_E_Err_UnKnown
		return
	}
	//换回默认荷官
	if recv.Id == 0 {
		ud.UsingDealer = recv.Id
		send.Id = recv.Id
		send.Err = msg.S2C_UsingOwnDealerSkins_E_Err_Succeed
		return
	}
	for _, i := range ud.Items {
		if i.TID != recv.Id {
			continue
		}

		ud.UsingDealer = recv.Id
		send.Id = recv.Id
		send.Err = msg.S2C_UsingOwnDealerSkins_E_Err_Succeed
	}
	if send.Err == msg.S2C_UsingOwnDealerSkins_E_Err_ {
		send.Err = msg.S2C_UsingOwnDealerSkins_E_Err_Not_Have
	}
}

func handleBuyItem(args []interface{}) {
	recv := args[0].(*msg.C2S_BuyItem)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleBuyItem return for nil session")
		return
	}

	send := msg.Get_S2C_BuyItem()
	defer sender.WriteMsg(send)
	defer session.Update()

	ud := session.UserData()
	if ud == nil {
		log.Error("[%s] userData in session:[%d] is nil", session.Sign(), session.ID())
		return
	}
	if recv.Id < 0 || recv.Num <= 0 {
		send.Err = msg.S2C_BuyItem_E_Err_UnKnown
		return
	}
	itemSd := sd.ItemMgr.Get(recv.Id)
	if itemSd == nil {
		log.Error("get item sd failed on handleBuyItem")
		send.Err = msg.S2C_BuyItem_E_Err_UnKnown
		return
	}

	var costs cost.Costs
	costs = append(costs, cost.CostItem{itemSd.BuyNeedID, itemSd.BuyCost})
	if err := cost.CanCost(ud, costs, 1); err != nil {
		send.Err = msg.S2C_BuyItem_E_Err_Not_Enough_Money
		return
	}
	_, _, err := ud.Gain(itemSd.IncomeID, itemSd.Income, false, nil)
	if err != nil {
		if err.Error() == "DealerSkin already exist" {
			send.Err = msg.S2C_BuyItem_E_Err_Already_Have
			return
		} else {
			send.Err = msg.S2C_BuyItem_E_Err_UnKnown
			return
		}
	}
	cost.Cost(ud, costs, 1, false, nil)
	if sd.E_Money(itemSd.IncomeID) == sd.E_Money_Gold {
		ud.GainCoin += itemSd.Income
	}
	send.Id = recv.Id
	send.Err = msg.S2C_BuyItem_E_Err_Succeed
}
