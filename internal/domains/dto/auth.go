package dto

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type RegisterRequest struct {
	NIK      string `json:"nik" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// StringOrNumber accepts either a JSON string or number and stores as string
type StringOrNumber string

func (s *StringOrNumber) UnmarshalJSON(b []byte) error {
	// try string
	var str string
	if err := json.Unmarshal(b, &str); err == nil {
		*s = StringOrNumber(str)
		return nil
	}
	// try number
	var num float64
	if err := json.Unmarshal(b, &num); err == nil {
		// format integer without decimal when appropriate
		if num == float64(int64(num)) {
			*s = StringOrNumber(strconv.FormatInt(int64(num), 10))
		} else {
			*s = StringOrNumber(fmt.Sprintf("%v", num))
		}
		return nil
	}
	return fmt.Errorf("invalid type for StringOrNumber")
}

func (s StringOrNumber) String() string { return string(s) }

type LoginRequest struct {
	NIK      StringOrNumber `json:"nik" binding:"required"`
	Password string         `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
