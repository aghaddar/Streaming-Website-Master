package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(server *gin.Engine, adminController *controllers.AdminController) {
	adminRoutes := server.Group("/api/admins")
	{
		adminRoutes.GET("/", adminController.GetAdmins)
		adminRoutes.GET("/:id", adminController.GetAdmin)
		adminRoutes.POST("/", adminController.CreateAdmin)
		adminRoutes.PUT("/:id", adminController.UpdateAdmin)
		adminRoutes.DELETE("/:id", adminController.DeleteAdmin)
	}
}
