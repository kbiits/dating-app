package main

import (
	cron_swipe "github.com/kbiits/dealls-take-home-test/adapters/cron/swipe"
	"github.com/robfig/cron/v3"
)

type CronHandlers struct {
	SwipeCronHandler cron_swipe.SwipeCronHandler
}

func RegisterCron(
	cronEngine *cron.Cron,
	cronHandlers CronHandlers,
) {
	cronEngine.AddFunc("0 0 * * *", cronHandlers.SwipeCronHandler.ClearYesterdayBlooms)
}
