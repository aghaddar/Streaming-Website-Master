package models

import "time"

type Comment struct {
	CommentID   uint64    `gorm:"primaryKey;autoIncrement" json:"comment_id"`
	UserID      uint64    `gorm:"not null" json:"user_id"`
	EpisodeID   *string   `gorm:"size:255" json:"episode_id"`
	CommentText string    `gorm:"type:text;not null" json:"comment_text"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
