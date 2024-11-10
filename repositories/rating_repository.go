package repositories

import (
	db "simpleapi/database"
	"simpleapi/models"
)

func CreateRating(rating *models.Rating) error {
	return db.DB.Create(rating).Error
}