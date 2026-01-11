package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatAhliWaris struct {
	ID          uuid.UUID `json:"id" gorm:"column:id"`
	IdAhliWaris uuid.UUID `json:"id_ahli_waris" gorm:"column:id_ahli_waris"`
	Chat        string    `json:"chat" gorm:"column:chat"`

	IDUser    uuid.UUID `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`

	ItsMe bool `gorm:"-" json:"its_me"`
}

func (ChatAhliWaris) TableName() string {
	return "chat_ahli_waris"
}
