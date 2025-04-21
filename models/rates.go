package models

import (
	"time"
)

type Rating struct {
	RateID    uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null"`
	AnimeID   string    `gorm:"type:varchar(255);not null"`
	Rating    int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	User      User      `gorm:"foreignKey:UserID"`
	Anime     Anime     `gorm:"foreignKey:AnimeID;references:AnimeID"`
}
