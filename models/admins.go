package models

import "time"

type Admin struct {
	AdminID      uint64    `gorm:"primaryKey;autoIncrement" json:"admin_id"`
	Username     string    `gorm:"size:50;not null" json:"username"`
	Email        string    `gorm:"size:100;not null" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"password_hash"`
	Role         string    `gorm:"size:20;default:moderator" json:"role"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
