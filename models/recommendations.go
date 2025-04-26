package models

import "time"

type Recommendation struct {
	RecommendationID uint   `gorm:"primaryKey;autoIncrement"`
	UserID           uint   `gorm:"not null"` // <-- Must be here!
	AnimeID          string `gorm:"type:varchar(255);not null"`
	Reason           string
	CreatedAt        time.Time `gorm:"default:current_timestamp"`
	User             User      `gorm:"foreignKey:UserID;references:UserID"` // <-- Explicit
	Anime            Anime     `gorm:"foreignKey:AnimeID;references:AnimeID"`
}
