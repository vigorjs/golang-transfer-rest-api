package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	Email        string `gorm:"unique"`
	Password     string `json:"-"`
	Balance      float64
	Tokens       []Token       `json:"-" gorm:"foreignKey:UserID"`
	Merchants    []Merchant    `json:"-" gorm:"foreignKey:UserID"`
	Transactions []Transaction `json:"-" gorm:"foreignKey:UserID"`
}