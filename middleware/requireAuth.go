package middleware

import (
	"bioskop_golang/initializers"
	"bioskop_golang/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	//get cookie dari req
	tokenString, err := c.Cookie("Auth_Token")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//validasi
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signin method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//cek expired
		if float64(time.Now().Unix()) > claims["expired"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//index query user
		var user models.User
		initializers.DB.First(&user, claims["subject"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//masukkan ke request
		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
