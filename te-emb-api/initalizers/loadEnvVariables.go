package initalizers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	run_mode := os.Getenv("MODE")
	if run_mode == "docker-mode" {
		return
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
