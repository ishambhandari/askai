package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const environment = "DEV"

func getEnv(key string) string {
	if environment == "DEV" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Error getting home directory: %v", err)
		}
		envFilePath := filepath.Join(homeDir, ".config", "askai.env")
		err = godotenv.Load(envFilePath)
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	value := os.Getenv(key)
	if value == "" {
		log.Printf("Warning: Environment variable %s is not set", key)
	}
	return value
}
