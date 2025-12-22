package routes

import (
	"readingtracker/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	bookHandler *handlers.BookHandler,
	progressHandler *handlers.ProgressHandler,
) {

	// Book CRUD
	r.POST("/books", bookHandler.CreateBook)
	r.GET("/books", bookHandler.GetBooks)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)

	// Reading Progress
	r.PUT("/books/:id/progress", progressHandler.UpdateProgress)
	r.GET("/books/:id/progress", progressHandler.GetProgress)
}
