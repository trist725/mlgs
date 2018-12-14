// Code generated by protoc-gen-mgo-go. DO NOT EDIT IT!!!
// source: card.proto

/*
It has these top-level messages:
	Card
	BestCombo
*/

package model

import "fmt"
import "encoding/json"
import "sync"
import "github.com/trist725/myleaf/db/mongodb"
import "gopkg.in/mgo.v2"

var _ = fmt.Sprintf
var _ = json.Marshal
var _ *sync.Pool
var _ *mongodb.DialContext
var _ *mgo.DBRef

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// collection [Card] begin

///牌
type Card struct {
	///花色,1-黑桃(Spade),2-红桃(Heart),3-方块(Diamond),4-梅花(Club)
	Color int32 `bson:"Color"`
	///牌值,2-14
	Num int32 `bson:"Num"`
}

func New_Card() *Card {
	m := &Card{}
	return m
}

func (m Card) String() string {
	ba, _ := json.Marshal(m)
	return fmt.Sprintf("{\"Card\":%s}", string(ba))
}

func (m *Card) Reset() {
	m.Color = 0
	m.Num = 0

}

func (m Card) Clone() *Card {
	n, ok := g_Card_Pool.Get().(*Card)
	if !ok || n == nil {
		n = &Card{}
	}

	n.Color = m.Color
	n.Num = m.Num

	return n
}

func Clone_Card_Slice(dst []*Card, src []*Card) []*Card {
	for _, i := range dst {
		Put_Card(i)
	}
	dst = []*Card{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

var g_Card_Pool = sync.Pool{}

func Get_Card() *Card {
	m, ok := g_Card_Pool.Get().(*Card)
	if !ok {
		m = New_Card()
	} else {
		if m == nil {
			m = New_Card()
		} else {
			m.Reset()
		}
	}
	return m
}

func Put_Card(i interface{}) {
	if m, ok := i.(*Card); ok && m != nil {
		g_Card_Pool.Put(i)
	}
}

// collection [Card] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// collection [BestCombo] begin

///最大牌型
type BestCombo struct {
	Cards []*Card `bson:"Cards"`
	///10-皇家同花顺,9-同花顺,8-四条(金刚),7-葫芦,6-通话 /5-顺子,4-三条,3-两队,2-对子,1-高牌
	Type int32 `bson:"Type"`
}

func New_BestCombo() *BestCombo {
	m := &BestCombo{
		Cards: []*Card{},
	}
	return m
}

func (m BestCombo) String() string {
	ba, _ := json.Marshal(m)
	return fmt.Sprintf("{\"BestCombo\":%s}", string(ba))
}

func (m *BestCombo) Reset() {

	for _, i := range m.Cards {
		Put_Card(i)
	}
	m.Cards = []*Card{}
	m.Type = 0

}

func (m BestCombo) Clone() *BestCombo {
	n, ok := g_BestCombo_Pool.Get().(*BestCombo)
	if !ok || n == nil {
		n = &BestCombo{}
	}

	if len(m.Cards) > 0 {
		for _, i := range m.Cards {
			if i != nil {
				n.Cards = append(n.Cards, i.Clone())
			} else {
				n.Cards = append(n.Cards, nil)
			}
		}
	} else {
		n.Cards = []*Card{}
	}

	n.Type = m.Type

	return n
}

func Clone_BestCombo_Slice(dst []*BestCombo, src []*BestCombo) []*BestCombo {
	for _, i := range dst {
		Put_BestCombo(i)
	}
	dst = []*BestCombo{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

var g_BestCombo_Pool = sync.Pool{}

func Get_BestCombo() *BestCombo {
	m, ok := g_BestCombo_Pool.Get().(*BestCombo)
	if !ok {
		m = New_BestCombo()
	} else {
		if m == nil {
			m = New_BestCombo()
		} else {
			m.Reset()
		}
	}
	return m
}

func Put_BestCombo(i interface{}) {
	if m, ok := i.(*BestCombo); ok && m != nil {
		g_BestCombo_Pool.Put(i)
	}
}

// collection [BestCombo] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
