package main

import (
	"Streaming-Website-Master/database"
	"Streaming-Website-Master/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load .env variables (for JWT_SECRET_KEY etc.)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system env vars")
	}

	// Connect to MySQL database
	db := database.ConnectDB()

	// Set Gin mode
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug"
	}
	gin.SetMode(mode)

	// Setup Gin router
	router := gin.Default()

	// Register auth routes: /api/auth/login, register, etc.
	routes.RegisterAuthRoutes(router, db)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Server running on http://localhost:%s", port)
	router.Run(":" + port)
}
