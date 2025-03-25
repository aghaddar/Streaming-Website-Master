package services

import (
	"Streaming-Website-Master/models"
	"errors"
	"gorm.io/gorm"
)

type AnimeService struct {
	db *gorm.DB
}

func NewAnimeService(db *gorm.DB) *AnimeService {
	return &AnimeService{db: db}
}

func (s *AnimeService) CreateAnime(anime *models.Anime) error {
	return s.db.Create(anime).Error
}

func (s *AnimeService) GetAnimeByID(id uint) (*models.Anime, error) {
	var anime models.Anime
	if err := s.db.First(&anime, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &anime, nil
}

func (s *AnimeService) UpdateAnime(anime *models.Anime) error {
	return s.db.Save(anime).Error
}

func (s *AnimeService) DeleteAnime(id uint) error {
	return s.db.Delete(&models.Anime{}, id).Error
}

func (s *AnimeService) ListAnimes() ([]models.Anime, error) {
	var animes []models.Anime
	if err := s.db.Find(&animes).Error; err != nil {
		return nil, err
	}
	return animes, nil
}
