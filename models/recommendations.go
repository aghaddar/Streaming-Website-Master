package models

import "time"

type Recommendation struct {
	RecommendationID uint      `gorm:"primaryKey;autoIncrement"`
	UserID           uint      `gorm:"not null;foreignKey:UserID"`
	AnimeID          uint      `gorm:"not null;foreignKey:AnimeID"`
	Reason           *string   `gorm:"type:text"`
	CreatedAt        time.Time `gorm:"default:current_timestamp"`

	User  User  `gorm:"foreignKey:UserID"`
	Anime Anime `gorm:"foreignKey:AnimeID"`
}
