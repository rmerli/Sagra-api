package service

import (
	"context"
	"sagre/src/database/model"
	"sagre/src/database/repository"

	"github.com/google/uuid"
)

type Section struct {
	Repo *repository.Section
}

func (s *Section) Get(ctx context.Context, id uuid.UUID) (model.Section, error) {
	return s.Repo.Get(ctx, id)
}

func (s *Section) Create(ctx context.Context, section model.Section) (model.Section, error) {
	return s.Repo.Create(ctx, section)
}

func (s *Section) Update(ctx context.Context, section model.Section) (model.Section, error) {
	return s.Repo.Update(ctx, section)
}

func (s *Section) GetAll(ctx context.Context) ([]model.Section, error) {
	return s.Repo.GetAll(ctx)
}

func NewSectionService(repo *repository.Section) Section {
	return Section{Repo: repo}
}
