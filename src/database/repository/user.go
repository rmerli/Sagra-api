package repository

import (
	"context"
	"gtmx/src/database/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func (v *User) Get(ctx context.Context, id uuid.UUID) (model.User, error) {
	user := model.User{ID: id}
	result := v.db.WithContext(ctx).First(&user)
	return user, result.Error
}

func (v *User) GetByEmail(ctx context.Context, email string) (model.User, error) {
	user := model.User{Email: email}
	result := v.db.WithContext(ctx).First(&user)
	return user, result.Error
}

func (v *User) Create(ctx context.Context, user model.User) (model.User, error) {
	result := v.db.WithContext(ctx).Create(&user)
	return user, result.Error
}

func (v *User) Update(ctx context.Context, user model.User) (model.User, error) {
	result := v.db.WithContext(ctx).Save(&user)
	return user, result.Error
}

func (v *User) GetAll(ctx context.Context) ([]model.User, error) {
	users := []model.User{}
	result := v.db.WithContext(ctx).Find(&users)
	return users, result.Error
}

func NewUserRepository(db *gorm.DB) User {
	return User{
		db: db,
	}
}
