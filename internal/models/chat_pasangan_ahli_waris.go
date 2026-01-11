package models

import (
	"time"
)

type ChatPasanganAhliWaris struct {
	ID                  string `json:"id" gorm:"column:id"`
	IdPasanganAhliWaris string `json:"id_pasangan_ahli_waris" gorm:"column:id_pasangan_ahli_waris"`
	Chat                string `json:"chat" gorm:"column:chat"`

	IDUser    string    `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (ChatPasanganAhliWaris) TableName() string {
	return "chat_pasangan_ahli_waris"
}
