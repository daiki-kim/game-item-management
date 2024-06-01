package models

type Trade struct {
	ID          uint `gorm:"primaryKey"`
	Is_Accepted bool
	ItemID      uint `gorm:"not null;foreignKey:ItemID"`
	FromUserID  uint `gorm:"not null;foreignKey:FromUserID"`
	ToUserID    uint `gorm:"not null;foreignKey:ToUserID"`
	// Item        Item `gorm:"foreignKey:ItemID"`
	// FromUser    User `gorm:"foreignKey:FromUserID"`
	// ToUser      User `gorm:"foreignKey:ToUserID"`
}
