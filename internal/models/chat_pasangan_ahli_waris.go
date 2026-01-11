package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatPasanganAhliWaris struct {
	ID                  uuid.UUID `json:"id" gorm:"column:id"`
	IdPasanganAhliWaris uuid.UUID `json:"id_pasangan_ahli_waris" gorm:"column:id_pasangan_ahli_waris"`
	Chat                string    `json:"chat" gorm:"column:chat"`

	IDUser    uuid.UUID `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (ChatPasanganAhliWaris) TableName() string {
	return "chat_pasangan_ahli_waris"
}
