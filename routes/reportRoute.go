package routes

import (
	"Streaming-Website-Master/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterReportRoutes(router *gin.Engine, reportController *controllers.ReportController) {
	reportRoutes := router.Group("/api/reports")
	{
		reportRoutes.POST("/", reportController.CreateReport)
		reportRoutes.GET("/:id", reportController.GetReportByID)
		reportRoutes.PUT("/:id", reportController.UpdateReport)
		reportRoutes.DELETE("/:id", reportController.DeleteReport)
		reportRoutes.GET("/", reportController.ListReports)
	}
}
