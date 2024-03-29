package service

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/database/repository"
)

type AuthService struct {
	Repository repository.UserRepository
}

func (s AuthService) RegisterUser(ctx context.Context, email string, password string) (database.User, error) {
	user := database.User{
		Email:    email,
		Password: password,
	}

	user, err := s.Repository.InsertUser(ctx, user)
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}
