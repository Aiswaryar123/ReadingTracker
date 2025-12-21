package models

import "time"

type ReadingProgress struct {
	ID     uint `json:"id" gorm:"primaryKey"`

	// Link to the Book (Foreign Key)
	BookID uint `json:"book_id" gorm:"not null"`
	Book   Book `json:"-" gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;"`

	// THE BOOKMARK 🔖
	// This stores "Page 45". Next time you open the app, 
	// we query this field to put you back on Page 45.
	CurrentPage int `json:"current_page" gorm:"default:0"`

	// READING STATUS
	// Helps us filter: "Show me only books I am currently reading"
	Status string `json:"status" gorm:"default:'Want to Read'"` // 'Reading', 'Completed'

	// RESUME TIME 🕒
	// "autoUpdateTime" is crucial. It updates automatically every time 
	// you change the page number. This lets us show "Recently Read" books first.
	LastUpdated time.Time `json:"last_updated" gorm:"autoUpdateTime"`
}