package repo

import (
	"time"

	"github.com/faizallmaullana/lenteng-agung/backend/internal/models"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/pkg/utils"
	"gorm.io/gorm"
)

type statusProses struct {
	PengisianData    string
	VerifikasiBerkas string
	DraftDisetujui   string
	Penandatanganan  string
	Selesai          string
}

var StatusProses = statusProses{
	PengisianData:    "Pengisian Data",
	VerifikasiBerkas: "Verifikasi Berkas",
	DraftDisetujui:   "Draft Disetujui",
	Penandatanganan:  "Penandatanganan",
	Selesai:          "Selesai",
}

type FormRepo struct {
	db *gorm.DB
}

func NewFormRepo(db *gorm.DB) *FormRepo {
	return &FormRepo{db: db}
}

func (r *FormRepo) CreateRequest(userID string) (*models.RegisterPernyataan, error) {
	models := &models.RegisterPernyataan{}
	models.ID = utils.GenerateUUIDV6()
	models.Status = StatusProses.PengisianData
	models.IDUser = userID
	models.Timestamp = time.Now()

	if err := r.db.Create(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (r *FormRepo) GetRequestByUserID(userID string) (*models.RegisterPernyataan, error) {
	var models models.RegisterPernyataan
	if err := r.db.Where("id_user = ?", userID).First(&models).Error; err != nil {
		return nil, err
	}
	return &models, nil
}
