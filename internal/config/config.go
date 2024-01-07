package config

import "skeleton/internal/settings"

type Config struct {
	ServerPort string
	DbConnStr  string
	EnvType    settings.EnvType
}

func Load() (*Config, error) {
	return &Config{}, nil
}
