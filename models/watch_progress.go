package models

import "time"

type Watch struct {
	WatchID     uint   `gorm:"primaryKey;autoIncrement"`
	UserID      uint   `gorm:"not null"`
	EpisodeID   string `gorm:"type:varchar(255);not null"`
	WatchedTime int
	Completed   bool
	UpdatedAt   time.Time `gorm:"default:current_timestamp"`

	User    User    `gorm:"foreignKey:UserID;references:UserID"` // ✅ Correct FK
	Episode Episode `gorm:"foreignKey:EpisodeID;references:EpisodeID"`
}
