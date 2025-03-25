package models

import "time"

type Report struct {
	ReportID   uint      `gorm:"primaryKey;autoIncrement"`
	ReporterID uint      `gorm:"not null;foreignKey:ReporterID"`
	ReportedID *uint     `gorm:"foreignKey:ReportedID"`
	CommentID  *uint     `gorm:"foreignKey:CommentID"`
	Reason     string    `gorm:"type:text;not null"`
	Status     string    `gorm:"size:20;default:'pending'"`
	CreatedAt  time.Time `gorm:"default:current_timestamp"`

	Reporter User      `gorm:"foreignKey:ReporterID"`
	Reported *User     `gorm:"foreignKey:ReportedID"`
	Comments *Comments `gorm:"foreignKey:CommentID"`
}
