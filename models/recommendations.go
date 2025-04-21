package models

import "time"

type Recommendation struct {
	RecommendationID uint      `gorm:"primaryKey;autoIncrement"`
	UserID           uint      `gorm:"not null"`
	AnimeID          string    `gorm:"type:text;not null"`
	Reason           *string   `gorm:"type:text"`
	CreatedAt        time.Time `gorm:"default:current_timestamp"`

	User  User  `gorm:"foreignKey:UserID"`
	Anime Anime `gorm:"foreignKey:AnimeID;references:AnimeID"`
}
