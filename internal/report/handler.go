package report

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ReportHandler struct {
	service *ReportService
}

func NewReportHandler(service *ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Report(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ReportHandler) Report(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	filter := ReportFilter{
		StartDate: startDate,
		EndDate:   endDate,
	}

	report, err := h.service.Report(filter)
	if err != nil {
		// Basic error mapping
		if strings.Contains(err.Error(), "no data") || strings.Contains(err.Error(), "required") {
			h.respondWithError(w, http.StatusBadRequest, err.Error())
		} else {
			h.respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	// 4. Send successful response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Use 200 for reports
	json.NewEncoder(w).Encode(report)
}

func (h *ReportHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
