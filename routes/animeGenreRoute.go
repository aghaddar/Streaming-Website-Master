package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAnimeGenreRoutes(router *gin.Engine, animeGenreController *controllers.AnimeGenreController) {
	animeGenreRoutes := router.Group("/api/anime_genres")
	{
		animeGenreRoutes.POST("/", animeGenreController.CreateAnimeGenre)
		animeGenreRoutes.GET("/:anime_id/:genre_id", animeGenreController.GetAnimeGenre)
		animeGenreRoutes.DELETE("/:anime_id/:genre_id", animeGenreController.DeleteAnimeGenre)
		animeGenreRoutes.GET("/", animeGenreController.ListAnimeGenres)
	}
}
