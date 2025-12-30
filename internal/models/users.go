package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/faizallmaullana/lenteng-agung/backend/internal/pkg/utils"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"-"`
	EncryptedID  string    `gorm:"-" json:"id,omitempty"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	Role         string    `gorm:"default:'masyarakat'" json:"role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = utils.GenerateUUIDV6()
	}
	return nil
}

func (u *User) BeforeFind(tx *gorm.DB) (err error) {
	if u.EncryptedID != "" {
		key, err := utils.GetEncryptKey()
		if err != nil {
			return err
		}
		id, err := utils.DecryptToUUID(u.EncryptedID, key)
		if err != nil {
			return err
		}
		// modify query to search by decrypted uuid
		tx.Where("id = ?", id)
	}
	return nil
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	key, err := utils.GetEncryptKey()
	if err != nil {
		return err
	}
	enc, err := utils.EncryptUUID(u.ID, key)
	if err != nil {
		return err
	}
	u.EncryptedID = enc
	return nil
}
