package tag

type Service interface {
	CreateTag(name string) (*Tag, error)
	GetAllTags() ([]Tag, error)
	GetTagsWithPagination(limit, offset int) ([]Tag, int64, error)
	UpdateTagPartial(id uint, fields map[string]interface{}) (*Tag, error)
	DeleteTag(id uint) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateTag(name string) (*Tag, error) {
	tag := &Tag{TagName: name}
	return s.repo.Create(tag)
}
func (s *service) GetAllTags() ([]Tag, error) {
	return s.repo.FindAll()
}

func (s *service) GetTagsWithPagination(limit, offset int) ([]Tag, int64, error) {
	return s.repo.FindWithPagination(limit, offset)
}

func (s *service) UpdateTagPartial(id uint, fields map[string]interface{}) (*Tag, error) {
	return s.repo.UpdateFields(id, fields)
}

func (s *service) DeleteTag(id uint) error {
	return s.repo.Delete(id)
}
