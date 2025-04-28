package models

import "time"

type Watch struct {
	WatchID     uint64    `gorm:"primaryKey;autoIncrement" json:"watch_id"`
	UserID      uint64    `gorm:"not null" json:"user_id"`
	EpisodeID   string    `gorm:"size:255;not null" json:"episode_id"`
	WatchedTime *int      `json:"watched_time"`
	Completed   bool      `gorm:"default:false" json:"completed"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
