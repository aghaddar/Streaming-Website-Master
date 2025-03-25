package routes

import (
	"Streaming-Website-Master/controllers"
	"Streaming-Website-Master/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterGenreRoute(router *gin.Engine, db *gorm.DB) {
	genreService := services.NewGenreService(db)
	genreController := controllers.NewGenreController(genreService)

	genreRoutes := router.Group("/api/genre")
	{
		genreRoutes.POST("/", genreController.CreateGenre)
		genreRoutes.GET("/:id", genreController.GetGenre)
		genreRoutes.PUT("/:id", genreController.UpdateGenre)
		genreRoutes.DELETE("/:id", genreController.DeleteGenre)
		genreRoutes.GET("/", genreController.ListGenres)
	}
}
