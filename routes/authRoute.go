package routes

import (
	"Streaming-Website-Master/controllers"
	"Streaming-Website-Master/middleware"

	"github.com/gin-gonic/gin"
)

// AuthRoutes sets up the authentication routes.
func AuthRoutes(router *gin.Engine) {
	authController := controllers.NewAuthController()

	// Public route for login.
	router.POST("/login", authController.Login)

	// Example: Protected route group
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Add protected routes here. For example, a user profile endpoint.
		protected.GET("/profile", func(c *gin.Context) {
			// Retrieve user info from the context set in the middleware.
			userID, _ := c.Get("userID")
			role, _ := c.Get("role")
			c.JSON(200, gin.H{
				"userID": userID,
				"role":   role,
			})
		})
	}
}
