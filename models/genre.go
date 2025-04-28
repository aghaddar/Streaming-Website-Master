package models

type Genre struct {
	GenreID uint64 `gorm:"primaryKey;autoIncrement" json:"genre_id"`
	Name    string `gorm:"size:50;not null" json:"name"`
}
