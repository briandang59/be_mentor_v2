package tag

type Service interface {
	CreateTag(name string) (*Tag, error)
	GetAllTags() ([]Tag, error)
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
