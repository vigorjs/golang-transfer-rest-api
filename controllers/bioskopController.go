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

func GetJadwalByFilmAndBioskop(c *gin.Context) {
	filmID := c.Param("filmID")
	bioskopID := c.Param("bioskopID")

	var jadwals []models.Jadwal

	result := initializers.DB.Preload("Film").Where("film_id = ? AND bioskop_id = ?", filmID, bioskopID).Find(&jadwals)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
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
		}
		response = append(response, jadwalData)
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func GetAllJadwals(c *gin.Context) {
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
