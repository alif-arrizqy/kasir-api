package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
	"time"
)

type ReportService interface {
	GetReportToday() (*models.ReportData, error)
	GetReportByDateRange(start, end *time.Time) (*models.ReportData, error)
}

type reportService struct {
	repo repositories.ReportRepository
}

func NewReportService(repo repositories.ReportRepository) ReportService {
	return &reportService{repo: repo}
}

func (s *reportService) GetReportToday() (*models.ReportData, error) {
	return s.repo.GetReport(nil, nil)
}

func (s *reportService) GetReportByDateRange(start, end *time.Time) (*models.ReportData, error) {
	return s.repo.GetReport(start, end)
}
