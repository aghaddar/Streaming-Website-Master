package models

import "time"

type User struct {
	UserID       uint64    `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username     string    `gorm:"size:50;not null" json:"username"`
	Email        string    `gorm:"size:100;not null" json:"email"`
	AvatarURL    *string   `json:"avatar_url"`
	HashPassword string    `gorm:"size:255;not null" json:"hash_password"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
