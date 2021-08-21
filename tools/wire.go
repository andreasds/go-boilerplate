// +build wireinject

package tools

import (
	"github.com/andreasds/go-boilerplate/configs"
	"github.com/andreasds/go-boilerplate/internal/application"
	"github.com/andreasds/go-boilerplate/internal/domain/repository"
	"github.com/andreasds/go-boilerplate/internal/infrastucture/database"
	"github.com/andreasds/go-boilerplate/internal/infrastucture/persistence"
	"github.com/google/wire"
)

// configurations wiring
var configurations = wire.NewSet(
	configs.GetInstance,
)

// databases wiring
var databases = wire.NewSet(
	database.NewMySQLConnection,
)

// fooDomain wiring domain Foo
var fooDomain = wire.NewSet(
	// Foo Application interface and implementation
	application.NewFooApplication,
	wire.Bind(new(application.FooApplication), new(*application.FooApp)),
	// Foo Repository interface and implementation
	persistence.NewFooRepository,
	wire.Bind(new(repository.FooRepository), new(*persistence.FooRepo)),
)

// domains wiring
var domains = wire.NewSet(
	fooDomain,
)

// InitServices wiring all services
func InitServices() *application.FooApp {
	wire.Build(
		// configurations
		configurations,
		// databases
		databases,
		// domains
		domains,
	)

	return &application.FooApp{}
}
