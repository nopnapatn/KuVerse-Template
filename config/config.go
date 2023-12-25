// config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}

func GetPort() string {
	return os.Getenv("PORT")
}

func GetApiKey() string {
	return os.Getenv("API_KEY")
}
