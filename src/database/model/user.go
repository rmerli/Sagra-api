package model

import ()

type User struct {
	Model
	Email    string
	Password string
	Salt     string
}
