package config

import (
	"log"
	"sync"

	lobbyconfig "github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
)

var (
	once         sync.Once
	cfg          *Config = &Config{}
	FeederConfig *feeder.DotEnv
)

func init() {
	FeederConfig = &feeder.DotEnv{
		Path: ".env",
	}
}

func GetConfig() *Config {
	once.Do(func() {
		initConfig()
	})

	return cfg
}

func initConfig() {
	lobbyconfig := lobbyconfig.New()
	lobbyconfig.AddFeeder(FeederConfig)
	lobbyconfig.AddStruct(cfg)

	if err := lobbyconfig.Feed(); err != nil {
		log.Fatalf("failed to parse config. err: %v\n", err)
	}
}
