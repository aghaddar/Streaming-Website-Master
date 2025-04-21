package routes

import (
	"Streaming-Website-Master/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterConsumetRoutes(router *gin.Engine, consumetController *controllers.ConsumetController) {
	consumetRoutes := router.Group("/api/consumet")
	{
		consumetRoutes.GET("/anime/:provider/:searchTerm", consumetController.SearchAnime)
		consumetRoutes.GET("/anime/:provider/info/:seriesID", consumetController.GetAnimeInfo)
		consumetRoutes.GET("/anime/:provider/watch", consumetController.WatchEpisode)
	}
}
