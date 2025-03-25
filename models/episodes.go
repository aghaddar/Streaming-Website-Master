package models

import "time"

type Episode struct {
	EpisodeID     uint       `gorm:"primaryKey;autoIncrement"`
	AnimeID       uint       `gorm:"not null;foreignKey:AnimeID"`
	Title         string     `gorm:"size:255;not null"`
	EpisodeNumber int        `gorm:"not null"`
	Duration      *int       `gorm:"size:11"`
	VideoURL      *string    `gorm:"size:255"`
	ReleaseDate   *time.Time `gorm:"type:date"`
	CreatedAt     time.Time  `gorm:"default:current_timestamp"`

	Anime Anime `gorm:"foreignKey:AnimeID"`
}
