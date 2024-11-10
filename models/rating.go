package models

import "time"

type Rating struct {
	ID 			uint `gorm:"primaryKey;autoIncrement;type:integer"`
	ProductID	uint
	Rate		float64 `json:"rate"`
	Count		int `json:"count"`

	createdAt time.Time `json:"-"`
	updatedAt time.Time `json:"-"`
}