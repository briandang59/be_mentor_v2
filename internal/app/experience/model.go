package experience

import (
	"time"

	"gorm.io/gorm"
)

type Experience struct {
	ID          int64          `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null"`
	Company     string         `json:"company" gorm:"type:varchar(255)"`
	Location    string         `json:"location" gorm:"type:varchar(255)"`
	From        string         `json:"from" gorm:"type:varchar(255)"`
	To          string         `json:"to" gorm:"type:varchar(255)"`
	Description string         `json:"description" gorm:"type:text"`
	IsCurrent   bool           `json:"is_current" gorm:"default:false"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
