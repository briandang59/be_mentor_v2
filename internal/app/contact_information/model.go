package contactinformation

import (
	"time"

	"gorm.io/gorm"
)

type ContactInformation struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	Phone     string         `json:"phone" gorm:"type:varchar(20)"`
	LinkedIn  string         `json:"linkedin" gorm:"type:varchar(255)"`
	Facebook  string         `json:"facebook" gorm:"type:varchar(255)"`
	Zalo      string         `json:"zalo" gorm:"type:varchar(255)"`
	Telegram  string         `json:"telegram" gorm:"type:varchar(255)"`
	X         string         `json:"x" gorm:"type:varchar(255)"`
	GitHub    string         `json:"github" gorm:"type:varchar(255)"`
	Twitter   string         `json:"twitter" gorm:"type:varchar(255)"`
	Website   string         `json:"website" gorm:"type:varchar(255)"`
	Location  string         `json:"location" gorm:"type:varchar(255)"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
