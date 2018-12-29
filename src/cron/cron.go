package cron

import "github.com/robfig/cron"

//https://godoc.org/github.com/robfig/cron

var gCrons []*cron.Cron

func Init() {
	gCrons = append(gCrons, DaySign())
	gCrons = append(gCrons, DayQuest())
}

func Dispose() {
	for _, c := range gCrons {
		c.Stop()
	}
	gCrons = nil
}
