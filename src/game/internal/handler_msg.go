package internal

import (
	"github.com/trist725/myleaf/gate"
	"github.com/trist725/myleaf/log"
	"mlgs/src/msg"
	s "mlgs/src/session"
	"reflect"
)

func init() {
	regiserMsgHandle(&msg.C2S_DaySign{}, handleDaySign)
}

func regiserMsgHandle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleDaySign(args []interface{}) {
	// 收到的消息
	recv := args[0].(*msg.C2S_DaySign)
	// 消息的发送者
	sender := args[1].(gate.Agent)
	send := msg.Get_S2C_DaySign()

	sid := sender.UserData().(uint64)
	session := s.GetSession(sid)
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
		log.Debug("sign day invaild")
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
		user.SignedDays = 0
	}
}
