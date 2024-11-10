package repositories

import (
	"simpleapi/config"
	"simpleapi/models"

	"gorm.io/gorm/clause"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Preload("Address.Geolocation").Preload(clause.Associations).Find(&users).Error
	return users, err
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := config.DB.Preload("Address.Geolocation").Preload(clause.Associations).First(&user, id).Error
	return user, err
}

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
