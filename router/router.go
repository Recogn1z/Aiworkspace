package router

import (
	"ai-workspace-backend/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.GET("/health", controller.HealthCheck)

		auth := api.Group("/auth")
		{
			auth.POST("/register",controller.Register)
		}
	}
	return r
}