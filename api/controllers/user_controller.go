package controllers

import (
//	"content-app-backend/api/models"
	"content-app-backend/api/services"
//	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

func (c *UserController) Auth (ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "auth"})
}


func (c *UserController) CreateUser (ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "create_user"})
}


func (c *UserController) GetUser (ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "get_user"})
}
