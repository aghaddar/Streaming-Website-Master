package models

import "time"

type Watchlist struct {
	WatchlistID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID             uint      `gorm:"not null;foreignKey:UserID"`
	AnimeID            uint      `gorm:"not null;foreignKey:AnimeID"`
	CreatedAt          time.Time `gorm:"default:current_timestamp"`
	Status             string    `gorm:"type:enum('Watching', 'Completed', 'On Hold', 'Dropped');default:'Watching'"`
	Priority           string    `gorm:"type:enum('High', 'Medium', 'Low');default:'Medium'"`
	LastWatchedEpisode *uint     `gorm:"default:NULL"`
	ProgressPercentage float64   `gorm:"type:decimal(5,2);default:0.00"`

	User  User  `gorm:"foreignKey:UserID"`
	Anime Anime `gorm:"foreignKey:AnimeID"`
}
