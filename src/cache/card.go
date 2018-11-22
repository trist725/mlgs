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
