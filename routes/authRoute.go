package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine, authController *controllers.AuthController) {
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/admin/login", authController.AdminLogin)
		authRoutes.POST("/change-password", authController.ChangePassword)
		authRoutes.GET("/user/:id", authController.GetUserProfile)
	}
}
