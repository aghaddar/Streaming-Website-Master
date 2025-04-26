package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(router *gin.Engine, commentController *controllers.CommentController) {
	commentRoutes := router.Group("/api/comments")
	{
		commentRoutes.POST("/", commentController.CreateComment)
		commentRoutes.GET("/:id", commentController.GetComment)
		commentRoutes.PUT("/:id", commentController.UpdateComment)
		commentRoutes.DELETE("/:id", commentController.DeleteComment)
		commentRoutes.GET("/", commentController.ListComments)
	}
}
