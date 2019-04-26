package internal

import (
	"encoding/base64"
	"encoding/json"
	"github.com/trist725/mgsu/util"
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"mlgs/src/cache"
	"mlgs/src/conf"
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

	regiserMsgHandle(&msg.C2S_SwitchHallRoleSex{}, handleSwitchHallRoleSex)

	regiserMsgHandle(&msg.C2S_GetNotices{}, handleGetNotices)

	regiserMsgHandle(&msg.C2S_SyncGameStatus{}, handleSyncGameStatus)

	regiserMsgHandle(&msg.C2S_Charge{}, handleCharge)
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

	if player.InRoom() && r.Stage() == 0 && sd.E_RoomType(r.RoomType()) == sd.E_RoomType_Match {
		send.Err = msg.S2C_PlayerLeaveRoom_E_Err_Playing
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
	//创建游戏内数据
	if player == nil {
		player = cache.NewPlayer(session.ID(), session.UserData().ID, recv.Type, session.UserData())
		session.SetPlayer(player)
	} else if player.InRoom() {
		log.Debug("player:[%d] already in room:[%d]", player.UserId(), player.RoomId())
		send.Err = msg.S2C_QuickMatchStart_E_Err_AlreadyInGame
		sender.WriteMsg(send)
		return
	}

	success := room.Mgr().PlayerJoin(player, uint32(recv.Type))
	//无空房,新建
	if !success {
		nr := room.Mgr().NewRoom(uint32(recv.Type), 1, recv.Type)
		if err := nr.PlayerJoin(player); err == nil {
			success = true
		}
	}
	//新建房间加入还是有可能失败
	if !success {
		send.Err = msg.S2C_QuickMatchStart_E_Err_Room
		log.Error("new room failed")
		sender.WriteMsg(send)
		return
	}
	log.Debug("[%s] roomId:[%d] type:[%d] pos:[%d]", session.Sign(), player.RoomId(), recv.Type, player.Pos())

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

	UpdateMails(ud)

	for _, mail := range ud.Mails {
		if mail.Received == false {
			nm := msg.Get_Mail()
			send.Mails = append(send.Mails, mail.ToMsg(nm))
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
				if _, _, err := ud.Gain(mail.RewardType, mail.RewardNum, false, nil); err != nil {
					send.Err = msg.S2C_GetMailReward_E_Err_UnKnown
					break
				}
				if sd.E_Money(mail.RewardType) == sd.E_Money_Gold {
					ud.GainCoin += mail.RewardNum
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
			if _, _, err := ud.Gain(mail.RewardType, mail.RewardNum, false, nil); err != nil {
				continue
			}
			if sd.E_Money(mail.RewardType) == sd.E_Money_Gold {
				ud.GainCoin += mail.RewardNum
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

func handleSwitchHallRoleSex(args []interface{}) {
	recv := args[0].(*msg.C2S_SwitchHallRoleSex)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleSwitchHallRoleSex return for nil session")
		return
	}

	send := msg.Get_S2C_SwitchHallRoleSex()
	defer sender.WriteMsg(send)
	defer session.Update()

	ud := session.UserData()
	if ud == nil {
		log.Error("[%s] userData in session:[%d] is nil", session.Sign(), session.ID())
		return
	}

	if recv.Sex < 0 || recv.Sex > 1 {
		send.Err = msg.S2C_SwitchHallRoleSex_E_Err_Invalid_Param
	}

	ud.HallRoleSex = recv.Sex
	send.Err = msg.S2C_SwitchHallRoleSex_E_Err_Success
}

func handleGetNotices(args []interface{}) {
	//recv := args[0].(*msg.C2S_GetNotices)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleGetNotices return for nil session")
		return
	}

	send := msg.Get_S2C_GetNotices()
	defer sender.WriteMsg(send)
	defer session.Update()

	GetNotices()

	send.Notices = append(send.Notices, ConvertNotices()...)

}

func handleSyncGameStatus(args []interface{}) {
	//recv := args[0].(*msg.C2S_GetNotices)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleSyncGameStatus return for nil session")
		return
	}

	send := msg.Get_S2C_SyncGameStatus()

	defer session.Update()

	player := session.Player()
	if player == nil {
		log.Debug("[%s], handleSyncGameStatus: player data is nil", session.Sign())
		return
	}
	r := room.Mgr().GetRoom(player.RoomId())
	if r == nil {
		log.Debug("[%s], handleSyncGameStatus: room data is nil", session.Sign())
		return
	}
	send.Room = msg.Get_Room()
	send.Room.Name = r.Name()
	send.Room.Id = r.Id()
	send.Room.Chip = r.Pot()
	send.Room.MaxBet = r.MaxBet()
	send.Room.RoomType = int32(r.RoomType())
	r.PlayerEach(func(p *cache.Player) {
		userData := p.UserData()
		if userData == nil {
			return
		}
		np := msg.Get_Player()
		np.UserId = p.UserId()
		np.NickName = userData.NickName
		np.AvatarURL = userData.AvatarURL
		np.Pos = p.Pos()
		np.Role = int32(p.Role())
		np.Chip = p.Chip()
		np.BetChip = p.TotalBet()
		np.Sex = userData.Sex
		np.Status = int32(p.Stat())

		//只发自己的手牌信息,防挂
		if p.UserId() == player.UserId() {
			for _, c := range p.Cards() {
				np.Cards = append(np.Cards, c.ToMsg(msg.Get_Card()))
			}
			for _, n := range p.Nuts() {
				np.BestCombo = msg.Get_BestCombo()
				np.BestCombo.Cards = append(np.BestCombo.Cards, n.ToMsg(msg.Get_Card()))
			}
			np.BestCombo.Type = p.NutsLevel()
		}

		send.Room.Players = append(send.Room.Players, np)
	})
	send.SmallBlind = r.SmallBlind()
	for _, c := range r.CommunityCards() {
		send.CommunityCards = append(send.CommunityCards, c.ToMsg(msg.Get_Card()))
	}

	if !player.InRoom() {
		send.GameStage = 6
	} else {
		send.GameStage = int32(r.Stage())
	}

	send.WinRound = int32(player.WinTimes())
	send.Round = int32(player.Round())
	send.CurTurnPos = int32(r.CurPos())
	send.Balances = r.MakeBalanceMsg()

	sender.WriteMsg(send)
	if send.GameStage == 6 {
		ud := session.UserData()
		if ud == nil {
			log.Error("[%s] userData in session:[%d] is nil", session.Sign(), session.ID())
			return
		}
		send := msg.Get_S2C_UpdateMoney()
		for _, nm := range ud.Monies {
			send.Monies = append(send.Monies, nm.ToMsg(msg.Get_Money()))
		}
		sender.WriteMsg(send)
	}
}

func handleCharge(args []interface{}) {
	recv := args[0].(*msg.C2S_Charge)
	sender := args[1].(gate.Agent)
	if sender.UserData() == nil {
		log.Debug("no session yet")
		return
	}
	sid := sender.UserData().(uint64)
	session := s.Mgr().GetSession(sid)
	if session == nil {
		log.Debug("handleCharge return for nil session")
		return
	}
	defer session.Update()
	ud := session.UserData()
	if ud == nil {
		log.Error("[%s] userData in session:[%d] is nil", session.Sign(), session.ID())
		return
	}
	send := msg.Get_S2C_Charge()
	defer sender.WriteMsg(send)

	itemSd := sd.ItemMgr.Get(recv.Id)
	if itemSd == nil {
		log.Error("fail to load item sd")
		send.Err = msg.S2C_Charge_E_Err_UnKnown
		return
	}
	if itemSd.Type != int32(sd.E_Item_Diamond_SupplyBag) {
		log.Error("invalid item id")
		send.Err = msg.S2C_Charge_E_Err_UnKnown
		return
	}

	pubKey, err := ioutil.ReadFile("publicKey.keystore.txt")
	if err != nil {
		log.Error("read public key file error")
		send.Err = msg.S2C_Charge_E_Err_UnKnown
		return
	}
	mer, err := util.RsaPubEncrypt(pubKey, []byte(conf.Server.MerchantCode))
	amount, err := util.RsaPubEncrypt(pubKey, []byte(strconv.Itoa(int(itemSd.BuyCost))))
	tranId, err := util.RsaPubEncrypt(pubKey, []byte(recv.TranCode))
	pc, err := util.RsaPubEncrypt(pubKey, []byte(conf.Server.ProductCode))
	pid, err := util.RsaPubEncrypt(pubKey, []byte(recv.AccountId))

	resp, err := resty.R().SetFormData(map[string]string{
		"Amount":        base64.StdEncoding.EncodeToString(amount),
		"MerchantCode":  base64.StdEncoding.EncodeToString(mer),
		"TransactionId": base64.StdEncoding.EncodeToString(tranId),
		"PlayerId":      base64.StdEncoding.EncodeToString(pid),
		"ProductCode":   base64.StdEncoding.EncodeToString(pc),
		"Token":         recv.Token,
	}).Post(conf.Server.UnionPlatUrl + "App/Account/Cost")
	var respCharge sd.ChargeResp
	if err = json.Unmarshal(resp.Body(), &respCharge); err != nil {
		log.Error("failed to unmarshal login resp")
		send.Err = msg.S2C_Charge_E_Err_UnKnown
		return
	}

	if 0 == respCharge.Code {
		if _, _, err := ud.Gain(itemSd.ID, 1, false, nil); err != nil {
			log.Error("[%s] gain diamond failed", session.Sign())
			send.Err = msg.S2C_Charge_E_Err_UnKnown
			return
		}

		diamond := ud.GetMoney(int32(itemSd.IncomeID))
		if diamond == nil {
			log.Error("[%s] get moeny failed", session.Sign())
			send.Err = msg.S2C_Charge_E_Err_UnKnown
			return
		}
		send.Diamond = diamond.Num
		send.Err = msg.S2C_Charge_E_Err_Succeed
		log.Debug("............%#v", send)
	} else if 512 == respCharge.Code {
		send.Err = msg.S2C_Charge_E_Err_TranCodeAlreadyExist
		return
	} else if 518 == respCharge.Code {
		send.Err = msg.S2C_Charge_E_Err_Not_Enough_Money
		return
	} else if respCharge.Code >= 522 && respCharge.Code <= 525 {
		send.Err = msg.S2C_Charge_E_Err_TokenInvalid
		return
	} else {
		send.Err = msg.S2C_Charge_E_Err_UnKnown
		return
	}

}
