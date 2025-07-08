package di

import (
	app_services "app/internal/app/services"
	"app/internal/infra/http/controllers"
	"app/internal/infra/repositories"
	infra_services "app/internal/infra/services"
)

type Container struct {
	AuthService    app_services.IAuthService
	AuthController controllers.IAuthController
}

func NewContainer() *Container {
	// user
	ur := repositories.NewUserRepository()

	// infra services
	hs := infra_services.NewBcryptHashService()
	jts := infra_services.NewJwtTokenService()

	// auth services
	as := app_services.NewAuthService(ur, jts, hs)
	ac := controllers.NewAuthController(as)

	return &Container{
		AuthService:    as,
		AuthController: ac,
	}
}
