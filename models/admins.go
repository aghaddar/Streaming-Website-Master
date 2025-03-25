package models

import (
	"time"
)

type Admin struct {
	AdminID      uint      `gorm:"primaryKey;autoIncrement"`
	Username     string    `gorm:"size:50;not null"`
	Email        string    `gorm:"size:100;not null"`
	PasswordHash string    `gorm:"size:255;not null"`
	Role         string    `gorm:"size:20;default:'moderator'"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
}
