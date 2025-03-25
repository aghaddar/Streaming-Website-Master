package routes

import (
	"Streaming-Website-Master/controllers"
	"Streaming-Website-Master/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterEpisodesRoutes(router *gin.Engine, db *gorm.DB) {
	episodeService := services.NewEpisodeService(db)
	episodeController := controllers.NewEpisodeController(episodeService)

	episodesRoutes := router.Group("/api/episodes")
	{
		episodesRoutes.POST("/", episodeController.CreateEpisode)
		episodesRoutes.GET("/", episodeController.ListEpisodes)
		episodesRoutes.GET("/:id", episodeController.GetEpisode)
		episodesRoutes.PUT("/:id", episodeController.UpdateEpisode)
		episodesRoutes.DELETE("/:id", episodeController.DeleteEpisode)
	}
}
