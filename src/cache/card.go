package cache

import (
	"github.com/trist725/myleaf/log"
	"mlgs/src/msg"
	"reflect"
	"sort"
)

//组牌数
const gGroupCardCount int = 5

func GroupCardCount() int {
	return gGroupCardCount
}

type Card struct {
	//花色,1-黑桃(Spade),2-红桃(Heart),3-方块(Diamond),4-梅花(Club)
	Color uint8
	//牌值,2-14
	Num uint8
}

func (c *Card) ToMsg(nc *msg.Card) *msg.Card {
	nc.Num = int32(c.Num)
	nc.Color = int32(c.Color)
	return nc
}

func (c *Card) Equal(card Card) bool {
	if c.Num == card.Num && c.Color == card.Color {
		return true
	}
	return false
}

func (c *Card) NumEqual(card Card) bool {
	if c.Num == card.Num {
		return true
	}
	return false
}

func (c *Card) ColorEqual(card Card) bool {
	if c.Color == card.Color {
		return true
	}
	return false
}

type CardSlice []Card

func (cs CardSlice) Len() int { // 重写 Len() 方法
	return len(cs)
}
func (cs CardSlice) Swap(i, j int) { // 重写 Swap() 方法
	cs[i], cs[j] = cs[j], cs[i]
}
func (cs CardSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return cs[j].Num < cs[i].Num
}

//最大牌型
const gMaxGroupLevel int32 = 10

func (cs CardSlice) CalLevel() int32 {
	if cs.Len() == 0 {
		return 0
	}
	for level := gMaxGroupLevel; level >= 2; level-- {
		switch level {
		case 10:
			if cs.IsRoyalFlush() {
				return level
			}
		case 9:
			if cs.IsStraightFlush() {
				return level
			}
		case 8:
			if cs.IsFourOfAKind() {
				return level
			}
		case 7:
			if cs.IsFullHouse() {
				return level
			}
		case 6:
			if cs.IsFlush() {
				return level
			}
		case 5:
			if cs.IsStraight() {
				return level
			}
		case 4:
			if cs.IsTriOfAKind() {
				return level
			}
		case 3:
			if cs.IsTwoPair() {
				return level
			}
		case 2:
			if cs.IsOnePair() {
				return level
			}
		}
	}
	//以上都不是,只能是高牌
	return 1
}

//需要先降序排序
//需要从高到低level顺序使用
//使用compare之前要先用is,因compare中不做边界检查
//皇家同花顺-10
func (cs CardSlice) IsRoyalFlush() bool {
	if len(cs) != gGroupCardCount {
		return false
	}

	if cs.IsStraightFlush() {
		if cs[0].Num == 14 {
			return true
		}
	}
	return false
}

//返回大的
func (cs CardSlice) RoyalFlushCompare() CardSlice {
	return nil
}

//同花顺-9
func (cs CardSlice) IsStraightFlush() bool {
	if len(cs) != gGroupCardCount {
		return false
	}
	if cs.IsStraight() && cs.IsFlush() {
		return true
	}
	return false
}

func (cs CardSlice) StraightFlushCompare(cs2 CardSlice) CardSlice {
	//就是比顺子
	return cs.StraightCompare(cs2)
}

//四条-8
func (cs CardSlice) IsFourOfAKind() bool {
	if len(cs) < 4 {
		return false
	}

	countMap := make(map[uint8]int)
	for _, v := range cs {
		countMap[v.Num]++
	}
	for _, v := range countMap {
		if v == 4 {
			return true
		}
	}

	return false
}

func (cs CardSlice) FourOfAKindCompare(cs2 CardSlice) CardSlice {
	csNum := cs.GetNOfAKindNum(4)
	cs2Num := cs2.GetNOfAKindNum(4)

	if csNum > cs2Num {
		return cs
	} else if csNum < cs2Num {
		return cs2
	}
	return nil
}

//葫芦-7
func (cs CardSlice) IsFullHouse() bool {
	if len(cs) != gGroupCardCount {
		return false
	}

	countMap := make(map[uint8]int)
	for _, v := range cs {
		countMap[v.Num]++
	}
	var tri, two int
	for _, v := range countMap {
		if v == 3 {
			tri++
		} else if v == 2 {
			two++
		}
	}
	if tri == 1 && two == 1 {
		return true
	}

	return false
}

func (cs CardSlice) FullHouseCompare(cs2 CardSlice) CardSlice {
	csTri := cs.GetNOfAKindNum(3)
	cs2Tri := cs2.GetNOfAKindNum(3)
	csTwo := cs.GetNOfAKindNum(2)
	cs2Two := cs2.GetNOfAKindNum(2)

	if csTri == cs2Tri {
		if csTwo > cs2Two {
			return cs
		} else if csTwo < cs2Two {
			return cs2
		}
		return nil
	} else {
		if csTri > cs2Tri {
			return cs
		} else if csTri < cs2Tri {
			return cs2
		}
	}
	return nil
}

//同花-6
func (cs CardSlice) IsFlush() bool {
	if len(cs) != gGroupCardCount {
		return false
	}

	countMap := make(map[uint8]int)
	for _, v := range cs {
		countMap[v.Color]++
	}
	for _, v := range countMap {
		if v == gGroupCardCount {
			return true
		}
	}

	return false
}

func (cs CardSlice) FlushCompare(cs2 CardSlice) CardSlice {
	for i := 0; i < gGroupCardCount; i++ {
		if cs[i].Num > cs2[i].Num {
			return cs
		} else if cs[i].Num < cs2[i].Num {
			return cs2
		} //else 一样大continue
	}
	//每个都一样大
	return nil
}

