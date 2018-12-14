package cost

import (
	"fmt"
	"github.com/trist725/myleaf/log"
	"github.com/trist725/myleaf/module"
	"mlgs/src/model"
	"mlgs/src/sd"
)

type CostItem [sd.E_CostItemField_Count]int64
type Costs []CostItem

func (i CostItem) Clone() (ni CostItem) {
	for j := 0; j < len(i); j++ {
		ni[j] = i[j]
	}
	return
}

func (i CostItem) TID() int64 {
	return i[sd.E_CostItemField_TID]
}

func (i CostItem) Num() int64 {
	return i[sd.E_CostItemField_Num]
}

func (c Costs) Clone() (nc Costs) {
	//fmt.Printf("Cost.Clone.c.p=[%p], Cost.Clone.nc.p=[%p]\n", c, nc)
	for _, i := range c {
		nc = append(nc, i.Clone())
	}
	return
}

func (c Costs) Each(fn func(i CostItem)) {
	for _, i := range c {
		fn(i)
	}
}

func (c Costs) IterateUntil(fn func(i CostItem) bool) {
	for _, i := range c {
		if fn(i) {
			return
		}
	}
}

//合并重复item
func (c Costs) Format() (nc Costs) {
	items := map[int64]int64{}
	for _, item := range c {
		if _, ok := items[item.TID()]; ok {
			items[item.TID()] += item.Num()
		} else {
			items[item.TID()] = item.Num()
		}
	}
	for tid, num := range items {
		nc = append(nc, CostItem{
			tid,
			num})
	}
	return
}

// 数量加倍
func (c Costs) Multiple(n int64) (nc Costs) {
	for _, item := range c {
		nc = append(nc, CostItem{item.TID(), item.Num() * n})
	}
	return
}

func (c Costs) Len() int {
	return len(c)
}

func (c Costs) Append(oc Costs) (nc Costs) {
	nc = append(c, oc...)
	return
}

func (c Costs) AppendItem(i CostItem) (nc Costs) {
	nc = append(c, i)
	return
}

func CanCost(user *model.User, cost Costs, n int64) (err error) {
	if cost.Len() == 0 {
		err = fmt.Errorf("empty cost")
		return
	}
	if n <= 0 {
		err = fmt.Errorf("cost %+v fail, n <= 0", cost)
		return
	}
	if n > 1 {
		cost = cost.Multiple(n)
	}

	cost = cost.Format()

	cost.Each(func(item CostItem) {
		ownNum := user.GetNum(item.TID())
		if ownNum < item.Num() {
			err = fmt.Errorf("can't cost %+v, ownNum[%d] < needNum[%d]", cost, ownNum, item.Num())
			return
		}
	})

	return
}

func Cost(user *model.User, cost Costs, n int64, notify bool, skeleton *module.Skeleton) (err error) {
	if cost.Len() == 0 {
		err = fmt.Errorf("[%d-%s] cost %+v fail, %v", user.ID, user.NickName, cost, err)
		return
	}
	if n <= 0 {
		err = fmt.Errorf("[%d-%s] cost %+v fail, %v", user.ID, user.NickName, cost, err)
		return
	}
	if n > 1 {
		cost = cost.Multiple(n)
	}

	cost = cost.Format()

	cost.Each(func(item CostItem) {
		_, _, err = user.Lost(item.TID(), item.Num(), notify, skeleton)
		if err != nil {
			log.Error("[%d-%s] cost %+v fail, %v", user.ID, user.NickName, cost, err)
			return
		}
	})

	return
}
