package repository

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/database/model"
)

type SectionRepository struct {
	db *database.Queries
}

func (r SectionRepository) List(ctx context.Context) ([]model.Section, error) {
	sections, err := r.db.ListSections(ctx)
	if err != nil {
		return []model.Section{}, err
	}

	return model.NewSectionList(sections), nil
}

func (r SectionRepository) Get(ctx context.Context, id int64) (model.Section, error) {
	section, err := r.db.GetSection(ctx, id)

	if err != nil {
		return model.Section{}, err

	}

	return model.NewSection(section.ID, section.Name), nil
}

func (r SectionRepository) Update(ctx context.Context, section model.Section) (model.Section, error) {
	_, err := r.db.UpdateSection(ctx, database.UpdateSectionParams{
		ID:   section.Id,
		Name: section.Name,
	})

	if err != nil {
		return model.Section{}, err
	}

	return section, nil
}

func (r SectionRepository) Insert(ctx context.Context, section model.Section) (model.Section, error) {
	insertedSection, err := r.db.CreateSection(ctx, section.Name)
	if err != nil {
		return model.Section{}, err
	}

	section.Id = insertedSection.ID
	return section, nil
}

func NewSectionRepository(db *database.Queries) SectionRepository {
	return SectionRepository{db: db}
}
