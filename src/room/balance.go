package room

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/cache"
	"sort"
)

//结算
func (r *Room) Balance() {
	var ps cache.PlayerSlice
	r.PlayerEach(func(player *cache.Player) {
		if player == nil {
			log.Error("Balance: invalid player")
			return
		}
		if player.Stat() != 3 || player.Stat() != 1 {
			return
		}

		player.SetRefundBet(player.TotalBet())
		ps = append(ps, player)
	})
	//要结算的玩家排序
	sort.Sort(cache.PlayerSlice(ps))

	for i := 0; i < ps.Len(); i++ {
		for j := i + 1; j < ps.Len(); j++ {
			//n人平牌
			var psDraw cache.PlayerSlice = nil
			if ps[i].CompareCards(ps[j].Nuts()) == nil {
				psDraw = append(psDraw, ps[i])
				psDraw = append(psDraw, ps[j])
				for k := j + 1; k < ps.Len(); k++ {
					if ps[j].CompareCards(ps[k].Nuts()) == nil {
						psDraw = append(psDraw, ps[k])
					}
				}
				//平分剩下人的筹码
				r.DivideChip(psDraw, ps[psDraw.Len():])
				ps = ps[psDraw.Len():]
				if ps.Len() >= 2 {
					i = -1
					j = 0

				} else { //结算结束
					i = ps.Len()
					j = ps.Len()
				}
				continue
			}

			//没钱的孩子
			if ps[j].RefundBet() == 0 {
				continue
			}
			r.GainBet(ps[i], ps[j])
		}
	}

	r.BoardCastBalance()
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

		for _, winner := range winners {
			r.GainDivide(winner, loser, divide, &remainder)
		}
	}
}

func (r *Room) GainDivide(big *cache.Player, small *cache.Player, divide int64, remainder *int64) int64 {
	if big.TotalBet() < divide {
		big.AddGain(big.TotalBet())
		small.AddGain(big.TotalBet() * -1)
		small.SetRefundBet(small.RefundBet() - big.TotalBet())
	}

	big.AddGain(divide)
	small.AddGain(divide * -1)
	return 0
}

func (r *Room) GainBet(big *cache.Player, small *cache.Player) {
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
