package services

import (
	"errors"
	"readingtracker/internal/models"
	"testing"
)

type FakeGoalRepo struct {
	Goal  *models.ReadingGoal
	Count int64
	Err   error
}

func (f *FakeGoalRepo) SaveGoal(goal *models.ReadingGoal) error {
	return nil
}

func (f *FakeGoalRepo) GetGoalByYear(year int) (*models.ReadingGoal, error) {
	if f.Err != nil {
		return nil, f.Err
	}
	return f.Goal, nil
}

func (f *FakeGoalRepo) GetFinishedBooksCount(year int) (int64, error) {
	return f.Count, nil
}

func TestGetGoalProgress_NotCompleted(t *testing.T) {

	repo := &FakeGoalRepo{
		Goal:  &models.ReadingGoal{Year: 2025, TargetBooks: 5},
		Count: 2,
	}
	service := NewGoalService(repo)

	result, err := service.GetGoalProgress(2025)

	if err != nil {
		t.Errorf("Did not expect an error, but got: %v", err)
	}
	if result.IsCompleted != false {
		t.Errorf("Goal should NOT be completed (2/5), but logic said it is true")
	}
	if result.Current != 2 {
		t.Errorf("Expected current count to be 2, but got %d", result.Current)
	}
}

func TestGetGoalProgress_Completed(t *testing.T) {

	repo := &FakeGoalRepo{
		Goal:  &models.ReadingGoal{Year: 2025, TargetBooks: 3},
		Count: 3,
	}
	service := NewGoalService(repo)

	result, err := service.GetGoalProgress(2025)

	if err != nil {
		t.Errorf("Did not expect an error, but got: %v", err)
	}
	if result.IsCompleted != true {
		t.Errorf("Goal SHOULD be completed (3/3), but logic said it is false")
	}
}

func TestGetGoalProgress_NotFound(t *testing.T) {

	repo := &FakeGoalRepo{
		Err: errors.New("record not found"),
	}
	service := NewGoalService(repo)

	_, err := service.GetGoalProgress(2099)

	if err == nil {
		t.Errorf("Expected an error for a non-existent goal, but got nil")
	}
}
