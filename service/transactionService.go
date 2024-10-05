package service

import (
	"errors"
	"golang-rest-api/initializers"
	"golang-rest-api/model"

	"gorm.io/gorm"
)

type TransactionService struct{
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (s *TransactionService) CreateTransaction(userID, merchantID uint, grossAmount float64) (*model.Transaction, error) {
	var sender, receiver model.User
	var merchant model.Merchant

	err := initializers.DB.Transaction(func(tx *gorm.DB) error {
		// Mengambil sender (user yang melakukan transfer) dan mengunci baris
		if err := tx.Set("gorm:pessimistic_lock", true).First(&sender, userID).Error; err != nil {
			return err
		}

		// Memeriksa saldo sender
		if sender.Balance < grossAmount {
			return errors.New("saldo tidak mencukupi")
		}

		// Mengambil merchant dan pemiliknya (receiver)
		if err := tx.Preload("User").First(&merchant, merchantID).Error; err != nil {
			return err
		}

		receiver = merchant.User

		// Membuat transaksi
		transaction := model.Transaction{
			UserID:      userID,
			MerchantID:  merchantID,
			GrossAmount: grossAmount,
			Status:      "success",
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		// Memperbarui saldo sender
		sender.Balance -= grossAmount
		if err := tx.Save(&sender).Error; err != nil {
			return err
		}

		// Memperbarui saldo receiver (pemilik merchant)
		receiver.Balance += grossAmount
		if err := tx.Save(&receiver).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Mengambil transaksi yang baru dibuat
	var transaction model.Transaction
	if err := initializers.DB.Preload("User").Preload("Merchant").Last(&transaction, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (s *TransactionService) GetTransactionByID(id uint) (*model.Transaction, error) {
	var transaction model.Transaction
	if err := initializers.DB.Preload("User").Preload("Merchant").First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (s *TransactionService) GetAllTransactions() ([]model.Transaction, error) {
	var transactions []model.Transaction
	if err := initializers.DB.Preload("User").Preload("Merchant").Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *TransactionService) DeleteTransaction(id uint) error {
	if err := initializers.DB.Delete(&model.Transaction{}, id).Error; err != nil {
		return err
	}
	return nil
}