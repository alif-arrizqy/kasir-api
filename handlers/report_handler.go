package handlers

import (
	"kasir-api/services"
	"kasir-api/utils"
	"net/http"
	"time"
)

type ReportHandler struct {
	service services.ReportService
}

func NewReportHandler(service services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

// GetReportToday handles GET /api/report/hari-ini
func (h *ReportHandler) GetReportToday(w http.ResponseWriter, r *http.Request) {
	report, err := h.service.GetReportToday()
	if err != nil {
		utils.ErrorResponse(w, "Failed to get report", http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, "OK", report, http.StatusOK)
}

// GetReportByDateRange handles GET /api/report with optional/required query start_date,end_date (YYYY-MM-DD)
func (h *ReportHandler) GetReportByDateRange(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	startStr := q.Get("start_date")
	endStr := q.Get("end_date")

	if startStr == "" || endStr == "" {
		utils.ErrorResponse(w, "start_date and end_date query parameters are required (format: YYYY-MM-DD)", http.StatusBadRequest)
		return
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	layout := "2006-01-02"
	s, err1 := time.ParseInLocation(layout, startStr, loc)
	e, err2 := time.ParseInLocation(layout, endStr, loc)
	if err1 != nil || err2 != nil {
		utils.ErrorResponse(w, "invalid date format, use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	// make end exclusive by adding one day
	endExclusive := e.Add(24 * time.Hour)

	report, err := h.service.GetReportByDateRange(&s, &endExclusive)
	if err != nil {
		utils.ErrorResponse(w, "Failed to get report", http.StatusInternalServerError)
		return
	}

	utils.SuccessResponse(w, "OK", report, http.StatusOK)
}
