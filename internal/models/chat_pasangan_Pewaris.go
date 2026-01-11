package models

import (
	"time"
)

type ChatPasanganPewaris struct {
	ID                string `json:"id" gorm:"column:id"`
	IdPasanganPewaris string `json:"id_pasangan_pewaris" gorm:"column:id_pasangan_pewaris"`
	Chat              string `json:"chat" gorm:"column:chat"`

	IDUser    string    `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (ChatPasanganPewaris) TableName() string {
	return "chat_pasangan_pewaris"
}
