package entities

import (
	nanoid "github.com/matoous/go-nanoid/v2"
)

type User struct {
	Id       string
	Username string
	Password string
}

func NewUser(username, password string) (*User, error) {
	id, err := nanoid.New(21)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:       id,
		Username: username,
		Password: password,
	}, nil
}
