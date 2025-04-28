package models

import "time"

type Rating struct {
	RateID    uint64    `gorm:"primaryKey;autoIncrement" json:"rate_id"`
	UserID    uint64    `gorm:"not null;index" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	AnimeID   string    `gorm:"size:255;not null;index" json:"anime_id"`
	Anime     Anime     `gorm:"foreignKey:AnimeID;references:AnimeID;constraint:OnDelete:CASCADE"`
	Rating    int       `gorm:"not null" json:"rating"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
