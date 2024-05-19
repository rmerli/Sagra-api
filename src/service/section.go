package service

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/database/repository"
)

type Section struct {
	Repo *repository.SectionRepository
}

func (s *Section) Get(ctx context.Context, id int64) (database.Section, error) {
	return s.Repo.Get(ctx, id)
}

func (s *Section) Create(ctx context.Context, section database.Section) (database.Section, error) {
	return s.Repo.Insert(ctx, section)
}

func (s *Section) Update(ctx context.Context, section database.Section) (database.Section, error) {
	return s.Repo.Update(ctx, section)
}

func (s *Section) GetAll(ctx context.Context) ([]database.Section, error) {
	return s.Repo.List(ctx)
}

func NewSectionService(repo *repository.SectionRepository) Section {
	return Section{Repo: repo}
}
