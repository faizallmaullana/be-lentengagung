package models

import (
	"time"
)

type ChatDocument struct {
	ID         string `json:"id" gorm:"column:id"`
	IdDocument string `json:"id_document" gorm:"column:id_document"`
	Chat       string `json:"chat" gorm:"column:chat"`

	IDUser    string    `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (ChatDocument) TableName() string {
	return "chat_document"
}
