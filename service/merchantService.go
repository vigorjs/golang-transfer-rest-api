package service

import (
	"golang-rest-api/initializers"
	"golang-rest-api/model"
)

type MerchantService struct{}

func NewMerchantService() *MerchantService {
	return &MerchantService{}
}

func (s *MerchantService) CreateMerchant(merchant model.Merchant) (*model.Merchant, error) {
	if err := initializers.DB.Create(&merchant).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (s *MerchantService) GetMerchantByID(id uint) (*model.Merchant, error) {
	var merchant model.Merchant
	if err := initializers.DB.First(&merchant, id).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (s *MerchantService) GetAllMerchants() ([]model.Merchant, error) {
	var merchants []model.Merchant
	if err := initializers.DB.Find(&merchants).Error; err != nil {
		return nil, err
	}
	return merchants, nil
}

func (s *MerchantService) UpdateMerchant(id uint, updatedMerchant model.Merchant) (*model.Merchant, error) {
	var merchant model.Merchant
	if err := initializers.DB.First(&merchant, id).Error; err != nil {
		return nil, err
	}

	merchant.Name = updatedMerchant.Name

	if err := initializers.DB.Save(&merchant).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (s *MerchantService) DeleteMerchant(id uint) error {
	if err := initializers.DB.Delete(&model.Merchant{}, id).Error; err != nil {
		return err
	}
	return nil
}
