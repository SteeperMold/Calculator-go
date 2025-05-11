package httpserver

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/bootstrap"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/service"
	"net/http"
	"time"
)

type UsersHandler struct {
	service        *service.UsersService
	contextTimeout time.Duration
}

func NewUsersHandler(db *sql.DB, config *bootstrap.Config) *UsersHandler {
	return &UsersHandler{
		service: service.NewUsersService(
			db,
			config.ContextTimeout,
			config.AccessTokenSecret,
			config.AccessTokenExpiryHours,
			config.RefreshTokenSecret,
			config.RefreshTokenExpiryHours,
		),
		contextTimeout: config.ContextTimeout,
	}
}

func (uh *UsersHandler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		r.Header.Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}

	err = uh.service.CreateUser(&req)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			http.Error(w, "invalid credentials", http.StatusUnprocessableEntity)
			return
		}
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			http.Error(w, "user already exists", http.StatusConflict)
			return
		}
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("OK"))
	if err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uh *UsersHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		r.Header.Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}

	accessToken, refreshToken, err := uh.service.Login(&req)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	res := &domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uh *UsersHandler) Profile(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("x-user-id").(int)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	profile, err := uh.service.GetProfileById(userId)
	if err != nil {
		http.Error(w, "failed to get profile", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uh *UsersHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		r.Header.Set("Allow", http.MethodPost)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.RefreshTokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}

	accessToken, refreshToken, err := uh.service.RefreshToken(req.RefreshToken)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	res := &domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "failed to encode data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
