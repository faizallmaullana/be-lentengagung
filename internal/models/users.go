package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/faizallmaullana/lenteng-agung/backend/internal/pkg/utils"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = utils.GenerateUUIDV6()
	}
	return nil
}
