package services

import (
	"readingtracker/internal/dto"
	"readingtracker/internal/models"
	"readingtracker/internal/repository"
)


type BookService interface {
	CreateBook(req dto.CreateBookRequest) (*models.Book, error)
	FetchBooks() ([]models.Book, error)
	UpdateBook(req dto.UpdateBookRequest) error
	DeleteBook(id uint) error 
}


type bookService struct {
	repo repository.BookRepository
}


func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}



func (s *bookService) CreateBook(req dto.CreateBookRequest) (*models.Book, error) {
	
	book := &models.Book{
		Title:           req.Title,
		Author:          req.Author,
		ISBN:            req.ISBN,
		Genre:           req.Genre,
		PublicationYear: req.PublicationYear,
		TotalPages:      req.TotalPages,
	}

	err := s.repo.CreateBook(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *bookService) FetchBooks() ([]models.Book, error) {
	return s.repo.GetAllBooks()
}

func (s *bookService) UpdateBook(req dto.UpdateBookRequest) error {
	book := &models.Book{
		ID:     req.ID,
		Title:  req.Title,
		Author: req.Author,
	}
	return s.repo.UpdateBook(book)
}

func (s *bookService) DeleteBook(id uint) error {
	return s.repo.DeleteBook(id)
}