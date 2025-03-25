package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterWatchlistRoutes(router *gin.Engine, controller *controllers.WatchlistController) {
	watchlistRoutes := router.Group("/api/watchlist")
	{
		watchlistRoutes.POST("/", controller.CreateWatchlist)
		watchlistRoutes.GET("/:id", controller.GetWatchlistByID)
		watchlistRoutes.PUT("/:id", controller.UpdateWatchlist)
		watchlistRoutes.DELETE("/:id", controller.DeleteWatchlist)
		watchlistRoutes.GET("/", controller.ListWatchlists)
	}
}
