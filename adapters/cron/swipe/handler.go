package cron_swipe

import (
	"context"
	"time"

	cron_handler "github.com/kbiits/dealls-take-home-test/adapters/cron"
	swipe_usecase "github.com/kbiits/dealls-take-home-test/usecases/swipe"
	"github.com/rs/zerolog/log"
)

var (
	logger = log.With().Str("module", "cron_swipe_handler").Caller().Logger()
)

type swipCronHandler struct {
	swipeUsecase swipe_usecase.SwipeUsecase
}

func NewSwipeCronHandler(swipeUsecase swipe_usecase.SwipeUsecase) cron_handler.SwipeCronHandler {
	return &swipCronHandler{
		swipeUsecase: swipeUsecase,
	}
}

func (s *swipCronHandler) ClearYesterdayBlooms() {
	const timeoutForClearYesterdayBlooms = 30 * time.Second
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	ctx, cancel := context.WithTimeout(context.Background(), timeoutForClearYesterdayBlooms)
	defer cancel()

	if err := s.swipeUsecase.ClearBloomsByDate(ctx, yesterday); err != nil {
		logger.Error().Err(err).Msg("failed to clear yesterday blooms")
		return
	}

	logger.Info().Msg("yesterday blooms cleared")
}
