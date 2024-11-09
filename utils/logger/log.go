package logger

import (
	"io"
	"os"
	"sync"

	"github.com/kbiits/dealls-take-home-test/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	logger    zerolog.Logger
	once      sync.Once
	logOutput io.Writer
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	logOutput = io.Discard
}

func GetLogger() zerolog.Logger {
	once.Do(func() {
		cfg := config.GetConfig()
		output := cfg.Logging.Output
		stream := os.Stdout
		if output != "stdout" {
			stream, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err == nil {
				logOutput = stream
			}

		} else {
			logOutput = stream
		}

		logger = zerolog.New(logOutput)
		log.Logger = logger
	})

	return logger
}
