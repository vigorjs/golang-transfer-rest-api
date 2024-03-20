package controllers

import (
	"bioskop_golang/initializers"
	"bioskop_golang/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateBioskop(c *gin.Context) {
	var body struct {
		Nama   string
		Lokasi string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "gagal get request body",
		})
		return
	}
	bioskop := models.Bioskop{Nama: body.Nama, Lokasi: body.Lokasi}

	create := initializers.DB.Create(&bioskop)
	if create.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "gagal create user",
		})
		return
	}
	c.JSON(http.StatusCreated, bioskop)
}

func CreateFilm(c *gin.Context) {
	var body struct {
		Judul     string `json:"judul"`
		BioskopID uint   `json:"bioskop_id"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get request body",
		})
		return
	}

	film := models.Film{Judul: body.Judul, BioskopID: body.BioskopID}
	result := initializers.DB.Create(&film)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, film)
}

func UpdateBioskop(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Bioskop tidak valid"})
		return
	}

	var reqBody struct {
		Nama   string `json:"nama"`
		Lokasi string `json:"lokasi"`
	}

	if err := c.Bind(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var bioskop models.Bioskop
	result := initializers.DB.First(&bioskop, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	if reqBody.Nama != "" {
		bioskop.Nama = reqBody.Nama
	}
	if reqBody.Lokasi != "" {
		bioskop.Lokasi = reqBody.Lokasi
	}

	initializers.DB.Save(&bioskop)
	c.JSON(http.StatusOK, bioskop)
}

func DeleteBioskop(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Bioskop tidak valid"})
		return
	}

	var bioskop models.Bioskop
	result := initializers.DB.First(&bioskop, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	initializers.DB.Delete(&bioskop)
	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil dihapus"})
}

func UpdateFilm(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Film tidak valid"})
		return
	}

	var reqBody struct {
		Judul     string `json:"judul"`
		BioskopID uint   `json:"bioskop_id"`
	}

	if err := c.Bind(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var film models.Film
	result := initializers.DB.First(&film, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Film tidak ditemukan"})
		return
	}

	if reqBody.Judul != "" {
		film.Judul = reqBody.Judul
	}
	if reqBody.BioskopID != 0 {
		film.BioskopID = reqBody.BioskopID
	}

	initializers.DB.Save(&film)
	c.JSON(http.StatusOK, film)
}

func DeleteFilm(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Film tidak valid"})
		return
	}

	var film models.Film
	result := initializers.DB.First(&film, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Film tidak ditemukan"})
		return
	}

	initializers.DB.Delete(&film)
	c.JSON(http.StatusOK, gin.H{"message": "Film berhasil dihapus"})
}

// jadwal
func CreateJadwal(c *gin.Context) {
	var body struct {
		FilmID     uint   `json:"film_id"`
		BioskopID  uint   `json:"bioskop_id"`
		Tanggal    string `json:"tanggal"`
		JamTayang  string `json:"jam_tayang"`
		JamSelesai string `json:"jam_selesai"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get request body",
		})
		return
	}

	tanggal, err := time.Parse("2006-01-02", body.Tanggal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid date format for tanggal",
		})
		return
	}

	jamTayang, err := time.Parse("15:04", body.JamTayang)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid time format for jam_tayang",
		})
		return
	}

	jamSelesai, err := time.Parse("15:04", body.JamSelesai)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid time format for jam_selesai",
		})
		return
	}

	jadwal := models.Jadwal{
		FilmID:     body.FilmID,
		BioskopID:  body.BioskopID,
		Tanggal:    tanggal,
		JamTayang:  jamTayang,
		JamSelesai: jamSelesai,
	}

	result := initializers.DB.Create(&jadwal)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	jadwalID := jadwal.ID
	kursis := []models.Kursi{
		{JadwalID: jadwalID, NomorKursi: "A1", IsAvailable: true},
		{JadwalID: jadwalID, NomorKursi: "A2", IsAvailable: true},
		{JadwalID: jadwalID, NomorKursi: "A3", IsAvailable: true},
		{JadwalID: jadwalID, NomorKursi: "A4", IsAvailable: true},
		{JadwalID: jadwalID, NomorKursi: "A5", IsAvailable: true},
		{JadwalID: jadwalID, NomorKursi: "B1", IsAvailable: true},
		{JadwalID: jadwalID, NomorKursi: "B2", IsAvailable: true},
		{JadwalID: jadwalID, NomorKursi: "B3", IsAvailable: true},
		{JadwalID: jadwalID, NomorKursi: "B4", IsAvailable: true},
		{JadwalID: jadwalID, NomorKursi: "B5", IsAvailable: true},
	}
	for _, kursi := range kursis {
		if err := initializers.DB.Create(&kursi).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusCreated, jadwal)
}

