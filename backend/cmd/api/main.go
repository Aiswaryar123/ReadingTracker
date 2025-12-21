package main

import (
	"readingtracker/configs"
	"readingtracker/internal/database"
	"readingtracker/internal/handlers"
	"readingtracker/internal/repository"
	"readingtracker/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.LoadConfig()
	database.ConnectDB(cfg)

	bookRepo := repository.NewBookRepository(database.DB)
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	progressRepo := repository.NewProgressRepository(database.DB)
	progressService := services.NewProgressService(progressRepo)
	progressHandler := handlers.NewProgressHandler(progressService)

	r := gin.Default()

	// Book CRUD
	r.POST("/books", bookHandler.CreateBook)
	r.GET("/books", bookHandler.GetBooks)
	r.PUT("/books", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)

	// Reading Progress
	r.PUT("/books/:id/progress", progressHandler.UpdateProgress)

	r.Run(":" + cfg.Port)
}
