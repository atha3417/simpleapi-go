package repositories

import (
	db "simpleapi/database"
	"simpleapi/models"

	"gorm.io/gorm/clause"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := db.DB.Preload(clause.Associations).Find(&products).Error
	return products, err
}
func CreateProduct(product *models.Product) error {
	return db.DB.Create(product).Error
}