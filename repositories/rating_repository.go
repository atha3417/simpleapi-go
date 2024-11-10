package repositories

import (
	"simpleapi/config"
	"simpleapi/models"
)

func CreateRating(rating *models.Rating) error {
	return config.DB.Create(rating).Error
}