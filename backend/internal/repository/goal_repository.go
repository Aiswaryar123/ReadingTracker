package repository

import (
	"readingtracker/internal/models"

	"gorm.io/gorm"
)

type GoalRepository interface {
	SaveGoal(goal *models.ReadingGoal) error
	GetGoalByYear(year int) (*models.ReadingGoal, error)
	GetFinishedBooksCount(year int) (int64, error)
}

type goalRepository struct {
	db *gorm.DB
}

func NewGoalRepository(db *gorm.DB) GoalRepository {
	return &goalRepository{db: db}
}

func (r *goalRepository) SaveGoal(goal *models.ReadingGoal) error {

	return r.db.Save(goal).Error
}

func (r *goalRepository) GetGoalByYear(year int) (*models.ReadingGoal, error) {
	var goal models.ReadingGoal
	err := r.db.Where("year = ?", year).First(&goal).Error
	return &goal, err
}

func (r *goalRepository) GetFinishedBooksCount(year int) (int64, error) {
	var count int64

	err := r.db.Model(&models.ReadingProgress{}).
		Where("status = ? AND EXTRACT(YEAR FROM last_updated) = ?", "Completed", year).
		Count(&count).Error
	return count, err
}
