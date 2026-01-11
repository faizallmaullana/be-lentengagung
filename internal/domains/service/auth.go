package service

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	database "github.com/faizallmaullana/lenteng-agung/backend/db/db_connection"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/domains/dto"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/domains/repo"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/models"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Register(c *gin.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
	Login(c *gin.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	ApproveRegistration(c *gin.Context, req dto.ApprovalRequest) (dto.ApprovalResponse, error)
}

type authService struct {
	repo     repo.AuthRepo
	provider database.DBProvider
	jwtSvc   *JWTService
}

func NewAuthService(r repo.AuthRepo, provider database.DBProvider, jwtSvc *JWTService) AuthService {
	return &authService{repo: r, provider: provider, jwtSvc: jwtSvc}
}

func (s *authService) Register(c *gin.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	if req.NIK == "" || req.Email == "" || req.Password == "" {
		return dto.RegisterResponse{}, errors.New("nik, email and password are required")
	}

	// uniqueness checks
	if ok, err := s.repo.IsEmailExists(req.Email); err != nil {
		return dto.RegisterResponse{}, err
	} else if ok {
		return dto.RegisterResponse{}, errors.New("email already registered")
	}

	if ok, err := s.repo.IsNIKExists(req.NIK); err != nil {
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
		if err := txRepo.CreateUser(u); err != nil {
			return err
		}

		p := &models.Profile{
			UserID: u.ID,
			NIK:    req.NIK,
			Phone:  req.Phone,
		}
		if err := txRepo.CreateProfile(p); err != nil {
			return err
		}

		user = u
		return nil
	})

	if err != nil {
		return dto.RegisterResponse{}, err
	}

	// // encrypt the user ID for response
	// key, err := utils.GetEncryptKey()
	// if err != nil {
	// 	return dto.RegisterResponse{ID: user.ID.String(), Email: user.Email, CreatedAt: user.CreatedAt}, nil
	// }
	// enc, err := utils.EncryptUUID(user.ID, key)
	// if err != nil {
	// 	return dto.RegisterResponse{}, err
	// }

	// Generate registration token
	token := utils.RandomString(6)
	fmt.Println("DEBUG: registration token:", token)

	payload := dto.JWTPayload{
		UserID: user.ID,
		Email:  user.Email,
		Token:  token,
	}

	regToken, err := s.jwtSvc.CreateRegistrationToken(payload)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	// Uncomment to send email
	// mailSender := mails.NewMailSender()
	// mail := mails.Mailer{
	// 	To:      user.Email,
	// 	Subject: "Registration Token",
	// 	Body:    "Your registration token: " + regToken,
	// }
	// if err := mailSender.SendMail(mail); err != nil {
	// 	fmt.Println("DEBUG: Error sending mail:", err)
	// 	return dto.RegisterResponse{}, err
	// }

	return dto.RegisterResponse{ID: user.ID, Email: user.Email, CreatedAt: time.Now(), RegistrationToken: regToken}, nil
}

func (s *authService) Login(c *gin.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	nik := req.NIK.String()
	if nik == "" || req.Password == "" {
		return dto.LoginResponse{}, errors.New("nik and password are required")
	}

	user, approvedAt, err := s.repo.GetUserByNIK(nik)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.LoginResponse{}, errors.New("invalid credentials")
		}
		return dto.LoginResponse{}, err
	}

	if !approvedAt.Valid {
		return dto.LoginResponse{}, errors.New("belum registrasi")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return dto.LoginResponse{}, errors.New("invalid credentials")
	}

	token, err := s.jwtSvc.CreateAccessToken(user.ID, user.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	p, err := s.repo.GetProfileByUserID(user.ID)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{AccessToken: token, TokenType: "bearer", User: user, Profile: p}, nil
}

func (s *authService) ApproveRegistration(c *gin.Context, req dto.ApprovalRequest) (dto.ApprovalResponse, error) {

	// ==============================
	// validate token from context and request
	token, ok := c.Get("token")
	if !ok {
		return dto.ApprovalResponse{}, errors.New("unauthorized")
	}
	if token.(string) == "" || req.Token == "" {
		fmt.Println("zero value token")
		return dto.ApprovalResponse{}, errors.New("zero value token")
	}
	fmt.Println(token.(string))
	fmt.Println(req.Token)
	if token.(string) != req.Token {
		fmt.Println("invalid token")
		return dto.ApprovalResponse{}, errors.New("invalid token")
	}

	fmt.Println(token)
	fmt.Println(token)
	// ===

	// ==============================
	// get id user
	user, ok := c.Get("id_user")
	if !ok {
		return dto.ApprovalResponse{}, errors.New("unauthorized")
	}
	// ===

	// ==============================
	// main functionality: approve user
	userID := user.(string)
	if err := s.repo.ApproveUser(userID); err != nil {
		return dto.ApprovalResponse{}, err
	}

	// fetch updated user and profile to return
	u, _, err := s.repo.GetUserByID(userID)
	if err != nil {
		return dto.ApprovalResponse{}, err
	}
	p, err := s.repo.GetProfileByUserID(userID)
	if err != nil {
		return dto.ApprovalResponse{}, err
	}

	return dto.ApprovalResponse{Message: "approved", User: *u, Profile: *p}, nil
}
