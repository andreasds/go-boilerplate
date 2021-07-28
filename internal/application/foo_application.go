package application

import (
	"github.com/andreasds/go-boilerplate/internal/domain/entity"
	"github.com/andreasds/go-boilerplate/internal/domain/repository"
	"github.com/gofrs/uuid"
)

type FooApplication interface {
	ResolveFooByID(id uuid.UUID) (entity.Foo, error)
}

type FooApp struct {
	FooRepository repository.FooRepository
}

func NewFooApplication(fooRepository repository.FooRepository) *FooApp {
	f := new(FooApp)
	f.FooRepository = fooRepository

	return f
}

// FooApp implements the FooApplication interface
var _ FooApplication = &FooApp{}

func (f *FooApp) ResolveFooByID(id uuid.UUID) (foo entity.Foo, err error) {
	return f.FooRepository.ResolveFooByID(id)
}
