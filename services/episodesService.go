package services

import (
	"Streaming-Website-Master/models"
	"errors"
	"gorm.io/gorm"
)

type EpisodeService struct {
	db *gorm.DB
}

func NewEpisodeService(db *gorm.DB) *EpisodeService {
	return &EpisodeService{db: db}
}

func (s *EpisodeService) CreateEpisode(episode *models.Episode) error {
	return s.db.Create(episode).Error
}

func (s *EpisodeService) GetEpisodeByID(id uint) (*models.Episode, error) {
	var episode models.Episode
	if err := s.db.First(&episode, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &episode, nil
}

func (s *EpisodeService) UpdateEpisode(episode *models.Episode) error {
	return s.db.Save(episode).Error
}

func (s *EpisodeService) DeleteEpisode(id uint) error {
	return s.db.Delete(&models.Episode{}, id).Error
}

func (s *EpisodeService) ListEpisodes() ([]models.Episode, error) {
	var episodes []models.Episode
	if err := s.db.Find(&episodes).Error; err != nil {
		return nil, err
	}
	return episodes, nil
}
