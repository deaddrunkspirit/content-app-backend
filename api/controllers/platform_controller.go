package controllers

import (
//	"content-app-backend/api/models"
	"content-app-backend/api/services"
//	"net/http"

	"github.com/gin-gonic/gin"
)

type PlatformController struct {
	service services.PlatformService
}

func NewPlatformController(service services.PlatformService) *PlatformController {
	return &PlatformController{service}
}

func (c *PlatformController) AddConnector (ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "add_connector"})
}


func (c *PlatformController) GetConnector (ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "get_connector"})
}
