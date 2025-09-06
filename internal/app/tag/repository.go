package tag

import "gorm.io/gorm"

type Repository interface {
	Create(tag *Tag) (*Tag, error)
	FindAll() ([]Tag, error)
	FindWithPagination(limit, offset int) ([]Tag, int64, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) Create(tag *Tag) (*Tag, error) {
	if err := r.db.Create(tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}
func (r *repo) FindAll() ([]Tag, error) {
	var tags []Tag
	if err := r.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *repo) FindWithPagination(limit, offset int) ([]Tag, int64, error) {
	var tags []Tag
	var total int64

	if err := r.db.Model(&Tag{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Limit(limit).Offset(offset).Find(&tags).Error; err != nil {
		return nil, 0, err
	}

	return tags, total, nil
}
