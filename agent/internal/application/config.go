package application

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	OrchestratorAddress string
	ComputingPower      int
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

func getString(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

func NewConfigFromEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using os env variables or defaults")
	}

	return &Config{
		OrchestratorAddress: getString("ORCHESTRATOR_ADDRESS", "localhost:50051"),
		ComputingPower:      getInt("COMPUTING_POWER", 5),
	}
}
