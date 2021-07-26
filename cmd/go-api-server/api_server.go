package main

import (
	"github.com/andreasds/go-boilerplate/configs"
	"github.com/andreasds/go-boilerplate/internal/shared/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	// Initializes logger
	logger.InitLogger()

	// Initializes configuration
	config := configs.GetInstance()

	// Sets log level
	logger.SetLogLevel(config.Server.LogLevel)

	log.Info().Msg("API Server application is starting.")
}
