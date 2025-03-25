package models

import "time"

type Comments struct {
	CommentID   uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null;foreignKey:UserID"`
	EpisodeID   *uint     `gorm:"foreignKey:EpisodeID"`
	CommentText string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	User        User      `gorm:"foreignKey:UserID"`
	Episode     Episode   `gorm:"foreignKey:EpisodeID"`
}
