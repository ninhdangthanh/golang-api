package services

import (
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
