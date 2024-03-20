package initializers

import (
	"bioskop_golang/models"
	"fmt"
)

func MigrateDb() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Bioskop{})
	DB.AutoMigrate(&models.Booking{})
	DB.AutoMigrate(&models.Film{})
	DB.AutoMigrate(&models.Jadwal{})
	DB.AutoMigrate(&models.Kursi{})
	DB.AutoMigrate(&models.Transaksi{})
}
func SeedBioskop() {
	bioskops := []models.Bioskop{
		{Nama: "CGV", Lokasi: "Surabaya"},
		{Nama: "XXI", Lokasi: "Semarang"},
		{Nama: "Samehadaku", Lokasi: "Jogja"},
		{Nama: "Rajawali", Lokasi: "Bandung"},
		{Nama: "Trans Studio", Lokasi: "Jakarta"},
	}

	// Loop untuk menyimpan data kursi ke dalam db
	for _, bioskop := range bioskops {
		if err := DB.Create(&bioskop).Error; err != nil {
			fmt.Printf("Failed to seed data bioskop: %v\n", err)
		}
	}
}

func SeedFilm() {
	var bioskops []models.Bioskop
	if err := DB.Find(&bioskops).Error; err != nil {
		fmt.Printf("Failed to retrieve bioskops: %v\n", err)
		return
	}

	for _, bioskop := range bioskops {
		films := []models.Film{
			{Judul: "Harry Potter", BioskopID: bioskop.ID},
			{Judul: "Fast And Furious", BioskopID: bioskop.ID},
			{Judul: "Spongebob", BioskopID: bioskop.ID},
			{Judul: "Insidious", BioskopID: bioskop.ID},
			{Judul: "Shutter Island", BioskopID: bioskop.ID},
		}
		for _, film := range films {
			if err := DB.Create(&film).Error; err != nil {
				fmt.Printf("Failed to seed data film: %v\n", err)
			}
		}
	}
}
