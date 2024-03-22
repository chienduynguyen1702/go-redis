package initialize

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVarFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
