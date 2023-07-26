package env

import (
	"log"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
