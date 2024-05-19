package repository

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/model"

	"github.com/jackc/pgx/v5/pgtype"
)

type CategoryRepository struct {
	db *database.Queries
}

func (r CategoryRepository) List(ctx context.Context) ([]model.Category, error) {
	categoriesWithSections, err := r.db.GetAllCategoryWithSection(ctx)
	if err != nil {
		return []model.Category{}, err

	}

	return model.NewCategoryListFromDatabase(categoriesWithSections), nil
}

func (r CategoryRepository) Get(ctx context.Context, id int64) (model.Category, error) {
	categoryWithSection, err := r.db.GetCategoryWithSection(ctx, id)
	if err != nil {
		return model.Category{}, err
	}

	return model.NewCategoryFromDatabase(categoryWithSection), nil
}

func (r CategoryRepository) Insert(ctx context.Context, category model.Category) (model.Category, error) {

	insertedCategory, err := r.db.CreateCategory(ctx, database.CreateCategoryParams{
		Name:      category.Name,
		SectionID: pgtype.Int8{Int64: category.Section.Id, Valid: true},
	})

	if err != nil {
		return model.Category{}, err
	}

	category.Id = insertedCategory.ID
	return category, nil
}

func (r CategoryRepository) Update(ctx context.Context, category model.Category) (model.Category, error) {
	_, err := r.db.UpdateCategory(ctx, database.UpdateCategoryParams{
		ID:        category.Id,
		Name:      category.Name,
		SectionID: pgtype.Int8{Int64: category.Section.Id, Valid: true},
	})

	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func NewCategoryRepository(db *database.Queries) CategoryRepository {
	return CategoryRepository{db: db}
}
