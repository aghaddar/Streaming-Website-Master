package models

import (
	"time"
)

type Anime struct {
	AnimeID          uint      `gorm:"primaryKey;autoIncrement"`
	Title            string    `gorm:"size:255;not null"`
	Description      *string   `gorm:"type:text"`
	ImgURL           *string   `gorm:"size:255"`
	TrailerURL       *string   `gorm:"size:255"`
	NumberOfEpisodes int       `gorm:"not null"`
	ReleaseYear      *int      `gorm:"size:4"`
	CreatedAt        time.Time `gorm:"default:current_timestamp"`
}
