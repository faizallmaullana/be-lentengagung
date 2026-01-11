package models

import (
	"time"
)

type RegisterPernyataan struct {
	ID             string `gorm:"primaryKey; column:id" json:"id"`
	KodeRegistrasi string `json:"kode_registrasi" gorm:"column:kode_registrasi"`
	Status         string `json:"status" gorm:"column:status"`

	IDUser         string     `json:"id_user" gorm:"column:id_user"`
	Pewaris        Pewaris    `gorm:"foreignKey:IdRegisterPernyataan;references:ID" json:"pewaris"`
	Dokumen        []Dokumen  `gorm:"foreignKey:IdRegisterPernyataan;references:ID" json:"dokumen"`
	ChatPernyataan []ChatMain `gorm:"foreignKey:IDRegisterPernyataan;references:ID" json:"chat_pernyataan"`

	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (RegisterPernyataan) TableName() string {
	return "register_pernyataan"
}
