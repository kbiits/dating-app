package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/kbiits/dealls-take-home-test/config"
	goredis "github.com/redis/go-redis/v9"
)

func ConnectToRedis(
	cfg config.RedisConfig,
) (*goredis.Client, error) {
	client := goredis.NewClient(&goredis.Options{
		Addr:     cfg.Host + ":" + fmt.Sprintf("%d", cfg.Port),
		Password: cfg.Password,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	pingRes := client.Ping(ctx)
	if pingRes.Err() != nil {
		return nil, pingRes.Err()
	}

	return client, nil
}
