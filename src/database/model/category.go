package model

import (
	"github.com/google/uuid"
)

type Category struct {
	Model
	Name      string
	SectionID uuid.UUID
	Section   Section
	Products  []Product
}
