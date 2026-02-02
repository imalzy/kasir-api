package kategori

type CategoryService struct {
	repo *CategoryRepository
}

func NewCategoryService(repo *CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Create(data *Category) (string, error) {
	return s.repo.Create(data)
}

func (s *CategoryService) GetByID(id string) (*Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Update(id string, product *Category) error {
	return s.repo.Update(id, product)
}

func (s *CategoryService) Delete(id string) error {
	return s.repo.Delete(id)
}
