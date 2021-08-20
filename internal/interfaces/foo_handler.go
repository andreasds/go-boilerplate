package interfaces

import (
	"fmt"

	"github.com/andreasds/go-boilerplate/internal/application"
	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"
)

// FooHandler provides foo handler use case
type FooHandler struct {
	FooApplication application.FooApplication
}

// NewFooHandler provides `FooHandler`
func NewFooHandler(fooApplication application.FooApplication) *FooHandler {
	h := new(FooHandler)
	h.FooApplication = fooApplication

	return h
}

// ResolveFooByID resolve foo by its identifier
func (h *FooHandler) ResolveFooByID(
	id uuid.UUID,
) {
	foo, err := h.FooApplication.ResolveFooByID(id)
	if err != nil {
		log.Error().Msg(fmt.Sprint(err.Error()))
		return
	}

	log.Info().Msg(fmt.Sprintf("Foo = %+v\n", foo))
}
