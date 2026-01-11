package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatPewaris struct {
	ID        uuid.UUID `json:"id" gorm:"column:id"`
	IdPewaris uuid.UUID `json:"id_pewaris" gorm:"column:id_pewaris"`
	Chat      string    `json:"chat" gorm:"column:chat"`

	IDUser    uuid.UUID `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (ChatPewaris) TableName() string {
	return "chat_pewaris"
}
