package services

import (
	"Streaming-Website-Master/models"
	"errors"
	"gorm.io/gorm"
)

type WatchProgressService struct {
	db *gorm.DB
}

func NewWatchProgressService(db *gorm.DB) *WatchProgressService {
	return &WatchProgressService{db: db}
}

func (s *WatchProgressService) CreateWatchProgress(watch *models.Watch) error {
	return s.db.Create(watch).Error
}

func (s *WatchProgressService) GetWatchProgressByID(id uint) (*models.Watch, error) {
	var watch models.Watch
	if err := s.db.First(&watch, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &watch, nil
}

func (s *WatchProgressService) UpdateWatchProgress(watch *models.Watch) error {
	return s.db.Save(watch).Error
}

func (s *WatchProgressService) DeleteWatchProgress(id uint) error {
	return s.db.Delete(&models.Watch{}, id).Error
}

func (s *WatchProgressService) ListWatchProgresses() ([]models.Watch, error) {
	var watches []models.Watch
	if err := s.db.Find(&watches).Error; err != nil {
		return nil, err
	}
	return watches, nil
}
