package app_services

import (
	"app/internal/domain/entities"
	"app/internal/infra/repositories"
	"errors"
)

type UserService struct {
	userRepo repositories.IUserRepository
}

var _ IUserService = (*UserService)(nil)

func NewUserService(userRepo repositories.IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) Create(username, password string) (*entities.User, error) {
	userExists := us.userRepo.FindByUsername(username)
	if userExists != nil {
		return nil, errors.New("user already exists")
	}

	user, err := entities.NewUser(username, password)
	if err != nil {
		return nil, err
	}

	us.userRepo.Save(user)

	return user, nil
}
