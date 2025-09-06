package tag

import "gorm.io/gorm"

type Repository interface {
	Create(tag *Tag) (*Tag, error)
	FindAll() ([]Tag, error)
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
