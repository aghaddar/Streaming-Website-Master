package routes

import (
	"Streaming-Website-Master/controllers"
	"Streaming-Website-Master/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAnimeGenreRoutes(router *gin.Engine, db *gorm.DB) {
	animeGenreService := services.NewAnimeGenreService(db)
	animeGenreController := controllers.NewAnimeGenreController(animeGenreService)

	animeGenreRoutes := router.Group("/api/anime_genres")
	{
		animeGenreRoutes.POST("/", animeGenreController.CreateAnimeGenre)
		animeGenreRoutes.GET("/:anime_id/:genre_id", animeGenreController.GetAnimeGenre)
		animeGenreRoutes.DELETE("/:anime_id/:genre_id", animeGenreController.DeleteAnimeGenre)
		animeGenreRoutes.GET("/", animeGenreController.ListAnimeGenres)
	}
}
