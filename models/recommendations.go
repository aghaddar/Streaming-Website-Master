package models

import "time"

type Recommendation struct {
	RecommendationID uint64    `gorm:"primaryKey;autoIncrement" json:"recommendation_id"`
	UserID           uint64    `gorm:"not null" json:"user_id"`
	AnimeID          string    `gorm:"size:255;not null" json:"anime_id"`
	Reason           *string   `gorm:"type:text" json:"reason"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
}
