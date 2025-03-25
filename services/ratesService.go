package services

import (
	"Streaming-Website-Master/models"
	"errors"
	"gorm.io/gorm"
)

type RatingService struct {
	db *gorm.DB
}

func NewRatingService(db *gorm.DB) *RatingService {
	return &RatingService{db: db}
}

func (s *RatingService) CreateRating(rating *models.Rating) error {
	return s.db.Create(rating).Error
}

func (s *RatingService) GetRatingByID(id uint) (*models.Rating, error) {
	var rating models.Rating
	if err := s.db.First(&rating, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &rating, nil
}

func (s *RatingService) UpdateRating(rating *models.Rating) error {
	return s.db.Save(rating).Error
}

func (s *RatingService) DeleteRating(id uint) error {
	return s.db.Delete(&models.Rating{}, id).Error
}

func (s *RatingService) ListRatings() ([]models.Rating, error) {
	var ratings []models.Rating
	if err := s.db.Find(&ratings).Error; err != nil {
		return nil, err
	}
	return ratings, nil
}
