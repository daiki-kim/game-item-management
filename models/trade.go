package models

import "gorm.io/gorm"

type Trade struct {
	gorm.Model
	ItemID     uint `gorm:"not null"`
	FromUserID uint `gorm:"not null"`
	ToUserID   uint `gorm:"not null"`
	Item       Item `gorm:"foreignKey:ItemID"`
	FromUser   User `gorm:"foreignKey:FromUserID"`
	ToUser     User `gorm:"foreignKey:ToUserID"`
}
