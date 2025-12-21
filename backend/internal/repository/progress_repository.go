package repository

import (
	"readingtracker/internal/models"

	"gorm.io/gorm"
)

type ProgressRepository interface {
	GetByBookID(bookID uint) (*models.ReadingProgress, error)
	Save(progress *models.ReadingProgress) error
}

type progressRepository struct {
	db *gorm.DB
}

func NewProgressRepository(db *gorm.DB) ProgressRepository {
	return &progressRepository{db: db}
}

func (r *progressRepository) GetByBookID(bookID uint) (*models.ReadingProgress, error) {
	var progress models.ReadingProgress

	err := r.db.Where("book_id = ?", bookID).First(&progress).Error
	if err != nil {
		return nil, err
	}
	return &progress, nil
}

func (r *progressRepository) Save(progress *models.ReadingProgress) error {

	return r.db.Save(progress).Error
}
