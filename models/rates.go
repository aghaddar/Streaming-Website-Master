package models

import (
	"time"
)

type Rating struct {
	RateID    uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null;foreignKey:UserID"`
	AnimeID   uint      `gorm:"not null;foreignKey:AnimeID"`
	Rating    int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	User      User      `gorm:"foreignKey:UserID"`
	Anime     Anime     `gorm:"foreignKey:AnimeID"`
}
