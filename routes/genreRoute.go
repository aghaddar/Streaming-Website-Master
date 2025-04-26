package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterGenreRoute(router *gin.Engine, genreController *controllers.GenreController) {
	genreRoutes := router.Group("/api/genre")
	{
		genreRoutes.POST("/", genreController.CreateGenre)
		genreRoutes.GET("/:id", genreController.GetGenre)
		genreRoutes.PUT("/:id", genreController.UpdateGenre)
		genreRoutes.DELETE("/:id", genreController.DeleteGenre)
		genreRoutes.GET("/", genreController.ListGenres)
	}
}
