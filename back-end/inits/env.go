package inits

import (
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	// Load connection string from .env file
	
	if os.Getenv("rabbitmq_url") == "" {
		godotenv.Load()
	}else if os.Getenv("mysql_url") == "" {
		godotenv.Load()
	}else if os.Getenv("jwt_secret") == "" {
		godotenv.Load()
	}else if os.Getenv("port") == "" {
		godotenv.Load()
	}

}
