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


	r := gin.Default()


	r.POST("/books", bookHandler.CreateBook)
    r.GET("/books", bookHandler.GetBooks)

	r.Run(":" + cfg.Port)
}