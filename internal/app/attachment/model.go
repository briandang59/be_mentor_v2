package attachment

import (
	"time"

	"gorm.io/gorm"
)

type Attachment struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	FileName     string         `json:"file_name" gorm:"not null"`
	Url          string         `json:"url" gorm:"not null"`
	MimeType     string         `json:"mime_type"`
	Size         int64          `json:"size"`
	PublicID     string         `json:"public_id" gorm:"not null"`
	ResourceType string         `json:"resource_type"`
	Folder       string         `json:"folder"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
