package repositories

import (
	"simpleapi/config"
	"simpleapi/models"
)

func CreateAddress(address *models.Address) error {
	return config.DB.Create(address).Error
}