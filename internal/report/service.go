package report

type ReportService struct {
	repo *ReportRepository
}

func NewProductService(repo *ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}
