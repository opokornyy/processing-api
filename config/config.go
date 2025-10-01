package config

import (
	"os"

	"github.com/rs/zerolog"
)

type Config struct {
	Port          string
	WeatherAPIKey string
}

func NewConfig() *Config {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	weatherAPIKey := os.Getenv("WEATHER_API_KEY")

	return &Config{
		Port:          port,
		WeatherAPIKey: weatherAPIKey,
	}
}
