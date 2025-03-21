package handler

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	envPath := ".env"

	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		if isRunningAsExecutable() {
			execPath, err := os.Executable()
			if err != nil {
				log.Fatalf("Error retrieving executable path: %v", err)
			}
			execDir := filepath.Dir(execPath)
			envPath = filepath.Join(execDir, ".env")
		}
	}

	if err := godotenv.Overload(envPath); err != nil {
		log.Fatalf("Error loading .env file from %s: %v", envPath, err)
	}
}

func isRunningAsExecutable() bool {
	exePath, err := os.Executable()
	if err != nil {
		return true
	}

	tempDir := os.TempDir()
	return !strings.HasPrefix(exePath, tempDir)
}
