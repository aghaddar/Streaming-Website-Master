package services

import (
	"Streaming-Website-Master/models"
	"errors"
	"gorm.io/gorm"
)

type WatchlistService struct {
	db *gorm.DB
}

func NewWatchlistService(db *gorm.DB) *WatchlistService {
	return &WatchlistService{db: db}
}

func (s *WatchlistService) CreateWatchlist(watchlist *models.Watchlist) error {
	return s.db.Create(watchlist).Error
}

func (s *WatchlistService) GetWatchlistByID(id uint) (*models.Watchlist, error) {
	var watchlist models.Watchlist
	if err := s.db.First(&watchlist, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &watchlist, nil
}

func (s *WatchlistService) UpdateWatchlist(watchlist *models.Watchlist) error {
	return s.db.Save(watchlist).Error
}

func (s *WatchlistService) DeleteWatchlist(id uint) error {
	return s.db.Delete(&models.Watchlist{}, id).Error
}

func (s *WatchlistService) ListWatchlists() ([]models.Watchlist, error) {
	var watchlists []models.Watchlist
	if err := s.db.Find(&watchlists).Error; err != nil {
		return nil, err
	}
	return watchlists, nil
}
