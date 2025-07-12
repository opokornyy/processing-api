package config

import "github.com/rs/zerolog"

type Config struct {
	// Add config fields
}

func NewConfig() *Config {
	// TODO: make it configurable with envs
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	return &Config{}
}
