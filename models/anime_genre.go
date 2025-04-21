package models

type Anime_Genre struct {
	AnimeID string `gorm:"primaryKey;type:text;not null"`
	GenreID uint   `gorm:"primaryKey;not null"`
}
