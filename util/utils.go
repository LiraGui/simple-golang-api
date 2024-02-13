package util

import (
	"log"

	"github.com/joho/godotenv"
)

func VerifyEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error to get .env --> ", err)
	}

	log.Println("=====> .env Loaded!")
}
