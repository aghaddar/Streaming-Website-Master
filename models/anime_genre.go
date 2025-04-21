package models

type Anime_Genre struct {
	AnimeID string `gorm:"primaryKey;type:varchar(255);not null"`
	GenreID uint   `gorm:"primaryKey;not null"`
}
