package controllers

import (
//	"content-app-backend/api/models"
	"content-app-backend/api/services"
//	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentController struct {
	service services.ContentService
}

func NewContentController(service services.ContentService) *ContentController {
	return &ContentController{service}
}

func (c *ContentController) CreateContent (ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "create_content"})
}


func (c *ContentController) PostClipToPlatform (ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "post_clip_to_platform"})
}


func (c *ContentController) GetClip (ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "get_clip"})
}
