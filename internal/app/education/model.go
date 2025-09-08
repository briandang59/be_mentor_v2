package education

import (
	"time"

	"gorm.io/gorm"
)

type Education struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null;unique"`
	Object    string         `json:"object"`
	From      string         `json:"from"`
	To        string         `json:"to"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
