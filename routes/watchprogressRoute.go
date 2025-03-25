package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterWatchProgressRoutes(router *gin.Engine, controller *controllers.WatchProgressController) {
	watchProgressRoutes := router.Group("/api/watch-progress")
	{
		watchProgressRoutes.POST("/", controller.CreateWatchProgress)
		watchProgressRoutes.GET("/:id", controller.GetWatchProgressByID)
		watchProgressRoutes.PUT("/:id", controller.UpdateWatchProgress)
		watchProgressRoutes.DELETE("/:id", controller.DeleteWatchProgress)
		watchProgressRoutes.GET("/", controller.ListWatchProgresses)
	}
}
