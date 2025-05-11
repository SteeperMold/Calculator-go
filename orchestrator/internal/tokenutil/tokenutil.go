package tokenutil

import (
	"fmt"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	now := time.Now()

	accessClaims := jwt.MapClaims{
		"login": user.Login,
		"id":    user.Id,
		"nbf":   now.Unix(),
		"exp":   now.Add(time.Duration(expiry) * time.Hour).Unix(),
		"iat":   now.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	now := time.Now()

	refreshClaims := jwt.MapClaims{
		"id":  user.Id,
		"nbf": now.Unix(),
		"exp": now.Add(time.Duration(expiry) * time.Hour).Unix(),
		"iat": now.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func IsAuthorized(authToken string, secret string) (bool, error) {
	_, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func ExtractIDFromToken(authToken string, secret string) (int, error) {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	return int(claims["id"].(float64)), nil
}
