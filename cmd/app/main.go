package main

import (
	"processing-api/config"
	"processing-api/internal/app"

	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.NewConfig()
	log.Info().Msgf("Config: %v", cfg)

	app.Run(*cfg)
}
