package models

type Genre struct {
	GenreID uint   `gorm:"primaryKey;autoIncrement"`
	Name    string `gorm:"size:50;not null"`
}
