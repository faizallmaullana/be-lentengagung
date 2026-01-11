package models

import (
	"time"
)

type ChatPewaris struct {
	ID        string `json:"id" gorm:"column:id"`
	IdPewaris string `json:"id_pewaris" gorm:"column:id_pewaris"`
	Chat      string `json:"chat" gorm:"column:chat"`

	IDUser    string    `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (ChatPewaris) TableName() string {
	return "chat_pewaris"
}
