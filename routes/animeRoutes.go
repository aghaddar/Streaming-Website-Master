package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAnimeRoutes(router *gin.Engine, animeController *controllers.AnimeController) {
	animeRoutes := router.Group("/api/anime")
	{
		animeRoutes.POST("/", animeController.CreateAnime)
		animeRoutes.GET("/:id", animeController.GetAnimeByID)
		animeRoutes.PUT("/:id", animeController.UpdateAnime)
		animeRoutes.DELETE("/:id", animeController.DeleteAnime)
		animeRoutes.GET("/", animeController.ListAnimes)
	}
}
