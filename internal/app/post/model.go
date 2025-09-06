package post

import (
	"mentors/internal/app/tag"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null"`
	Slug      string         `json:"slug"`
	Content   string         `json:"content" gorm:"not null"`
	Tags      []*tag.Tag     `json:"tags,omitempty" gorm:"many2many:post_tags;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
