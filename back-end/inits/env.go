package inits

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	// Load connection string from .env file
	godotenv.Load()

	if os.Getenv("rabbitmq_url") == "" {
		log.Fatal("rabbitmq_url is not set")
	}
	if os.Getenv("mysql_url") == "" {
		log.Fatal("mysql_url is not set")
	}
	if os.Getenv("jwt_secret") == "" {
		log.Fatal("jwt_secret is not set")
	}
	if os.Getenv("port") == "" {
		log.Fatal("port is not set")
	}

}
