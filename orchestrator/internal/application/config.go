package application

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port                 string
	TimeAdditionMs       int
	TimeSubtractionMs    int
	TimeMultiplicationMs int
	TimeDivisionMs       int
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port:                 port,
		TimeAdditionMs:       getInt("TIME_ADDITION_MS", 1000),
		TimeSubtractionMs:    getInt("TIME_SUBTRACTION_MS", 1000),
		TimeMultiplicationMs: getInt("TIME_MULTIPLICATION_MS", 1000),
		TimeDivisionMs:       getInt("TIME_DIVISION_MS", 1000),
	}
}
