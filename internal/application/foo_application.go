package application

import (
	"github.com/andreasds/go-boilerplate/internal/domain/entity"
	"github.com/andreasds/go-boilerplate/internal/domain/repository"
	"github.com/gofrs/uuid"
)

// FooApplication represent application of the foo.
// Expect implementation by the application layer
type FooApplication interface {
	ResolveFooByID(id uuid.UUID) (entity.Foo, error)
}

// FooApp provides foo application use case
type FooApp struct {
	FooRepository repository.FooRepository
}

// NewFooApplication provides `FooApp`
func NewFooApplication(
	fooRepository repository.FooRepository,
) *FooApp {
	f := new(FooApp)
	f.FooRepository = fooRepository

	return f
}

// FooApp implements the FooApplication interface
var _ FooApplication = &FooApp{}

// ResolveFooByID resolve foo by its identifier
func (f *FooApp) ResolveFooByID(
	id uuid.UUID,
) (foo entity.Foo, err error) {
	foo, err = f.FooRepository.ResolveFooByID(id)
	return
}
