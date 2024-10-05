package model

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UserID    uint
	User      User `json:"-" gorm:"foreignKey:UserID"`
	Token     string
	Status    string
	ExpiresAt time.Time
}