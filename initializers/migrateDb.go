package initializers

import (
	"golang-rest-api/model"
)

func MigrateDb() {
	DB.AutoMigrate(&model.User{}, &model.Merchant{}, &model.Token{}, &model.Transaction{})
}

// func SeedBioskop() {
// 	bioskops := []model.Bioskop{
// 		{Username: "CGV", Lokasi: "Surabaya"},
// 		{Username: "XXI", Lokasi: "Semarang"},
// 		{Username: "Samehadaku", Lokasi: "Jogja"},
// 		{Username: "Rajawali", Lokasi: "Bandung"},
// 		{Username: "Trans Studio", Lokasi: "Jakarta"},
// 	}

// 	// Loop untuk menyimpan data kursi ke dalam db
// 	for _, bioskop := range bioskops {
// 		if err := DB.Create(&bioskop).Error; err != nil {
// 			fmt.Printf("Failed to seed data bioskop: %v\n", err)
// 		}
// 	}
// }

// func SeedFilm() {
// 	var bioskops []model.Bioskop
// 	if err := DB.Find(&bioskops).Error; err != nil {
// 		fmt.Printf("Failed to retrieve bioskops: %v\n", err)
// 		return
// 	}

// 	for _, bioskop := range bioskops {
// 		films := []model.Film{
// 			{Judul: "Harry Potter", BioskopID: bioskop.ID},
// 			{Judul: "Fast And Furious", BioskopID: bioskop.ID},
// 			{Judul: "Spongebob", BioskopID: bioskop.ID},
// 			{Judul: "Insidious", BioskopID: bioskop.ID},
// 			{Judul: "Shutter Island", BioskopID: bioskop.ID},
// 		}
// 		for _, film := range films {
// 			if err := DB.Create(&film).Error; err != nil {
// 				fmt.Printf("Failed to seed data film: %v\n", err)
// 			}
// 		}
// 	}
// }
