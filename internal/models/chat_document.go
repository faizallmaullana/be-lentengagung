package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatDocument struct {
	ID         uuid.UUID `json:"id" gorm:"column:id"`
	IdDocument uuid.UUID `json:"id_document" gorm:"column:id_document"`
	Chat       string    `json:"chat" gorm:"column:chat"`

	IDUser    uuid.UUID `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (ChatDocument) TableName() string {
	return "chat_document"
}
