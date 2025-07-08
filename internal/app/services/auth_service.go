package app_services

import (
	"app/internal/domain/entities"
	"app/internal/infra/repositories"
	infra_services "app/internal/infra/services"
	"errors"
	"os"
	"time"
)

type AuthService struct {
	userRepo     repositories.IUserRepository
	tokenService infra_services.ITokenService
	hashService  infra_services.IHashService
}

var _ IAuthService = (*AuthService)(nil)

func NewAuthService(
	userRepo repositories.IUserRepository,
	tokenService infra_services.ITokenService,
	hashService infra_services.IHashService,
) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		tokenService: tokenService,
		hashService:  hashService,
	}
}

var ACCESS_TOKEN_EXPIRATION = 15 * time.Minute
var ACCESS_TOKEN_SECRET = "03492jasdbsadgyuwgu23"
var REFRESH_TOKEN_EXPIRATION = 7 * 24 * time.Hour
var REFRESH_TOKEN_SECRET = "dsajhnkdashkm78y"

func (as *AuthService) Login(username string, password string) (map[string]string, error) {
	user := as.userRepo.FindByUsername(username)
	if user == nil {
		return nil, errors.New("username or password is incorrect [username]")
	}

	if !as.hashService.Compare(user.Password, password) {
		return nil, errors.New("username or password is incorrect [password]")
	}

	token, err := as.tokenService.Generate(user.Id, os.Getenv("JWT_SECRET"), ACCESS_TOKEN_EXPIRATION)
	if err != nil {
		return nil, err
	}

	refreshToken, err := as.tokenService.Generate(user.Id, os.Getenv("JWT_SECRET"), REFRESH_TOKEN_EXPIRATION)

	if err != nil {
		return nil, err
	}

	return map[string]string{"token": token, "refresh_token": refreshToken}, nil

}

func (as *AuthService) Register(username string, password string) error {
	if as.userRepo.FindByUsername(username) != nil {
		return errors.New("user already exists")
	}

	passwordHash, err := as.hashService.Hash(password)
	if err != nil {
		return err
	}

	user, err := entities.NewUser(username, passwordHash)
	if err != nil {
		return err
	}

	as.userRepo.Save(user)
	return nil
}
