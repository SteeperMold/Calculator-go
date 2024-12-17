package application

import "os"

type Config struct {
	Address string
}

func NewConfigFromEnv() *Config {
	address := os.Getenv("PORT")
	if address == "" {
		address = "8080"
	}

	return &Config{
		Address: address,
	}
}
