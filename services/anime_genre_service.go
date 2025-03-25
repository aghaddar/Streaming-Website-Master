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

func (s *AnimeGenreService) CreateAnimeGenre(ag *models.Anime_Genre) error {
	return s.db.Create(ag).Error
}

func (s *AnimeGenreService) GetAnimeGenre(animeID, genreID uint) (*models.Anime_Genre, error) {
	var ag models.Anime_Genre
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
	return s.db.Delete(&models.Anime_Genre{}, "anime_id = ? AND genre_id = ?", animeID, genreID).Error
}

func (s *AnimeGenreService) ListAnimeGenres() ([]models.Anime_Genre, error) {
	var ags []models.Anime_Genre
	if err := s.db.Find(&ags).Error; err != nil {
		return nil, err
	}
	return ags, nil
}
