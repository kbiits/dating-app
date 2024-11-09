package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	cron_handler "github.com/kbiits/dealls-take-home-test/adapters/cron"
	cron_swipe "github.com/kbiits/dealls-take-home-test/adapters/cron/swipe"
	"github.com/kbiits/dealls-take-home-test/config"
	swipe_service "github.com/kbiits/dealls-take-home-test/domain/services/swipe"
	"github.com/kbiits/dealls-take-home-test/infrastructure/postgres"
	"github.com/kbiits/dealls-take-home-test/infrastructure/redis"
	pg_repo "github.com/kbiits/dealls-take-home-test/repositories/postgres"
	premium_package_pg_repo "github.com/kbiits/dealls-take-home-test/repositories/postgres/premium_package"
	profile_pg_repo "github.com/kbiits/dealls-take-home-test/repositories/postgres/profile"
	swipe_pg_repo "github.com/kbiits/dealls-take-home-test/repositories/postgres/swipe"
	pg_tx_repo "github.com/kbiits/dealls-take-home-test/repositories/postgres/tx"
	user_pg_repo "github.com/kbiits/dealls-take-home-test/repositories/postgres/user"
	redis_swipe_repo "github.com/kbiits/dealls-take-home-test/repositories/redis/swipe"
	swipe_usecase "github.com/kbiits/dealls-take-home-test/usecases/swipe"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

var (
	logger = log.With().Str("app", "cron").Logger()
)

func main() {

	handlers := bootstrapCron()
	cronEngine := cron.New(
		cron.WithLocation(time.Local),
	)

	cron_handler.RegisterCron(cronEngine, handlers)

	go func() {
		log.Info().Msg("starting cron")
		cronEngine.Start()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-c

	logger.Info().Msg("shutting down cron")
	cronEngine.Stop()
}

func bootstrapCron() cron_handler.CronHandlers {
	cfg := config.GetConfig()

	db, err := postgres.ConnectToPostgres(cfg.Database)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to postgres")
	}

	redis, err := redis.ConnectToRedis(cfg.RedisConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to redis")
	}

	baseRepo := pg_repo.NewBaseRepo(db)

	txRepo := pg_tx_repo.NewTxRepo(db)

	profileRepo := profile_pg_repo.NewProfileRepository(baseRepo)
	swipeRepo := swipe_pg_repo.NewSwipeRepository(baseRepo)
	swipeCacheRepo := redis_swipe_repo.NewRedisSwipeRepo(redis)

	userRepo := user_pg_repo.NewUserRepository(baseRepo)
	premiumPackageRepo := premium_package_pg_repo.NewPremiumPackageRepo(baseRepo)
	swipeService := swipe_service.NewSwipeService(userRepo, swipeRepo, premiumPackageRepo)

	swipeUsecase := swipe_usecase.NewSwipeUsecase(profileRepo, swipeRepo, swipeCacheRepo, txRepo, swipeService)

	return cron_handler.CronHandlers{
		SwipeCronHandler: cron_swipe.NewSwipeCronHandler(swipeUsecase),
	}
}
