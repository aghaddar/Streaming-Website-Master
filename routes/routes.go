package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/api/auth/register") //register api
	server.POST("/api/auth/login")
	server.POST("/api/auth/logout")
	server.GET("/api/user/profile")
	server.PUT("/api/user/profile")
	server.DELETE("/api/user/delete")
}
