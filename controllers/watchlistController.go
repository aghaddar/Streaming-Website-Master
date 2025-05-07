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
	watchlist.WatchlistID = uint64(id)
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
func (wc *WatchlistController) GetUserWatchlist(c *gin.Context) {
	userIDParam := c.Query("userId")
	if userIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return
	}
	watchlists, err := wc.watchlistService.GetUserWatchlist(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, watchlists)
}

// AddAnimeToWatchlist adds an anime to the user's watchlist.
// It expects a JSON body with userId and animeId.
func (wc *WatchlistController) AddAnimeToWatchlist(c *gin.Context) {
	var watchlist models.Watchlist
	if err := c.ShouldBindJSON(&watchlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := wc.watchlistService.AddAnimeToWatchlist(&watchlist); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, watchlist)
}

// RemoveAnimeFromWatchlist removes an anime from the user's watchlist.
// It expects the animeId as a URL parameter and a query or context parameter for userId.
func (wc *WatchlistController) RemoveAnimeFromWatchlist(c *gin.Context) {
	animeID := c.Param("animeId")
	userIDParam := c.Query("userId")
	if userIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return
	}
	if err := wc.watchlistService.RemoveAnimeFromWatchlist(userID, animeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Anime removed from watchlist"})
}

// CheckAnimeInWatchlist checks if an anime is in the user's watchlist.
// It expects the animeId as a URL parameter and a query or context parameter for userId.
func (wc *WatchlistController) CheckAnimeInWatchlist(c *gin.Context) {
	animeID := c.Param("animeId")
	userIDParam := c.Query("userId")
	if userIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return
	}
	exists, err := wc.watchlistService.CheckAnimeInWatchlist(userID, animeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"exists": exists})
}
