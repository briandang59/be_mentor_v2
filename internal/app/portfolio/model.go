package portfolio

import (
	"mentors/internal/app/attachment"
	"time"

	"gorm.io/gorm"
)

type Portfolio struct {
	ID          int64                  `json:"id" gorm:"primaryKey"`
	Title       string                 `json:"title" gorm:"type:varchar(255);not null"`
	Description string                 `json:"description" gorm:"type:text"`
	URL         string                 `json:"url" gorm:"type:varchar(255)"`
	JobTitle    string                 `json:"job_title" gorm:"type:varchar(255)"`
	TimePeriod  string                 `json:"time_period" gorm:"type:varchar(255)"`
	IsDraft     bool                   `json:"is_draft" gorm:"default:true"`
	ThumbnailID uint                   `json:"thumbnail_id" gorm:"type:int"`
	Thumbnail   *attachment.Attachment `json:"thumbnail" gorm:"foreignKey:ThumbnailID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	DeletedAt   gorm.DeletedAt         `json:"-" gorm:"index"`
}
