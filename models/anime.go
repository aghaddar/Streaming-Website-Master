package models

import (
	"time"
)

type Anime struct {
	AnimeID          string  `gorm:"primaryKey;type:varchar(255);not null"`
	Title            string  `gorm:"type:text;not null"`
	Description      *string `gorm:"type:text"`
	ImgURL           *string `gorm:"type:text"`
	TrailerURL       *string `gorm:"type:text"`
	NumberOfEpisodes int     `gorm:"not null"`
	ReleaseYear      *int
	CreatedAt        time.Time `gorm:"default:current_timestamp"`
}
