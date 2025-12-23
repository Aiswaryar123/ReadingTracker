package services

import (
	"errors"
	"readingtracker/internal/dto"
	"readingtracker/internal/models"
	"testing"
)

type FakeBookRepoForReview struct {
	Exists bool
}

func (f *FakeBookRepoForReview) GetBookByID(id uint) (*models.Book, error) {
	if !f.Exists {
		return nil, errors.New("book not found")
	}
	return &models.Book{ID: id}, nil
}

func (f *FakeBookRepoForReview) CreateBook(b *models.Book) error     { return nil }
func (f *FakeBookRepoForReview) GetAllBooks() ([]models.Book, error) { return nil, nil }
func (f *FakeBookRepoForReview) UpdateBook(b *models.Book) error     { return nil }
func (f *FakeBookRepoForReview) DeleteBook(id uint) error            { return nil }
func (f *FakeBookRepoForReview) GetDashboardStats() (dto.DashboardStats, error) {
	return dto.DashboardStats{}, nil
}

type FakeReviewRepo struct {
	Reviews []models.Review
}

func (f *FakeReviewRepo) CreateReview(r *models.Review) error {
	f.Reviews = append(f.Reviews, *r)
	return nil
}

func (f *FakeReviewRepo) GetReviewsByBookID(bookID uint) ([]models.Review, error) {
	return f.Reviews, nil
}

func TestAddReview_BookNotFound(t *testing.T) {

	bookRepo := &FakeBookRepoForReview{Exists: false}
	reviewRepo := &FakeReviewRepo{}
	service := NewReviewService(reviewRepo, bookRepo)

	req := dto.CreateReviewRequest{Rating: 5, Comment: "Great book!"}

	err := service.AddReview(99, req)

	if err == nil {
		t.Errorf("Expected an error because book 99 does not exist, but got nil")
	}
}

func TestAddReview_Success(t *testing.T) {

	bookRepo := &FakeBookRepoForReview{Exists: true}
	reviewRepo := &FakeReviewRepo{}
	service := NewReviewService(reviewRepo, bookRepo)

	req := dto.CreateReviewRequest{Rating: 4, Comment: "Nice read"}

	err := service.AddReview(1, req)

	if err != nil {
		t.Errorf("Expected success, but got error: %v", err)
	}
	if len(reviewRepo.Reviews) != 1 {
		t.Errorf("Expected 1 review to be saved, but found %d", len(reviewRepo.Reviews))
	}
	if reviewRepo.Reviews[0].Rating != 4 {
		t.Errorf("Expected rating 4, but got %d", reviewRepo.Reviews[0].Rating)
	}
}

func TestGetBookReviews_Success(t *testing.T) {

	reviewRepo := &FakeReviewRepo{
		Reviews: []models.Review{{BookID: 1, Rating: 5, Comment: "Amazing"}},
	}
	bookRepo := &FakeBookRepoForReview{}
	service := NewReviewService(reviewRepo, bookRepo)

	reviews, err := service.GetBookReviews(1)

	if err != nil {
		t.Errorf("Expected success, but got error: %v", err)
	}
	if len(reviews) != 1 {
		t.Errorf("Expected 1 review, but got %d", len(reviews))
	}
}
