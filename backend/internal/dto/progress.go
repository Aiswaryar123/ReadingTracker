package dto

type UpdateProgressRequest struct {
	Status      string `json:"status" binding:"required"`
	CurrentPage int    `json:"current_page"`
}
