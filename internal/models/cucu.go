package models

import (
	"time"

	"github.com/google/uuid"
)

type Cucu struct {
	ID          uuid.UUID `gorm:"primaryKey; column:id" json:"id"`
	Name        string    `json:"name" gorm:"column:name"`
	NIK         string    `json:"nik" gorm:"column:nik"`
	Phone       string    `json:"phone" gorm:"column:phone"`
	RT          string    `json:"rt" gorm:"column:rt"`
	RW          string    `json:"rw" gorm:"column:rw"`
	Kelurahan   string    `json:"kelurahan" gorm:"column:kelurahan"`
	Kecamatan   string    `json:"kecamatan" gorm:"column:kecamatan"`
	Kabupaten   string    `json:"kabupaten" gorm:"column:kabupaten"`
	Province    string    `json:"province" gorm:"column:province"`
	Address     string    `json:"address" gorm:"column:address"`
	Religion    string    `json:"religion" gorm:"column:religion"`
	Work        string    `json:"work" gorm:"column:work"`
	DateOfBirth string    `json:"date_of_birth" gorm:"column:date_of_birth"`
	Gender      string    `json:"gender" gorm:"column:gender"`
	BloodType   string    `json:"blood_type" gorm:"column:blood_type"`
	IdLampiran  string    `json:"id_lampiran" gorm:"column:id_lampiran"`

	UrutanPasangan     int8   `json:"urutan_pasangan" gorm:"column:urutan_pasangan"`
	IsDead             bool   `json:"is_dead" gorm:"column:is_dead"`
	NoAktaKematian     string `json:"no_akta_kematian" gorm:"column:no_akta_kematian"`
	KeteranganKematian string `json:"keterangan_kematian" gorm:"column:keterangan_kematian"`

	IdPasanganAhliWaris uuid.UUID `json:"id_pasangan_ahli_waris" gorm:"column:id_pasangan_ahli_waris"`

	Timestamp time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (Cucu) TableName() string {
	return "cucu"
}
