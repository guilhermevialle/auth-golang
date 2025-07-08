package repositories

import "app/internal/domain/entities"

type IUserRepository interface {
	Save(user *entities.User)
	FindById(id string) *entities.User
	FindByUsername(username string) *entities.User
}
