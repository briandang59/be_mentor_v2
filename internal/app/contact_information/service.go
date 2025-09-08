package contactinformation

type Service interface {
	Create(contactInformation *ContactInformation) (*ContactInformation, error)
	FindWithPagination(limit, offset int) ([]ContactInformation, int64, error)
	FindByID(id uint) (*ContactInformation, error)
	UpdateFields(id uint, fields map[string]interface{}) (*ContactInformation, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(contactInformation *ContactInformation) (*ContactInformation, error) {
	return s.repo.Create(contactInformation)
}

func (s *service) UpdateFields(id uint, fields map[string]interface{}) (*ContactInformation, error) {
	return s.repo.UpdateFields(id, fields)
}

func (s *service) FindWithPagination(limit, offset int) ([]ContactInformation, int64, error) {
	return s.repo.FindWithPagination(limit, offset)
}

func (s *service) FindByID(id uint) (*ContactInformation, error) {
	return s.repo.FindByID(id)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
