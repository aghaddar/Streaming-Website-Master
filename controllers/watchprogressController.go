package controllers

import (
	"net/http"
	"strconv"

	"Streaming-Website-Master/models"
	"Streaming-Website-Master/services"
	"github.com/gin-gonic/gin"
)

type WatchProgressController struct {
	watchProgressService *services.WatchProgressService
}

func NewWatchProgressController(service *services.WatchProgressService) *WatchProgressController {
	return &WatchProgressController{watchProgressService: service}
}

func (wc *WatchProgressController) CreateWatchProgress(c *gin.Context) {
	var watch models.Watch
	if err := c.ShouldBindJSON(&watch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := wc.watchProgressService.CreateWatchProgress(&watch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, watch)
}

func (wc *WatchProgressController) GetWatchProgressByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid watch progress ID"})
		return
	}
	watch, err := wc.watchProgressService.GetWatchProgressByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if watch == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Watch progress not found"})
		return
	}
	c.JSON(http.StatusOK, watch)
}

func (wc *WatchProgressController) UpdateWatchProgress(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid watch progress ID"})
		return
	}
	var watch models.Watch
	if err := c.ShouldBindJSON(&watch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	watch.WatchID = uint(id)
	if err := wc.watchProgressService.UpdateWatchProgress(&watch); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, watch)
}

func (wc *WatchProgressController) DeleteWatchProgress(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid watch progress ID"})
		return
	}
	if err := wc.watchProgressService.DeleteWatchProgress(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Watch progress deleted"})
}

func (wc *WatchProgressController) ListWatchProgresses(c *gin.Context) {
	watches, err := wc.watchProgressService.ListWatchProgresses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, watches)
}
