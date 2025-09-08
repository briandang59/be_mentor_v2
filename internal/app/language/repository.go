package language

import "gorm.io/gorm"

type Repository interface {
	Create(language *Language) (*Language, error)
	FindAll() ([]Language, error)
	FindWithPagination(limit, offset int) ([]Language, int64, error)
	FindByID(id uint) (*Language, error)
	UpdateFields(id uint, fields map[string]interface{}) (*Language, error)
	Delete(id uint) error
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) Create(language *Language) (*Language, error) {
	if err := r.db.Create(language).Error; err != nil {
		return nil, err
	}
	return language, nil
}

func (r *repo) FindAll() ([]Language, error) {
	var languages []Language
	if err := r.db.Find(&languages).Error; err != nil {
		return nil, err
	}
	return languages, nil
}

func (r *repo) FindWithPagination(limit, offset int) ([]Language, int64, error) {
	var languages []Language
	var total int64
	if err := r.db.Model(&Language{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Limit(limit).Offset(offset).Find(&languages).Error; err != nil {
		return nil, 0, err
	}
	return languages, total, nil
}

func (r *repo) FindByID(id uint) (*Language, error) {
	var language Language
	if err := r.db.First(&language, id).Error; err != nil {
		return nil, err
	}
	return &language, nil
}

func (r *repo) UpdateFields(id uint, fields map[string]interface{}) (*Language, error) {
	var language Language
	if err := r.db.Model(&language).Where("id = ?", id).Updates(fields).Error; err != nil {
		return nil, err
	}
	return &language, nil
}

func (r *repo) Delete(id uint) error {
	if err := r.db.Delete(&Language{}, id).Error; err != nil {
		return err
	}
	return nil
}
