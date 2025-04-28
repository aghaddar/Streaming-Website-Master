package models

import "time"

type Watchlist struct {
	WatchlistID        uint64    `gorm:"primaryKey;autoIncrement" json:"watchlist_id"`
	UserID             uint64    `gorm:"not null" json:"user_id"`
	AnimeID            string    `gorm:"size:255;not null" json:"anime_id"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
	Status             string    `gorm:"type:enum('Watching','Completed','On Hold','Dropped');default:Watching" json:"status"`
	Priority           string    `gorm:"type:enum('High','Medium','Low');default:Medium" json:"priority"`
	LastWatchedEpisode *string   `gorm:"size:255" json:"last_watched_episode"`
	ProgressPercentage float64   `gorm:"type:decimal(5,2);default:0.00" json:"progress_percentage"`
}
