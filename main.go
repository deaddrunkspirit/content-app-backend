package main

import (
	"content-app-backend/api/config"
	"content-app-backend/api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	api := r.Group("/api")
	{
		userController := new(controllers.UserController)
		platformController := new(controllers.PlatformController)
		contentController := new(controllers.ContentController)

		// User routes
		api.GET("/user", userController.GetUser)
		api.POST("/user", userController.CreateUser)

		// Content requests
		api.GET("/clip", contentController.GetClip)
		api.POST("/post_clip_to_platform", contentController.PostClipToPlatform)
		api.POST("/content", contentController.CreateContent)

		// Platform requests
		api.GET("/platform_connector", platformController.GetConnector)
		api.POST("/platform_connector", platformController.AddConnector)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
