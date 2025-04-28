package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine, authController *controllers.AuthController, adminsController *controllers.AdminController) {
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/admin/login", authController.AdminLogin)
		authRoutes.POST("/register/admin", adminsController.CreateAdmin)
		authRoutes.POST("/change/Password", authController.ChangePassword)
		authRoutes.GET("/user/:id", authController.GetUserProfile)
	}
}
