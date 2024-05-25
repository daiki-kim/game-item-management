package main

import (
	"game-item-management/config"
	"game-item-management/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
	config.ConnectDatabase()

	r := routes.TestRouter()
	r.Run("localhost:8080")
}
