package repositories

import (
	db "simpleapi/database"
	"simpleapi/models"
)

func CreateAddress(address *models.Address) error {
	return db.DB.Create(address).Error
}