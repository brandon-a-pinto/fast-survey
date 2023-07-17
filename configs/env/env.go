package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
