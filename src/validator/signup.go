package validator

import (
	"context"
	"errors"
	"sagre/src/service"
)

type SignUp struct {
	Email    string
	Password string
	Errors   map[string]string
}

func (v *SignUp) Validate(ctx context.Context, userService *service.User) error {
	v.Errors = make(map[string]string)
	_, err := userService.GetByEmail(ctx, v.Email)

	if err == nil {
		v.Errors["email"] = "Email address not available"
		return errors.New("Email address already exists")
	}

	return nil
}
