package tag

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	TagName   string         `json:"tag_name" gorm:"not null"`
	PostID    uint           `json:"post_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
