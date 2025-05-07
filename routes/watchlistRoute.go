package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterWatchlistRoutes(router *gin.Engine, controller *controllers.WatchlistController) {
	watchlistRoutes := router.Group("/api/watchlist")
	{
		// User-specific watchlist endpoints
		watchlistRoutes.GET("/user", controller.GetUserWatchlist)
		watchlistRoutes.POST("/add", controller.AddAnimeToWatchlist)
		watchlistRoutes.DELETE("/remove/:animeId", controller.RemoveAnimeFromWatchlist)
		watchlistRoutes.GET("/check/:animeId", controller.CheckAnimeInWatchlist)

		// General watchlist endpoints (if needed)
		watchlistRoutes.POST("/", controller.CreateWatchlist)
		watchlistRoutes.GET("/:id", controller.GetWatchlistByID)
		watchlistRoutes.PUT("/:id", controller.UpdateWatchlist)
		watchlistRoutes.DELETE("/:id", controller.DeleteWatchlist)
		watchlistRoutes.GET("/", controller.ListWatchlists)
	}
}
