package portfolio

import "gorm.io/gorm"

type Repository interface {
	Create(portfolio *Portfolio) (*Portfolio, error)
	FindWithPagination(limit, offset int) ([]Portfolio, int64, error)
	UpdateFields(id uint, fields map[string]interface{}) (*Portfolio, error)
	Delete(id uint) error
	FindByID(id uint) (*Portfolio, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) Create(portfolio *Portfolio) (*Portfolio, error) {
	if err := r.db.Create(portfolio).Error; err != nil {
		return nil, err
	}
	return portfolio, nil
}

func (r *repo) FindWithPagination(limit, offset int) ([]Portfolio, int64, error) {
	var portfolios []Portfolio
	var total int64
	if err := r.db.Model(&Portfolio{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Preload("Thumbnail").Limit(limit).Offset(offset).Find(&portfolios).Error; err != nil {
		return nil, 0, err
	}
	return portfolios, total, nil
}

func (r *repo) UpdateFields(id uint, fields map[string]interface{}) (*Portfolio, error) {
	var portfolio Portfolio
	if err := r.db.Model(&portfolio).Where("id = ?", id).Updates(fields).Error; err != nil {
		return nil, err
	}
	if err := r.db.Preload("Thumbnail").First(&portfolio, id).Error; err != nil {
		return nil, err
	}

	return &portfolio, nil
}

func (r *repo) Delete(id uint) error {
	if err := r.db.Delete(&Portfolio{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *repo) FindByID(id uint) (*Portfolio, error) {
	var portfolio Portfolio
	if err := r.db.Preload("Thumbnail").First(&portfolio, id).Error; err != nil {
		return nil, err
	}
	return &portfolio, nil
}
