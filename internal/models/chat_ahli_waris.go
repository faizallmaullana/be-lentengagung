package models

import (
	"time"
)

type ChatAhliWaris struct {
	ID          string `json:"id" gorm:"column:id"`
	IdAhliWaris string `json:"id_ahli_waris" gorm:"column:id_ahli_waris"`
	Chat        string `json:"chat" gorm:"column:chat"`

	IDUser    string    `json:"id_user" gorm:"column:id_user"`
	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`

	ItsMe bool `gorm:"-" json:"its_me"`
}

func (ChatAhliWaris) TableName() string {
	return "chat_ahli_waris"
}
