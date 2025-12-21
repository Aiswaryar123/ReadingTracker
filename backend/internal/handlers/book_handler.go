package handlers

import (
	"net/http"
	"readingtracker/internal/dto"
	"readingtracker/internal/services"
	"strconv"

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

func (h *BookHandler) UpdateBook(c *gin.Context) {
	var req dto.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateBook(req); err != nil {
		c.JSON(500, gin.H{"error": "Failed to update"})
		return
	}
	c.JSON(200, gin.H{"message": "Book updated"})
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	if err := h.service.DeleteBook(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}
	c.JSON(200, gin.H{"message": "Book deleted"})
}
