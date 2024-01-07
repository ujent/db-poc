package config

import "skeleton/internal/settings"

type Config struct {
	EnvType settings.EnvType

	ServerPort string

	DbConnStr string
	Dialect   settings.DBType
}

func Load() (*Config, error) {
	return &Config{}, nil
}
