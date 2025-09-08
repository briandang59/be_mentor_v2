package portfolio

type Service interface {
	CreatePortfolio(portfolio *Portfolio) (*Portfolio, error)
	GetPortfoliosWithPagination(limit, offset int) ([]Portfolio, int64, error)
	UpdatePortfolioPartial(id uint, fields map[string]interface{}) (*Portfolio, error)
	DeletePortfolio(id uint) error
	GetPortfolioByID(id uint) (*Portfolio, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreatePortfolio(portfolio *Portfolio) (*Portfolio, error) {
	return s.repo.Create(portfolio)
}
func (s *service) GetPortfoliosWithPagination(limit, offset int) ([]Portfolio, int64, error) {
	return s.repo.FindWithPagination(limit, offset)
}
func (s *service) UpdatePortfolioPartial(id uint, fields map[string]interface{}) (*Portfolio, error) {
	return s.repo.UpdateFields(id, fields)
}
func (s *service) DeletePortfolio(id uint) error {
	return s.repo.Delete(id)
}
func (s *service) GetPortfolioByID(id uint) (*Portfolio, error) {
	return s.repo.FindByID(id)
}
