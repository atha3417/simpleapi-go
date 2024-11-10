package repositories

import (
	"simpleapi/config"
	"simpleapi/models"

	"gorm.io/gorm/clause"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Preload(clause.Associations).Find(&products).Error
	return products, err
}
func CreateProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}