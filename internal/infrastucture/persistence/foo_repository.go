package persistence

import (
	"github.com/andreasds/go-boilerplate/internal/domain/entity"
	"github.com/andreasds/go-boilerplate/internal/domain/repository"
	"github.com/andreasds/go-boilerplate/internal/infrastucture/database"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

type FooRepo struct {
	db *database.MySQLConnection
}

func NewFooRepository(db *database.MySQLConnection) *FooRepo {
	r := new(FooRepo)
	r.db = db

	return r
}

// FooRepo implements the repository.FooRepository interface
var _ repository.FooRepository = &FooRepo{}

func (r *FooRepo) ResolveFooByID(id uuid.UUID) (foo entity.Foo, err error) {
	err = r.db.Read.Get(
		&foo,
		"SELECT "+fooQueries.All+fooQueries.From+"WHERE id = UUID_TO_BIN(?)",
		id,
	)
	if err != nil {
		log.Error().Err(err)
		return
	}

	return
}
