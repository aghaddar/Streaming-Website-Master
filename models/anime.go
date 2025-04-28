package models

import "time"

type Anime struct {
	AnimeID          string    `gorm:"primaryKey;size:255;not null" json:"anime_id"`
	Title            string    `gorm:"type:text;not null" json:"title"`
	Description      *string   `gorm:"type:text" json:"description"`
	ImgURL           *string   `gorm:"type:text" json:"img_url"`
	TrailerURL       *string   `gorm:"type:text" json:"trailer_url"`
	NumberOfEpisodes int       `gorm:"not null" json:"number_of_episodes"`
	ReleaseYear      *int      `json:"release_year"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
}
