package services

import (
	"Streaming-Website-Master/models"
	"errors"
	"gorm.io/gorm"
)

type RecommendationService struct {
	db *gorm.DB
}

func NewRecommendationService(db *gorm.DB) *RecommendationService {
	return &RecommendationService{db: db}
}

func (s *RecommendationService) CreateRecommendation(recommendation *models.Recommendation) error {
	return s.db.Create(recommendation).Error
}

func (s *RecommendationService) GetRecommendationByID(id uint) (*models.Recommendation, error) {
	var recommendation models.Recommendation
	if err := s.db.First(&recommendation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &recommendation, nil
}

func (s *RecommendationService) UpdateRecommendation(recommendation *models.Recommendation) error {
	return s.db.Save(recommendation).Error
}

func (s *RecommendationService) DeleteRecommendation(id uint) error {
	return s.db.Delete(&models.Recommendation{}, id).Error
}

func (s *RecommendationService) ListRecommendations() ([]models.Recommendation, error) {
	var recommendations []models.Recommendation
	if err := s.db.Find(&recommendations).Error; err != nil {
		return nil, err
	}
	return recommendations, nil
}
