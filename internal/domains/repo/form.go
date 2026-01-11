package repo

import (
	"github.com/faizallmaullana/lenteng-agung/backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FormRepo struct {
	db *gorm.DB
}

func NewFormRepo(db *gorm.DB) *FormRepo {
	return &FormRepo{db: db}
}

func (r *FormRepo) CreateRequest(userID uuid.UUID) (*models.RegisterPernyataan, error) {
	models := &models.RegisterPernyataan{}
	models.IDUser = userID

	if err := r.db.Create(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}
