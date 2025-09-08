package language

type Service interface {
	CreateLanguage(language *Language) (*Language, error)
	GetLanguagesWithPagination(limit, offset int) ([]Language, int64, error)
	UpdateLanguagePartial(id uint, fields map[string]interface{}) (*Language, error)
	DeleteLanguage(id uint) error
	GetLanguageByID(id uint) (*Language, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateLanguage(language *Language) (*Language, error) {
	return s.repo.Create(language)
}

func (s *service) GetLanguagesWithPagination(limit, offset int) ([]Language, int64, error) {
	return s.repo.FindWithPagination(limit, offset)
}

func (s *service) UpdateLanguagePartial(id uint, fields map[string]interface{}) (*Language, error) {
	return s.repo.UpdateFields(id, fields)
}

func (s *service) DeleteLanguage(id uint) error {
	return s.repo.Delete(id)
}
func (s *service) GetLanguageByID(id uint) (*Language, error) {
	return s.repo.FindByID(id)
}
