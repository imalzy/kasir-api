package report

import (
	"errors"
	"log"
)

type ReportService struct {
	repo *ReportRepository
}

func NewReportService(repo *ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) Report(filter ReportFilter) (*ReportResponse, error) {
	// Check if dates are provided
	if filter.StartDate == "" || filter.EndDate == "" {
		return nil, errors.New("date range is required")
	}

	// Pass the object down to the repository
	report, err := s.repo.GetReportData(filter)
	log.Print(err)
	if err != nil {
		return nil, err
	}

	return report, nil
}
