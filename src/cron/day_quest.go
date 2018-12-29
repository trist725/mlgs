package cron

import (
	"github.com/robfig/cron"
	"github.com/trist725/myleaf/log"
	"gopkg.in/mgo.v2/bson"
	"mlgs/src/model"
)

func DayQuest() *cron.Cron {
	dbSession := model.GetSession()
	defer model.PutSession(dbSession)
	c := cron.New()

	err := c.AddFunc("@midnight", func() {
		some, err := model.FindSome_User(dbSession, bson.M{"AllocDayQuest": true})
		if err != nil {
			log.Error("cron DayQuest: %s", err.Error())
			return
		}
		for _, one := range some {
			user := one
			user.AllocDayQuest = false
			if err := user.UpdateByID(dbSession); err != nil {
				log.Error("cron DayQuest: UpdateByID error:[%s]", err)
			}
		}
	})

	if err != nil {
		log.Error("cron DayQuest AddFunc..%s", err.Error())
	}

	c.Start()

	return c
}
