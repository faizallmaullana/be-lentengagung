package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatPasanganPewaris struct {
	ID                uuid.UUID `json:"id" gorm:"column:id"`
	IdPasanganPewaris uuid.UUID `json:"id_pasangan_pewaris" gorm:"column:id_pasangan_pewaris"`
	Chat              string    `json:"chat" gorm:"column:chat"`

	IDUser    uuid.UUID `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (ChatPasanganPewaris) TableName() string {
	return "chat_pasangan_pewaris"
}
