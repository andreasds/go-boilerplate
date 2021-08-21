package main

import (
	"github.com/andreasds/go-boilerplate/configs"
	"github.com/andreasds/go-boilerplate/internal/interfaces"
	"github.com/andreasds/go-boilerplate/internal/shared/logger"
	"github.com/andreasds/go-boilerplate/tools"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

func main() {
	// Initializes logger
	logger.InitLogger()

	// Initializes configuration
	config := configs.GetInstance()

	// Sets log level
	logger.SetLogLevel(config.Server.LogLevel)

	// Initializes database
	fooApplication := tools.InitServices()
	foo := interfaces.NewFooHandler(fooApplication)
	foo.ResolveFooByID(uuid.FromStringOrNil("f5610c25-91ed-4c18-9beb-4272a2e7ca90"))

	log.Info().Msg("API Server application is starting.")
}
