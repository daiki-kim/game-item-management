package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	UserID      uint `gorm:"not null;foreignKey:UserID"`
	// User        User `gorm:"foreignKey:UserID"`
}
