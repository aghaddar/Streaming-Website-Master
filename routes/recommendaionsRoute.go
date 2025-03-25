package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRecommendationRoutes(router *gin.Engine, recommendationController *controllers.RecommendationController) {
	recommendationRoutes := router.Group("/api/recommendations")
	{
		recommendationRoutes.POST("/", recommendationController.CreateRecommendation)
		recommendationRoutes.GET("/:id", recommendationController.GetRecommendationByID)
		recommendationRoutes.PUT("/:id", recommendationController.UpdateRecommendation)
		recommendationRoutes.DELETE("/:id", recommendationController.DeleteRecommendation)
		recommendationRoutes.GET("/", recommendationController.ListRecommendations)
	}
}
