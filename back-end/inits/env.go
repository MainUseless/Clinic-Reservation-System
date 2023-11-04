package inits

import (
	"log"

	"github.com/joho/godotenv"
)


func InitEnv(){
	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}
}