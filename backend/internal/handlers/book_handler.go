package handlers

import (
	"net/http"

	"readingtracker/internal/dto"
	"readingtracker/internal/services"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service services.BookService
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.service.FetchBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}
func NewBookHandler(service services.BookService) *BookHandler {
	return &BookHandler{service: service}
}


func (h *BookHandler) CreateBook(c *gin.Context) {
	var req dto.CreateBookRequest

	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	book, err := h.service.CreateBook(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	
	c.JSON(http.StatusCreated, book)
}