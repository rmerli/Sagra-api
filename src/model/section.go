package model

import (
	"gtmx/src/database"
)

type Section struct {
	Id   int64
	Name string
}

func NewSection(id int64, name string) Section {
	return Section{
		Id:   id,
		Name: name,
	}
}

func NewSectionList(dbSections []database.Section) []Section {
	sections := make([]Section, len(dbSections))
	for i, section := range dbSections {
		sections[i] = Section{
			Id:   section.ID,
			Name: section.Name,
		}
	}
	return sections
}
