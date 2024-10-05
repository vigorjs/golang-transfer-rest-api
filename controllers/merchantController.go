package controllers

import (
	"golang-rest-api/initializers"
	"golang-rest-api/model"
	"golang-rest-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var merchantService = service.NewMerchantService()

type MerchantRequest struct {
	Name string `json:"name" binding:"required"`
}

// CreateMerchant godoc
// @Summary Create a new merchant
// @Description Create a new merchant for an authenticated user
// @Tags Merchant
// @Accept json
// @Produce json
// @Param body body MerchantRequest true "Merchant request body"
// @Param Authorization header string true "Bearer token"
// @Success 200 "Merchant created successfully"
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
// @Router /api/v1/merchants [post]
func CreateMerchant(c *gin.Context) {
	var req MerchantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal mendapatkan request body"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user tidak terautentikasi"})
		return
	}

	merchant := model.Merchant{
		Name:   req.Name,
		UserID: userID.(uint),
	}

	createdMerchant, err := merchantService.CreateMerchant(merchant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var fullMerchant model.Merchant
	if err := initializers.DB.Preload("User").First(&fullMerchant, createdMerchant.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengambil data merchant lengkap"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": fullMerchant, "message": "Merchant created successfully"})
}


// GetMerchantByID godoc
// @Summary Get a merchant by ID
// @Description Retrieve a merchant by its ID
// @Tags Merchant
// @Produce json
// @Param id path string true "Merchant ID"
// @Success 200 "Merchant found"
// @Failure 400 "Invalid ID"
// @Failure 404 "Merchant not found"
// @Router /api/v1/merchants/{id} [get]
func GetMerchantByID(c *gin.Context) {
	id := c.Param("id")
	merchantID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	merchant, err := merchantService.GetMerchantByID(uint(merchantID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "merchant tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"merchant": merchant})
}

// GetAllMerchants godoc
// @Summary Get all merchants
// @Description Retrieve a list of all merchants
// @Tags Merchant
// @Produce json
// @Success 200 "List of merchants"
// @Failure 500 "Internal server error"
// @Router /api/v1/merchants [get]
func GetAllMerchants(c *gin.Context) {
	merchants, err := merchantService.GetAllMerchants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"merchants": merchants})
}

// UpdateMerchant godoc
// @Summary Update a merchant by ID
// @Description Update an existing merchant by its ID
// @Tags Merchant
// @Accept json
// @Produce json
// @Param id path string true "Merchant ID"
// @Param body body MerchantRequest true "Merchant update request body"
// @Success 200 "Merchant updated successfully"
// @Failure 400 "Invalid ID or request body"
// @Failure 404 "Merchant not found"
// @Router /api/v1/merchants/{id} [put]
func UpdateMerchant(c *gin.Context) {
	id := c.Param("id")
	merchantID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var req MerchantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal mendapatkan request body"})
		return
	}

	updatedMerchant := model.Merchant{
		Name: req.Name,
	}

	merchant, err := merchantService.UpdateMerchant(uint(merchantID), updatedMerchant)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "merchant tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": merchant, "message": "Merchant updated successfully"})
}


// DeleteMerchant godoc
// @Summary Delete a merchant by ID
// @Description Delete a merchant by its ID
// @Tags Merchant
// @Param id path string true "Merchant ID"
// @Success 200 "Merchant deleted successfully"
// @Failure 400 "Invalid ID"
// @Failure 404 "Merchant not found"
// @Router /api/v1/merchants/{id} [delete]
func DeleteMerchant(c *gin.Context) {
	id := c.Param("id")
	merchantID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := merchantService.DeleteMerchant(uint(merchantID)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "merchant tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "merchant berhasil dihapus"})
}
