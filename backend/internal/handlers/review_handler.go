package handlers

import (
	"net/http"
	"readingtracker/internal/dto"
	"readingtracker/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	service services.ReviewService
}

func NewReviewHandler(service services.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {

	idParam := c.Param("id")
	bookID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var req dto.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AddReview(uint(bookID), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Review added successfully"})
}
func (h *ReviewHandler) GetReviews(c *gin.Context) {
	idParam := c.Param("id")
	bookID, _ := strconv.ParseUint(idParam, 10, 32)

	reviews, err := h.service.GetBookReviews(uint(bookID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}
