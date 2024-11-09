package config

import "fmt"

// singleton config

type Config struct {
	Logging        LoggingConfig
	Http           HttpConfig
	Database       DatabaseConfig
	JwtConfig      JwtConfig
	RedisConfig    RedisConfig
	InternalConfig InternalConfig
}

type HttpConfig struct {
	Address     string `env:"HTTP_ADDRESS"`
	ReadTimeout int    `env:"HTTP_READ_TIMEOUT"`
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

type JwtConfig struct {
	Secret             string `env:"JWT_SECRET"`
	ExpirationDuration int    `env:"JWT_EXPIRATION_DURATION"` // in seconds
}

type LoggingConfig struct {
	Level  string `env:"LOGGING_LEVEL"`
	Output string `env:"LOGGING_OUTPUT"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     int    `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
}

type InternalConfig struct {
	APIKey string `env:"API_KEY"`
}

func (c *Config) Setup() error {
	if c.Http.Address == "" {
		c.Http.Address = "0.0.0.0:8080"
	}
	if c.Http.ReadTimeout == 0 {
		c.Http.ReadTimeout = 30 // just wait 30 seconds for reading request, to prevent resource exhaustion
	}

	if c.Database.Host == "" {
		return fmt.Errorf("missing database host")
	}
	if c.Database.Port == 0 {
		return fmt.Errorf("missing database port")
	}
	if c.Database.Username == "" {
		return fmt.Errorf("missing database username")
	}
	if c.Database.Password == "" {
		return fmt.Errorf("missing database password")
	}
	if c.Database.Name == "" {
		return fmt.Errorf("missing database name")
	}

	if c.JwtConfig.Secret == "" {
		return fmt.Errorf("missing jwt secret")
	}

	if c.Logging.Level == "" {
		c.Logging.Level = "info"
	}

	if c.Logging.Output == "" {
		c.Logging.Output = "stdout"
	}

	if c.RedisConfig.Host == "" {
		return fmt.Errorf("missing redis host")
	}
	if c.RedisConfig.Port == 0 {
		return fmt.Errorf("missing redis port")
	}

	return nil
}
