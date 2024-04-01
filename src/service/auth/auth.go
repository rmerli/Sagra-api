package auth

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/database/repository"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type AuthService struct {
	repository repository.UserRepository
}

func (s AuthService) RegisterUser(ctx context.Context, email string, password string) (database.User, error) {
	user := database.User{
		Email:    email,
		Password: password,
	}

	user, err := s.repository.InsertUser(ctx, user)
	if err != nil {
		return database.User{}, err
	}

	return user, nil
}

func (s AuthService) GetRepository() repository.UserRepository {
	return s.repository
}

func NewAuthService(repository repository.UserRepository) AuthService {
	return AuthService{repository: repository}
}

func GetUser(c echo.Context) (database.User, error) {
	session, _ := session.Get("session-key", c)
	user := session.Values["user"]
	return user.(database.User), nil
}
