package main

import (
	"Streaming-Website-Master/controllers"
	"Streaming-Website-Master/database"
	"Streaming-Website-Master/models"
	"Streaming-Website-Master/routes"
	"Streaming-Website-Master/services"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default configuration")
	}

	// Set Gin mode
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// Connect to database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connection established successfully")

	// Disable foreign key constraints during migration to avoid constraint errors
	db.DisableForeignKeyConstraintWhenMigrating = true

	// Auto migrate database models
	log.Println("Running database migrations...")
	db.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Anime{},
		&models.Episode{},
		&models.Genre{},
		&models.AnimeGenre{},
		&models.Comment{},
		&models.Rating{},
		&models.Recommendation{},
		&models.Report{},
		&models.Watch{},
		&models.Watchlist{},
	)
	log.Println("Database migrations completed")

	// Create a new Gin router
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	// Initialize services
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
	adminService := services.NewAdminService(db)
	authService := services.NewAuthService(db)

	// Initialize controllers
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
	// Note: watchlistService removed as watchlist routes are defined in routes/watchlistRoute.go
	adminController := controllers.NewAdminController(adminService)
	authController := controllers.NewAuthController(authService)

	// Register auth routes (include both auth and admin controllers as required)
	routes.RegisterAuthRoutes(router, authController, adminController)

	// Register other routes
	userRoutes := router.Group("/api/users")
	{
		userRoutes.GET("/", userController.ListUsers)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	animeRoutes := router.Group("/api/anime")
	{
		animeRoutes.GET("/", animeController.ListAnimes)
		animeRoutes.GET("/:id", animeController.GetAnimeByID)
		animeRoutes.POST("/", animeController.CreateAnime)
		animeRoutes.PUT("/:id", animeController.UpdateAnime)
		animeRoutes.DELETE("/:id", animeController.DeleteAnime)
	}

	episodeRoutes := router.Group("/api/episodes")
	{
		episodeRoutes.GET("/", episodeController.ListEpisodes)
		episodeRoutes.GET("/:id", episodeController.GetEpisode)
		episodeRoutes.POST("/", episodeController.CreateEpisode)
		episodeRoutes.PUT("/:id", episodeController.UpdateEpisode)
		episodeRoutes.DELETE("/:id", episodeController.DeleteEpisode)
	}

	genreRoutes := router.Group("/api/genre")
	{
		genreRoutes.GET("/", genreController.ListGenres)
		genreRoutes.GET("/:id", genreController.GetGenre)
		genreRoutes.POST("/", genreController.CreateGenre)
		genreRoutes.PUT("/:id", genreController.UpdateGenre)
		genreRoutes.DELETE("/:id", genreController.DeleteGenre)
	}

	animeGenreRoutes := router.Group("/api/anime_genres")
	{
		animeGenreRoutes.GET("/", animeGenreController.ListAnimeGenres)
		animeGenreRoutes.GET("/:anime_id/:genre_id", animeGenreController.GetAnimeGenre)
		animeGenreRoutes.POST("/", animeGenreController.CreateAnimeGenre)
		animeGenreRoutes.DELETE("/:anime_id/:genre_id", animeGenreController.DeleteAnimeGenre)
	}

	commentRoutes := router.Group("/api/comments")
	{
		commentRoutes.GET("/", commentController.ListComments)
		commentRoutes.GET("/:id", commentController.GetComment)
		commentRoutes.POST("/", commentController.CreateComment)
		commentRoutes.PUT("/:id", commentController.UpdateComment)
		commentRoutes.DELETE("/:id", commentController.DeleteComment)
	}

	ratingRoutes := router.Group("/api/ratings")
	{
		ratingRoutes.GET("/", ratingController.ListRatings)
		ratingRoutes.GET("/:id", ratingController.GetRatingByID)
		ratingRoutes.POST("/", ratingController.CreateRating)
		ratingRoutes.PUT("/:id", ratingController.UpdateRating)
		ratingRoutes.DELETE("/:id", ratingController.DeleteRating)
	}

	recommendationRoutes := router.Group("/api/recommendations")
	{
		recommendationRoutes.GET("/", recommendationController.ListRecommendations)
		recommendationRoutes.GET("/:id", recommendationController.GetRecommendationByID)
		recommendationRoutes.POST("/", recommendationController.CreateRecommendation)
		recommendationRoutes.PUT("/:id", recommendationController.UpdateRecommendation)
		recommendationRoutes.DELETE("/:id", recommendationController.DeleteRecommendation)
	}

	reportRoutes := router.Group("/api/reports")
	{
		reportRoutes.GET("/", reportController.ListReports)
		reportRoutes.GET("/:id", reportController.GetReportByID)
		reportRoutes.POST("/", reportController.CreateReport)
		reportRoutes.PUT("/:id", reportController.UpdateReport)
		reportRoutes.DELETE("/:id", reportController.DeleteReport)
	}

	watchProgressRoutes := router.Group("/api/watch-progress")
	{
		watchProgressRoutes.GET("/", watchProgressController.ListWatchProgresses)
		watchProgressRoutes.GET("/:id", watchProgressController.GetWatchProgressByID)
		watchProgressRoutes.POST("/", watchProgressController.CreateWatchProgress)
		watchProgressRoutes.PUT("/:id", watchProgressController.UpdateWatchProgress)
		watchProgressRoutes.DELETE("/:id", watchProgressController.DeleteWatchProgress)
	}

	adminRoutes := router.Group("/api/admins")
	{
		adminRoutes.GET("/", adminController.GetAdmins)
		adminRoutes.GET("/:id", adminController.GetAdmin)
		adminRoutes.POST("/", adminController.CreateAdmin)
		adminRoutes.PUT("/:id", adminController.UpdateAdmin)
		adminRoutes.DELETE("/:id", adminController.DeleteAdmin)
	}

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server starting on port %s...\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
