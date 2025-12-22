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

	bookRepo := repository.NewBookRepository(database.DB)
	progressRepo := repository.NewProgressRepository(database.DB)

	bookService := services.NewBookService(bookRepo)
	progressService := services.NewProgressService(progressRepo)

	bookHandler := handlers.NewBookHandler(bookService)
	progressHandler := handlers.NewProgressHandler(progressService)

	r := gin.Default()

	routes.RegisterRoutes(r, bookHandler, progressHandler)

	r.Run(":" + cfg.Port)
}
