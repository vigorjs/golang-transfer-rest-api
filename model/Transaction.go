package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID      uint
	User        User     `json:"-" gorm:"foreignKey:UserID"`
	MerchantID  uint
	Merchant    Merchant `json:"-" gorm:"foreignKey:MerchantID"`
	GrossAmount float64
	Status      string
}