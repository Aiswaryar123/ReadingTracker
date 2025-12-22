package services

import (
	"errors"
	"readingtracker/internal/models"
	"testing"
)

type MockGoalRepo struct {
	mockGoal  *models.ReadingGoal
	mockCount int64
	mockErr   error
}

func (m *MockGoalRepo) SaveGoal(goal *models.ReadingGoal) error { return nil }

func (m *MockGoalRepo) GetGoalByYear(year int) (*models.ReadingGoal, error) {
	if m.mockErr != nil {
		return nil, m.mockErr
	}
	return m.mockGoal, nil
}

func (m *MockGoalRepo) GetFinishedBooksCount(year int) (int64, error) {
	return m.mockCount, nil
}

func TestGetGoalProgress(t *testing.T) {

	t.Run("Should correctly calculate goal as NOT completed", func(t *testing.T) {

		mockRepo := &MockGoalRepo{
			mockGoal:  &models.ReadingGoal{Year: 2025, TargetBooks: 5},
			mockCount: 2,
		}
		service := NewGoalService(mockRepo)

		result, err := service.GetGoalProgress(2025)

		if err != nil {
			t.Errorf("Did not expect error, but got: %v", err)
		}
		if result.IsCompleted != false {
			t.Errorf("Expected IsCompleted to be false (2/5), but got true")
		}
	})

	t.Run("Should correctly calculate goal as COMPLETED", func(t *testing.T) {

		mockRepo := &MockGoalRepo{
			mockGoal:  &models.ReadingGoal{Year: 2025, TargetBooks: 3},
			mockCount: 3,
		}
		service := NewGoalService(mockRepo)

		result, err := service.GetGoalProgress(2025)

		if err != nil {
			t.Errorf("Did not expect error, but got: %v", err)
		}
		if result.IsCompleted != true {
			t.Errorf("Expected IsCompleted to be true (3/3), but got false")
		}
	})

	t.Run("Should return error if goal for year is not found", func(t *testing.T) {

		mockRepo := &MockGoalRepo{
			mockErr: errors.New("record not found"),
		}
		service := NewGoalService(mockRepo)

		_, err := service.GetGoalProgress(2030)

		if err == nil {
			t.Errorf("Expected an error for missing goal, but got nil")
		}
	})
}
