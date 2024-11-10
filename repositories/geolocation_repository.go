package repositories

import (
	"simpleapi/config"
	"simpleapi/models"
)

func CreateGeolocation(geolocation *models.Geolocation) error {
	return config.DB.Create(geolocation).Error
}