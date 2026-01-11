package repo

import (
	"database/sql"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	database "github.com/faizallmaullana/lenteng-agung/backend/db/db_connection"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/models"
)

type authRepo struct {
	provider database.DBProvider
	tx       *gorm.DB
}

// AuthRepo defines low-level data operations only.
type AuthRepo interface {
	IsEmailExists(email string) (bool, error)
	IsNIKExists(nik string) (bool, error)
	CreateUser(u *models.User) error
	GetUserByNIK(nik string) (*models.User, sql.NullTime, error)
	GetUserByID(userID string) (*models.User, sql.NullTime, error)
	ApproveUser(userID string) error
	CreateProfile(p *models.Profile) error
	GetProfileByUserID(userID string) (*models.Profile, error)
	WithTx(tx *gorm.DB) AuthRepo
}

func NewAuthRepo(provider database.DBProvider) AuthRepo {
	return &authRepo{provider: provider}
}

func (r *authRepo) WithTx(tx *gorm.DB) AuthRepo {
	return &authRepo{provider: r.provider, tx: tx}
}

func (r *authRepo) IsEmailExists(email string) (bool, error) {
	var cnt int64
	db := r.provider.DB()
	if r.tx != nil {
		db = r.tx
	}
	if err := db.Model(&models.User{}).Where("email = ? AND approved_at IS NOT NULL", email).Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (r *authRepo) IsNIKExists(nik string) (bool, error) {
	var cnt int64
	db := r.provider.DB()
	if r.tx != nil {
		db = r.tx
	}
	if err := db.Model(&models.Profile{}).
		Joins("JOIN users ON users.id = profiles.user_id").
		Where("profiles.nik = ? AND users.approved_at IS NOT NULL", nik).
		Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (r *authRepo) CreateUser(u *models.User) error {
	db := r.provider.DB()
	if r.tx != nil {
		db = r.tx
	}

	// if an existing user with this email is already approved, do nothing
	var ex struct{ ApprovedAt sql.NullTime }
	if err := db.Table("users").Select("approved_at").Where("email = ?", u.Email).First(&ex).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else {
		if ex.ApprovedAt.Valid {
			return nil
		}
	}

	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"password_hash", "is_active"}),
	}).Create(u).Error; err != nil {
		return err
	}

	// ensure u.ID is populated (upsert path may not populate struct fields)
	var id struct{ ID string }
	if err := db.Table("users").Select("id").Where("email = ?", u.Email).First(&id).Error; err != nil {
		return err
	}
	u.ID = id.ID
	return nil
}

func (r *authRepo) CreateProfile(p *models.Profile) error {
	db := r.provider.DB()
	if r.tx != nil {
		db = r.tx
	}

	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "nik"}},
		DoUpdates: clause.AssignmentColumns([]string{"user_id", "phone", "religion", "address", "work", "name"}),
	}).Create(p).Error
}

func (r *authRepo) GetUserByNIK(nik string) (*models.User, sql.NullTime, error) {
	users := &models.User{}
	db := r.provider.DB()
	if r.tx != nil {
		db = r.tx
	}
	if err := db.Joins("JOIN profiles ON profiles.user_id = users.id").Preload("Profile").Where("profiles.nik = ?", nik).First(users).Error; err != nil {
		return nil, sql.NullTime{}, err
	}

	approvedAt := sql.NullTime{}
	if !users.ApprovedAt.IsZero() {
		approvedAt = sql.NullTime{Time: users.ApprovedAt, Valid: true}
	}

	return users, approvedAt, nil
}

func (r *authRepo) GetUserByID(userID string) (*models.User, sql.NullTime, error) {
	db := r.provider.DB()
	if r.tx != nil {
		db = r.tx
	}

	var dest struct {
		ID         string       `gorm:"column:id"`
		Email      string       `gorm:"column:email"`
		CreatedAt  time.Time    `gorm:"column:created_at"`
		IsActive   bool         `gorm:"column:is_active"`
		ApprovedAt sql.NullTime `gorm:"column:approved_at"`
	}

	if err := db.Table("users").Select("id, email, created_at, is_active, approved_at").Where("users.id = ?", userID).Limit(1).Scan(&dest).Error; err != nil {
		return nil, sql.NullTime{}, err
	}

	u := &models.User{
		ID:        dest.ID,
		Email:     dest.Email,
		CreatedAt: dest.CreatedAt,
		IsActive:  dest.IsActive,
	}
	return u, dest.ApprovedAt, nil
}

func (r *authRepo) GetProfileByUserID(userID string) (*models.Profile, error) {
	db := r.provider.DB()
	if r.tx != nil {
		db = r.tx
	}

	var p models.Profile
	if err := db.Table("profiles").Where("user_id = ?", userID).Limit(1).Scan(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *authRepo) ApproveUser(userID string) error {
	db := r.provider.DB()
	if r.tx != nil {
		db = r.tx
	}
	return db.Table("users").Where("id = ?", userID).Update("approved_at", time.Now()).Error
}
