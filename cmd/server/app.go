package server

import (
	"app/internal/app/di"

	"github.com/gin-gonic/gin"
)

func NewApp() *gin.Engine {
	r := gin.Default()
	c := di.NewContainer()

	r.POST("/auth/register", c.AuthController.RegisterUser)
	r.POST("/auth/login", c.AuthController.LoginUser)

	return r
}
