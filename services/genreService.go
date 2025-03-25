package services

import (
	"Streaming-Website-Master/models"
	"errors"
	"gorm.io/gorm"
)

type GenreService struct {
	db *gorm.DB
}

func NewGenreService(db *gorm.DB) *GenreService {
	return &GenreService{db: db}
}

func (s *GenreService) CreateGenre(genre *models.Genre) error {
	return s.db.Create(genre).Error
}

func (s *GenreService) GetGenreByID(id uint) (*models.Genre, error) {
	var genre models.Genre
	if err := s.db.First(&genre, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &genre, nil
}

func (s *GenreService) UpdateGenre(genre *models.Genre) error {
	return s.db.Save(genre).Error
}

func (s *GenreService) DeleteGenre(id uint) error {
	return s.db.Delete(&models.Genre{}, id).Error
}

func (s *GenreService) ListGenres() ([]models.Genre, error) {
	var genres []models.Genre
	if err := s.db.Find(&genres).Error; err != nil {
		return nil, err
	}
	return genres, nil
}
