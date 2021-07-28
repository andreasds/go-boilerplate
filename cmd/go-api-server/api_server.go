package main

import (
	"github.com/andreasds/go-boilerplate/configs"
	"github.com/andreasds/go-boilerplate/internal/application"
	"github.com/andreasds/go-boilerplate/internal/infrastucture/persistence"
	"github.com/andreasds/go-boilerplate/internal/interfaces"
	"github.com/andreasds/go-boilerplate/internal/shared/logger"
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

	db := "Hello World!"
	fooRepository := persistence.NewFooRepository(&db)
	fooApplication := application.NewFooApplication(fooRepository)
	foo := interfaces.NewFooHandler(fooApplication)
	id, _ := uuid.NewV4()
	foo.ResolveFooByID(id)

	log.Info().Msg("API Server application is starting.")
}
