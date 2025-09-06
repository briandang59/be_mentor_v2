package contract

import (
	"mentors/internal/enum"
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	ID         uint                `json:"id" gorm:"primaryKey"`
	MentorID   uint                `json:"mentor_id"`
	MenteeID   uint                `json:"mentee_id"`
	HourlyRate float64             `json:"hourly_rate"`
	TotalHours float64             `json:"total_hours"`
	Status     enum.ContractStatus `json:"status" gorm:"type:varchar(20);not null;default:PENDING"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
	DeletedAt  gorm.DeletedAt      `json:"-" gorm:"index"`
}
