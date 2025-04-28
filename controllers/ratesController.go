package controllers

import (
	"net/http"
	"strconv"

	"Streaming-Website-Master/models"
	"Streaming-Website-Master/services"
	"github.com/gin-gonic/gin"
)

type RateController struct {
	ratingService *services.RatingService
}

func NewRateController(ratingService *services.RatingService) *RateController {
	return &RateController{ratingService: ratingService}
}

func (rc *RateController) CreateRating(c *gin.Context) {
	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := rc.ratingService.CreateRating(&rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, rating)
}

func (rc *RateController) GetRatingByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rating ID"})
		return
	}
	rating, err := rc.ratingService.GetRatingByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rating == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rating not found"})
		return
	}
	c.JSON(http.StatusOK, rating)
}

func (rc *RateController) UpdateRating(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rating ID"})
		return
	}
	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rating.RateID = uint64(id)
	if err := rc.ratingService.UpdateRating(&rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rating)
}

func (rc *RateController) DeleteRating(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rating ID"})
		return
	}
	if err := rc.ratingService.DeleteRating(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Rating deleted"})
}

func (rc *RateController) ListRatings(c *gin.Context) {
	ratings, err := rc.ratingService.ListRatings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ratings)
}
