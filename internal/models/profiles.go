package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/faizallmaullana/lenteng-agung/backend/internal/pkg/utils"
)

type Profile struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID   uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	NIK      string    `gorm:"unique;not null" json:"nik"`
	Phone    string    `json:"phone"`
	Religion string    `json:"religion"`
	Address  string    `json:"address"`
	Work     string    `json:"work"`
	Name     string    `json:"name"`
}

func (p *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = utils.GenerateUUIDV6()
	}
	return nil
}
