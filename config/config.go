package config

import "github.com/rs/zerolog"

type Config struct {
	Port string
}

func NewConfig() *Config {
	// TODO: load configuration form .env
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	return &Config{
		Port: "8080",
	}
}
