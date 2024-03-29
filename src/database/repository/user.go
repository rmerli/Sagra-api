package repository

import (
	"context"
	"gtmx/src/database"
)

type UserRepository struct {
	Db *database.Queries
}

func (r UserRepository) InsertUser(ctx context.Context, user database.User) (database.User, error) {

	insertedUser, err := r.Db.CreateUser(ctx, database.CreateUserParams{
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		return database.User{}, err
	}

	return insertedUser, nil
}
