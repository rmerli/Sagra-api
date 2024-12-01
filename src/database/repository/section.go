package repository

import (
	"context"
	"gtmx/src/database/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Section struct {
	db *gorm.DB
}

func (s *Section) Get(ctx context.Context, id uuid.UUID) (model.Section, error) {
	section := model.Section{ID: id}
	result := s.db.WithContext(ctx).First(&section)
	return section, result.Error
}

func (s *Section) Create(ctx context.Context, section model.Section) (model.Section, error) {
	result := s.db.WithContext(ctx).Create(&section)
	return section, result.Error
}

func (s *Section) Update(ctx context.Context, section model.Section) (model.Section, error) {
	result := s.db.WithContext(ctx).Save(&section)
	return section, result.Error
}

func (s *Section) GetAll(ctx context.Context) ([]model.Section, error) {
	sections := []model.Section{}
	result := s.db.WithContext(ctx).Find(&sections)
	return sections, result.Error
}

func NewSectionRepository(db *gorm.DB) Section {
	return Section{
		db: db,
	}
}
