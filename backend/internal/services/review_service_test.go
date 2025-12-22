package services

import (
	"errors"
	"readingtracker/internal/dto"
	"readingtracker/internal/models"
	"testing"
)

type MockBookRepoForReview struct {
	shouldFound bool
}

func (m *MockBookRepoForReview) GetBookByID(id uint) (*models.Book, error) {
	if m.shouldFound {
		return &models.Book{ID: id}, nil
	}
	return nil, errors.New("book not found")
}

func (m *MockBookRepoForReview) CreateBook(book *models.Book) error  { return nil }
func (m *MockBookRepoForReview) GetAllBooks() ([]models.Book, error) { return nil, nil }
func (m *MockBookRepoForReview) UpdateBook(book *models.Book) error  { return nil }
func (m *MockBookRepoForReview) DeleteBook(id uint) error            { return nil }

type MockReviewRepo struct{}

func (m *MockReviewRepo) CreateReview(review *models.Review) error                { return nil }
func (m *MockReviewRepo) GetReviewsByBookID(bookID uint) ([]models.Review, error) { return nil, nil }

func TestAddReview(t *testing.T) {
	t.Run("Should fail if book does not exist", func(t *testing.T) {

		mockBookRepo := &MockBookRepoForReview{shouldFound: false}
		mockReviewRepo := &MockReviewRepo{}
		service := NewReviewService(mockReviewRepo, mockBookRepo)

		req := dto.CreateReviewRequest{Rating: 5, Comment: "Great!"}
		err := service.AddReview(999, req)

		if err == nil {
			t.Errorf("Expected error when book is missing, but got nil")
		}
	})

	t.Run("Should succeed if book exists", func(t *testing.T) {

		mockBookRepo := &MockBookRepoForReview{shouldFound: true}
		mockReviewRepo := &MockReviewRepo{}
		service := NewReviewService(mockReviewRepo, mockBookRepo)

		req := dto.CreateReviewRequest{Rating: 5, Comment: "Great!"}
		err := service.AddReview(1, req)

		if err != nil {
			t.Errorf("Expected success, but got error: %v", err)
		}
	})
}
