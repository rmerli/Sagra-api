package auth

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gtmx/src/database"
	"gtmx/src/database/repository"
	"math/big"
	"math/rand"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repository repository.UserRepository
}
type ValidationError struct {
	Key     string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (s AuthService) RegisterUser(ctx context.Context, email string, password string) (database.User, error) {
	hash := md5.Sum(big.NewInt(rand.Int63()).Bytes())
	salt := hex.EncodeToString(hash[:])
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%s%s", password, salt)), 8)

	if err != nil {
		log.Error(err)
		return database.User{}, err
	}

	user := database.User{
		Email:    email,
		Password: string(hashedPassword),
		Salt:     salt,
	}

	user, err = s.repository.InsertUser(ctx, user)
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
