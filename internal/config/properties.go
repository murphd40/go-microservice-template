package config

import (
	"github.com/caarlos0/env/v6"
	"time"
)

type ServerProperties struct {
	Port         int           `env:"SERVER_PORT"`
	ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT"`
	WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT"`
}

func NewServerProperties() ServerProperties {
	return ServerProperties{
		Port:         9080,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}

func PopulateConfig[T any](environment map[string]string, target *T) {
	env.Parse(target, env.Options{
		Environment: environment,
	})
}
