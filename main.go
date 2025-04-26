package main

import (
	"log"
	"os"

	"Streaming-Website-Master/controllers"
	"Streaming-Website-Master/database"
	"Streaming-Website-Master/models"
	"Streaming-Website-Master/routes"
	"Streaming-Website-Master/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env; if not found, continue with defaults.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default configuration")
	}

	// Set Gin mode from environment or default to "debug" mode.
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// Connect to the database.
	// Assuming database.ConnectDB returns (db, error).
	db, err := database.ConnectDB()
	log.Println("Database connection established successfully")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established successfully")

	// Auto-migrate database models.
	log.Println("Running database migrations...")
	if err := db.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Anime{},
		&models.Episode{},
		&models.Genre{},
		&models.Anime_Genre{},
		&models.Comments{},
		&models.Rating{},
		&models.Recommendation{},
		&models.Report{},
		&models.Watch{},
		&models.Watchlist{},
	); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Database migrations completed")

	// Create a new Gin router.
	router := gin.Default()

	// Configure CORS.
	corsConfig := cors.DefaultConfig()
	// Adjust allowed origins as needed (use AllowAllOrigins only for development)
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "https://your-production-domain.com"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	router.Use(cors.New(corsConfig))

	// Initialize services.
	userService := services.NewUserService(db)
	animeService := services.NewAnimeService(db)
	episodeService := services.NewEpisodeService(db)
	genreService := services.NewGenreService(db)
	animeGenreService := services.NewAnimeGenreService(db)
	commentService := services.NewCommentService(db)
	ratingService := services.NewRatingService(db)
	recommendationService := services.NewRecommendationService(db)
	reportService := services.NewReportService(db)
	watchProgressService := services.NewWatchProgressService(db)
	watchlistService := services.NewWatchlistService(db)
	adminService := services.NewAdminService(db)
	authService := services.NewAuthService(db)

	// Initialize controllers.
	userController := controllers.NewUserController(userService)
	animeController := controllers.NewAnimeController(animeService)
	episodeController := controllers.NewEpisodeController(episodeService)
	genreController := controllers.NewGenreController(genreService)
	animeGenreController := controllers.NewAnimeGenreController(animeGenreService)
	commentController := controllers.NewCommentController(commentService)
	ratingController := controllers.NewRateController(ratingService)
	recommendationController := controllers.NewRecommendationController(recommendationService)
	reportController := controllers.NewReportController(reportService)
	watchProgressController := controllers.NewWatchProgressController(watchProgressService)
	watchlistController := controllers.NewWatchlistController(watchlistService)
	adminController := controllers.NewAdminController(adminService)
	authController := controllers.NewAuthController(authService)

	// Register authentication routes.
	routes.RegisterAuthRoutes(router, authController)

	// Register other API routes.
	routes.RegisterUserRoutes(router, userController)
	routes.RegisterAnimeRoutes(router, animeController)
	routes.RegisterEpisodesRoutes(router, episodeController)
	routes.RegisterGenreRoute(router, genreController)
	routes.RegisterAnimeGenreRoutes(router, animeGenreController)
	routes.RegisterCommentRoutes(router, commentController)
	// Assuming there is a separate function for ratings; if not, adjust accordingly.
	routes.RegisterRatingRoutes(router, ratingController)
	routes.RegisterRecommendationRoutes(router, recommendationController)
	routes.RegisterReportRoutes(router, reportController)
	routes.RegisterWatchProgressRoutes(router, watchProgressController)
	routes.RegisterWatchlistRoutes(router, watchlistController)
	routes.RegisterAdminRoutes(router, adminController)

	// Serve static files.
	// Use the URL path ("/static") and the local directory ("./static") where your assets are stored.
	router.Static("/static", "./static")

	// Start the server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
