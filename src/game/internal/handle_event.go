package internal

import (
	"github.com/trist725/myleaf/gate"
	"mlgs/src/model"
	"mlgs/src/msg"
)

//每轮签到天数
const signCountPerRound = 14

func handleAfterLoginAuthPass(args []interface{}) {
	sender := args[0].(gate.Agent)
	user := args[1].(*model.User)

	send := &msg.S2C_LoginInfo{
		ID:         user.ID,
		NickName:   user.NickName,
		AvatarURL:  user.AvatarURL,
		DaySigned:  user.DaySigned,
		SignedDays: user.SignedDays,
	}

	for _, m := range user.Monies {
		nm := msg.Get_Money()
		send.Monies = append(send.Monies, m.ToMsg(nm))
	}

	//暂写死,构造14天签到奖励
	//todo：动态获取14天签到奖励
	for i := 1; i <= signCountPerRound; i++ {
		item := &msg.Item{
			TID: 1,
			Num: int32(1000 * i),
		}
		send.SignRewards = append(send.SignRewards, item)
	}

	sender.WriteMsg(send)
}
