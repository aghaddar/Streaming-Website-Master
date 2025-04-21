package models

import "time"

type Watch struct {
	WatchID     uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null"`
	EpisodeID   string    `gorm:"type:text;not null"`
	WatchedTime *uint     `gorm:"default:NULL"`
	Completed   bool      `gorm:"default:0"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp"`

	User    User    `gorm:"foreignKey:UserID"`
	Episode Episode `gorm:"foreignKey:EpisodeID;references:EpisodeID"`
}
