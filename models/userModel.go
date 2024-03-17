package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	Email    string `gorm:"unique"`
	Password string
	Bookings []Booking // Relasi one-to-many ke tabel Booking
}
