package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App       App
		Log       Log
		GoogleMap GoogleMap
	}

	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
		Env     string `env:"ENV,required"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	GoogleMap struct {
		ApiKey string `env:"GOOGLE_MAPS_API_KEY,required"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
