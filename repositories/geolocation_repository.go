package repositories

import (
	db "simpleapi/database"
	"simpleapi/models"
)

func CreateGeolocation(geolocation *models.Geolocation) error {
	return db.DB.Create(geolocation).Error
}