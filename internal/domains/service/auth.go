package service

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	database "github.com/faizallmaullana/lenteng-agung/backend/db/db_connection"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/domains/dto"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/domains/repo"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/models"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/pkg/utils"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
}

type authService struct {
	repo     repo.AuthRepo
	provider database.DBProvider
}

func NewAuthService(r repo.AuthRepo, provider database.DBProvider) AuthService {
	return &authService{repo: r, provider: provider}
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	if req.NIK == "" || req.Email == "" || req.Password == "" {
		return dto.RegisterResponse{}, errors.New("nik, email and password are required")
	}

	// uniqueness checks
	if ok, err := s.repo.IsEmailExists(ctx, req.Email); err != nil {
		return dto.RegisterResponse{}, err
	} else if ok {
		return dto.RegisterResponse{}, errors.New("email already registered")
	}

	if ok, err := s.repo.IsNIKExists(ctx, req.NIK); err != nil {
		return dto.RegisterResponse{}, err
	} else if ok {
		return dto.RegisterResponse{}, errors.New("nik already registered")
	}

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	var user *models.User

	err = s.provider.DB().Transaction(func(tx *gorm.DB) error {
		txRepo := s.repo.WithTx(tx)

		u := &models.User{
			Email:        req.Email,
			PasswordHash: string(hashed),
		}
		if err := txRepo.CreateUser(ctx, u); err != nil {
			return err
		}

		p := &models.Profile{
			UserID: u.ID,
			NIK:    req.NIK,
			Phone:  req.Phone,
		}
		if err := txRepo.CreateProfile(ctx, p); err != nil {
			return err
		}

		user = u
		// profile := p
		return nil
	})

	if err != nil {
		return dto.RegisterResponse{}, err
	}

	// encrypt the user ID for response
	key, err := utils.GetEncryptKey()
	if err != nil {
		// if key not set, still return plain UUID string
		return dto.RegisterResponse{ID: user.ID.String(), Email: user.Email, CreatedAt: user.CreatedAt}, nil
	}
	enc, err := utils.EncryptUUID(user.ID, key)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{ID: enc, Email: user.Email, CreatedAt: time.Now()}, nil
}
