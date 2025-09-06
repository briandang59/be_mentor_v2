package post

import (
	"mentors/internal/dto"
	"mentors/internal/utils"
)

type Service interface {
	CreatePost(req dto.PostCreateRequest) (*Post, error)
	GetAllPosts() ([]Post, error)
	GetPostsWithPagination(limit, offset int, populates []string) ([]Post, int64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreatePost(req dto.PostCreateRequest) (*Post, error) {
	post := &Post{
		Title:   req.Title,
		Slug:    utils.GenerateSlug(req.Title),
		Content: req.Content,
	}

	if len(req.TagIDs) > 0 {
		tags, err := s.repo.FindTagsByIDs(req.TagIDs)
		if err != nil {
			return nil, err
		}
		post.Tags = tags
	}

	return s.repo.Create(post)
}

func (s *service) GetAllPosts() ([]Post, error) {
	return s.repo.FindAll()
}

func (s *service) GetPostsWithPagination(limit, offset int, populates []string) ([]Post, int64, error) {
	return s.repo.FindWithPagination(limit, offset, populates)
}
