package persistence

import (
	"github.com/andreasds/go-boilerplate/internal/domain/entity"
	"github.com/andreasds/go-boilerplate/internal/domain/repository"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type FooRepo struct {
	db *sqlx.DB
}

func NewFooRepository(db *sqlx.DB) *FooRepo {
	r := new(FooRepo)
	r.db = db

	return r
}

// FooRepo implements the repository.FooRepository interface
var _ repository.FooRepository = &FooRepo{}

func (r *FooRepo) ResolveFooByID(id uuid.UUID) (foo entity.Foo, err error) {
	err = r.db.Get(&foo, "SELECT "+fooQueries.All+fooQueries.From+"WHERE id = UUID_TO_BIN(?)", id)
	if err != nil {
		log.Error().Err(err)
		return
	}

	return
}
