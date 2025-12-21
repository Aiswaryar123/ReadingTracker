package services_test

import (
	"errors"
	"testing"

	"readingtracker/internal/dto"
	"readingtracker/internal/models"
	"readingtracker/internal/services"

	"github.com/stretchr/testify/assert"
)

type MockProgressRepository struct {
	Progress *models.ReadingProgress
	Err      error
}

func (m *MockProgressRepository) GetByBookID(bookID uint) (*models.ReadingProgress, error) {
	if m.Progress == nil {
		return nil, errors.New("record not found")
	}
	return m.Progress, nil
}

func (m *MockProgressRepository) Save(p *models.ReadingProgress) error {
	m.Progress = p
	return nil
}

func TestUpdateProgress_NewEntry(t *testing.T) {
	mockRepo := &MockProgressRepository{}
	service := services.NewProgressService(mockRepo)

	req := dto.UpdateProgressRequest{
		Status:      "Reading",
		CurrentPage: 50,
	}

	err := service.UpdateProgress(1, req)

	assert.NoError(t, err)
	assert.Equal(t, 50, mockRepo.Progress.CurrentPage)
	assert.Equal(t, "Reading", mockRepo.Progress.Status)
}
func TestUpdateProgress_ExistingEntry(t *testing.T) {
	existing := &models.ReadingProgress{
		ID:          1,
		BookID:      1,
		Status:      "Reading",
		CurrentPage: 20,
	}

	mockRepo := &MockProgressRepository{
		Progress: existing,
	}

	service := services.NewProgressService(mockRepo)

	req := dto.UpdateProgressRequest{
		Status:      "Completed",
		CurrentPage: 200,
	}

	err := service.UpdateProgress(1, req)

	assert.NoError(t, err)
	assert.Equal(t, "Completed", mockRepo.Progress.Status)
	assert.Equal(t, 200, mockRepo.Progress.CurrentPage)
}
