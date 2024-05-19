package model

import (
	"gtmx/src/database"
)

type Category struct {
	Id      int64
	Name    string
	Section Section
}

func NewCategoryFromDatabase(category database.GetCategoryWithSectionRow) Category {
	return Category{
		Id:   category.Category.ID,
		Name: category.Category.Name,
		Section: Section{
			Id:   category.Section.ID,
			Name: category.Section.Name,
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

func NewCategoryListFromDatabase(categoriesWithSections []database.GetAllCategoryWithSectionRow) []Category {
	categories := make([]Category, len(categoriesWithSections))

	for i, category := range categoriesWithSections {
		categories[i] = Category{
			Id:   category.Category.ID,
			Name: category.Category.Name,
			Section: Section{
				Id:   category.Section.ID,
				Name: category.Section.Name,
			},
		}
	}
	return categories
}
