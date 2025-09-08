package education

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(education *Education) (*Education, error)
	FindWithPagination(limit, offset int) ([]Education, int64, error)
	FindByID(id uint) (*Education, error)
	UpdateFields(id uint, fields map[string]interface{}) (*Education, error)
	Delete(id uint) error
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) Create(education *Education) (*Education, error) {
	if err := r.db.Create(education).Error; err != nil {
		return nil, err
	}
	return education, nil
}

func (r *repo) UpdateFields(id uint, fields map[string]interface{}) (*Education, error) {
	var education Education
	if err := r.db.Model(&education).Where("id = ?", id).Updates(fields).Error; err != nil {
		return nil, err
	}
	if err := r.db.First(&education, id).Error; err != nil {
		return nil, err
	}
	return &education, nil
}

func (r *repo) FindWithPagination(limit, offset int) ([]Education, int64, error) {
	var educations []Education
	var total int64

	if err := r.db.Model(&Education{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Limit(limit).Offset(offset).Find(&educations).Error; err != nil {
		return nil, 0, err
	}
	return educations, total, nil
}

func (r *repo) FindByID(id uint) (*Education, error) {
	var education Education
	if err := r.db.First(&education, id).Error; err != nil {
		return nil, err
	}
	return &education, nil
}

func (r *repo) Delete(id uint) error {
	if err := r.db.Delete(&Education{}, id).Error; err != nil {
		return err
	}
	return nil
}
