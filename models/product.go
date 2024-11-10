package models

import "time"

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Title       string  `json:"title"`
	Price       int		`json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rating  `json:"rating" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	createdAt time.Time `json:"-"`
	updatedAt time.Time `json:"-"`
}