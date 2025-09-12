package experience

import "gorm.io/gorm"

type Repository interface {
	FindWithPagination(offset, limit int) ([]Experience, int64, error)
	GetByID(id uint) (*Experience, error)
	Create(experience *Experience) error
	UpdateFields(id uint, fields map[string]interface{}) (*Experience, error)
	Delete(experience *Experience) error
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) FindWithPagination(offset, limit int) ([]Experience, int64, error) {
	var experiences []Experience
	var total int64
	if err := r.db.Model(&Experience{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Offset(offset).Limit(limit).Find(&experiences).Error; err != nil {
		return nil, 0, err
	}
	return experiences, total, nil
}

func (r *repo) GetByID(id uint) (*Experience, error) {
	var experience Experience
	if err := r.db.First(&experience, id).Error; err != nil {
		return nil, err
	}
	return &experience, nil
}

func (r *repo) Create(experience *Experience) error {
	return r.db.Create(experience).Error
}

func (r *repo) UpdateFields(id uint, fields map[string]interface{}) (*Experience, error) {
	var experience Experience
	if err := r.db.Model(&experience).Where("id = ?", id).Updates(fields).Error; err != nil {
		return nil, err
	}
	return &experience, nil
}
func (r *repo) Delete(experience *Experience) error {
	return r.db.Delete(experience).Error
}
