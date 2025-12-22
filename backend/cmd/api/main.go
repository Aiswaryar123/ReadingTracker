package main

import (
	"readingtracker/configs"
	"readingtracker/internal/database"
	"readingtracker/internal/handlers"
	"readingtracker/internal/repository"
	"readingtracker/internal/routes"
	"readingtracker/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.LoadConfig()
	database.ConnectDB(cfg)

	// Repositories
	bookRepo := repository.NewBookRepository(database.DB)
	progressRepo := repository.NewProgressRepository(database.DB)
	reviewRepo := repository.NewReviewRepository(database.DB)
	goalRepo := repository.NewGoalRepository(database.DB)

	//  Services
	bookService := services.NewBookService(bookRepo)
	progressService := services.NewProgressService(progressRepo)
	reviewService := services.NewReviewService(reviewRepo, bookRepo)
	goalService := services.NewGoalService(goalRepo)

	//  Handlers
	bookHandler := handlers.NewBookHandler(bookService)
	progressHandler := handlers.NewProgressHandler(progressService)
	reviewHandler := handlers.NewReviewHandler(reviewService)
	goalHandler := handlers.NewGoalHandler(goalService)

	r := gin.Default()

	// Register Routes
	routes.RegisterRoutes(r, bookHandler, progressHandler, reviewHandler, goalHandler)

	r.Run(":" + cfg.Port)
}
