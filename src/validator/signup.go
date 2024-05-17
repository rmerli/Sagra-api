package validator

import (
	"context"
	"errors"
	"gtmx/src/database/repository"
)

type SignUp struct {
	Email    string
	Password string
	Errors   map[string]string
}

func (v *SignUp) Validate(ctx context.Context, repo repository.UserRepository) error {
	v.Errors = make(map[string]string)
	_, err := repo.GetUser(ctx, v.Email)

	if err == nil {
		v.Errors["email"] = "Email address not available"
		return errors.New("Email address already exists")
	}

	return nil
}
