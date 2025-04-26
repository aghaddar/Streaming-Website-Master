package models

import "time"

type Comments struct {
	CommentID   uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null"`
	EpisodeID   *string   `gorm:"type:varchar(255)"`
	CommentText string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	User        User      `gorm:"foreignKey:UserID;references:UserID"`
	Episode     Episode   `gorm:"foreignKey:EpisodeID;references:EpisodeID"`
}
