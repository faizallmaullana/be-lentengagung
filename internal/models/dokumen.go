package models

import (
	"time"

	"github.com/google/uuid"
)

type Dokumen struct {
	ID       uuid.UUID `gorm:"primaryKey; column:id" json:"id"`
	Name     string    `json:"name" gorm:"column:name"`
	FilePath string    `json:"file_path" gorm:"column:file_path"`

	IdRegisterPernyataan uuid.UUID `json:"id_register_pernyataan" gorm:"column:id_register_pernyataan"`
	ApprovedAt           time.Time `json:"approved_at" gorm:"column:approved_at"`

	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}
