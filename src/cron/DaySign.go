package cron

import (
	"fmt"
	"github.com/trist725/myleaf/timer"
)

func init() {
	fmt.Println("My name is Leaf")
}

func DaySign() {
	d := timer.NewDispatcher(10)

	// cron expr
	cronExpr, err := timer.NewCronExpr("2 * * * * *")
	if err != nil {
		return
	}

	// cron
	var c *timer.Cron
	c = d.CronFunc(cronExpr, func() {
		fmt.Println("My name is Leaf")
		c.Stop()
	})

	// dispatch
	(<-d.ChanTimer).Cb()

	// Output:
	// My name is Leaf
}
