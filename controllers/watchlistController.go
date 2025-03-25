package controllers

import (
	"net/http"
	"strconv"

	"Streaming-Website-Master/models"
	"Streaming-Website-Master/services"
	"github.com/gin-gonic/gin"
)

type WatchlistController struct {
	watchlistService *services.WatchlistService
}

func NewWatchlistController(service *services.WatchlistService) *WatchlistController {
	return &WatchlistController{watchlistService: service}
}

func (wc *WatchlistController) CreateWatchlist(c *gin.Context) {
	var watchlist models.Watchlist
	if err := c.ShouldBindJSON(&watchlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := wc.watchlistService.CreateWatchlist(&watchlist); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, watchlist)
}

func (wc *WatchlistController) GetWatchlistByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid watchlist ID"})
		return
	}
	watchlist, err := wc.watchlistService.GetWatchlistByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if watchlist == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Watchlist not found"})
		return
	}
	c.JSON(http.StatusOK, watchlist)
}

func (wc *WatchlistController) UpdateWatchlist(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid watchlist ID"})
		return
	}
	var watchlist models.Watchlist
	if err := c.ShouldBindJSON(&watchlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	watchlist.WatchlistID = uint(id)
	if err := wc.watchlistService.UpdateWatchlist(&watchlist); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, watchlist)
}

func (wc *WatchlistController) DeleteWatchlist(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid watchlist ID"})
		return
	}
	if err := wc.watchlistService.DeleteWatchlist(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Watchlist deleted"})
}

func (wc *WatchlistController) ListWatchlists(c *gin.Context) {
	watchlists, err := wc.watchlistService.ListWatchlists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, watchlists)
}
