package services

import (
	"errors"
	"readingtracker/internal/dto"
	"readingtracker/internal/models"
	"readingtracker/internal/repository"
)

type ReviewService interface {
	AddReview(bookID uint, req dto.CreateReviewRequest) error
	GetBookReviews(bookID uint) ([]models.Review, error)
}

type reviewService struct {
	repo     repository.ReviewRepository
	bookRepo repository.BookRepository
}

func NewReviewService(repo repository.ReviewRepository, bookRepo repository.BookRepository) ReviewService {
	return &reviewService{
		repo:     repo,
		bookRepo: bookRepo,
	}
}

func (s *reviewService) AddReview(bookID uint, req dto.CreateReviewRequest) error {

	_, err := s.bookRepo.GetBookByID(bookID)
	if err != nil {
		return errors.New("cannot review: book not found")
	}

	review := &models.Review{
		BookID:  bookID,
		Rating:  req.Rating,
		Comment: req.Comment,
	}

	return s.repo.CreateReview(review)
}

func (s *reviewService) GetBookReviews(bookID uint) ([]models.Review, error) {
	return s.repo.GetReviewsByBookID(bookID)
}
