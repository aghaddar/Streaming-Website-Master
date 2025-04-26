package models

import "time"

type Watchlist struct {
	WatchlistID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID             uint      `gorm:"not null"`
	AnimeID            string    `gorm:"type:varchar(255);not null"`
	CreatedAt          time.Time `gorm:"default:current_timestamp"`
	Status             string    `gorm:"type:enum('Watching','Completed','On Hold','Dropped');default:'Watching'"`
	Priority           string    `gorm:"type:enum('High','Medium','Low');default:'Medium'"`
	LastWatchedEpisode *string   `gorm:"type:varchar(255)"`
	ProgressPercentage float64   `gorm:"type:decimal(5,2);default:0.00"`

	User  User  `gorm:"foreignKey:UserID;references:UserID"` // Correct FK
	Anime Anime `gorm:"foreignKey:AnimeID;references:AnimeID"`
}
