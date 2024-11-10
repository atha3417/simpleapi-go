package repositories

import (
	db "simpleapi/database"
	"simpleapi/models"
)

func CreateName(name *models.Name) error {
	return db.DB.Create(name).Error
}