package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatCucu struct {
	ID     uuid.UUID `json:"id" gorm:"column:id"`
	IdCucu uuid.UUID `json:"id_cucu" gorm:"column:id_cucu"`
	Chat   string    `json:"chat" gorm:"column:chat"`

	IDUser    uuid.UUID `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (ChatCucu) TableName() string {
	return "chat_cucu"
}
