package repository

import (
	"context"
	"gtmx/src/database"
)

type SectionRepository struct {
	db *database.Queries
}

func (r SectionRepository) List(ctx context.Context) ([]database.Section, error) {
	sections, err := r.db.ListSections(ctx)
	if err != nil {
		return []database.Section{}, err

	}

	return sections, nil
}

func (r SectionRepository) Get(ctx context.Context, id int64) (database.Section, error) {
	section, err := r.db.GetSection(ctx, id)

	if err != nil {
		return database.Section{}, err

	}

	return section, nil
}

func (r SectionRepository) Update(ctx context.Context, section database.Section) (database.Section, error) {
	section, err := r.db.UpdateSection(ctx, database.UpdateSectionParams{
		ID:   section.ID,
		Name: section.Name,
	})

	if err != nil {
		return database.Section{}, err
	}

	return section, nil
}

func (r SectionRepository) Insert(ctx context.Context, section database.Section) (database.Section, error) {

	insertedSection, err := r.db.CreateSection(ctx, section.Name)
	if err != nil {
		return database.Section{}, err
	}

	return insertedSection, nil
}

func NewSectionRepository(db *database.Queries) SectionRepository {
	return SectionRepository{db: db}
}
