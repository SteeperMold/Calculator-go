package domain

import (
	"fmt"
	"time"
)

var (
	ErrUserAlreadyExists  = fmt.Errorf("user already exists")
	ErrInvalidCredentials = fmt.Errorf("invalid credentials")
	ErrUserDoesntExist    = fmt.Errorf("user doesn't exist")
)

type User struct {
	Id           int64
	Login        string
	PasswordHash string
	CreationTime time.Time
}

type Profile struct {
	Login string `json:"login"`
}

type SignupRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
