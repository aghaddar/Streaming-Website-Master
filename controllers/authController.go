package controllers

import (
	"net/http"
	"time"

	"Streaming-Website-Master/services"
	"github.com/gin-gonic/gin"
)

// AuthController handles authentication endpoints.
type AuthController struct{}

// NewAuthController creates a new instance of AuthController.
func NewAuthController() *AuthController {
	return &AuthController{}
}

// LoginInput represents the expected payload for a login request.
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Login handles user login by validating credentials and generating a JWT token.
func (a *AuthController) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real-world application, replace this static verification with a proper database lookup.
	if input.Email != "test@test.com" || input.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Generate a JWT token for the authenticated user.
	// Here, UserID is set to 1 and role to "user" as placeholders.
	token, err := services.GeneratesJWT(1, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"expires": time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})
}
