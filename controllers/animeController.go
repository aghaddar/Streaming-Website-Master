package controllers

import (
	"Streaming-Website-Master/models"
	"Streaming-Website-Master/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AnimeController struct {
	service *services.AnimeService
}

func NewAnimeController(as *services.AnimeService) *AnimeController {
	return &AnimeController{service: as}
}

func (ac *AnimeController) CreateAnime(c *gin.Context) {
	var anime models.Anime
	if err := c.ShouldBindJSON(&anime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ac.service.CreateAnime(&anime); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, anime)
}

func (ac *AnimeController) GetAnimeByID(c *gin.Context) {
	idParam := c.Param("id") // AnimeID is a string
	anime, err := ac.service.GetAnimeByID(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if anime == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Anime not found"})
		return
	}
	c.JSON(http.StatusOK, anime)
}

func (ac *AnimeController) UpdateAnime(c *gin.Context) {
	idParam := c.Param("id") // AnimeID is a string
	var anime models.Anime
	if err := c.ShouldBindJSON(&anime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	anime.AnimeID = idParam // Set AnimeID from URL parameter
	if err := ac.service.UpdateAnime(&anime); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, anime)
}

func (ac *AnimeController) DeleteAnime(c *gin.Context) {
	idParam := c.Param("id") // AnimeID is a string
	if err := ac.service.DeleteAnime(idParam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Anime deleted successfully"})
}

func (ac *AnimeController) ListAnimes(c *gin.Context) {
	animes, err := ac.service.ListAnimes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, animes)
}
