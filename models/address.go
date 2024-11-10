package models

import "time"

type Address struct {
	ID				uint `json:"id" gorm:"primaryKey"` // primary key	
	UserID 			uint `json:"user_id"`
	City 			string `json:"city"`
	Street 			string `json:"street"`
	Number 			int `json:"number"`
	ZipCode 		int `json:"zipcode"`
	Geolocation 	Geolocation `json:"geolocation" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	createdAt 		time.Time `json:"-"`
	updatedAt 		time.Time `json:"-"`
}