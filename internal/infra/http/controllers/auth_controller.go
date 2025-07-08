package controllers

import (
	app_services "app/internal/app/services"
	"app/internal/interfaces/dtos"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService app_services.IAuthService
}

var _ IAuthController = (*AuthController)(nil)

func NewAuthController(
	authService app_services.IAuthService,
) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) LoginUser(ctx *gin.Context) {
	var req dtos.LoginUserDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tokens, err := ac.authService.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, tokens)
}

func (ac *AuthController) RegisterUser(ctx *gin.Context) {
	var req dtos.RegisterUserDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := ac.authService.Register(req.Username, req.Password)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "user created"})
}
