package app

import (
	"processing-api/config"

	"github.com/rs/zerolog/log"
)

func Run(cfg config.Config) {
	log.Info().Msg("Starting the application")
}
