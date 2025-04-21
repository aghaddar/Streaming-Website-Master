package main

import (
	"fmt"
	"log"
	"os"

	"Streaming-Website-Master/controllers"
	"Streaming-Website-Master/database"
	"Streaming-Website-Master/routes"
	"Streaming-Website-Master/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, relying on system environment variables.")
	}

	// Connect to DB
	db := database.ConnectDB()

	// Set Gin mode
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug"
	}
	gin.SetMode(mode)

	// Initialize router
	router := gin.Default()

	// Services & Controllers
	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)

	animeService := services.NewAnimeService(db)
	animeController := controllers.NewAnimeController(animeService)

	// ✅ Register Routes
	routes.RegisterUserRoutes(router, userController)
	routes.RegisterAnimeRoutes(router, animeController)
	routes.RegisterAuthRoutes(router, db) // <-- Make sure this is called!

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
	router.Run(":" + port)
}
