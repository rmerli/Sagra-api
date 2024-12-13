package service

import (
	"context"
	"sagre/src/database/model"
	"sagre/src/database/repository"

	"github.com/google/uuid"
)

type User struct {
	Repo *repository.User
}

func (u *User) Get(ctx context.Context, id uuid.UUID) (model.User, error) {
	return u.Repo.Get(ctx, id)
}

func (u *User) GetByEmail(ctx context.Context, email string) (model.User, error) {
	return u.Repo.GetByEmail(ctx, email)
}

func (u *User) Create(ctx context.Context, User model.User) (model.User, error) {
	return u.Repo.Create(ctx, User)
}

func (u *User) Update(ctx context.Context, User model.User) (model.User, error) {
	return u.Repo.Update(ctx, User)
}

func (u *User) GetAll(ctx context.Context) ([]model.User, error) {
	return u.Repo.GetAll(ctx)
}

func NewUserService(repo *repository.User) User {
	return User{Repo: repo}
}
