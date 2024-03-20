package models

import (
	"time"

	"gorm.io/gorm"
)

type Bioskop struct {
	gorm.Model
	Nama    string   // Bioskop's name
	Lokasi  string   // Bioskop's location
	Films   []Film   // One-to-many relationship to films
	Jadwals []Jadwal // One-to-many relationship to jadwals
}

type Film struct {
	gorm.Model
	Judul     string   // Film's title
	BioskopID uint     // Foreign key to Bioskop's ID
	Bioskop   Bioskop  // Many-to-one relationship to Bioskop
	Jadwals   []Jadwal // One-to-many relationship to jadwals
}

type Jadwal struct {
	gorm.Model
	FilmID     uint      // Foreign key to Film's ID
	BioskopID  uint      // Foreign key to Bioskop's ID
	Tanggal    time.Time // Jadwal's date
	JamTayang  time.Time // Jadwal's start time
	JamSelesai time.Time // Jadwal's end time
	Kursi      []Kursi   // One-to-many relationship to kursis
	Film       Film
	Bioskop    Bioskop
}

type Kursi struct {
	gorm.Model
	JadwalID    uint   // Foreign key to Jadwal's ID
	NomorKursi  string // Kursi's seat number
	IsAvailable bool   `gorm:"default:true"` // Kursi's availability
	Jadwal      Jadwal
}

type Booking struct {
	gorm.Model
	UserID        uint      // Foreign key to User's ID
	JadwalID      uint      // Foreign key to Jadwal's ID
	PaymentStatus string    // Booking's payment status
	BookingDate   time.Time // Booking's date and time
	User          User      // Many-to-one relationship to User
	Kursi         []Kursi   `gorm:"many2many:booking_kursis"` // Many-to-many relationship to kursis
}

type Transaksi struct {
	gorm.Model
	BookingID     uint      // Foreign key to Booking's ID
	Total         float64   // Transaksi's total
	PaymentDate   time.Time // Transaksi's payment date and time
	PaymentMethod string    // Transaksi's payment method
}
