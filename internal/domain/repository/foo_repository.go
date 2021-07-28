package repository

import (
	"github.com/andreasds/go-boilerplate/internal/domain/entity"
	"github.com/gofrs/uuid"
)

type FooRepository interface {
	ResolveFooByID(id uuid.UUID) (entity.Foo, error)
}
