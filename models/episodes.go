package models

import "time"

type Episode struct {
	EpisodeID     string     `gorm:"primaryKey;type:text;not null"`
	AnimeID       string     `gorm:"type:text;not null"`
	Title         string     `gorm:"size:255;not null"`
	EpisodeNumber int        `gorm:"not null"`
	Duration      *int       `gorm:"size:11"`
	VideoURL      *string    `gorm:"type:text"`
	ReleaseDate   *time.Time `gorm:"type:date"`
	CreatedAt     time.Time  `gorm:"default:current_timestamp"`

	Anime Anime `gorm:"foreignKey:AnimeID;references:AnimeID"`
}
