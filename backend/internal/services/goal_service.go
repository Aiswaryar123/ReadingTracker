package services

import (
	"readingtracker/internal/dto"
	"readingtracker/internal/models"
	"readingtracker/internal/repository"
)

type GoalService interface {
	SetGoal(req dto.SetGoalRequest) error
	GetGoalProgress(year int) (*dto.GoalProgressResponse, error)
}

type goalService struct {
	repo repository.GoalRepository
}

func NewGoalService(repo repository.GoalRepository) GoalService {
	return &goalService{repo: repo}
}

func (s *goalService) SetGoal(req dto.SetGoalRequest) error {
	goal := &models.ReadingGoal{
		Year:        req.Year,
		TargetBooks: req.TargetBooks,
	}
	return s.repo.SaveGoal(goal)
}

func (s *goalService) GetGoalProgress(year int) (*dto.GoalProgressResponse, error) {

	goal, err := s.repo.GetGoalByYear(year)
	if err != nil {
		return nil, err
	}

	currentCount, err := s.repo.GetFinishedBooksCount(year)
	if err != nil {
		return nil, err
	}

	return &dto.GoalProgressResponse{
		Year:        goal.Year,
		Target:      goal.TargetBooks,
		Current:     int(currentCount),
		IsCompleted: int(currentCount) >= goal.TargetBooks,
	}, nil
}
