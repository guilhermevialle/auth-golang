package controllers

import "github.com/gin-gonic/gin"

type IUserController interface {
	CreateUser(ctx *gin.Context)
}

type IAuthController interface {
	LoginUser(ctx *gin.Context)
	RegisterUser(ctx *gin.Context)
}
