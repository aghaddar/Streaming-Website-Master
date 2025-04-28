package models

import "time"

type Report struct {
	ReportID   uint64    `gorm:"primaryKey;autoIncrement" json:"report_id"`
	ReporterID uint64    `gorm:"not null" json:"reporter_id"`
	ReportedID *uint64   `json:"reported_id"`
	CommentID  *uint64   `json:"comment_id"`
	Reason     string    `gorm:"type:text;not null" json:"reason"`
	Status     string    `gorm:"size:20;default:pending" json:"status"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}
