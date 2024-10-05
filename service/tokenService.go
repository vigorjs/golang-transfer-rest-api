package service

import (
	"errors"
	"golang-rest-api/initializers"
	"golang-rest-api/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type TokenService struct{}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (s *TokenService) CreateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": userID,
		"expired": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	// Nonaktifkan token lama yang masih aktif
	if err := initializers.DB.Model(&model.Token{}).Where("user_id = ? AND status = ?", userID, "active").Update("status", "inactive").Error; err != nil {
		return "", err
	}

	// Simpan token baru
	tokenRecord := model.Token{
		UserID:    userID,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
		Status:    "active",
	}
	if err := initializers.DB.Create(&tokenRecord).Error; err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *TokenService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return token, err
}

func (s *TokenService) Logout(tokenString string) error {
	var tokenRecord model.Token
	err := initializers.DB.Where("token = ? AND status = ?", tokenString, "active").First(&tokenRecord).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("active token not found")
		}
		return err
	}

	return initializers.DB.Model(&tokenRecord).Update("status", "inactive").Error
}