package models

import "time"

type Geolocation struct {
	ID				uint `json:"id" gorm:"primaryKey"` // primary key
	AddressID 		uint `json:"address_id"`
	Latitude 		float64 `json:"lat"`
	Longitude 		float64 `json:"long"`

	createdAt 		time.Time `json:"-"`
	updatedAt 		time.Time `json:"-"`
}