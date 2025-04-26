package models

import "time"

type Report struct {
	ReportID   uint `gorm:"primaryKey;autoIncrement"`
	ReporterID uint `gorm:"not null"`
	ReportedID *uint
	CommentID  *uint
	Reason     string    `gorm:"type:text;not null"`
	Status     string    `gorm:"size:20;default:'pending'"`
	CreatedAt  time.Time `gorm:"default:current_timestamp"`

	Reporter User      `gorm:"foreignKey:ReporterID;references:UserID"`
	Reported *User     `gorm:"foreignKey:ReportedID;references:UserID"`
	Comment  *Comments `gorm:"foreignKey:CommentID;references:CommentID"`
}
