package entity

import "github.com/gofrs/uuid"

type Foo struct {
	ID          uuid.UUID
	Name        string
	Description string
}
