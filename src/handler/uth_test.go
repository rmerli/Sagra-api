package handler_test

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/database/repository"
	"gtmx/src/handler"
	"gtmx/src/service/auth"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = `{"email":"jon@labstack.com", "password": "1234"}`
)

func TestRegisterUser(t *testing.T) {
	e := echo.New()

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgresql://sagra:sagra@localhost/sagra_go?sslmode=disable")
	if err != nil {
		return
	}
	db := database.New(conn)

	userRepo := repository.NewUserRepository(db)
	authService := auth.NewAuthService(userRepo)
	authHandler := handler.AuthHandler{AuthService: authService}

	f := make(url.Values)
	f.Set("email", "jon@labstack.com")
	f.Set("password", "1234")

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	c := e.NewContext(req, rec)

	if assert.NoError(t, authHandler.HandleSignUp(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}

}
