package post

import (
	"mentors/internal/app/tag"

	"gorm.io/gorm"
)

type Repository interface {
	Create(post *Post) (*Post, error)
	FindTagsByIDs(ids []uint) ([]*tag.Tag, error)
	FindAll() ([]Post, error)
	FindWithPagination(limit, offset int, populates []string) ([]Post, int64, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) Create(post *Post) (*Post, error) {
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (r *repo) FindTagsByIDs(ids []uint) ([]*tag.Tag, error) {
	var tags []*tag.Tag
	if err := r.db.Find(&tags, ids).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *repo) FindAll() ([]Post, error) {
	var posts []Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *repo) FindWithPagination(limit, offset int, populates []string) ([]Post, int64, error) {
	var posts []Post
	var total int64

	r.db.Model(&Post{}).Count(&total)

	query := r.db.Model(&Post{})
	for _, p := range populates {
		query = query.Preload(p)
	}

	err := query.Limit(limit).Offset(offset).Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}
	return posts, total, nil
}
