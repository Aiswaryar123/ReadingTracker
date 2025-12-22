package models

import "time"

type ReadingGoal struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Year        int       `json:"year" gorm:"not null;unique"` // One goal per year
	TargetBooks int       `json:"target_books" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
