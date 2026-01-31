package produk

type ProductService struct {
	repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]ProductResponse, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(data *Product) (string, error) {
	return s.repo.Create(data)
}

func (s *ProductService) GetByID(id string) (*ProductResponse, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(id string, product *Product) error {
	return s.repo.Update(id, product)
}

func (s *ProductService) Delete(id string) error {
	return s.repo.Delete(id)
}
