package models

import (
	"time"
)

type ChatCucu struct {
	ID     string `json:"id" gorm:"column:id"`
	IdCucu string `json:"id_cucu" gorm:"column:id_cucu"`
	Chat   string `json:"chat" gorm:"column:chat"`

	IDUser    string    `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (ChatCucu) TableName() string {
	return "chat_cucu"
}
