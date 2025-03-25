package models

import "time"

type Watch struct {
	WatchID     uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null;foreignKey:UserID"`
	EpisodeID   uint      `gorm:"not null;foreignKey:EpisodeID"`
	WatchedTime *uint     `gorm:"default:NULL"`
	Completed   bool      `gorm:"default:0"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp"`

	User    User    `gorm:"foreignKey:UserID"`
	Episode Episode `gorm:"foreignKey:EpisodeID"`
}
