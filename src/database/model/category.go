package model

import (
	"gtmx/src/database"
)

type Category struct {
	Id      int64
	Name    string
	Section Section
}

func NewCategoryFromDatabase[T database.CategoryWithSection](category T) Category {
	cat, sec := category.Get()
	return Category{
		Id:   cat.ID,
		Name: cat.Name,
		Section: Section{
			Id:   sec.ID,
			Name: sec.Name,
		},
	}
}

func NewCategory(id int64, name string, section Section) Category {
	return Category{
		Id:      id,
		Name:    name,
		Section: section,
	}
}

func NewCategoryListFromDatabase[T database.CategoryWithSection](categoriesWithSections []T) []Category {
	categories := make([]Category, len(categoriesWithSections))

	for i, category := range categoriesWithSections {
		categories[i] = NewCategoryFromDatabase(category)
	}
	return categories
}
