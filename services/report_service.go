package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetTodayReport() (*models.TodayReport, error) {
	return s.repo.GetTodayReport()
}

func (s *ReportService) GetReportByDateRange(startDate, endDate string) (*models.TodayReport, error) {
	return s.repo.GetReportByDateRange(startDate, endDate)
}
