package main

import (
	"github.com/andreasds/go-boilerplate/internal/shared/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	// Initialize logger
	logger.InitLogger()

	log.Info().Msg("API Server application is starting.")
}
