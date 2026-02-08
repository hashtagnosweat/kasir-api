package handlers

import (
	"encoding/json"
	"kasir-api/services"
	"net/http"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

// HandleReport godoc
// @Summary Get reports
// @Description Get today report or report by date range
// @Tags report
// @Accept  json
// @Produce  json
// @Param start_date query string false "Start date in YYYY-MM-DD format"
// @Param end_date query string false "End date in YYYY-MM-DD format"
// @Success 200 {object} models.TodayReport
// @Failure 400 {object} map[string]string
// @Failure 404 {string} string
// @Failure 500 {object} map[string]string
// @Router /api/report [get]
func (h *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.routeReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ReportHandler) routeReport(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch path {
	case "/api/report/today":
		h.GetTodayReport(w, r)
		return

	case "/api/report":
		startDate := r.URL.Query().Get("start_date")
		endDate := r.URL.Query().Get("end_date")

		if startDate == "" || endDate == "" {
			http.Error(w, "start_date and end_date are required", http.StatusBadRequest)
			return
		}

		h.GetDateRangeReport(w, r, startDate, endDate)
		return

	default:
		http.NotFound(w, r)
	}
}

// GetTodayReport godoc
// @Summary Get today's report
// @Description Returns total revenue, total transactions, and best-selling product for today
// @Tags report
// @Produce json
// @Success 200 {object} models.TodayReport
// @Failure 500 {object} map[string]string
// @Router /api/report/today [get]
func (h *ReportHandler) GetTodayReport(w http.ResponseWriter, r *http.Request) {
	report, err := h.service.GetTodayReport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

// GetDateRangeReport godoc
// @Summary Get report by date range
// @Description Returns total revenue, total transactions, and best-selling product between start_date and end_date
// @Tags report
// @Produce json
// @Param start query string true "Start date in YYYY-MM-DD format"
// @Param end query string true "End date in YYYY-MM-DD format"
// @Success 200 {object} models.TodayReport
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/report [get]
func (h *ReportHandler) GetDateRangeReport(
	w http.ResponseWriter,
	r *http.Request,
	start, end string,
) {
	report, err := h.service.GetReportByDateRange(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
