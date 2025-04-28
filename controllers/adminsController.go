// File: controllers/adminsController.go
package controllers

import (
	"Streaming-Website-Master/services"
	"net/http"
	"strconv"

	"Streaming-Website-Master/database"
	"Streaming-Website-Master/models"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminsService *services.AdminService
}

// NewAdminController returns a new AdminController instance.
func NewAdminController(adminService *services.AdminService) *AdminController {
	return &AdminController{
		adminsService: adminService,
	}
}

func (a *AdminController) GetAdmins(c *gin.Context) {
	var admins []models.Admin
	if err := database.DB.Find(&admins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admins)
}

func (a *AdminController) GetAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	var admin models.Admin
	if err := database.DB.First(&admin, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (c *AdminController) CreateAdmin(ctx *gin.Context) {
	var request struct {
		Username     string `json:"username" binding:"required"`
		Email        string `json:"email" binding:"required,email"`
		PasswordHash string `json:"passwordHash" binding:"required,min=6"`
		Role         string `json:"role" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err := c.adminsService.Register(request.Username, request.Email, request.PasswordHash, request.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, admin)
}

func (a *AdminController) UpdateAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	var admin models.Admin
	if err := database.DB.First(&admin, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}
	var updatedAdmin models.Admin
	if err := c.ShouldBindJSON(&updatedAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	admin.Username = updatedAdmin.Username
	admin.Email = updatedAdmin.Email
	admin.PasswordHash = updatedAdmin.PasswordHash
	admin.Role = updatedAdmin.Role
	if err := database.DB.Save(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (a *AdminController) DeleteAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	if err := database.DB.Delete(&models.Admin{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admin deleted successfully"})
}
