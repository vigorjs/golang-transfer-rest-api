package model

import (
	"gorm.io/gorm"
)

type Merchant struct {
	gorm.Model
	UserID       uint		   `json:"user_id" gorm:"unique"`
	User         User          `json:"-" gorm:"foreignKey:UserID"`
	Name         string		   `json:"name"`
	Transactions []Transaction `json:"-" gorm:"foreignKey:MerchantID"`
}