package service

import (
	"os"
	"time"

	"github.com/faizallmaullana/lenteng-agung/backend/internal/domains/dto"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey            string
	accessTokenTTL       time.Duration
	refreshTokenTTL      time.Duration
	tokenRegistrationTTL time.Duration
}

func NewJWTService() *JWTService {
	return &JWTService{
		secretKey:            os.Getenv("JWT_SECRET_KEY"),
		accessTokenTTL:       1 * 24 * time.Hour,
		refreshTokenTTL:      7 * 24 * time.Hour,
		tokenRegistrationTTL: 3 * time.Minute,
	}
}

type TokenPayload struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token,omitempty"`
	jwt.RegisteredClaims
}

func (j *JWTService) CreateAccessToken(userID, email string) (string, error) {
	claims := TokenPayload{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) CreateRefreshToken(payload dto.JWTPayload) (string, error) {
	claims := TokenPayload{
		UserID: payload.UserID,
		Email:  payload.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) CreateRegistrationToken(payload dto.JWTPayload) (string, error) {
	claims := TokenPayload{
		UserID: payload.UserID,
		Email:  payload.Email,
		Token:  payload.Token,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenRegistrationTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) ValidateToken(payload dto.JWTPayload) (*jwt.Token, error) {
	return jwt.ParseWithClaims(payload.Token, &TokenPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
}
