package routes

import (
	"Streaming-Website-Master/controllers"
	"Streaming-Website-Master/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCommentRoutes(router *gin.Engine, db *gorm.DB) {
	commentService := services.NewCommentService(db)
	commentController := controllers.NewCommentController(commentService)

	commentRoutes := router.Group("/api/comments")
	{
		commentRoutes.POST("/", commentController.CreateComment)
		commentRoutes.GET("/:id", commentController.GetComment)
		commentRoutes.PUT("/:id", commentController.UpdateComment)
		commentRoutes.DELETE("/:id", commentController.DeleteComment)
		commentRoutes.GET("/", commentController.ListComments)
	}
}
