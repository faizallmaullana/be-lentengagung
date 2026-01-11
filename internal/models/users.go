package models

import (
	"time"
)

type User struct {
	ID          string `gorm:"type:uuid;primaryKey;column:id" json:"-"`
	EncryptedID string `gorm:"-" json:"id,omitempty"`

	Email        string    `gorm:"unique;not null;column:email" json:"email"`
	PasswordHash string    `gorm:"not null;column:password_hash" json:"-"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	IsActive     bool      `gorm:"default:true;column:is_active" json:"is_active"`
	ApprovedAt   time.Time `gorm:"column:approved_at" json:"approved_at"`
	Role         string    `gorm:"default:'masyarakat';column:role" json:"role"`

	Profile Profile `gorm:"foreignKey:UserID;references:ID" json:"profile"`
}

func (User) TableName() string {
	return "users"
}

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	if u.ID == uuid.Nil {
// 		u.ID = utils.GenerateUUIDV6()
// 	}
// 	return nil
// }

// func (u *User) BeforeFind(tx *gorm.DB) (err error) {
// 	if u.EncryptedID != "" {
// 		key, err := utils.GetEncryptKey()
// 		if err != nil {
// 			return err
// 		}
// 		id, err := utils.DecryptToUUID(u.EncryptedID, key)
// 		if err != nil {
// 			return err
// 		}
// 		// modify query to search by decrypted uuid
// 		tx.Where("id = ?", id)
// 	}
// 	return nil
// }

// func (u *User) AfterFind(tx *gorm.DB) (err error) {
// 	key, err := utils.GetEncryptKey()
// 	if err != nil {
// 		return err
// 	}
// 	enc, err := utils.EncryptUUID(u.ID, key)
// 	if err != nil {
// 		return err
// 	}
// 	u.EncryptedID = enc
// 	return nil
// }
