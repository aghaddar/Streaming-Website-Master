package services

import (
	"Streaming-Website-Master/models"
	"errors"
	"gorm.io/gorm"
)

type ReportService struct {
	db *gorm.DB
}

func NewReportService(db *gorm.DB) *ReportService {
	return &ReportService{db: db}
}

func (s *ReportService) CreateReport(report *models.Report) error {
	return s.db.Create(report).Error
}

func (s *ReportService) GetReportByID(id uint) (*models.Report, error) {
	var report models.Report
	if err := s.db.First(&report, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &report, nil
}

func (s *ReportService) UpdateReport(report *models.Report) error {
	return s.db.Save(report).Error
}

func (s *ReportService) DeleteReport(id uint) error {
	return s.db.Delete(&models.Report{}, id).Error
}

func (s *ReportService) ListReports() ([]models.Report, error) {
	var reports []models.Report
	if err := s.db.Find(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}
