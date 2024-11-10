package repositories

import (
	"simpleapi/config"
	"simpleapi/models"
)

func CreateName(name *models.Name) error {
	return config.DB.Create(name).Error
}