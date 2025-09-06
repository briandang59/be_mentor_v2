package coverletter

import (
	"mentors/internal/app/attachment"
	"mentors/internal/app/user"
	"time"

	"gorm.io/gorm"
)

type CoverLetter struct {
	ID           uint                   `json:"id" gorm:"primaryKey"`
	Content      string                 `json:"content" gorm:"not null"`
	AttachmentID uint                   `json:"attachment_id"`
	Attachment   *attachment.Attachment `json:"attachment"`
	PostID       uint                   `json:"post_id"`
	ApplierID    uint                   `json:"applier_id"`
	Applier      *user.User             `json:"applier"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	DeletedAt    gorm.DeletedAt         `json:"-" gorm:"index"`
}
