package models

import (
	"time"

	"github.com/google/uuid"
)

type RegisterPernyataan struct {
	ID             uuid.UUID `gorm:"primaryKey; column:id" json:"id"`
	KodeRegistrasi string    `json:"kode_registrasi" gorm:"column:kode_registrasi"`
	Status         string    `json:"status" gorm:"column:status"`

	IDUser  uuid.UUID `json:"id_user" gorm:"column:id_user"`
	Pewaris Pewaris   `gorm:"foreignKey:IdRegisterPernyataan;references:ID;omitempty" json:"pewaris"`
	Dokumen []Dokumen `gorm:"foreignKey:IdRegisterPernyataan;references:ID;omitempty" json:"dokumen"`

	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (RegisterPernyataan) TableName() string {
	return "register_pernyataan"
}
