package controllers

import (
	"bioskop_golang/initializers"
	"bioskop_golang/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//get requestnya dan validasi
	var body struct {
		Nama     string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "gagal get request body",
		})
		return
	}

	//hash pw req
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "gagal hash password",
		})
		return
	}

	//buat user di db
	user := models.User{Nama: body.Nama, Email: body.Email, Password: string(hash)}
	create := initializers.DB.Create(&user)
	if create.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "gagal create user",
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"message": "data berhasil dibuat",
	})
}

func Login(c *gin.Context) {
	// Get request dan validasi
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "gagal get request body",
		})
		return
	}

	//index query user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email tidak ditemukan",
		})
		return
	}

	// compare pw
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password Salah",
		})
		return
	}

	//generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.ID,
		"expired": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal generate token!",
		})
		return
	}

	//response
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth_Token", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil Login",
	})

}

func Validasi(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
