package handlers

import (
	"net/http"
	"readingtracker/internal/dto"
	"readingtracker/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoalHandler struct {
	service services.GoalService
}

func NewGoalHandler(service services.GoalService) *GoalHandler {
	return &GoalHandler{service: service}
}

func (h *GoalHandler) SetGoal(c *gin.Context) {
	var req dto.SetGoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.SetGoal(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set goal"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reading goal saved"})
}

func (h *GoalHandler) GetGoalProgress(c *gin.Context) {
	yearStr := c.Param("year")
	year, _ := strconv.Atoi(yearStr)

	progress, err := h.service.GetGoalProgress(year)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No goal set for this year"})
		return
	}
	c.JSON(http.StatusOK, progress)
}
