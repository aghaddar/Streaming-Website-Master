package routes

import (
	"Streaming-Website-Master/controllers"
	"Streaming-Website-Master/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(router *gin.Engine, db *gorm.DB) {
	authService := services.NewAuthService(db)
	authController := controllers.NewAuthController(authService)

	// Auth routes
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/admin/login", authController.AdminLogin)
		authRoutes.POST("/change-password", authController.ChangePassword)
		authRoutes.GET("/user/:id", authController.GetUserProfile)
	}
}
