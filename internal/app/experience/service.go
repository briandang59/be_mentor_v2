package experience

type Service interface {
	CreateExperience(experience *Experience) error
	GetExperiencesWithPagination(offset, limit int) ([]Experience, int64, error)
	UpdateExperiencePartial(id uint, fields map[string]interface{}) (*Experience, error)
	DeleteExperience(experience *Experience) error
	GetExperienceByID(id uint) (*Experience, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateExperience(experience *Experience) error {
	return s.repo.Create(experience)
}

func (s *service) GetExperiencesWithPagination(offset, limit int) ([]Experience, int64, error) {
	return s.repo.FindWithPagination(offset, limit)
}

func (s *service) UpdateExperiencePartial(id uint, fields map[string]interface{}) (*Experience, error) {
	return s.repo.UpdateFields(id, fields)
}

func (s *service) DeleteExperience(experience *Experience) error {
	return s.repo.Delete(experience)
}

func (s *service) GetExperienceByID(id uint) (*Experience, error) {
	return s.repo.GetByID(id)
}
