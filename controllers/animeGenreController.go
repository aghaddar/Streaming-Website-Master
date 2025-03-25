package controllers

import (
	"Streaming-Website-Master/models"
	"Streaming-Website-Master/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AnimeGenreController struct {
	service *services.AnimeGenreService
}

func NewAnimeGenreController(service *services.AnimeGenreService) *AnimeGenreController {
	return &AnimeGenreController{service: service}
}

func (c *AnimeGenreController) CreateAnimeGenre(ctx *gin.Context) {
	var ag models.Anime_Genre
	if err := ctx.ShouldBindJSON(&ag); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateAnimeGenre(&ag); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, ag)
}

func (c *AnimeGenreController) GetAnimeGenre(ctx *gin.Context) {
	animeID, err := strconv.ParseUint(ctx.Param("anime_id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anime ID"})
		return
	}
	genreID, err := strconv.ParseUint(ctx.Param("genre_id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}
	ag, err := c.service.GetAnimeGenre(uint(animeID), uint(genreID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if ag == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Anime genre not found"})
		return
	}
	ctx.JSON(http.StatusOK, ag)
}

func (c *AnimeGenreController) DeleteAnimeGenre(ctx *gin.Context) {
	animeID, err := strconv.ParseUint(ctx.Param("anime_id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anime ID"})
		return
	}
	genreID, err := strconv.ParseUint(ctx.Param("genre_id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid genre ID"})
		return
	}
	if err := c.service.DeleteAnimeGenre(uint(animeID), uint(genreID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *AnimeGenreController) ListAnimeGenres(ctx *gin.Context) {
	ags, err := c.service.ListAnimeGenres()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ags)
}
