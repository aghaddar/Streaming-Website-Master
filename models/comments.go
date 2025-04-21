package models

import "time"

type Comments struct {
	CommentID   uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null"`
	EpisodeID   *string   `gorm:"type:text"`
	CommentText string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	User        User      `gorm:"foreignKey:UserID"`
	Episode     Episode   `gorm:"foreignKey:EpisodeID;references:EpisodeID"`
}
