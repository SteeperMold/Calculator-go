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

func NewConfigFromEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using os env variables or defaults")
	}

	orchestratorAddress := os.Getenv("ORCHESTRATOR_ADDRESS")
	if orchestratorAddress == "" {
		orchestratorAddress = "http://localhost:8080"
	}

	return &Config{
		OrchestratorAddress: orchestratorAddress,
		ComputingPower:      getInt("COMPUTING_POWER", 5),
	}
}
