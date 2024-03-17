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
	initializers.SeedKursi()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validasi", middleware.RequireAuth, controllers.Validasi)
	r.GET("/api/jadwal/:filmID/:bioskopID", middleware.RequireAuth, controllers.GetJadwalByFilmAndBioskop)
	r.GET("/api/jadwal", middleware.RequireAuth, controllers.GetAllJadwals)
	r.POST("/api/bioskop", middleware.RequireAuth, controllers.CreateBioskop)
	r.POST("/api/film", middleware.RequireAuth, controllers.CreateFilm)
	r.POST("/api/jadwal", middleware.RequireAuth, controllers.CreateJadwal)
	r.PUT("/api/jadwal/:id_jadwal", middleware.RequireAuth, controllers.UpdateJadwal)
	r.DELETE("/api/jadwal/:id_jadwal", middleware.RequireAuth, controllers.DeleteJadwal)
	r.PUT("/api/film/:id", middleware.RequireAuth, controllers.UpdateFilm)
	r.DELETE("/api/film/:id", middleware.RequireAuth, controllers.DeleteFilm)
	r.PUT("/api/bioskop/:id", middleware.RequireAuth, controllers.UpdateBioskop)
	r.DELETE("/api/bioskop/:id", middleware.RequireAuth, controllers.DeleteBioskop)
	r.Run()
}
