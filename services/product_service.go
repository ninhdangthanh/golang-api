package services

import (
	"errors"

	"github.com/example/intern/models"
	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db}
}

func (s *ProductService) CreateProduct(product *models.ProductModel) error {
	if err := s.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (s *ProductService) GetOwnProducts(userID uint) ([]models.ProductModel, error) {
	var products []models.ProductModel
	if err := s.db.Where("user_id = ?", userID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) DeleteOwnProduct(userID uint, productID uint) error {
	result := s.db.Where("user_id = ? AND id = ?", userID, productID).Delete(&models.ProductModel{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("product not found or does not belong to the user")
	}

	return nil
}

func (s *ProductService) UpdateOwnProduct(userID uint, productID uint, updatedProduct models.ProductModel) (*models.ProductModel, error) {
	var product models.ProductModel

	if err := s.db.Where("user_id = ? AND id = ?", userID, productID).First(&product).Error; err != nil {
		return nil, errors.New("product not found or does not belong to the user")
	}

	if err := s.db.Model(&product).Updates(updatedProduct).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
