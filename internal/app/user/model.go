package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username" gorm:"unique"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"-"`
	IsVerified bool   `json:"is_verified" gorm:"default:false"`
}
