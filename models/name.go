package models

import "time"

type Name struct {
	ID        uint `json:"id" gorm:"primaryKey"` // primary key
	UserID    uint
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	createdAt time.Time `json:"-"`
	updatedAt time.Time `json:"-"`
}