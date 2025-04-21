package controllers

import (
	"Streaming-Website-Master/models"
	"Streaming-Website-Master/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EpisodeController struct {
	service *services.EpisodeService
}

func NewEpisodeController(service *services.EpisodeService) *EpisodeController {
	return &EpisodeController{service: service}
}

func (c *EpisodeController) CreateEpisode(ctx *gin.Context) {
	var episode models.Episode
	if err := ctx.ShouldBindJSON(&episode); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.CreateEpisode(&episode); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, episode)
}

func (c *EpisodeController) GetEpisode(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid episode ID"})
		return
	}
	episode, err := c.service.GetEpisodeByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if episode == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Episode not found"})
		return
	}
	ctx.JSON(http.StatusOK, episode)
}

/*
	func (c *EpisodeController) UpdateEpisode(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid episode ID"})
			return
		}
		var episode models.Episode
		if err := ctx.ShouldBindJSON(&episode); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		episode.EpisodeID = id
		if err := c.service.UpdateEpisode(&episode); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, episode)
	}
*/
func (c *EpisodeController) UpdateEpisode(ctx *gin.Context) {
	id := ctx.Param("id") // EpisodeID is now a string
	var episode models.Episode
	if err := ctx.ShouldBindJSON(&episode); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	episode.EpisodeID = id // Assign the string ID directly
	if err := c.service.UpdateEpisode(&episode); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, episode)
}
func (c *EpisodeController) DeleteEpisode(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid episode ID"})
		return
	}
	if err := c.service.DeleteEpisode(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *EpisodeController) ListEpisodes(ctx *gin.Context) {
	episodes, err := c.service.ListEpisodes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, episodes)
}
