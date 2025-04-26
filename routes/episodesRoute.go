package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterEpisodesRoutes(router *gin.Engine, episodeController *controllers.EpisodeController) {
	episodesRoutes := router.Group("/api/episodes")
	{
		episodesRoutes.POST("/", episodeController.CreateEpisode)
		episodesRoutes.GET("/", episodeController.ListEpisodes)
		episodesRoutes.GET("/:id", episodeController.GetEpisode)
		episodesRoutes.PUT("/:id", episodeController.UpdateEpisode)
		episodesRoutes.DELETE("/:id", episodeController.DeleteEpisode)
	}
}
