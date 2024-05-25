package database

type CategoryWithSection interface {
	Get() (Category, Section)
}

func (c GetAllCategoryWithSectionRow) Get() (Category, Section) {
	return c.Category, c.Section
}

func (c GetCategoryWithSectionRow) Get() (Category, Section) {
	return c.Category, c.Section
}
