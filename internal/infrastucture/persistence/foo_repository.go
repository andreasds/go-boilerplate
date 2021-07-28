package persistence

import (
	"github.com/andreasds/go-boilerplate/internal/domain/entity"
	"github.com/andreasds/go-boilerplate/internal/domain/repository"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

type FooRepo struct {
	db *string
}

func NewFooRepository(db *string) *FooRepo {
	r := new(FooRepo)
	r.db = db

	return r
}

// FooRepo implements the repository.FooRepository interface
var _ repository.FooRepository = &FooRepo{}

func (r *FooRepo) ResolveFooByID(id uuid.UUID) (foo entity.Foo, err error) {
	log.Info().Msg("Resolve Foo by ID in " + *r.db)

	foo.ID = id
	foo.Name = "Foo"
	foo.Description = "Foo Description"

	return
}
