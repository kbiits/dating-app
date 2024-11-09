package cron_handler

import "github.com/robfig/cron/v3"

type SwipeCronHandler interface {
	ClearYesterdayBlooms()
}

type CronHandlers struct {
	SwipeCronHandler SwipeCronHandler
}

func RegisterCron(
	cronEngine *cron.Cron,
	cronHandlers CronHandlers,
) {
	cronEngine.AddFunc("0 0 * * *", cronHandlers.SwipeCronHandler.ClearYesterdayBlooms)
}
