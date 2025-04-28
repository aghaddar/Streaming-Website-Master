package models

import "time"

type Episode struct {
	EpisodeID     string     `gorm:"primaryKey;size:255;not null" json:"episode_id"`
	AnimeID       string     `gorm:"size:255;not null" json:"anime_id"`
	Title         string     `gorm:"size:255;not null" json:"title"`
	EpisodeNumber int        `gorm:"not null" json:"episode_number"`
	Duration      *int       `json:"duration"` // duration in minutes?
	VideoURL      *string    `gorm:"type:text" json:"video_url"`
	ReleaseDate   *time.Time `json:"release_date"`
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"created_at"`
}
