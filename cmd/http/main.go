package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	http_controllers "github.com/kbiits/dealls-take-home-test/adapters/http"
	http_auth "github.com/kbiits/dealls-take-home-test/adapters/http/auth"
	http_profile "github.com/kbiits/dealls-take-home-test/adapters/http/profile"
	http_swipe "github.com/kbiits/dealls-take-home-test/adapters/http/swipe"
	http_transaction "github.com/kbiits/dealls-take-home-test/adapters/http/transaction"
	"github.com/kbiits/dealls-take-home-test/config"
	swipe_service "github.com/kbiits/dealls-take-home-test/domain/services/swipe"
	"github.com/kbiits/dealls-take-home-test/infrastructure/postgres"
	"github.com/kbiits/dealls-take-home-test/infrastructure/redis"
	postgres_repo "github.com/kbiits/dealls-take-home-test/repositories/postgres"
	premium_package_pg_repo "github.com/kbiits/dealls-take-home-test/repositories/postgres/premium_package"
	profile_pg_repo "github.com/kbiits/dealls-take-home-test/repositories/postgres/profile"
	swipe_pg_repo "github.com/kbiits/dealls-take-home-test/repositories/postgres/swipe"
	pg_tx_repo "github.com/kbiits/dealls-take-home-test/repositories/postgres/tx"
	user_pg_repo "github.com/kbiits/dealls-take-home-test/repositories/postgres/user"
	redis_swipe_repo "github.com/kbiits/dealls-take-home-test/repositories/redis/swipe"
	auth_usecase "github.com/kbiits/dealls-take-home-test/usecases/auth"
	profile_usecase "github.com/kbiits/dealls-take-home-test/usecases/profile"
	swipe_usecase "github.com/kbiits/dealls-take-home-test/usecases/swipe"
	transaction_usecase "github.com/kbiits/dealls-take-home-test/usecases/transaction"
	jwt_util "github.com/kbiits/dealls-take-home-test/utils/jwt"
)

func main() {

	var configPath string
	flag.StringVar(&configPath, "config", "./.env", "specify path to dotenv file (default: .env)")
	flag.Parse()

	config.SetConfigFilePath(configPath)
	cfg := config.GetConfig()

	controllers := bootstrapApp(cfg)

	gin := gin.Default()
	http_controllers.RegisterRoutes(gin, cfg, controllers)

	srv := http.Server{
		Addr:        cfg.Http.Address,
		Handler:     gin,
		ReadTimeout: time.Duration(cfg.Http.ReadTimeout) * time.Second,
	}

	// graceful shutdown handling
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("server started at %s\n", cfg.Http.Address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server. err: %v\n", err)
		}
	}()

	<-exitChan

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown server. err: %v\n", err)
	}

	log.Println("server shutdown gracefully")
}

func bootstrapApp(cfg *config.Config) http_controllers.Controllers {
	jwtUtil := jwt_util.NewJwtUtil(cfg.JwtConfig)

	db, err := postgres.ConnectToPostgres(cfg.Database)
	if err != nil {
		log.Fatalf("failed to connect to postgres. err: %v\n", err)
	}

	redis, err := redis.ConnectToRedis(cfg.RedisConfig)
	if err != nil {
		log.Fatalf("failed to connect to redis. err: %v\n", err)
	}

	swipeCacheRepo := redis_swipe_repo.NewRedisSwipeRepo(redis)

	baseRepo := postgres_repo.NewBaseRepo(db)
	txRepo := pg_tx_repo.NewTxRepo(db)

	userRepo := user_pg_repo.NewUserRepository(baseRepo)
	profileRepo := profile_pg_repo.NewProfileRepository(baseRepo)
	swipeRepo := swipe_pg_repo.NewSwipeRepository(baseRepo)
	premiumPackageRepo := premium_package_pg_repo.NewPremiumPackageRepo(baseRepo)

	authUsecase := auth_usecase.NewAuthUsecase(txRepo, userRepo, profileRepo, jwtUtil)
	authController := http_auth.NewAuthController(authUsecase)

	profileUsecase := profile_usecase.NewProfileUsecase(profileRepo)
	profileController := http_profile.NewProfileController(profileUsecase)

	swipeService := swipe_service.NewSwipeService(userRepo, swipeRepo, premiumPackageRepo)
	swipeUsecase := swipe_usecase.NewSwipeUsecase(profileRepo, swipeRepo, swipeCacheRepo, txRepo, swipeService)
	swipeController := http_swipe.NewSwipeController(swipeUsecase)

	return http_controllers.Controllers{
		AuthController:    authController,
		ProfileController: profileController,
		SwipeController:   swipeController,
		TransactionController: http_transaction.NewTransactionController(
			transaction_usecase.NewAuthUsecase(userRepo, premiumPackageRepo),
		),
		JwtUtil: jwtUtil,
	}
}
