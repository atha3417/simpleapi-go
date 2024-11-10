package models

import "time"

type Rating struct {
	ID 			uint `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID	uint
	Rate		float64 `json:"rate"`
	Count		int `json:"count"`

	createdAt time.Time `json:"-"`
	updatedAt time.Time `json:"-"`
}