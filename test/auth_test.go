package test

import (
	"gtmx/src/database"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

func (s *IntegrationTestSuite) TestRegisterUser() {
	f := make(url.Values)
	f.Set("email", "jon5@labstack.com")
	f.Set("password", "1234")
	req, err := http.NewRequest(echo.POST, "http://localhost:8080/signup", strings.NewReader(f.Encode()))
	s.NoError(err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusMovedPermanently, response.StatusCode)
}

func (s *IntegrationTestSuite) TestEmailAlreadyExists() {
	user := database.User{
		Email:    "jon5@labstack.com",
		Password: "1234",
		Salt:     "1234",
	}

	s.db.CreateUser(s.context, database.CreateUserParams{
		Email:    user.Email,
		Password: user.Password,
		Salt:     user.Salt,
	})

	f := make(url.Values)
	f.Set("email", user.Email)
	f.Set("password", user.Password)
	req, err := http.NewRequest(echo.POST, "http://localhost:8080/signup", strings.NewReader(f.Encode()))
	s.NoError(err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusBadRequest, response.StatusCode)

}
