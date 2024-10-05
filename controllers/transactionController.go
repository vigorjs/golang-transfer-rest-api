package controllers

import (
	"golang-rest-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService *service.TransactionService
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		transactionService: service.NewTransactionService(),
	}
}

type TransactionRequest struct {
	MerchantID  uint    `json:"merchant_id" binding:"required"`
	GrossAmount float64 `json:"gross_amount" binding:"required,gt=0"`
}

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description Create a new transaction for a merchant
// @Tags Transaction
// @Accept json
// @Produce json
// @Param body body TransactionRequest true "Transaction request body"
// @Param Authorization header string true "Bearer token"
// @Success 200 "Transaction created successfully"
// @Failure 400 "Bad request"
// @Failure 401 "Unauthorized"
// @Router /api/v1/transaction [post]
func (tc *TransactionController) CreateTransaction(c *gin.Context) {
	var transactionRequest TransactionRequest

	if err := c.ShouldBindJSON(&transactionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user tidak terautentikasi"})
		return
	}

	transaction, err := tc.transactionService.CreateTransaction(userID.(uint), transactionRequest.MerchantID, transactionRequest.GrossAmount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, gin.H{"data": transaction, "message": "Transaction created successfully"})
}

// GetTransactionByID godoc
// @Summary Get a transaction by ID
// @Description Retrieve a transaction by its ID
// @Tags Transaction
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 "Sucess"
// @Failure 400 "Error"
// @Failure 404 "Error"
// @Router /api/v1/transaction/{id} [get]
func (tc *TransactionController) GetTransactionByID(c *gin.Context) {
	id := c.Param("id")
	transactionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	transaction, err := tc.transactionService.GetTransactionByID(uint(transactionID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "transaction tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}

// GetAllTransactions godoc
// @Summary Get all transactions
// @Description Retrieve a list of all transactions
// @Tags Transaction
// @Produce json
// @Success 200 "Success"
// @Failure 500 "Error"
// @Router /api/v1/transactions [get]
func (tc *TransactionController) GetAllTransactions(c *gin.Context) {
	transactions, err := tc.transactionService.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions})
}

// DeleteTransaction godoc
// @Summary Delete a transaction by ID
// @Description Delete a transaction by its ID
// @Tags Transaction
// @Param id path string true "Transaction ID"
// @Success 200 "Success Delete Transaction"
// @Failure 400 "Error"
// @Failure 404 "Error"
// @Router /api/v1/transaction/{id} [delete]
func (tc *TransactionController) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	transactionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := tc.transactionService.DeleteTransaction(uint(transactionID)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "transaction tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}