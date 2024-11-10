package repositories

import (
	db "simpleapi/database"
	"simpleapi/models"

	"gorm.io/gorm/clause"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := db.DB.Preload("Address.Geolocation").Preload(clause.Associations).Find(&users).Error
	return users, err
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := db.DB.Preload("Address.Geolocation").Preload(clause.Associations).First(&user, id).Error
	return user, err
}

func CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func UpdateUser(user *models.User) error {
	return db.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return db.DB.Delete(&models.User{}, id).Error
}
