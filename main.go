package main

import (
	"bioskop_golang/controllers"
	"bioskop_golang/initializers"
	"bioskop_golang/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Conn()
	initializers.MigrateDb()
	// hilangkan comment untuk seed data awal setelah itu command lg
	// initializers.SeedBioskop()
	// initializers.SeedFilm()
}

func main() {
	r := gin.Default()

	//auth
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validasi", middleware.RequireAuth, controllers.Validasi)
	//jadwal
	r.GET("/api/jadwal", middleware.RequireAuth, controllers.GetAllJadwal)
	r.GET("/api/jadwal/id/:jadwalID", middleware.RequireAuth, controllers.GetJadwalByID)
	r.GET("/api/jadwal/:filmID/:bioskopID", middleware.RequireAuth, controllers.GetJadwalWithKursiByFilmAndBioskop)
	r.POST("/api/jadwal", middleware.RequireAuth, controllers.CreateJadwal)
	r.PUT("/api/jadwal/:id_jadwal", middleware.RequireAuth, controllers.UpdateJadwal)
	r.DELETE("/api/jadwal/:id_jadwal", middleware.RequireAuth, controllers.DeleteJadwal)
	//Booking
	r.POST("/api/booking", middleware.RequireAuth, controllers.CreateBooking)
	r.GET("/api/booking", middleware.RequireAuth, controllers.GetAllBooking)
	//bioskop
	r.POST("/api/bioskop", middleware.RequireAuth, controllers.CreateBioskop)
	r.PUT("/api/bioskop/:id", middleware.RequireAuth, controllers.UpdateBioskop)
	r.DELETE("/api/bioskop/:id", middleware.RequireAuth, controllers.DeleteBioskop)
	//film
	r.POST("/api/film", middleware.RequireAuth, controllers.CreateFilm)
	r.PUT("/api/film/:id", middleware.RequireAuth, controllers.UpdateFilm)
	r.DELETE("/api/film/:id", middleware.RequireAuth, controllers.DeleteFilm)
	r.Run()
}
