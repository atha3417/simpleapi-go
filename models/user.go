package models

import "time"

type User struct {
	ID				uint `json:"id" gorm:"primaryKey"`
	Name 			Name `json:"name" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Email 			string `json:"email" gorm:"unique"`
	Username 		string `json:"username" gorm:"unique"`
	Password 		string `json:"password"`
	Phone 			string `json:"phone"`
	Address 		Address `json:"address" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	createdAt 		time.Time `json:"-"`
	updatedAt 		time.Time `json:"-"`
}
