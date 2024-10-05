package service

import (
	"errors"
	"golang-rest-api/initializers"
	"golang-rest-api/model"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	tokenService *TokenService
}

func NewAuthService() *AuthService {
	return &AuthService{
		tokenService: NewTokenService(),
	}
}

func (s *AuthService) SignUp(username, email, password string) (string, model.User, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    if err != nil {
        return "", model.User{}, errors.New("gagal hash password")
    }

    user := model.User{
        Username: username,
        Email:    email,
        Password: string(hash),
        Balance:  9999999, // Set default balance
    }

    if err := initializers.DB.Create(&user).Error; err != nil {
        return "", model.User{}, err
    }

    token, err := s.tokenService.CreateToken(user.ID)
    return token, user, err
}


func (s *AuthService) Login(email, password string) (string, model.User, error) {
    var user model.User
    if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return "", model.User{}, errors.New("email tidak ditemukan")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", model.User{}, errors.New("password salah")
    }

    token, err := s.tokenService.CreateToken(user.ID)
    return token, user, err 
}

func (s *AuthService) Logout(tokenString string) error {
	return s.tokenService.Logout(tokenString)
}