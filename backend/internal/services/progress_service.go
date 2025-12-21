package services

import (
	"readingtracker/internal/dto"
	"readingtracker/internal/models"
	"readingtracker/internal/repository"
)

type ProgressService interface {
	UpdateProgress(bookID uint, req dto.UpdateProgressRequest) error
}

type progressService struct {
	repo repository.ProgressRepository
}

func NewProgressService(repo repository.ProgressRepository) ProgressService {
	return &progressService{repo: repo}
}

func (s *progressService) UpdateProgress(bookID uint, req dto.UpdateProgressRequest) error {

	progress, err := s.repo.GetByBookID(bookID)

	if err != nil {

		progress = &models.ReadingProgress{
			BookID: bookID,
		}
	}

	progress.Status = req.Status
	progress.CurrentPage = req.CurrentPage

	return s.repo.Save(progress)
}
