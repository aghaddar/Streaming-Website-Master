package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRatingRoute(router *gin.Engine, rateController *controllers.RateController) {
	rateRoutes := router.Group("/api/ratings")
	{
		rateRoutes.POST("/", rateController.CreateRating)
		rateRoutes.GET("/:id", rateController.GetRatingByID)
		rateRoutes.PUT("/:id", rateController.UpdateRating)
		rateRoutes.DELETE("/:id", rateController.DeleteRating)
		rateRoutes.GET("/", rateController.ListRatings)
	}
}
