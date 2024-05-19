package repository

import (
	"context"
	"gtmx/src/database"
)

type CategoryRepository struct {
	db *database.Queries
}

func (r CategoryRepository) List(ctx context.Context) ([]database.Category, error) {
	categories, err := r.db.ListCategories(ctx)
	if err != nil {
		return []database.Category{}, err

	}

	return categories, nil
}

func (r CategoryRepository) Get(ctx context.Context, id int64) (database.Category, error) {
	product, err := r.db.GetCategory(ctx, id)

	if err != nil {
		return database.Category{}, err

	}

	return product, nil
}

func (r CategoryRepository) Insert(ctx context.Context, category database.Category) (database.Category, error) {

	insertedCategory, err := r.db.CreateCategory(ctx, database.CreateCategoryParams{
		Name:      category.Name,
		SectionID: category.SectionID,
	})
	if err != nil {
		return database.Category{}, err
	}

	return insertedCategory, nil
}

func (r CategoryRepository) Update(ctx context.Context, category database.Category) (database.Category, error) {

	category, err := r.db.UpdateCategory(ctx, database.UpdateCategoryParams{
		Name:      category.Name,
		SectionID: category.SectionID,
	})

	if err != nil {
		return database.Category{}, err
	}

	return category, nil
}

func NewCategoryRepository(db *database.Queries) CategoryRepository {
	return CategoryRepository{db: db}
}
