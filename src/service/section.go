package service

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/database/repository"
)

type Section struct {
	Repo *repository.SectionRepository
}

func (s *Section) Update(ctx context.Context, product database.Product) (database.Section, error) {
	// product, err := s.Repo.UpdateProduct(ctx, product)
	// if err != nil {
	// 	return database.Product{}, err
	// }
	// return product, nil
}

func (s *Section) Get(ctx context.Context, id int64) (database.Section, error) {
	// return s.Repo.GetOneProductById(ctx, id)
}

func NewSectionService(repo *repository.SectionRepository) Section {
	// return Section{Repo: repo}
}
