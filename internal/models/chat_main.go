package models

import (
	"time"
)

type ChatMain struct {
	ID                   string `json:"id" gorm:"column:id"`
	IDRegisterPernyataan string `json:"id_register_pernyataan" gorm:"column:id_register_pernyataan"`
	Chat                 string `json:"chat" gorm:"column:chat"`

	IDUser    string    `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (ChatMain) TableName() string {
	return "chat_main"
}
