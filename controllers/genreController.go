package controllers

import (
	"Streaming-Website-Master/models"
	"Streaming-Website-Master/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GenreController struct {
	service *services.GenreService
}

func NewGenreController(service *services.GenreService) *GenreController {
	return &GenreController{service: service}
}

// CreateGenre creates a new genre.
func (gc *GenreController) CreateGenre(c *gin.Context) {
	var genre models.Genre
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := gc.service.CreateGenre(&genre); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, genre)
}

// GetGenre retrieves a genre by its id.
func (gc *GenreController) GetGenre(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre id"})
		return
	}
	genre, err := gc.service.GetGenreByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if genre == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Genre not found"})
		return
	}
	c.JSON(http.StatusOK, genre)
}

// UpdateGenre updates an existing genre.
func (gc *GenreController) UpdateGenre(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre id"})
		return
	}
	var genre models.Genre
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	genre.GenreID = uint64(id)
	if err := gc.service.UpdateGenre(&genre); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, genre)
}

// DeleteGenre deletes a genre by its id.
func (gc *GenreController) DeleteGenre(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre id"})
		return
	}
	if err := gc.service.DeleteGenre(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Genre deleted successfully"})
}

// ListGenres returns all genres.
func (gc *GenreController) ListGenres(c *gin.Context) {
	genres, err := gc.service.ListGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, genres)
}
