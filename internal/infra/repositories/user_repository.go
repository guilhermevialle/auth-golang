package repositories

import "app/internal/domain/entities"

type UserRepository struct {
	users []*entities.User
}

// implements IUserRepository
var _ IUserRepository = (*UserRepository)(nil)

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]*entities.User, 0),
	}
}

func (ur *UserRepository) Save(user *entities.User) {
	ur.users = append(ur.users, user)
}

func (ur *UserRepository) FindById(id string) *entities.User {
	for _, user := range ur.users {
		if user.Id == id {
			return user
		}
	}
	return nil
}

func (ur *UserRepository) FindByUsername(username string) *entities.User {
	for _, user := range ur.users {
		if user.Username == username {
			return user
		}
	}
	return nil
}
