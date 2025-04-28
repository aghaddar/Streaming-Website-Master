package controllers

import (
	"net/http"
	"strconv"

	"Streaming-Website-Master/models"
	"Streaming-Website-Master/services"
	"github.com/gin-gonic/gin"
)

type RecommendationController struct {
	recommendationService *services.RecommendationService
}

func NewRecommendationController(service *services.RecommendationService) *RecommendationController {
	return &RecommendationController{recommendationService: service}
}

func (rc *RecommendationController) CreateRecommendation(c *gin.Context) {
	var recommendation models.Recommendation
	if err := c.ShouldBindJSON(&recommendation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := rc.recommendationService.CreateRecommendation(&recommendation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, recommendation)
}

func (rc *RecommendationController) GetRecommendationByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recommendation ID"})
		return
	}
	recommendation, err := rc.recommendationService.GetRecommendationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if recommendation == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recommendation not found"})
		return
	}
	c.JSON(http.StatusOK, recommendation)
}

func (rc *RecommendationController) UpdateRecommendation(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recommendation ID"})
		return
	}
	var recommendation models.Recommendation
	if err := c.ShouldBindJSON(&recommendation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recommendation.RecommendationID = uint64(id)
	if err := rc.recommendationService.UpdateRecommendation(&recommendation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recommendation)
}

func (rc *RecommendationController) DeleteRecommendation(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recommendation ID"})
		return
	}
	if err := rc.recommendationService.DeleteRecommendation(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Recommendation deleted"})
}

func (rc *RecommendationController) ListRecommendations(c *gin.Context) {
	recommendations, err := rc.recommendationService.ListRecommendations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recommendations)
}
