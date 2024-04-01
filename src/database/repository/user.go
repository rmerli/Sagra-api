package repository

import (
	"context"
	"gtmx/src/database"
)

type UserRepository struct {
	db *database.Queries
}

func (r UserRepository) InsertUser(ctx context.Context, user database.User) (database.User, error) {

	insertedUser, err := r.db.CreateUser(ctx, database.CreateUserParams{
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		return database.User{}, err
	}

	return insertedUser, nil
}

func (r UserRepository) GetUser(ctx context.Context, email string) (database.User, error) {
	return r.db.GetUser(ctx, email)
}

func NewUserRepository(db *database.Queries) UserRepository {
	return UserRepository{db: db}
}
