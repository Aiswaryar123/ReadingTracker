package services

import (
	"errors"
	"readingtracker/internal/dto"
	"readingtracker/internal/models"
	"testing"
)

// dummy repo
type FakeBookRepo struct {
	Books []models.Book
	Err   error
}

func (f *FakeBookRepo) CreateBook(book *models.Book) error {
	if f.Err != nil {
		return f.Err
	}
	f.Books = append(f.Books, *book)
	return nil
}

func (f *FakeBookRepo) GetAllBooks() ([]models.Book, error) {
	if f.Err != nil {
		return nil, f.Err
	}
	return f.Books, nil
}

func (f *FakeBookRepo) GetBookByID(id uint) (*models.Book, error) {
	if f.Err != nil {
		return nil, f.Err
	}
	return &models.Book{ID: id}, nil
}

func (f *FakeBookRepo) UpdateBook(book *models.Book) error {
	return f.Err
}

func (f *FakeBookRepo) DeleteBook(id uint) error {
	return f.Err
}

func (f *FakeBookRepo) GetDashboardStats() (dto.DashboardStats, error) {
	return dto.DashboardStats{}, f.Err
}

// create
func TestCreateBook_Success(t *testing.T) {
	repo := &FakeBookRepo{}
	service := NewBookService(repo)
	req := dto.CreateBookRequest{Title: "Clean Code", Author: "Robert Martin"}

	book, err := service.CreateBook(req)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if book.Title != "Clean Code" {
		t.Errorf("Expected title 'Clean Code', but got: %s", book.Title)
	}
	if len(repo.Books) != 1 {
		t.Errorf("Expected 1 book in repo, but found %d", len(repo.Books))
	}
}

func TestCreateBook_DatabaseError(t *testing.T) {
	repo := &FakeBookRepo{Err: errors.New("db error")}
	service := NewBookService(repo)
	req := dto.CreateBookRequest{Title: "Fail Book"}

	_, err := service.CreateBook(req)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

// fetch
func TestFetchBooks_Success(t *testing.T) {
	repo := &FakeBookRepo{
		Books: []models.Book{
			{Title: "Book 1"},
			{Title: "Book 2"},
		},
	}
	service := NewBookService(repo)

	books, err := service.FetchBooks()

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if len(books) != 2 {
		t.Errorf("Expected 2 books, but got: %d", len(books))
	}
}

func TestFetchBooks_DatabaseError(t *testing.T) {
	repo := &FakeBookRepo{Err: errors.New("connection failed")}
	service := NewBookService(repo)

	_, err := service.FetchBooks()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

// update
func TestUpdateBook_Success(t *testing.T) {
	repo := &FakeBookRepo{}
	service := NewBookService(repo)
	req := dto.UpdateBookRequest{ID: 1, Title: "Updated Title"}

	err := service.UpdateBook(req)

	if err != nil {
		t.Errorf("Expected success, but got error: %v", err)
	}
}

// delete
func TestDeleteBook_Success(t *testing.T) {
	repo := &FakeBookRepo{}
	service := NewBookService(repo)

	err := service.DeleteBook(1)

	if err != nil {
		t.Errorf("Expected success, but got error: %v", err)
	}
}

// dashboard stats
func TestGetStats_Success(t *testing.T) {
	repo := &FakeBookRepo{}
	service := NewBookService(repo)

	_, err := service.GetStats()

	if err != nil {
		t.Errorf("Expected stats to work, but got error: %v", err)
	}
}
