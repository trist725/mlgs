package room

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
	"sort"
)

//结算
func (r *Room) Balance() {
	var ps, psFold cache.PlayerSlice
	r.PlayerEach(func(player *cache.Player) {
		if player == nil {
			log.Error("Balance: invalid player")
			return
		}

		//下过注但弃牌的
		if player.Stat() == 2 && player.TotalBet() > 0 {
			psFold = append(psFold, player)
		}
		//最后留下的和allin的
		if player.Stat() == 3 || player.Stat() == 1 {
			ps = append(ps, player)
		}
		player.SetRefundBet(player.TotalBet())
	})
	//非弃牌的玩家排序
	sort.Sort(cache.PlayerSlice(ps))
	//非弃牌的先分了弃牌的钱
	r.DivideLoser(ps, psFold, true)

	r.DivideLoser(ps, nil, false)

	r.PlayerEach(func(player *cache.Player) {
		player.SetChip(player.Chip() + player.Gain() + player.RefundBet())
	})

	r.BoardCastBalance()
}

func (r *Room) DivideLoser(winners cache.PlayerSlice, losers cache.PlayerSlice, flop bool) {
	for i := 0; i < winners.Len(); i++ {
		//n人平牌
		var psDraw cache.PlayerSlice = nil
		if i+1 < winners.Len() && winners[i].CompareCards(winners[i+1].Nuts()) == nil {
			psDraw = append(psDraw, winners[i])
			psDraw = append(psDraw, winners[i+1])
			for k := i + 2; k < winners.Len(); k++ {
				if winners[k].CompareCards(winners[k].Nuts()) == nil {
					psDraw = append(psDraw, winners[k])
				}
			}
			if !flop {
				if psDraw.Len()+i >= winners.Len() {
					return
				}
				losers = winners[psDraw.Len()+i:]
			}
			//平分失败者的筹码
			r.DivideChip(psDraw, losers)
			//截取剩余的
			if psDraw.Len()+i >= winners.Len() {
				return
			}
			winners = winners[psDraw.Len()+i:]
			i = -1
		} else {
			//非平牌
			if flop {
				for _, loser := range losers {
					r.GainBet(winners[i], loser)
				}
			} else {
				for j := i + 1; j < winners.Len(); j++ {
					r.GainBet(winners[i], winners[j])
				}
			}
		}
	}

	//if !flop {
	//	return
	//}
	////弃牌的人还有钱,递归继续分
	//var newLosers cache.PlayerSlice = nil
	//for _, loser := range losers {
	//	if loser.RefundBet() > 0 {
	//		newLosers = append(newLosers, loser)
	//	}
	//}
	//if newLosers.Len() > 0 {
	//	r.DivideLoser(winners, newLosers, true)
	//}
}

func (r *Room) DivideChip(winners cache.PlayerSlice, losers cache.PlayerSlice) {
	var wLen, lLen int
	//没有牌更小的玩家
	if lLen = losers.Len(); lLen == 0 {
		return
	}
	wLen = winners.Len()

	for _, loser := range losers {
		if loser.RefundBet() == 0 {
			continue
		}

		divide := loser.RefundBet() / int64(wLen)
		remainder := loser.RefundBet() % int64(wLen)

		for i := 0; i < winners.Len(); i++ {
			if remainder > 0 {
				//前面的人多给1
				r.GainDivide(winners[i], loser, divide, 1)
				remainder--
			} else {
				r.GainDivide(winners[i], loser, divide, 0)
			}
		}
	}
}

func (r *Room) GainDivide(big *cache.Player, small *cache.Player, divide int64, remainder int64) int64 {
	if big.TotalBet() < divide {
		big.AddGain(big.TotalBet())
		small.AddGain(big.TotalBet() * -1)
		small.SetRefundBet(small.RefundBet() - big.TotalBet())
	}

	big.AddGain(divide)
	small.AddGain(divide * -1)
	small.SetRefundBet(small.RefundBet() - divide)

	if remainder > 0 {
		big.AddGain(remainder)
		small.AddGain(remainder * -1)
		small.SetRefundBet(small.RefundBet() - remainder)
	}
	return 0
}

func (r *Room) GainBet(big *cache.Player, small *cache.Player) {
	if small.RefundBet() == 0 {
		return
	}
	if big.TotalBet() < small.RefundBet() {
		big.AddGain(big.TotalBet())
		small.AddGain(big.TotalBet() * -1)
		small.SetRefundBet(small.RefundBet() - big.TotalBet())
	} else {
		big.AddGain(small.RefundBet())
		small.AddGain(small.RefundBet() * -1)
		small.SetRefundBet(0)
	}
}
