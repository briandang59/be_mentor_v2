package user

import "gorm.io/gorm"

type Repository interface {
	Create(user *User) error
	FindByEmail(email string) (User, error)
	FindByID(id uint) (User, error)
	UpdatePassword(id uint, newHash string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}
func (r *repo) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repo) FindByEmail(email string) (User, error) {
	var u User
	err := r.db.Where("email = ?", email).First(&u).Error
	return u, err
}

func (r *repo) FindByID(id uint) (User, error) {
	var u User
	err := r.db.First(&u, id).Error
	return u, err
}
func (r *repo) UpdatePassword(id uint, newHash string) error {
	return r.db.Model(&User{}).Where("id = ?", id).Update("password", newHash).Error
}
