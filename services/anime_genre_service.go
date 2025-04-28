package services

import (
	"Streaming-Website-Master/models"
	"errors"
	"gorm.io/gorm"
)

type AnimeGenreService struct {
	db *gorm.DB
}

func NewAnimeGenreService(db *gorm.DB) *AnimeGenreService {
	return &AnimeGenreService{db: db}
}

func (s *AnimeGenreService) CreateAnimeGenre(ag *models.AnimeGenre) error {
	return s.db.Create(ag).Error
}

func (s *AnimeGenreService) GetAnimeGenre(animeID, genreID uint) (*models.AnimeGenre, error) {
	var ag models.AnimeGenre
	if err := s.db.
		Where("anime_id = ? AND genre_id = ?", animeID, genreID).
		First(&ag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &ag, nil
}

func (s *AnimeGenreService) DeleteAnimeGenre(animeID, genreID uint) error {
	return s.db.Delete(&models.AnimeGenre{}, "anime_id = ? AND genre_id = ?", animeID, genreID).Error
}

func (s *AnimeGenreService) ListAnimeGenres() ([]models.AnimeGenre, error) {
	var ags []models.AnimeGenre
	if err := s.db.Find(&ags).Error; err != nil {
		return nil, err
	}
	return ags, nil
}
