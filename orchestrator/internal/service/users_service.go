package service

import (
	"context"
	"database/sql"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/repository"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/tokenutil"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UsersService struct {
	repository              *repository.UsersRepository
	contextTimeout          time.Duration
	accessTokenSecret       string
	accessTokenExpiryHours  int
	refreshTokenSecret      string
	refreshTokenExpiryHours int
}

func NewUsersService(
	db *sql.DB,
	contextTimeout time.Duration,
	accessSecret string,
	accessExpiry int,
	refreshSecret string,
	refreshExpiry int,
) *UsersService {
	return &UsersService{
		repository:              repository.NewUsersRepository(db),
		contextTimeout:          contextTimeout,
		accessTokenSecret:       accessSecret,
		accessTokenExpiryHours:  accessExpiry,
		refreshTokenSecret:      refreshSecret,
		refreshTokenExpiryHours: refreshExpiry,
	}
}

func (us *UsersService) CreateUser(userData *domain.SignupRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), us.contextTimeout)
	defer cancel()

	if len(userData.Login) < 3 || len(userData.Login) > 32 {
		return domain.ErrInvalidCredentials
	}

	if len(userData.Password) < 5 || len(userData.Password) > 255 {
		return domain.ErrInvalidCredentials
	}

	_, err := us.repository.GetUserByLogin(ctx, userData.Login)
	if err == nil {
		return domain.ErrUserAlreadyExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &domain.User{
		Login:        userData.Login,
		PasswordHash: string(passwordHash),
	}

	err = us.repository.CreateUser(ctx, newUser)
	return err
}

func (us *UsersService) Login(userData *domain.LoginRequest) (accessToken string, refreshToken string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), us.contextTimeout)
	defer cancel()

	user, err := us.repository.GetUserByLogin(ctx, userData.Login)
	if err != nil {
		return "", "", domain.ErrUserDoesntExist
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userData.Password))
	if err != nil {
		return "", "", domain.ErrInvalidCredentials
	}

	accessToken, err = tokenutil.CreateAccessToken(user, us.accessTokenSecret, us.accessTokenExpiryHours)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = tokenutil.CreateRefreshToken(user, us.refreshTokenSecret, us.refreshTokenExpiryHours)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (us *UsersService) GetProfileById(userId int) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), us.contextTimeout)
	defer cancel()

	user, err := us.repository.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &domain.Profile{Login: user.Login}, nil
}

func (us *UsersService) RefreshToken(oldRefreshToken string) (accessToken string, refreshToken string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), us.contextTimeout)
	defer cancel()

	id, err := tokenutil.ExtractIDFromToken(oldRefreshToken, us.accessTokenSecret)
	if err != nil {
		return "", "", err
	}

	user, err := us.repository.GetUserById(ctx, id)
	if err != nil {
		return "", "", err
	}

	accessToken, err = tokenutil.CreateAccessToken(user, us.accessTokenSecret, us.accessTokenExpiryHours)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = tokenutil.CreateRefreshToken(user, us.refreshTokenSecret, us.refreshTokenExpiryHours)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
