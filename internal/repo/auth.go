package repo

import (
	"context"

	"gorm.io/gorm"

	"github.com/faizallmaullana/lenteng-agung/backend/database"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/models"
)

type authRepo struct {
	provider database.DBProvider
	tx       *gorm.DB
}

// AuthRepo defines low-level data operations only.
type AuthRepo interface {
	IsEmailExists(ctx context.Context, email string) (bool, error)
	IsNIKExists(ctx context.Context, nik string) (bool, error)
	CreateUser(ctx context.Context, u *models.User) error
	CreateProfile(ctx context.Context, p *models.Profile) error
	WithTx(tx *gorm.DB) AuthRepo
}

func NewAuthRepo(provider database.DBProvider) AuthRepo {
	return &authRepo{provider: provider}
}

func (r *authRepo) WithTx(tx *gorm.DB) AuthRepo {
	return &authRepo{provider: r.provider, tx: tx}
}

func (r *authRepo) IsEmailExists(ctx context.Context, email string) (bool, error) {
	var cnt int64
	db := r.provider.DB()
	if r.tx != nil {
		db = r.tx
	}
	if err := db.Model(&models.User{}).Where("email = ?", email).Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (r *authRepo) IsNIKExists(ctx context.Context, nik string) (bool, error) {
	var cnt int64
	db := r.provider.DB()
	if r.tx != nil {
		db = r.tx
	}
	if err := db.Model(&models.Profile{}).Where("nik = ?", nik).Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (r *authRepo) CreateUser(ctx context.Context, u *models.User) error {
	db := r.provider.DB()
	if r.tx != nil {
		db = r.tx
	}
	return db.Create(u).Error
}

func (r *authRepo) CreateProfile(ctx context.Context, p *models.Profile) error {
	db := r.provider.DB()
	if r.tx != nil {
		db = r.tx
	}
	return db.Create(p).Error
}
