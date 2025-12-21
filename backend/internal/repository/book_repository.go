package repository

import (
	"readingtracker/internal/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	GetAllBooks() ([]models.Book, error) 
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) CreateBook(book *models.Book) error {
	return r.db.Create(book).Error
}


func (r *bookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book

	err := r.db.Find(&books).Error
	return books, err
}