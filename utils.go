package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func loadEnvFile(filepath string) error {
	err := godotenv.Load(filepath)
	if err != nil {
		return fmt.Errorf("error loading environment file: %w", err)
	}
	return nil
}
