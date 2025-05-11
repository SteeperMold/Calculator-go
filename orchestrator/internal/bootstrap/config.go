package bootstrap

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	HTTPPort                string
	GRPCPort                string
	ContextTimeout          time.Duration
	TimeAdditionMs          int
	TimeSubtractionMs       int
	TimeMultiplicationMs    int
	TimeDivisionMs          int
	AccessTokenSecret       string
	AccessTokenExpiryHours  int
	RefreshTokenSecret      string
	RefreshTokenExpiryHours int
}

func getString(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

func getInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return number
}

func NewConfigFromEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using os env vars or defaults")
	}

	return &Config{
		HTTPPort:                getString("HTTP_PORT", "8080"),
		GRPCPort:                getString("GRPC_PORT", "50051"),
		ContextTimeout:          time.Duration(getInt("CONTEXT_TIMEOUT_MS", 2000)) * time.Millisecond,
		TimeAdditionMs:          getInt("TIME_ADDITION_MS", 1000),
		TimeSubtractionMs:       getInt("TIME_SUBTRACTION_MS", 1000),
		TimeMultiplicationMs:    getInt("TIME_MULTIPLICATION_MS", 1000),
		TimeDivisionMs:          getInt("TIME_DIVISION_MS", 1000),
		AccessTokenSecret:       getString("ACCESS_TOKEN_SECRET", "not_very_secret"),
		AccessTokenExpiryHours:  getInt("ACCESS_TOKEN_EXPIRY_HOURS", 2),
		RefreshTokenSecret:      getString("REFRESH_TOKEN_SECRET", "not_very_secret"),
		RefreshTokenExpiryHours: getInt("REFRESH_TOKEN_EXPIRY_HOURS", 720),
	}
}
