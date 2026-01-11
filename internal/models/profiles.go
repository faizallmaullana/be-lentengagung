package models

type Profile struct {
	ID          string `gorm:"type:uuid;primaryKey;column:id" json:"-"`
	EncryptedID string `gorm:"-" json:"id,omitempty"`
	UserID      string `gorm:"type:uuid;not null;column:user_id" json:"user_id"`
	NIK         string `gorm:"unique;not null;column:nik" json:"nik"`
	Phone       string `gorm:"column:phone" json:"phone"`
	Religion    string `gorm:"column:religion" json:"religion"`
	Address     string `gorm:"column:address" json:"address"`
	Work        string `gorm:"column:work" json:"work"`
	Name        string `gorm:"column:name" json:"name"`
}

// func (p *Profile) BeforeCreate(tx *gorm.DB) (err error) {
// 	if p.ID == uuid.Nil {
// 		p.ID = utils.GenerateUUIDV6()
// 	}
// 	return nil
// }

// func (p *Profile) BeforeFind(tx *gorm.DB) (err error) {
// 	if p.EncryptedID != "" {
// 		key, err := utils.GetEncryptKey()
// 		if err != nil {
// 			return err
// 		}
// 		id, err := utils.DecryptToUUID(p.EncryptedID, key)
// 		if err != nil {
// 			return err
// 		}
// 		tx.Where("id = ?", id)
// 	}
// 	return nil
// }

// func (p *Profile) AfterFind(tx *gorm.DB) (err error) {
// 	key, err := utils.GetEncryptKey()
// 	if err != nil {
// 		return err
// 	}
// 	enc, err := utils.EncryptUUID(p.ID, key)
// 	if err != nil {
// 		return err
// 	}
// 	p.EncryptedID = enc
// 	return nil
// }
