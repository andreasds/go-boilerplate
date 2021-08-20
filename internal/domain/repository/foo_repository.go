package repository

import (
	"github.com/andreasds/go-boilerplate/internal/domain/entity"
	"github.com/gofrs/uuid"
)

// FooRepository represent repository of the foo.
// Expect implementation by the infrastructure layer
type FooRepository interface {
	ResolveFooByID(id uuid.UUID) (entity.Foo, error)
}
