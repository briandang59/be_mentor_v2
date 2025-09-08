package language

import (
	"time"

	"gorm.io/gorm"
)

type Language struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null;uniqueIndex"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
