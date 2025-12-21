package repository

import (
	"readingtracker/internal/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	GetAllBooks() ([]models.Book, error)

	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
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

func (r *bookRepository) UpdateBook(book *models.Book) error {
	return r.db.Model(book).Updates(book).Error
}

func (r *bookRepository) DeleteBook(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}
