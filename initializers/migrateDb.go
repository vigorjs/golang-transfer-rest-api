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
func SeedKursi() {
	// Data kursi yang akan di-generate
	kursis := []models.Kursi{
		{JadwalID: 1, NomorKursi: "A1", IsAvailable: true},
		{JadwalID: 1, NomorKursi: "A2", IsAvailable: true},
		{JadwalID: 1, NomorKursi: "A3", IsAvailable: true},
		{JadwalID: 2, NomorKursi: "B1", IsAvailable: true},
		{JadwalID: 2, NomorKursi: "B2", IsAvailable: true},
		{JadwalID: 1, NomorKursi: "A4", IsAvailable: true},
		{JadwalID: 1, NomorKursi: "A5", IsAvailable: true},
		{JadwalID: 1, NomorKursi: "A6", IsAvailable: true},
		{JadwalID: 2, NomorKursi: "B3", IsAvailable: true},
		{JadwalID: 2, NomorKursi: "B4", IsAvailable: true},
	}

	// Loop untuk menyimpan data kursi ke dalam database
	for _, kursi := range kursis {
		if err := DB.Create(&kursi).Error; err != nil {
			fmt.Printf("Failed to create kursi: %v\n", err)
		}
	}

	fmt.Println("Seeder for Kursi completed")
}
