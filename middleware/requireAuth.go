package middleware

import (
	"golang-rest-api/initializers"
	"golang-rest-api/model"
	"golang-rest-api/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := tokenParts[1]
	tokenService := service.NewTokenService()
	token, err := tokenService.ValidateToken(tokenString)

	if err != nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user model.User
	if err := initializers.DB.First(&user, claims["subject"]).Error; err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var tokenRecord model.Token
	if err := initializers.DB.Where("user_id = ? AND token = ? AND status = ?", user.ID, tokenString, "active").First(&tokenRecord).Error; err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user", user)
	c.Set("userID", user.ID)
	c.Set("token", tokenString)
	c.Next()
}