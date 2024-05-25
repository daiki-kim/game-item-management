package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
}
