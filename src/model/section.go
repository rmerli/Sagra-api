package model

import (
	"gtmx/src/database"
)

type Section struct {
	Id   int64
	Name string
}

func (p Section) FromDatabase(section database.Section) (Section, error) {
	return Section{
		Id:   section.ID,
		Name: section.Name,
	}, nil
}

func (p Section) ToDatabase(section Section) database.Section {
	return database.Section{
		ID:   section.Id,
		Name: section.Name,
	}
}
