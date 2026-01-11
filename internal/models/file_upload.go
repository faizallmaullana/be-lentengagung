package models

import "time"

type FileUpload struct {
	ID       string `json:"id" gorm:"column:id"`
	FileName string `json:"file_name" gorm:"column:file_name"`
	FilePath string `json:"file_path" gorm:"column:file_path"`

	Status     string    `json:"status" gorm:"column:status"`
	ApprovedAt time.Time `json:"approved_at" gorm:"column:approved_at"`

	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}
