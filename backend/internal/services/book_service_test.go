package services_test

import (
	"errors"
	"testing"

	"readingtracker/internal/dto"
	"readingtracker/internal/models"
	"readingtracker/internal/services"

	"github.com/stretchr/testify/assert"
)


// Mock Repository


type MockBookRepository struct {
	Books []models.Book
	Err   error
}

func (m *MockBookRepository) CreateBook(book *models.Book) error {
	if m.Err != nil {
		return m.Err
	}
	m.Books = append(m.Books, *book)
	return nil
}

func (m *MockBookRepository) GetAllBooks() ([]models.Book, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Books, nil
}
func (m *MockBookRepository) UpdateBook(book *models.Book) error {
	if m.Err != nil { return m.Err }
	return nil // Pretend it worked
}


func (m *MockBookRepository) DeleteBook(id uint) error { return nil }



// Create Book 


func TestCreateBook_Success(t *testing.T) {
	mockRepo := &MockBookRepository{}
	bookService := services.NewBookService(mockRepo)

	req := dto.CreateBookRequest{
		Title:           "Clean Code",
		Author:          "Robert C. Martin",
		Genre:           "Programming",
		PublicationYear: 2008,
		TotalPages:      464,
	}

	book, err := bookService.CreateBook(req)

	assert.NoError(t, err)
	assert.NotNil(t, book)
	assert.Equal(t, "Clean Code", book.Title)
	assert.Equal(t, 1, len(mockRepo.Books))
}

func TestCreateBook_Failure(t *testing.T) {
	mockRepo := &MockBookRepository{
		Err: errors.New("database error"),
	}
	bookService := services.NewBookService(mockRepo)

	req := dto.CreateBookRequest{
		Title:  "Fail Book",
		Author: "Unknown",
	}

	book, err := bookService.CreateBook(req)

	assert.Error(t, err)
	assert.Nil(t, book)
}

//fetch book

func TestFetchBooks_Success(t *testing.T) {
	mockRepo := &MockBookRepository{
		Books: []models.Book{
			{Title: "Book One", Author: "Author One"},
			{Title: "Book Two", Author: "Author Two"},
		},
	}

	bookService := services.NewBookService(mockRepo)

	books, err := bookService.FetchBooks()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(books))
	assert.Equal(t, "Book One", books[0].Title)
}

func TestFetchBooks_Failure(t *testing.T) {
	mockRepo := &MockBookRepository{
		Err: errors.New("fetch error"),
	}
	bookService := services.NewBookService(mockRepo)

	books, err := bookService.FetchBooks()

	assert.Error(t, err)
	assert.Nil(t, books)
}


//update test
func TestUpdateBook_Success(t *testing.T) {
	mockRepo := &MockBookRepository{}
	service := services.NewBookService(mockRepo)

	
	req := dto.UpdateBookRequest{
		ID:    1,
		Title: "Clean Architecture",
	}

	err := service.UpdateBook(req)

	assert.NoError(t, err)
}
//delete test
func TestDeleteBook_Success(t *testing.T) {
	mockRepo := &MockBookRepository{}
	service := services.NewBookService(mockRepo)

	// Try to delete Book #1
	err := service.DeleteBook(1)

	assert.NoError(t, err)
}