func UpdateJadwal(c *gin.Context) {
	id := c.Param("id_jadwal")

	var body struct {
		FilmID     uint   `json:"film_id"`
		BioskopID  uint   `json:"bioskop_id"`
		Tanggal    string `json:"tanggal"`
		JamTayang  string `json:"jam_tayang"`
		JamSelesai string `json:"jam_selesai"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to get request body",
		})
		return
	}

	tanggal, err := time.Parse("2006-01-02", body.Tanggal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid date format for tanggal",
		})
		return
	}

	jamTayang, err := time.Parse("15:04", body.JamTayang)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid time format for jam_tayang",
		})
		return
	}

	jamSelesai, err := time.Parse("15:04", body.JamSelesai)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid time format for jam_selesai",
		})
		return
	}

	var jadwal models.Jadwal
	result := initializers.DB.First(&jadwal, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "jadwal not found",
		})
		return
	}

	jadwal.FilmID = body.FilmID
	jadwal.BioskopID = body.BioskopID
	jadwal.Tanggal = tanggal
	jadwal.JamTayang = jamTayang
	jadwal.JamSelesai = jamSelesai

	result = initializers.DB.Save(&jadwal)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update jadwal",
		})
		return
	}

	c.JSON(http.StatusOK, jadwal)
}

func DeleteJadwal(c *gin.Context) {
	id := c.Param("id_jadwal")

	var jadwal models.Jadwal
	result := initializers.DB.Commit().Delete(&jadwal, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "jadwal not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "jadwal deleted successfully",
	})
}

func GetJadwalWithKursiByFilmAndBioskop(c *gin.Context) {
	filmID := c.Param("filmID")
	bioskopID := c.Param("bioskopID")

	var jadwals []models.Jadwal

	result := initializers.DB.Preload("Kursi").Preload("Film").Where("film_id = ? AND bioskop_id = ?", filmID, bioskopID).Find(&jadwals)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var response []gin.H
	for _, jadwal := range jadwals {
		kursiAvailable := make([]interface{}, 0)
		for _, kursi := range jadwal.Kursi {
			kursiData := []interface{}{kursi.NomorKursi, kursi.IsAvailable}
			kursiAvailable = append(kursiAvailable, kursiData)
		}
		jadwalData := gin.H{
			"ID":             jadwal.ID,
			"Tanggal":        jadwal.Tanggal,
			"JamTayang":      jadwal.JamTayang,
			"JamSelesai":     jadwal.JamSelesai,
			"JudulFilm":      jadwal.Film.Judul,
			"KursiAvailable": kursiAvailable,
		}
		response = append(response, jadwalData)
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func GetAllJadwal(c *gin.Context) {
	var jadwals []models.Jadwal

	if err := initializers.DB.Preload("Film").Preload("Bioskop").Find(&jadwals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H
	for _, jadwal := range jadwals {
		jadwalData := gin.H{
			"ID":         jadwal.ID,
			"Tanggal":    jadwal.Tanggal,
			"JamTayang":  jadwal.JamTayang,
			"JamSelesai": jadwal.JamSelesai,
			"JudulFilm":  jadwal.Film.Judul,
			"Bioskop":    jadwal.Bioskop.Nama,
		}
		response = append(response, jadwalData)
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func GetJadwalByID(c *gin.Context) {
	jadwalID := c.Param("jadwalID")

	var jadwal models.Jadwal

	result := initializers.DB.Preload("Film").Preload("Bioskop").Preload("Kursi").First(&jadwal, jadwalID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	kursiAvailable := make([]interface{}, 0)
	for _, kursi := range jadwal.Kursi {
		kursiData := []interface{}{kursi.NomorKursi, kursi.IsAvailable}
		kursiAvailable = append(kursiAvailable, kursiData)
	}

	response := gin.H{
		"ID":             jadwal.ID,
		"Tanggal":        jadwal.Tanggal,
		"JamTayang":      jadwal.JamTayang,
		"JamSelesai":     jadwal.JamSelesai,
		"JudulFilm":      jadwal.Film.Judul,
		"KursiAvailable": kursiAvailable,
		"Bioskop":        jadwal.Bioskop.Nama,
		"Lokasi":         jadwal.Bioskop.Lokasi,
	}

	c.JSON(http.StatusOK, response)
}

// Booking
func CreateBooking(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	userID := user.(models.User).ID

	var body struct {
		UserID        uint   `json:"user_id"`
		JadwalID      uint   `json:"jadwal_id"`
		BookingDate   string `json:"booking_date"`
		PaymentStatus string `json:"payment_status"`
		KursiID       []uint `json:"kursi_id"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to send Book request",
		})
		return
	}
	// Validasi jika kursi yg di req diluar dari kursi_id yang sesuai dengan jadwal yang di request
	for _, kursiID := range body.KursiID {
		var count int64
		if err := initializers.DB.Model(&models.Kursi{}).Where("id = ? AND jadwal_id = ?", kursiID, body.JadwalID).Count(&count).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check kursi"})
			return
		}
		if count == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "kursi yg di req diluar dari kursi_id yang sesuai dengan jadwal yang di request"})
			return
		}
	}
	bookingdate, err := time.Parse("2006-01-02", body.BookingDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid date format for tanggal",
		})
		return
	}

	booking := models.Booking{
		UserID:        userID,
		JadwalID:      body.JadwalID,
		PaymentStatus: body.PaymentStatus,
		BookingDate:   bookingdate,
	}

	// Menyimpan booking ke dalam database
	if err := initializers.DB.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create booking"})
		return
	}

	// Menyimpan kursi-kursi yang dipilih ke dalam tabel booking_kursis
	for _, kursiID := range body.KursiID {
		var kursi models.Kursi
		//validasi ketersediaan kursi
		if err := initializers.DB.Where("id = ? AND jadwal_id = ?", kursiID, body.JadwalID).First(&kursi).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to find kursi"})
			return
		}
		if !kursi.IsAvailable {
			c.JSON(http.StatusBadRequest, gin.H{"error": "kursi is not available"})
			return
		}

		if err := initializers.DB.Exec("INSERT INTO booking_kursis (booking_id, kursi_id) VALUES (?, ?)", booking.ID, kursiID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create booking_kursis"})
			return
		}
		if err := initializers.DB.Model(&models.Kursi{}).Where("id = ? AND jadwal_id = ?", kursiID, booking.JadwalID).Update("is_available", false).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update kursi status"})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Booking created successfully", "data": booking})
}

func GetAllBooking(c *gin.Context) {
	var bookings []models.Booking

	if err := initializers.DB.Preload("Kursi").Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H
	for _, booking := range bookings {

		bookedKursi := make([]string, 0)
		for _, kursi := range booking.Kursi {
			bookedKursi = append(bookedKursi, kursi.NomorKursi)
		}

		bookingData := gin.H{
			"ID":            booking.ID,
			"UserID":        booking.UserID,
			"JadwalID":      booking.JadwalID,
			"BookingDate":   booking.BookingDate,
			"PaymentStatus": booking.PaymentStatus,
			"Kursi":         bookedKursi,
		}
		response = append(response, bookingData)
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
