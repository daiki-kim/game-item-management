package main

import (
	"game-item-management/config"
	"game-item-management/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
	db := config.ConnectDatabase()

	router := gin.Default()
	routes.SetupRouter(db, router)
	router.Run("localhost:8080")
}
