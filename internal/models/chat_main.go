package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatMain struct {
	ID                   uuid.UUID `json:"id" gorm:"column:id"`
	IDRegisterPernyataan uuid.UUID `json:"id_register_pernyataan" gorm:"column:id_register_pernyataan"`
	Chat                 string    `json:"chat" gorm:"column:chat"`

	IDUser    uuid.UUID `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (ChatMain) TableName() string {
	return "chat_main"
}