//顺子-5
func (cs CardSlice) IsStraight() bool {
	if len(cs) != gGroupCardCount {
		return false
	}

	//有A特殊处理
	if cs[0].Num == 14 {
		if cs[1].Num == 5 || cs[1].Num == 13 {
			for i := 1; i < gGroupCardCount-1; i++ {
				if cs[i].Num-cs[i+1].Num != 1 {
					return false
				}
			}
			return true
		} else {
			return false
		}
	} else {
		for i := 0; i < gGroupCardCount-1; i++ {
			if cs[i].Num-cs[i+1].Num != 1 {
				return false
			}
		}
		return true
	}

	return false
}

func (cs CardSlice) StraightCompare(cs2 CardSlice) CardSlice {
	//A特例处理
	if cs[0].Num == 14 || cs2[0].Num == 14 {
		//第二张是5,是最小
		if cs[1].Num == 5 {
			if cs2[1].Num == 5 {
				return nil
			} else {
				return cs2
			}
		} else {
			if cs2[1].Num == 5 {
				return cs
			}
		}
	}

	//特例之外比第二大那张
	if cs[1].Num > cs2[1].Num {
		return cs
	} else if cs[1].Num < cs2[1].Num {
		return cs2
	}
	return nil
}

//三条-4
func (cs CardSlice) IsTriOfAKind() bool {
	if len(cs) < 3 {
		return false
	}

	countMap := make(map[uint8]int)
	for _, v := range cs {
		countMap[v.Num]++
	}
	for _, v := range countMap {
		if v == 3 {
			return true
		}
	}

	return false
}

func (cs CardSlice) TriOfAKindCompare(cs2 CardSlice) CardSlice {
	csTri := cs.GetNOfAKindNum(3)
	cs2Tri := cs2.GetNOfAKindNum(3)

	if csTri == cs2Tri {
		return cs.HighCardCompare(cs2)
	} else if csTri > cs2Tri {
		return cs
	} else {
		return cs2
	}
}

//两对-3
func (cs CardSlice) IsTwoPair() bool {
	if len(cs) < 4 {
		return false
	}

	countMap := make(map[uint8]int)
	for _, v := range cs {
		countMap[v.Num]++
	}
	var count int
	for _, v := range countMap {
		if v == 2 {
			count++
		}
	}
	if count == 2 {
		return true
	}

	return false
}

func (cs CardSlice) TwoPairCompare(cs2 CardSlice) CardSlice {
	var csBig, cs2Big uint8
	var csSmall, cs2Small uint8
	var csOne, cs2One uint8
	countMap := make(map[uint8]int)
	for _, v := range cs {
		countMap[v.Num]++
	}
	for k, v := range countMap {
		if v == 2 {
			if csBig == 0 {
				csBig = k
				csSmall = k
			} else {
				if k > csBig {
					csBig = k
				} else {
					csSmall = k
				}
			}
		} else if v == 1 {
			csOne = k
		}
	}
	countMap = make(map[uint8]int)
	for _, v := range cs2 {
		countMap[v.Num]++
	}
	for k, v := range countMap {
		if v == 2 {
			if cs2Big == 0 {
				cs2Big = k
				cs2Small = k
			} else {
				if k > cs2Big {
					cs2Big = k
				} else {
					cs2Small = k
				}
			}
		} else if v == 1 {
			cs2One = k
		}
	}

	if csBig == cs2Big {
		if csSmall == cs2Small {
			if csOne > cs2One {
				return cs
			} else if csOne < cs2One {
				return cs2
			} else {
				return nil
			}
		} else if csSmall > cs2Small {
			return cs
		} else {
			return cs2
		}
	} else if csBig > cs2Big {
		return cs
	} else { //if csBig < cs2Big
		return cs2
	}
}

//一对-2
func (cs CardSlice) IsOnePair() bool {
	if len(cs) < 2 {
		return false
	}

	countMap := make(map[uint8]int)
	for _, v := range cs {
		countMap[v.Num]++
	}
	for _, v := range countMap {
		if v == 2 {
			return true
		}
	}

	return false
}

func (cs CardSlice) OnePairCompare(cs2 CardSlice) CardSlice {
	csTwo := cs.GetNOfAKindNum(2)
	cs2Two := cs2.GetNOfAKindNum(2)

	if csTwo == cs2Two {
		var hc, hc2 CardSlice
		for _, v := range cs {
			if v.Num != csTwo {
				hc = append(hc, v)
			}
		}
		for _, v := range cs2 {
			if v.Num != cs2Two {
				hc2 = append(hc2, v)
			}
		}
		sort.Sort(hc)
		sort.Sort(hc2)
		if nil == hc.HighCardCompare(hc2) {
			return nil
		} else if reflect.DeepEqual(hc, hc.HighCardCompare(hc2)) {
			return cs
		} else {
			return cs2
		}
	} else if csTwo > cs2Two {
		return cs
	}
	return cs2
}

//高牌-1
//func (cs CardSlice) IsHighCard() bool {
//	return true
//}

func (cs CardSlice) HighCardCompare(cs2 CardSlice) CardSlice {
	for i := 0; i < cs.Len(); i++ {
		if cs[i].Num > cs2[i].Num {
			return cs
		} else if cs[i].Num < cs2[i].Num {
			return cs2
		} //else == continue
	}
	return nil
}

func (cs CardSlice) GetNOfAKindNum(n int) uint8 {
	var csN uint8
	countMap := make(map[uint8]int)
	for _, v := range cs {
		countMap[v.Num]++
	}
	for k, v := range countMap {
		if v == n {
			csN = k
		}
	}
	if csN == 0 {
		log.Error("GetNOfAKindNum failed")
	}
	return csN
}
