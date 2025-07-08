package app_services

import "app/internal/domain/entities"

type IUserService interface {
	Create(username, password string) (*entities.User, error)
}

type IAuthService interface {
	Login(username string, password string) (map[string]string, error)
	Register(username string, password string) error
}
