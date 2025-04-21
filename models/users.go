package models

import (
	"time"
)

type User struct {
	UserID       uint      `gorm:"primaryKey;autoIncrement"`
	Username     string    `gorm:"size:50;not null"`
	Email        string    `gorm:"size:100;not null"`
	AvatarURL    *string   `gorm:"type:text"`
	HashPassword string    `gorm:"size:255;not null"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp"`
}
