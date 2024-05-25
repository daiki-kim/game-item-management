package models

import "gorm.io/gorm"

type Trade struct {
	gorm.Model
	ItemID     uint
	FromUserID uint
	ToUserID   uint
	Item       Item `gorm:"foreignKey:ItemID"`
	FromUser   User `gorm:"foreignKey:FromUserID"`
	ToUser     User `gorm:"foreignKey:ToUserID"`
}
