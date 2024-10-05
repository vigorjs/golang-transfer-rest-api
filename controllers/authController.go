package controllers

import (
	"golang-rest-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: service.NewAuthService(),
	}
}

type SignUpRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// SignUp godoc
// @Summary Register a new user
// @Description Create a new user and return a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body SignUpRequest true "Sign Up request body"
// @Success 201 "User successfully created"
// @Failure 400 "Bad request"
// @Router /api/v1/auth/signup [post]
func (ctrl *AuthController) SignUp(c *gin.Context) {
	var body SignUpRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal mendapatkan request body"})
		return
	}

	token, user,  err := ctrl.authService.SignUp(body.Username, body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
        "message": "user berhasil dibuat",
        "data": gin.H{
            "token": token,
            "user":  user,
        },
    })
}

// Login godoc
// @Summary User login
// @Description Authenticate a user and return a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body LoginRequest true "Login request body"
// @Success 200 "Successfully logged in"
// @Failure 400 "Bad request"
// @Router /api/v1/auth/login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var body LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal mendapatkan request body"})
		return
	}

	token, user, err := ctrl.authService.Login(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
        "message": "berhasil login",
        "data": gin.H{
            "token": token,
            "user":  user,
        },
    })
}

// Logout godoc
// @Summary Logout a user
// @Description Logout a user with token
// @Tags Auth
// @Produce json
// @Param token header string true "Bearer token"
// @Success 200 "Logout successful"
// @Failure 401 "Unauthorized"
// @Failure 400 "Bad request"
// @Router /api/v1/logout [post]
func (ctrl *AuthController) Logout(c *gin.Context) {
	tokenString, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token tidak ditemukan"})
		return
	}

	if err := ctrl.authService.Logout(tokenString.(string)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "berhasil logout"})
}