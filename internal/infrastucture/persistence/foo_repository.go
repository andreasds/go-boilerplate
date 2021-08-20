package persistence

import (
	"github.com/andreasds/go-boilerplate/internal/domain/entity"
	"github.com/andreasds/go-boilerplate/internal/domain/repository"
	"github.com/andreasds/go-boilerplate/internal/infrastucture/database"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

// FooRepo provides foo repository use case
type FooRepo struct {
	db *database.MySQLConnection
}

// NewFooRepository provides `FooRepo`
func NewFooRepository(
	db *database.MySQLConnection,
) *FooRepo {
	r := new(FooRepo)
	r.db = db

	return r
}

// FooRepo implements the repository.FooRepository interface
var _ repository.FooRepository = &FooRepo{}

// ResolveFooByID resolve foo by its identifier
func (r *FooRepo) ResolveFooByID(
	id uuid.UUID,
) (foo entity.Foo, err error) {
	err = r.db.Read.Get(
		&foo,
		"SELECT"+fooQueries.AllField+fooQueries.From+" WHERE id = UUID_TO_BIN(?)",
		id,
	)
	if err != nil {
		log.Error().Err(err)
		return
	}

	return
}
