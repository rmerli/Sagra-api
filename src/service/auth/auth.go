package auth

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gtmx/src/database/model"
	"gtmx/src/service"
	"math/big"
	"math/rand"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserService *service.User
}
type ValidationError struct {
	Key     string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (s AuthService) RegisterUser(ctx context.Context, email string, password string) (model.User, error) {
	hash := md5.Sum(big.NewInt(rand.Int63()).Bytes())
	salt := hex.EncodeToString(hash[:])
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%s%s", password, salt)), 8)

	if err != nil {
		log.Error(err)
		return model.User{}, err
	}

	user, err := s.UserService.Create(ctx, model.User{
		Email:    email,
		Password: string(hashedPassword),
		Salt:     salt,
	})

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func NewAuthService(userService *service.User) AuthService {
	return AuthService{UserService: userService}
}

func GetUser(c echo.Context) (model.User, error) {
	session, _ := session.Get("session-key", c)
	user := session.Values["user"]
	return user.(model.User), nil
}
