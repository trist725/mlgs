package cache

type Card struct {
	//花色,1-黑桃(Spade),2-红桃(Heart),3-方块(Diamond),4-梅花(Club)
	Color uint8
	//牌值,2-14
	Num uint8
}

func (c *Card) Equal(card Card) bool {
	if c.Num == card.Num && c.Color == card.Color {
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

//皇家同花顺
func (cs CardSlice) IsRoyalFlush() (bool, CardSlice) {
	return false, cs
}

//同花顺
func (cs CardSlice) IsStraightFlush() (bool, CardSlice) {
	return false, cs
}

//四条
func (cs CardSlice) IsFourOfAKind() (bool, CardSlice) {
	return false, cs
}

//葫芦
func (cs CardSlice) IsFullHouse() (bool, CardSlice) {
	return false, cs
}

//同花
func (cs CardSlice) IsFlush() (bool, CardSlice) {
	return false, cs
}

//顺子
func (cs CardSlice) IsStraight() (bool, CardSlice) {
	return false, cs
}

//三条
func (cs CardSlice) IsTriOfAKind() (bool, CardSlice) {
	return false, cs
}

//两对
func (cs CardSlice) IsTwoPair() (bool, CardSlice) {
	return false, cs
}

//一对
func (cs CardSlice) IsOnePair() (bool, CardSlice) {
	return false, cs
}

//高牌
func (cs CardSlice) IsHighCard() (bool, CardSlice) {
	return false, cs
}
