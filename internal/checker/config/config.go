package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
)

type Config struct {
	Addr            string        `env:"SERVER_ADDR"`
	TelegramTimeout time.Duration `env:"TELEGRAM_TIMEOUT"`
	DatabaseURL     string        `env:"DATABASE_URL"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error parse to cfg: %w", err)
	}

	return &cfg, nil
}
