package availabletime

import (
	"mentors/internal/enum"
	"time"

	"gorm.io/gorm"
)

type AvailableTime struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	Day       *enum.Day      `json:"day" gorm:"type:varchar(20);not null;uniqueIndex:idx_day_time"`
	From      string         `json:"from" gorm:"type:varchar(20);not null;uniqueIndex:idx_day_time"`
	To        string         `json:"to" gorm:"type:varchar(20);not null;uniqueIndex:idx_day_time"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
