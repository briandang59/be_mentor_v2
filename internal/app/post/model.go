package post

import (
	coverletter "mentors/internal/app/cover_letter"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID           uint                       `json:"id" gorm:"primaryKey"`
	Title        string                     `json:"title" gorm:"not null"`
	Slug         string                     `json:"slug"`
	Content      string                     `json:"content" gorm:"not null"`
	CoverLetters []*coverletter.CoverLetter `json:"cover_letters,omitempty" gorm:"foreignKey:PostID"`
	CreatedAt    time.Time                  `json:"created_at"`
	UpdatedAt    time.Time                  `json:"updated_at"`
	DeletedAt    gorm.DeletedAt             `json:"-" gorm:"index"`
}
