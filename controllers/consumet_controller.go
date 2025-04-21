package controllers

import (
	"net/http"
	"strconv"

	"Streaming-Website-Master/services"

	"github.com/gin-gonic/gin"
)

type ConsumetController struct {
	consumetService *services.ConsumetService
}

func NewConsumetController(service *services.ConsumetService) *ConsumetController {
	return &ConsumetController{consumetService: service}
}

func (c *ConsumetController) SearchAnime(ctx *gin.Context) {
	provider := ctx.Param("provider")
	searchTerm := ctx.Param("searchTerm")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	result, err := c.consumetService.SearchAnime(provider, searchTerm, page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *ConsumetController) GetAnimeInfo(ctx *gin.Context) {
	provider := ctx.Param("provider")
	seriesID := ctx.Param("seriesID")
	episodePage, _ := strconv.Atoi(ctx.DefaultQuery("episodePage", "1"))

	result, err := c.consumetService.GetAnimeInfo(provider, seriesID, episodePage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *ConsumetController) WatchEpisode(ctx *gin.Context) {
	provider := ctx.Param("provider")
	episodeID := ctx.Query("episodeId")

	result, err := c.consumetService.WatchEpisode(provider, episodeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
