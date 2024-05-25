package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	UserID      uint `gorm:"not null"`
	User        User `gorm:"foreignKey:UserID"`
}
