package handlers

import (
	"net/http"
	"strconv"

	"readingtracker/internal/dto"
	"readingtracker/internal/services"

	"github.com/gin-gonic/gin"
)

type ProgressHandler struct {
	service services.ProgressService
}

func NewProgressHandler(service services.ProgressService) *ProgressHandler {
	return &ProgressHandler{service: service}
}

func (h *ProgressHandler) UpdateProgress(c *gin.Context) {

	bookIDStr := c.Param("id")
	bookID, err := strconv.ParseUint(bookIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var req dto.UpdateProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.UpdateProgress(uint(bookID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update progress"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reading progress updated"})
}
