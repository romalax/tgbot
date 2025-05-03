package config

import (
	"bot1/internal/adapters/api"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Adapters Adapters
}

type Adapters struct {
	ApiAdapter api.Config
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
