package models

import "time"

// Book represents the permanent details of a book
type Book struct {
	// ID is the unique fingerprint (Primary Key)
	ID uint `json:"id" gorm:"primaryKey"`

	// Basic Details (From your Add Book Form)
	Title           string `json:"title" gorm:"not null"`
	Author          string `json:"author" gorm:"not null"`
	ISBN            string `json:"isbn"`
	Genre           string `json:"genre"`
	PublicationYear int    `json:"publication_year"`
	TotalPages      int    `json:"total_pages"`
	

	// Timestamps (System handles these)
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}