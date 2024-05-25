package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"game-item-management/models"
)

func ConnectDatabase() (db *gorm.DB) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %s", err.Error())
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Trade{},
	)
	if err != nil {
		log.Fatalf("Error migrating database:%s", err.Error())
	}
	return db
}
