package education

type Service interface {
	CreateEducation(education *Education) (*Education, error)
	GetEducationsWithPagination(limit, offset int) ([]Education, int64, error)
	UpdateEducationPartial(id uint, fields map[string]interface{}) (*Education, error)
	DeleteEducation(id uint) error
	GetEducationByID(id uint) (*Education, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateEducation(education *Education) (*Education, error) {
	return s.repo.Create(education)
}

func (s *service) GetEducationsWithPagination(limit, offset int) ([]Education, int64, error) {
	return s.repo.FindWithPagination(limit, offset)
}

func (s *service) UpdateEducationPartial(id uint, fields map[string]interface{}) (*Education, error) {
	return s.repo.UpdateFields(id, fields)
}

func (s *service) DeleteEducation(id uint) error {
	return s.repo.Delete(id)
}
func (s *service) GetEducationByID(id uint) (*Education, error) {
	return s.repo.FindByID(id)
}
