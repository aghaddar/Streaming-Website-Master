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
	// Load environment variables from .env file if present.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default configuration")
	}

	// Set Gin mode from environment variable or default to debug mode.
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// Connect to the database.
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established successfully")

	// Auto-migrate database models in the correct order
	log.Println("Running database migrations...")

	// First, migrate tables without foreign keys or with simple dependencies
	log.Println("Migrating Admin model...")
	if err := db.AutoMigrate(&models.Admin{}); err != nil {
		log.Fatalf("Failed to migrate Admin model: %v", err)
	}

	log.Println("Migrating User model...")
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate User model: %v", err)
	}

	log.Println("Migrating Genre model...")
	if err := db.AutoMigrate(&models.Genre{}); err != nil {
		log.Fatalf("Failed to migrate Genre model: %v", err)
	}

	log.Println("Migrating Anime model...")
	if err := db.AutoMigrate(&models.Anime{}); err != nil {
		log.Fatalf("Failed to migrate Anime model: %v", err)
	}

	// Then migrate tables with foreign keys in the correct order
	log.Println("Migrating AnimeGenre model...")
	if err := db.AutoMigrate(&models.AnimeGenre{}); err != nil {
		log.Fatalf("Failed to migrate AnimeGenre model: %v", err)
	}

	log.Println("Migrating Episode model...")
	if err := db.AutoMigrate(&models.Episode{}); err != nil {
		log.Fatalf("Failed to migrate Episode model: %v", err)
	}

	log.Println("Migrating Comment model...")
	if err := db.AutoMigrate(&models.Comment{}); err != nil {
		log.Fatalf("Failed to migrate Comment model: %v", err)
	}

	log.Println("Migrating Rating model...")
	if err := db.AutoMigrate(&models.Rating{}); err != nil {
		log.Fatalf("Failed to migrate Rating model: %v", err)
	}

	log.Println("Migrating Recommendation model...")
	if err := db.AutoMigrate(&models.Recommendation{}); err != nil {
		log.Fatalf("Failed to migrate Recommendation model: %v", err)
	}

	log.Println("Migrating Watch model...")
	if err := db.AutoMigrate(&models.Watch{}); err != nil {
		log.Fatalf("Failed to migrate Watch model: %v", err)
	}

	log.Println("Migrating Watchlist model...")
	if err := db.AutoMigrate(&models.Watchlist{}); err != nil {
		log.Fatalf("Failed to migrate Watchlist model: %v", err)
	}

	// Migrate Report model last as it depends on both User and Comment
	log.Println("Migrating Report model...")
	if err := db.AutoMigrate(&models.Report{}); err != nil {
		log.Fatalf("Failed to migrate Report model: %v", err)
	}

	log.Println("Database migrations completed")

	// Create a new Gin router.
	router := gin.Default()

	// Configure CORS.
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "http://localhost:8080", "https://your-production-domain.com"}
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
	routes.RegisterAuthRoutes(router, authController, adminController)

	// Register other API routes.
	routes.RegisterUserRoutes(router, userController)
	routes.RegisterAnimeRoutes(router, animeController)
	routes.RegisterEpisodesRoutes(router, episodeController)
	routes.RegisterGenreRoute(router, genreController)
	routes.RegisterAnimeGenreRoutes(router, animeGenreController)
	routes.RegisterCommentRoutes(router, commentController)
	routes.RegisterRatingRoutes(router, ratingController)
	routes.RegisterRecommendationRoutes(router, recommendationController)
	routes.RegisterReportRoutes(router, reportController)
	routes.RegisterWatchProgressRoutes(router, watchProgressController)
	routes.RegisterWatchlistRoutes(router, watchlistController)
	routes.RegisterAdminRoutes(router, adminController)

	// Start the server on the specified port, defaulting to 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
