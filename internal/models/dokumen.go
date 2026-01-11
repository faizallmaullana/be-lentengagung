package models

import (
	"time"
)

type Dokumen struct {
	ID       string `gorm:"primaryKey; column:id" json:"id"`
	Name     string `json:"name" gorm:"column:name"`
	FilePath string `json:"file_path" gorm:"column:file_path"`

	IdRegisterPernyataan string    `json:"id_register_pernyataan" gorm:"column:id_register_pernyataan"`
	ApprovedAt           time.Time `json:"approved_at" gorm:"column:approved_at"`

	ChatDokumen []ChatDocument `gorm:"foreignKey:IdDocument;references:ID" json:"chat_dokumen"`

	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (Dokumen) TableName() string {
	return "dokumen"
}
