package models

type Anime_Genre struct {
	AnimeID uint `gorm:"primaryKey;not null"`
	GenreID uint `gorm:"primaryKey;not null"`
}
