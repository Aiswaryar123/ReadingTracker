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

	// 1. Repositories
	bookRepo := repository.NewBookRepository(database.DB)
	progressRepo := repository.NewProgressRepository(database.DB)
	reviewRepo := repository.NewReviewRepository(database.DB)

	// 2. Services
	bookService := services.NewBookService(bookRepo)

	reviewService := services.NewReviewService(reviewRepo, bookRepo)
	progressService := services.NewProgressService(progressRepo)

	// 3. Handlers
	bookHandler := handlers.NewBookHandler(bookService)
	progressHandler := handlers.NewProgressHandler(progressService)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	r := gin.Default()

	routes.RegisterRoutes(r, bookHandler, progressHandler, reviewHandler)

	r.Run(":" + cfg.Port)
}
