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

// GetUserWatchlist retrieves all watchlist items for a specific user.
func (s *WatchlistService) GetUserWatchlist(userID uint64) ([]models.Watchlist, error) {
	var watchlists []models.Watchlist
	if err := s.db.Where("user_id = ?", userID).Find(&watchlists).Error; err != nil {
		return nil, err
	}
	return watchlists, nil
}

// AddAnimeToWatchlist creates a new watchlist record for the user if it doesn't exist.
func (s *WatchlistService) AddAnimeToWatchlist(watchlist *models.Watchlist) error {
	// Check if the anime already exists in the user's watchlist.
	var existing models.Watchlist
	err := s.db.Where("user_id = ? AND anime_id = ?", watchlist.UserID, watchlist.AnimeID).First(&existing).Error
	if err == nil {
		return errors.New("anime already in watchlist")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.db.Create(watchlist).Error
}

// RemoveAnimeFromWatchlist deletes the anime from the user's watchlist.
func (s *WatchlistService) RemoveAnimeFromWatchlist(userID uint64, animeID string) error {
	result := s.db.Where("user_id = ? AND anime_id = ?", userID, animeID).Delete(&models.Watchlist{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("anime not found in watchlist")
	}
	return nil
}

// CheckAnimeInWatchlist checks if an anime exists in the user's watchlist.
func (s *WatchlistService) CheckAnimeInWatchlist(userID uint64, animeID string) (bool, error) {
	var count int64
	err := s.db.Model(&models.Watchlist{}).
		Where("user_id = ? AND anime_id = ?", userID, animeID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
