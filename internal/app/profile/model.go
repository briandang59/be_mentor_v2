package profile

import (
	"mentors/internal/app/attachment"
	availabletime "mentors/internal/app/available_time"
	contactinformation "mentors/internal/app/contact_information"
	"mentors/internal/app/education"
	"mentors/internal/app/experience"
	"mentors/internal/app/language"
	"mentors/internal/app/portfolio"
	"mentors/internal/app/tag"
	"mentors/internal/app/user"
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID                   int64                                  `json:"id" gorm:"primaryKey"`
	Fullname             string                                 `json:"fullname" gorm:"type:varchar(255);not null"`
	Title                string                                 `json:"title" gorm:"type:varchar(255)"`
	ProfileOverview      string                                 `json:"profile_overview" gorm:"type:text"`
	HourlyRate           float64                                `json:"hourly_rate" gorm:"type:decimal(10,2)"`
	QCoin                float64                                `json:"q_coin" gorm:"type:decimal(10,2);default:0"`
	MaxHoursPerWeek      int                                    `json:"max_hours_per_week" gorm:"type:int; default:0"`
	Location             string                                 `json:"location" gorm:"type:varchar(255)"`
	ProfileImageID       uint                                   `json:"profile_image_id" gorm:"type:int"`
	EducationID          uint                                   `json:"education_id" gorm:"not null;uniqueIndex"`
	UserID               uint                                   `json:"user_id" gorm:"not null;uniqueIndex"`
	LanguageIDs          []uint                                 `json:"language_ids" gorm:"-"`
	PortfolioIDs         []uint                                 `json:"portfolio_ids" gorm:"-"`
	ExperienceIDs        []uint                                 `json:"experience_ids" gorm:"-"`
	SkillIDs             []uint                                 `json:"skill_ids" gorm:"-"`
	ContactInformationID uint                                   `json:"contact_information_id" gorm:"type:int"`
	AvailableTimeIDs     []uint                                 `json:"available_time_ids" gorm:"-"`
	AvailableTimes       []*availabletime.AvailableTime         `json:"available_times" gorm:"many2many:profile_available_times;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Skills               []*tag.Tag                             `json:"skills" gorm:"many2many:profile_skills;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Experiences          []*experience.Experience               `json:"experiences" gorm:"foreignKey:ProfileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Portfolios           []*portfolio.Portfolio                 `json:"portfolios" gorm:"foreignKey:ProfileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Languages            []*language.Language                   `json:"languages" gorm:"many2many:profile_languages;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ContactInformation   *contactinformation.ContactInformation `json:"contact_information" gorm:"foreignKey:ContactInformationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProfileImage         *attachment.Attachment                 `json:"profile_image" gorm:"foreignKey:ProfileImageID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Education            *education.Education                   `json:"education" gorm:"foreignKey:EducationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User                 *user.User                             `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt            time.Time                              `json:"created_at"`
	UpdatedAt            time.Time                              `json:"updated_at"`
	DeletedAt            gorm.DeletedAt                         `json:"-" gorm:"index"`
}
