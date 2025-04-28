package models

type AnimeGenre struct {
	AnimeID string `gorm:"primaryKey;size:255;not null" json:"anime_id"`
	GenreID uint64 `gorm:"primaryKey;not null" json:"genre_id"`
}
